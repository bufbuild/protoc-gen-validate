// Copyright 2023 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package migrator

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bufbuild/protocompile/ast"
	"github.com/bufbuild/protocompile/parser"
	"github.com/bufbuild/protocompile/reporter"
	"golang.org/x/sync/errgroup"
)

type Migrator struct {
	print   bool
	write   bool
	output  string
	cfg     Config
	Handler *reporter.Handler
}

func NewMigrator(cfg Config) *Migrator {
	handler := reporter.NewHandler(reporter.NewReporter(
		func(err reporter.ErrorWithPos) error { return err },
		func(err reporter.ErrorWithPos) {
			if cfg.Verbose {
				log.Println("warning:", err)
			}
		},
	))
	return &Migrator{
		cfg:     cfg,
		Handler: handler,
	}
}

func (m *Migrator) Migrate(paths ...string) error {
	if m.print && (len(paths) > 1 || filepath.Ext(paths[0]) != ".proto") {
		return errors.New("only a single proto file path can be supplied in print mode")
	}

	grp, _ := errgroup.WithContext(context.Background())

	for _, root := range paths {
		err := filepath.WalkDir(root, func(srcPath string, stat fs.DirEntry, walkErr error) error {
			rootPath := root
			if walkErr != nil {
				return fmt.Errorf("error walking %s: %w", srcPath, walkErr)
			} else if stat.IsDir() || filepath.Ext(srcPath) != ".proto" {
				return nil
			}
			grp.Go(func() error { return m.migrate(rootPath, srcPath, stat) })
			return nil
		})
		if err != nil {
			return err
		}
	}

	return grp.Wait()
}

func (m *Migrator) migrate(rootPath, srcPath string, stat fs.DirEntry) error {
	switch {
	case m.print:
		return m.PrintMigrate(srcPath)
	case m.write:
		fileInfo, err := stat.Info()
		if err != nil {
			return fmt.Errorf("failed to stat %s: %w", srcPath, err)
		}
		return m.InPlaceMigrate(fileInfo, srcPath)
	case m.output != "":
		fileInfo, err := stat.Info()
		if err != nil {
			return fmt.Errorf("failed to stat %s: %w", srcPath, err)
		}
		dstPath := filepath.Join(m.output, strings.TrimPrefix(srcPath, rootPath))
		return m.OutputMigrate(fileInfo, srcPath, dstPath)
	default:
		return errors.New("nothing to migrate")
	}
}

func (m *Migrator) PrintMigrate(srcPath string) error {
	file, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("failed to open %s: %w", srcPath, err)
	}
	defer file.Close()
	return m.MigrateFile(srcPath, file, os.Stdout)
}

func (m *Migrator) InPlaceMigrate(src os.FileInfo, srcPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("failed to open source file %s: %w", srcPath, err)
	}
	defer srcFile.Close()

	dstFile, err := os.CreateTemp("", src.Name())
	if err != nil {
		return fmt.Errorf("failed to create tmp output file: %w", err)
	}
	defer os.Remove(dstFile.Name())
	defer dstFile.Close()

	err = dstFile.Chmod(src.Mode())
	if err != nil {
		return fmt.Errorf("failed to set mode on output file %s: %w", dstFile.Name(), err)
	}

	err = m.MigrateFile(srcPath, srcFile, dstFile)
	if err != nil {
		return err
	}
	srcFile.Close()
	dstFile.Close()

	err = os.Rename(dstFile.Name(), srcFile.Name())
	if err != nil {
		return fmt.Errorf("failed to overwrite source file %s: %w", srcPath, err)
	}

	return nil
}

func (m *Migrator) OutputMigrate(src os.FileInfo, srcPath, dstPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("failed to open source file %s: %w", srcPath, err)
	}
	defer srcFile.Close()

	err = os.MkdirAll(filepath.Dir(dstPath), 0o755)
	if err != nil {
		return fmt.Errorf("failed to create output directories for %s: %w", dstPath, err)
	}

	dstFile, err := os.OpenFile(dstPath, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, src.Mode())
	if err != nil {
		return fmt.Errorf("failed to create destination file %s: %w", dstPath, err)
	}
	defer dstFile.Close()

	return m.MigrateFile(srcPath, srcFile, dstFile)
}

func (m *Migrator) MigrateFile(name string, in io.Reader, out io.Writer) error {
	fileNode, err := parser.Parse(name, in, m.Handler)
	if err != nil {
		return fmt.Errorf("failed to parse %s: %w", name, err)
	}
	err = ast.Visit(fileNode, New(m.cfg, fileNode, out))
	if err != nil {
		return fmt.Errorf("failed to rewrite %s: %w", name, err)
	}
	return nil
}
