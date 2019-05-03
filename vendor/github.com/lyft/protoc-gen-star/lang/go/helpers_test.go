package pgsgo

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/golang/protobuf/proto"
	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
	pgs "github.com/lyft/protoc-gen-star"
	"github.com/stretchr/testify/require"
)

func readCodeGenReq(t *testing.T, dir ...string) *plugin_go.CodeGeneratorRequest {
	dirs := append(append([]string{"testdata"}, dir...), "code_generator_request.pb.bin")
	filename := filepath.Join(dirs...)

	data, err := ioutil.ReadFile(filename)
	require.NoError(t, err, "unable to read CDR at %q", filename)

	req := &plugin_go.CodeGeneratorRequest{}
	err = proto.Unmarshal(data, req)
	require.NoError(t, err, "unable to unmarshal CDR data at %q", filename)

	return req
}

func buildGraph(t *testing.T, dir ...string) pgs.AST {
	d := pgs.InitMockDebugger()
	ast := pgs.ProcessCodeGeneratorRequest(d, readCodeGenReq(t, dir...))
	require.False(t, d.Failed(), "failed to build graph (see previous log statements)")
	return ast
}

func loadContext(t *testing.T, dir ...string) Context {
	dirs := append(append([]string{"testdata"}, dir...), "params")
	filename := filepath.Join(dirs...)

	data, err := ioutil.ReadFile(filename)
	require.NoError(t, err, "unable to read params at %q", filename)

	params := pgs.ParseParameters(strings.TrimSpace(string(data)))
	return InitContext(params)
}
