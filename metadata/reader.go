package metadata

import (
	_ "embed"
)

//go:embed base.txt
var baseInfo string

//go:embed build.txt
var buildInfo string

// BaseInfo returns basic metadata information.
func BaseInfo() string {
	return baseInfo
}

// BuildInfo return current build information.
func BuildInfo() string {
	return buildInfo
}
