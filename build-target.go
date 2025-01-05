//go:generate stringer -type BuildTarget -linecomment -output build-target_string.go

package buildtarget

//──────────────────────────────────────────────────────────────────────────────────────────────────

type BuildTarget uint8

const (
	_       BuildTarget = iota
	RELEASE             // RELEASE
	DEVELOP             // DEVELOP
)

//──────────────────────────────────────────────────────────────────────────────────────────────────

var Status BuildTarget
