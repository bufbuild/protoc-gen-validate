// Copyright 2023 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate/conformance/harness"
	protovalidate "github.com/bufbuild/protovalidate/go/v2"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/dynamicpb"
	"google.golang.org/protobuf/types/known/anypb"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("[protovalidate-go] ")

	req := &harness.TestConformanceRequest{}
	if data, err := io.ReadAll(os.Stdin); err != nil {
		log.Fatalf("failed to read input from stdin: %v", err)
	} else if err = proto.Unmarshal(data, req); err != nil {
		log.Fatalf("failed to unmarshal conformance request: %v", err)
	}

	resp, err := TestConformance(req)
	if err != nil {
		log.Fatalf("unable to test conformance: %v", err)
	} else if data, err := proto.Marshal(resp); err != nil {
		log.Fatalf("unable to marshal conformance response: %v", err)
	} else if _, err = os.Stdout.Write(data); err != nil {
		log.Fatalf("unable to write output to stdout: %v", err)
	}
}

func TestConformance(req *harness.TestConformanceRequest) (*harness.TestConformanceResponse, error) {
	files, err := protodesc.NewFiles(req.GetFdset())
	if err != nil {
		err = fmt.Errorf("failed to parse file descriptors: %w", err)
		return nil, err
	}
	val, err := protovalidate.New()
	if err != nil {
		err = fmt.Errorf("failed to initialize validator: %w", err)
		return nil, err
	}
	resp := &harness.TestConformanceResponse{Results: map[string]*harness.TestResult{}}
	for caseName, testCase := range req.GetCases() {
		resp.Results[caseName] = TestCase(val, files, testCase)
	}
	return resp, nil
}

func TestCase(val *protovalidate.Validator, files *protoregistry.Files, testCase *anypb.Any) *harness.TestResult {
	urlParts := strings.Split(testCase.GetTypeUrl(), "/")
	fullName := protoreflect.FullName(urlParts[len(urlParts)-1])
	desc, err := files.FindDescriptorByName(fullName)
	if err != nil {
		return unexpectedErrorResult("unable to find descriptor: %v", err)
	}
	msgDesc, ok := desc.(protoreflect.MessageDescriptor)
	if !ok {
		return unexpectedErrorResult("expected message descriptor, got %T", desc)
	}

	dyn := dynamicpb.NewMessage(msgDesc)
	if err = anypb.UnmarshalTo(testCase, dyn, proto.UnmarshalOptions{}); err != nil {
		return unexpectedErrorResult("unable to unmarshal test case: %v", err)
	}

	err = val.Validate(dyn)
	if err == nil {
		return &harness.TestResult{
			Result: &harness.TestResult_Success{
				Success: true,
			},
		}
	}
	switch res := err.(type) {
	case *protovalidate.ValidationError:
		return &harness.TestResult{
			Result: &harness.TestResult_ValidationError{
				ValidationError: res.ToProto(),
			},
		}
	case *protovalidate.RuntimeError:
		return &harness.TestResult{
			Result: &harness.TestResult_RuntimeError{
				RuntimeError: res.Error(),
			},
		}
	case *protovalidate.CompilationError:
		return &harness.TestResult{
			Result: &harness.TestResult_CompilationError{
				CompilationError: res.Error(),
			},
		}
	default:
		return unexpectedErrorResult("unknown error: %v", err)
	}
}

func unexpectedErrorResult(format string, args ...any) *harness.TestResult {
	return &harness.TestResult{
		Result: &harness.TestResult_UnexpectedError{
			UnexpectedError: fmt.Sprintf(format, args...),
		},
	}
}
