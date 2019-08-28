package proto

import (
	"os"
	"testing"
)

type counter struct {
	counts map[string]int
}

func (c counter) handleService(s *Service) {
	c.counts["service"] = c.counts["service"] + 1
}

func (c counter) handleRPC(r *RPC) {
	c.counts["rpc"] = c.counts["rpc"] + 1
}

func TestWalkGoogleApisDLP(t *testing.T) {
	if len(os.Getenv("PB")) == 0 {
		t.Skip("PB test not run")
	}
	proto := fetchAndParse(t, "https://raw.githubusercontent.com/gogo/protobuf/master/test/theproto3/theproto3.proto")
	count := counter{counts: map[string]int{}}
	Walk(proto, WithService(count.handleService), WithRPC(count.handleRPC))
	t.Logf("%#v", count)
}
