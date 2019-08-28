package protolock

import (
	"path/filepath"
	"strings"
)

const (
	// FileSep is the string representation of the OS-specific path separator.
	FileSep = string(filepath.Separator)

	// ProtoSep is an OS-ambiguous path separator to encode into the proto.lock
	// file. Use OsPath and ProtoPath funcs to convert.
	ProtoSep = ":/:"
)

// Protopath is a type to assist in OS filepath transformations
type Protopath string

// OSPath converts a path in the Protopath format to the OS path format
func OSPath(ProtoPath Protopath) Protopath {
	return Protopath(
		strings.Replace(string(ProtoPath), ProtoSep, FileSep, -1),
	)
}

// ProtoPath converts a path in the OS path format to Protopath format
func ProtoPath(OSPath Protopath) Protopath {
	return Protopath(
		strings.Replace(string(OSPath), FileSep, ProtoSep, -1),
	)
}

func (p Protopath) String() string {
	return string(p)
}
