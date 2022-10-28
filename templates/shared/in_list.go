package shared

import (
	"fmt"
	"strings"

	pgs "github.com/lyft/protoc-gen-star"
)

// inList - returns list between box brackets, if type is ENUM, instead of numbers enum values are used
func inList(f pgs.Field, list []int32) string {
	stringList := make([]string, 0, len(list))
	if enum := f.Type().Enum(); enum != nil {
		for _, n := range list {
			stringList = append(stringList, enum.Values()[n].Name().String())
		}
	} else {
		for _, n := range list {
			stringList = append(stringList, fmt.Sprint(n))
		}
	}
	return "[" + strings.Join(stringList, " ") + "]"
}
