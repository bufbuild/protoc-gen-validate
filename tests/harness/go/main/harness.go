package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/go"
	_ "github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/go"
	_ "github.com/envoyproxy/protoc-gen-validate/tests/harness/cases/other_package/go"
	"github.com/envoyproxy/protoc-gen-validate/tests/harness/go"
	"google.golang.org/protobuf/proto"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	checkErr(err)

	tc := new(harness.TestCase)
	checkErr(proto.Unmarshal(b, tc))

	msg, err := tc.Message.UnmarshalNew()
	checkErr(err)

	_, isIgnored := msg.(*cases.MessageIgnored)

	vmsg, hasValidate := msg.(interface {
		Validate() error
	})

	if isIgnored {
		// confirm that ignored messages don't have a validate method
		if hasValidate {
			checkErr(fmt.Errorf("ignored message %T has Validate() method", msg))
		}
	} else if !hasValidate {
		checkErr(fmt.Errorf("non-ignored message %T is missing Validate()", msg))
	} else {
		err = vmsg.Validate()
	}
	checkValid(err)
}

func checkValid(err error) {
	if err == nil {
		resp(&harness.TestResult{Valid: true})
	} else {
		resp(&harness.TestResult{Reason: err.Error()})
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
