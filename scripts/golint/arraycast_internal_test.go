package main

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestArrayCast(t *testing.T) {
	t.Parallel()

	content := `package main

type Array [2]int

func main() {
	var slice []int
	a := Array(slice) // want "panicable array cast"
	_ = Array(a[:]) // want "panicable array cast"
	_ = Array(append(slice, 0)) // want "panicable array cast"
	_ = [3]int(slice) // want "panicable array cast"

	_ = Array([2]int{1, 2}) // OK
	_ = Array(Array{}) // OK
}
`

	dir, cleanup, err := analysistest.WriteFiles(map[string]string{"main/main.go": content})
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()
	_ = analysistest.Run(t, dir, arrayCastAnalyzer, "main")
}
