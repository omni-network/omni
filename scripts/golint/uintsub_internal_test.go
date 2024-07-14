package main

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestUintSub(t *testing.T) {
	t.Parallel()

	content := `package main

func main() {
	var a uint64 = 1
	c := a - b()    // want "uint subtraction"
	d := c + b()    // This is ok
	_ = 5 / (d - 3) // want "uint subtraction"
}

func b() uint64 {
	return 3
}

// uintsub bypasses the linter by using 'a-b' variable names.
func uintsub(a, b uint64) uint64 {
	return a - b // The linter ignores 'a-b' specifically
}
`

	dir, cleanup, err := analysistest.WriteFiles(map[string]string{"main/main.go": content})
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()
	_ = analysistest.Run(t, dir, uintSubtractAnalyzer, "main")
}
