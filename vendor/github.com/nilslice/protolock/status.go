package protolock

import (
	"errors"
	"os"
)

// ErrOutOfDate indicates that the locked definitions are ahead or behind of
// the source proto definitions within the tree.
var ErrOutOfDate = errors.New("proto.lock file is not up-to-date with source")

// Status will report on any issues encountered when comparing the updated tree
// of parsed proto files and the current proto.lock file.
func Status(cfg Config) (*Report, error) {
	updated, err := getUpdatedLock(cfg)
	if err != nil {
		return nil, err
	}

	lockFile, err := openLockFile(cfg)
	if err != nil {
		if os.IsNotExist(err) {
			msg := `no "proto.lock" file found, first run "init"`
			return nil, errors.New(msg)
		}
		return nil, err
	}
	defer lockFile.Close()

	current, err := FromReader(lockFile)
	if err != nil {
		return nil, err
	}

	report, err := Compare(current, *updated)
	if err != nil {
		return report, err
	}

	// only check if current and updated are equal if up-to-date flag is true
	if cfg.UpToDate && !current.Equal(updated) {
		err = ErrOutOfDate
	}
	return report, err
}
