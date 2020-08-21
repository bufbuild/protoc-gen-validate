package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/reflect/protoreflect"

	_ "github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/go_proto_field_name"
	_ "github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/other_package/go_proto_field_name"
	harness "github.com/envoyproxy/protoc-gen-validate/tests/harness/go"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	checkErr(err)

	tc := new(harness.TestCase)
	checkErr(proto.Unmarshal(b, tc))

	da := new(ptypes.DynamicAny)
	checkErr(ptypes.UnmarshalAny(tc.Message, da))
	v := da.Message.(interface {
		Validate() error
	})

	checkMsg(da.Message, v.Validate())
}

func checkMsg(message proto.Message, err error) {
	if err == nil {
		resp(&harness.TestResult{Valid: true}) // no error means no error message to be checked
	}
	fieldError, ok := err.(interface{ Field() string })
	if !ok {
		resp(&harness.TestResult{Error: true,
			Reason: fmt.Sprintf("error does not implement the `interface{ Field() string }`")})
	}
	field := fieldError.Field()
	if field == "" {
		resp(&harness.TestResult{Error: true,
			Reason: fmt.Sprintf("error does not contain the field")})
	}
	msg := proto.MessageReflect(message)
	fieldNames := listAllFields(msg)
	if !contains(field, fieldNames) {
		resp(&harness.TestResult{Error: true,
			Reason: fmt.Sprintf("error message '%s' does not contain the field '%s'", err.Error(), field)})
	}
	resp(&harness.TestResult{Reason: err.Error()})
}

func checkErr(err error) {
	if err == nil {
		return
	}

	resp(&harness.TestResult{
		Error:  true,
		Reason: err.Error(),
	})
}

func resp(result *harness.TestResult) {
	if b, err := proto.Marshal(result); err != nil {
		log.Fatalf("could not marshal response: %v", err)
	} else if _, err = os.Stdout.Write(b); err != nil {
		log.Fatalf("could not write response: %v", err)
	}
	os.Exit(0)
}

func listAllFields(msg protoreflect.Message) []string {
	var fieldNames []string
	fields := msg.Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		oneOf := fields.Get(i).ContainingOneof()
		if oneOf == nil {
			fieldNames = append(fieldNames, string(fields.Get(i).Name()))
		} else {
			fieldNames = append(fieldNames, string(oneOf.Name()))
			for j := 0; j < oneOf.Fields().Len(); j++ {
				fieldNames = append(fieldNames, string(oneOf.Fields().Get(i).Name()))
			}
		}
	}
	return fieldNames
}

func contains(s string, slice []string) bool {
	for _, contained := range slice {
		if strings.Contains(s, contained) {
			return true
		}
	}
	return false
}
