//go:build tools
// +build tools

// tools is a dummy package that will be ignored for builds, but included for dependencies
package registry

import (
	_ "github.com/google/wire/cmd/wire"
)
