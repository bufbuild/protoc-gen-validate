package main

import (
	"fmt"
	"github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/go"
	"io/ioutil"
	"log"
	"os"

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

	msg, hasValidate := da.Message.(interface {
		Validate(bool) error
	})

	var multierr error
	if isIgnored {
		// confirm that ignored messages don't have a validate method
		if hasValidate {
			err = fmt.Errorf("ignored message has Validate(bool) method")
		}
	} else if !hasValidate {
		err = fmt.Errorf("non-ignored message is missing Validate(bool)")
	} else {
		err = msg.Validate(false)
		multierr = msg.Validate(true)
	}
	checkValid(err, multierr)
}

func checkValid(err, multierr error) {
	if err == nil && multierr == nil {
		resp(&harness.TestResult{Valid: true})
	} else {
		resp(&harness.TestResult{Reason: err.Error(), AllReasons: mergeReasons(nil, multierr)})
	}
}

func mergeReasons(reasons []string, err error) []string {
	multi, ok := err.(interface{ AllErrors() []error })
	if !ok {
		caused, ok := err.(interface{ Cause() error })
		if !ok || caused.Cause() == nil {
			return append(reasons, err.Error())
		}
		return mergeReasons(reasons, caused.Cause())
	}
	for _, err := range multi.AllErrors() {
		reasons = mergeReasons(reasons, err)
	}
	return reasons
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
