// commenttypes comments out type definitions in a Go file
// Usage: go run commenttypes.go -- <file> <type1> <type2> ...

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/omni-network/omni/lib/errors"
)

func main() {
	err := run()

	if err != nil {
		fmt.Printf("commenttypes failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("commenttypes succeeded")
}

func run() error {
	if len(os.Args) < 3 {
		return errors.New("usage: go run commenttypes.go -- <file> <type1> <type2> ...")
	}

	// os.Args[0] is the name of the command
	// os.Args[1] is either <file> or  "--" separator

	firstArgIdx := 1
	if os.Args[1] == "--" {
		firstArgIdx = 2
	}

	filepath := os.Args[firstArgIdx]
	types := os.Args[firstArgIdx+1:]

	if len(types) == 0 {
		return errors.New("no types")
	}

	inputF, err := os.Open(filepath)
	if err != nil {
		return errors.Wrap(err, "open file")
	}
	defer inputF.Close()

	scanner := bufio.NewScanner(inputF)
	var content string
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	modified := commentTypes(content, types)

	err = os.WriteFile(filepath, []byte(modified), 0600)
	if err != nil {
		return errors.Wrap(err, "write file")
	}

	return nil
}

// commentTypes comments out type definitions in a Go file.
func commentTypes(content string, types []string) string {
	var modifiedLines []string
	lines := strings.Split(content, "\n")

	// tracks whether we're in a type definition
	//
	// type Name {
	// ...
	// }
	inTypeDef := false

	for _, line := range lines {
		commentline := func() {
			modifiedLines = append(modifiedLines, "// "+line)
		}

		if inTypeDef && isTypeDefEnd(line) {
			inTypeDef = false
			commentline()

			continue
		}

		if inTypeDef {
			commentline()

			continue
		}

		if isTypeDefStart(line, types) {
			inTypeDef = true
			modifiedLines = append(modifiedLines, "// autocommented by commenttypes.go")
			commentline()

			// single line type definition
			if isTypeDefEnd(line) {
				inTypeDef = false
			}

			continue
		}

		// no type definition, just copy the line
		modifiedLines = append(modifiedLines, line)
	}

	return strings.Join(modifiedLines, "\n")
}

func isTypeDefStart(line string, types []string) bool {
	for _, t := range types {
		if strings.HasPrefix(strings.TrimSpace(line), "type "+t) {
			return true
		}
	}

	return false
}

func isTypeDefEnd(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), "}")
}
