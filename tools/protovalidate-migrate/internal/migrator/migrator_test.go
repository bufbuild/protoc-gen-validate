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
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMigrator(t *testing.T) {
	t.Parallel()

	tests := []struct {
		srcPath string
		dstPath string
		cfg     Config
	}{
		{
			srcPath: "none.source.proto",
			dstPath: "none.migrated.proto",
		},
		{
			srcPath: "message.source.proto",
			dstPath: "message.migrated.proto",
		},
		{
			srcPath: "message.source.proto",
			dstPath: "message.removed.proto",
			cfg:     Config{RemovePGV: true},
		},
		{
			srcPath: "oneof.source.proto",
			dstPath: "oneof.migrated.proto",
		},
		{
			srcPath: "oneof.source.proto",
			dstPath: "oneof.removed.proto",
			cfg:     Config{RemovePGV: true},
		},
		{
			srcPath: "field_leaves.source.proto",
			dstPath: "field_leaves.migrated.proto",
		},
		{
			srcPath: "field_leaves.source.proto",
			dstPath: "field_leaves.removed.proto",
			cfg:     Config{RemovePGV: true},
		},
		{
			srcPath: "field_msglit.source.proto",
			dstPath: "field_msglit.migrated.proto",
		},
		{
			srcPath: "field_msglit.source.proto",
			dstPath: "field_msglit.removed.proto",
			cfg:     Config{RemovePGV: true},
		},
		{
			srcPath: "field_required.source.proto",
			dstPath: "field_required.migrated.proto",
		},
		{
			srcPath: "field_required.source.proto",
			dstPath: "field_required.removed.proto",
			cfg:     Config{RemovePGV: true},
		},
		{
			srcPath: "field_skipped.source.proto",
			dstPath: "field_skipped.migrated.proto",
		},
		{
			srcPath: "field_skipped.source.proto",
			dstPath: "field_skipped.removed.proto",
			cfg:     Config{RemovePGV: true},
		},
		{
			srcPath: "field_wkt.source.proto",
			dstPath: "field_wkt.migrated.proto",
		},
		{
			srcPath: "field_wkt.source.proto",
			dstPath: "field_wkt.removed.proto",
			cfg:     Config{RemovePGV: true},
		},
		{
			srcPath: "field_no_sparse.source.proto",
			dstPath: "field_no_sparse.migrated.proto",
		},
		{
			srcPath: "field_no_sparse.source.proto",
			dstPath: "field_no_sparse.removed.proto",
			cfg:     Config{RemovePGV: true},
		},
		{
			srcPath: "mixed.source.proto",
			dstPath: "mixed.migrated.proto",
		},
		{
			srcPath: "mixed.source.proto",
			dstPath: "mixed.removed.proto",
			cfg:     Config{RemovePGV: true},
		},
		{
			srcPath: "mixed.source.proto",
			dstPath: "mixed.replaced.proto",
			cfg:     Config{ReplacePV: true},
		},
		{
			srcPath: "mixed.source.proto",
			dstPath: "mixed.removed_replaced.proto",
			cfg:     Config{RemovePGV: true, ReplacePV: true},
		},
		{
			srcPath: "field_ignore_empty.source.proto",
			dstPath: "field_ignore_empty.migrated.proto",
		},
	}

	pathPrefix := "../../../../proto/testing/tests/migrate/"
	for _, tc := range tests {
		test := tc
		t.Run(test.dstPath, func(t *testing.T) {
			t.Parallel()
			test.cfg.PGVImport = "validate/validate.proto"
			test.cfg.PVImport = "buf/validate/validate.proto"
			m := NewMigrator(test.cfg)

			srcFile, err := os.Open(filepath.Join(pathPrefix, test.srcPath))
			require.NoError(t, err)
			defer srcFile.Close()

			outBuf := &bytes.Buffer{}
			err = m.MigrateFile(test.srcPath, srcFile, outBuf)
			require.NoError(t, err)

			diffBuf := &bytes.Buffer{}
			cmd := exec.Command("git", "diff",
				"--src-prefix=[actual]",
				"--dst-prefix=[expected] ",
				"--color",
				"--no-index",
				"--minimal",
				"--",
				"-",
				filepath.Join(pathPrefix, test.dstPath),
			)
			cmd.Stdin = outBuf
			cmd.Stdout = diffBuf
			cmd.Stderr = os.Stderr
			require.NoErrorf(t, cmd.Run(),
				"actual output does not match expected\n%s",
				diffBuf.String())
		})
	}
}
