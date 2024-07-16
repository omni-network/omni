// Command golint is a custom Go linter.
// It doesn't replace golang-ci-lint which runs all industry standard linters.
// It merely runs a few custom omni-specific linters.
package main

import (
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(uintSubtractAnalyzer)
}
