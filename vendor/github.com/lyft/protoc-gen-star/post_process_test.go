package pgs

type mockPP struct {
	match bool
	out   []byte
	err   error
}

func (pp mockPP) Match(a Artifact) bool             { return pp.match }
func (pp mockPP) Process(in []byte) ([]byte, error) { return pp.out, pp.err }
