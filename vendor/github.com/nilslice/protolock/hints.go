package protolock

import (
	"errors"
	"fmt"
	"strings"

	"github.com/emicklei/proto"
)

const (
	// CommentSkip tells the parse step to skip the comparable entity.
	CommentSkip = "@protolock:skip"

	// commentInternal is used for tests
	commentInternal = "@protolock:internal"
)

var (
	// ErrSkipEntry indicates that the CommentSkip hint was found.
	ErrSkipEntry = errors.New("protolock: skip entry hint encountered")

	// errInternalTest indicates that the internal test hint was found.
	errInternalTest = errors.New("protolock: internal hint encountered")
)

func checkComments(v interface{}) []error {
	var errs []error
	switch v.(type) {
	case *proto.Enum:
		e := v.(*proto.Enum)
		errs = append(errs, hints(e.Comment)...)

	case *proto.Message:
		m := v.(*proto.Message)
		errs = append(errs, hints(m.Comment)...)

	case *proto.Service:
		s := v.(*proto.Service)
		errs = append(errs, hints(s.Comment)...)
	}

	return errs
}

func hints(c *proto.Comment) []error {
	if c == nil {
		return nil
	}

	var errs []error
	for _, line := range c.Lines {
		if strings.Contains(line, CommentSkip) {
			debugHint(c, CommentSkip)
			errs = append(errs, ErrSkipEntry)
		}

		if strings.Contains(line, commentInternal) {
			debugHint(c, commentInternal)
			errs = append(errs, errInternalTest)
		}
	}

	return errs
}

func debugHint(c *proto.Comment, hint string) {
	if debug {
		fmt.Println(
			"HINT:", hint,
			fmt.Sprintf(
				"%s:%d:%d",
				c.Position.Filename, c.Position.Line, c.Position.Column,
			),
		)
	}
}
