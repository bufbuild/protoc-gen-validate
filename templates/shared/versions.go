package shared

import (
	"runtime/debug"
	"strings"
)

// Versions read build information embedded in the running binary, will be output into generated file
func Versions() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}
	var b strings.Builder
	b.WriteString("// versions:")
	for _, v := range info.Deps {
		b.WriteString("\n")
		b.WriteString("//  ")
		b.WriteString(v.Path)
		b.WriteString(" ")
		b.WriteString(v.Version)
	}
	return b.String()
}