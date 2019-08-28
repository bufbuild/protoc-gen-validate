package protolock

import (
	"os"
	"path/filepath"
)

type Config struct {
	LockDir   string
	ProtoRoot string
	Ignore    string
	UpToDate  bool
}

func NewConfig(lockDir, protoRoot, ignores string, upToDate bool) (*Config, error) {
	l, err := filepath.Abs(lockDir)
	if err != nil {
		return nil, err
	}
	p, err := filepath.Abs(protoRoot)
	if err != nil {
		return nil, err
	}

	return &Config{
		LockDir:   l,
		ProtoRoot: p,
		Ignore:    ignores,
		UpToDate:  upToDate,
	}, nil
}

func (cfg *Config) LockFileExists() bool {
	_, err := os.Stat(cfg.LockFilePath())
	return err == nil && !os.IsNotExist(err)
}

func (cfg *Config) LockFilePath() string {
	return filepath.Join(cfg.LockDir, LockFileName)
}
