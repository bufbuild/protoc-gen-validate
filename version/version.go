package version

import "fmt"

var (
	// Version is the current version of protoc-gen-validate.
	// This can be set at build time using ldflags:
	// go build -ldflags "-X github.com/envoyproxy/protoc-gen-validate/version.Version=v1.0.0"
	Version = "dev"
	
	// Commit is the git commit hash.
	Commit = "unknown"
	
	// BuildDate is the build date.
	BuildDate = "unknown"
)

// String returns the version string.
func String() string {
	return fmt.Sprintf("protoc-gen-validate %s (commit: %s, built: %s)", Version, Commit, BuildDate)
}

