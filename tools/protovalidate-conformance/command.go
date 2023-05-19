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
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate/conformance/harness"
	"google.golang.org/protobuf/proto"
)

func runCommand(cfg *config, req *harness.TestConformanceRequest) (*harness.TestConformanceResponse, error) {
	input, err := proto.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to marshal conformance request: %w", err)
	}
	output := &bytes.Buffer{}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.suiteTimeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, cfg.cmd, cfg.args...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = bytes.NewReader(input)
	cmd.Stdout = output

	if err = cmd.Run(); err != nil {
		return nil, fmt.Errorf(
			"failed to execute conformance test: %w", err)
	}

	resp := &harness.TestConformanceResponse{}
	if err := proto.Unmarshal(output.Bytes(), resp); err != nil {
		return nil, fmt.Errorf(
			"failed to unmarshal conformance response: %w", err)
	}
	return resp, nil
}
