// Package version provides some injected build time information.
package version

var (
	// BuildSha is the commit hash while building.
	BuildSha string = "unknown"

	// BuildTag is the git tag from which this build was created.
	BuildTag string = "unknown"
)
