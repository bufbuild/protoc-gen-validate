package main

import (
	"fmt"
	"github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/go"
	"github.com/hashicorp/go-multierror"
	"io/ioutil"
	"log"
	"os"
	"strings"

	_ "github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/go"
	_ "github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/other_package/go"
	"github.com/envoyproxy/protoc-gen-validate/tests/harness/go"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	checkErr(err)

	tc := new(harness.TestCase)
	checkErr(proto.Unmarshal(b, tc))

	da := new(ptypes.DynamicAny)
	checkErr(ptypes.UnmarshalAny(tc.Message, da))

	_, isIgnored := da.Message.(*cases.MessageIgnored)

	msgValidate, hasValidate := da.Message.(interface {
		Validate() error
	})

	msgAllErrors, hasAllErrors := da.Message.(interface {
		AllErrors() error
	})

	if isIgnored {
		// confirm that ignored messages don't have a validate method
		if hasValidate || hasAllErrors {
			err = fmt.Errorf("ignored message has Validate() or AllErrors method")
		}
	} else if !hasValidate || !hasAllErrors {
		err = fmt.Errorf("non-ignored message is missing Validate()")
	} else {
		if tc.TestType == harness.TestType_TestTypeValidate {
			err = msgValidate.Validate()
		} else {
			err = msgAllErrors.AllErrors()
		}
	}
	checkValid(err)
}

func checkValid(err error) {
	if err == nil {
		resp(&harness.TestResult{Valid: true, ErrorCount: 0})
	} else {
		var reason string
		var errorCount int
		if multiErr, ok := err.(*multierror.Error); ok {
			errorCount = len(multiErr.Errors)
			reasons := make([]string, errorCount)
			for i, e := range multiErr.Errors {
				reasons[i] = e.Error()
			}
			reason = strings.Join(reasons, ",")
		} else {
			errorCount = 1
			reason = err.Error()
		}
		resp(&harness.TestResult{Reason: reason, ErrorCount: int32(errorCount)})
	}
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
