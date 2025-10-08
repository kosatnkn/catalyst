package metadata

import (
	_ "embed"
	"strings"
)

//go:embed name.txt
var name string

//go:embed base.txt
var baseInfo string

//go:embed build.txt
var buildInfo string

// Name returns the name of the repository.
func Name() string {
	return strings.TrimSpace(name)
}

// BaseInfo returns basic metadata information.
func BaseInfo() string {
	return strings.TrimSpace(baseInfo)
}

// BuildInfo return current build information.
func BuildInfo() []string {
	return strings.Split(strings.TrimSpace(buildInfo), "\n")
}
