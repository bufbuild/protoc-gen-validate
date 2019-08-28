package proto

import (
	"net/http"
	"os"
	"testing"
)

func fetchAndParse(t *testing.T, url string) *Proto {
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(url, err)
	}
	defer resp.Body.Close()
	parser := NewParser(resp.Body)
	def, err := parser.Parse()
	if err != nil {
		t.Fatal(url, err)
	}
	t.Log("elements:", len(def.Elements))
	return def
}

// PB=y go test -v -run ^TestPublicProtoDefinitions$
func TestPublicProtoDefinitions(t *testing.T) {
	if len(os.Getenv("PB")) == 0 {
		t.Skip("PB test not run")
	}
	for _, each := range []string{
		"https://raw.githubusercontent.com/gogo/protobuf/master/test/thetest.proto",
		"https://raw.githubusercontent.com/gogo/protobuf/master/test/theproto3/theproto3.proto",
		"https://raw.githubusercontent.com/googleapis/googleapis/master/google/privacy/dlp/v2/dlp.proto",
		// "https://raw.githubusercontent.com/envoyproxy/data-plane-api/master/envoy/api/v2/auth/cert.proto",
	} {
		def := fetchAndParse(t, each)
		checkParent(def, t)
	}
}
