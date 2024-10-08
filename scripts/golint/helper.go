package main

import (
	"strings"

	"golang.org/x/tools/go/analysis"
)

// isTestPass returns true if the package being analyzed is a test package.
func isTestPass(p *analysis.Pass) bool {
	if strings.HasSuffix(p.Pkg.Name(), "_test") {
		return true
	}

	for _, file := range p.Files {
		pos := p.Fset.Position(file.Package)
		if strings.HasSuffix(pos.Filename, "_test.go") {
			return true
		}
	}

	return false
}
