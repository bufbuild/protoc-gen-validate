package shared

import (
	"fmt"
	"strings"

	pgs "github.com/lyft/protoc-gen-star"
)

func isEnum(f pgs.Field) bool {
	return f.Type().IsEnum()
}

// enumList - if type is ENUM, enum values are returned
func enumList(f pgs.Field, list []int32) string {
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

// enumVal - if type is ENUM, enum value is returned
func enumVal(f pgs.Field, val int32) string {
	if enum := f.Type().Enum(); enum != nil {
		return enum.Values()[val].Name().String()
	}
	return fmt.Sprint(val)
}
