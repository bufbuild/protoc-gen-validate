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
			checkErr(fmt.Errorf("ignored message %T has Validate(bool) method", da.Message))
		}
	} else if !hasValidate {
		checkErr(fmt.Errorf("non-ignored message %T is missing Validate(bool)", da.Message))
	} else {
		err = msg.Validate(false)
		multierr = msg.Validate(true)
	}
	checkValid(err, multierr)
}

type hasAllErrors interface{ AllErrors() []error }
type hasCause interface{ Cause() error }

func checkValid(err, multierr error) {
	if err == nil && multierr == nil {
		resp(&harness.TestResult{Valid: true})
		return
	}
	if (err != nil) != (multierr != nil) {
		checkErr(fmt.Errorf("different verdict of Validate(false) [%v] vs. Validate(true) [%v]", err, multierr))
		return
	}

	// Extract the message from "lazy" Validate(false), for comparison with Validate(true)
	rootCause := err
	for {
		caused, ok := rootCause.(hasCause)
		if !ok || caused.Cause() == nil {
			break
		}
		rootCause = caused.Cause()
	}

	// Retrieve the messages from "extensive" Validate(true) and compare first one with the "lazy" message
	m, ok := multierr.(hasAllErrors)
	if !ok {
		checkErr(fmt.Errorf("Validate(true) returned error without AllErrors() method: %#v", multierr))
		return
	}
	reasons := mergeReasons(nil, m)
	if rootCause.Error() != reasons[0] {
		checkErr(fmt.Errorf("different first message, Validate(false)==%q, Validate(true)==%q", rootCause.Error(), reasons[0]))
		return
	}

	resp(&harness.TestResult{Reasons: reasons})
}

func mergeReasons(reasons []string, multi hasAllErrors) []string {
	for _, err := range multi.AllErrors() {
		caused, ok := err.(hasCause)
		if ok && caused.Cause() != nil {
			err = caused.Cause()
		}
		multi, ok := err.(hasAllErrors)
		if ok {
			reasons = mergeReasons(reasons, multi)
		} else {
			reasons = append(reasons, err.Error())
		}
	}
	return reasons
}

func checkErr(err error) {
	if err == nil {
		return
	}

	resp(&harness.TestResult{
		Error:   true,
		Reasons: []string{err.Error()},
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
