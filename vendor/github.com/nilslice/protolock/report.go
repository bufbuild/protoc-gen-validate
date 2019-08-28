package protolock

import (
	"fmt"
	"io"
	"sort"
)

// HandleReport checks a report for warnigs and writes warnings to an io.Writer.
// The returned int (an exit code) is 1 if warnings are encountered.
func HandleReport(report *Report, w io.Writer, err error) (int, error) {
	if len(report.Warnings) > 0 {
		// sort the warnings so they are grouped by file location
		orderByPathAndMessage(report.Warnings)

		for _, warning := range report.Warnings {
			fmt.Fprintf(
				w,
				"CONFLICT: %s [%s]\n",
				warning.Message, warning.Filepath,
			)
		}
		return 1, err
	}

	return 0, err
}

func orderByPathAndMessage(warnings []Warning) {
	sort.Slice(warnings, func(i, j int) bool {
		if warnings[i].Filepath < warnings[j].Filepath {
			return true
		}
		if warnings[i].Filepath > warnings[j].Filepath {
			return false
		}
		return warnings[i].Message < warnings[j].Message
	})
}
