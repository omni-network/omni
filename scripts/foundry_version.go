package scripts

import (
	_ "embed"
)

//go:embed foundry_version.txt
var foundryVersion string

func FoundryVersion() string {
	return foundryVersion
}
