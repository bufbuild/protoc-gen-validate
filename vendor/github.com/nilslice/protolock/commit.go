package protolock

import (
	"fmt"
	"io"
	"os"
)

// Commit will return an io.Reader with the lock representation data for caller to
// use as needed.
func Commit(cfg Config) (io.Reader, error) {
	if !cfg.LockFileExists() {
		fmt.Println(`no "proto.lock" file found, first run "init"`)
		os.Exit(1)
	}

	updated, err := getUpdatedLock(cfg)
	if err != nil {
		return nil, err
	}

	return readerFromProtolock(updated)
}
