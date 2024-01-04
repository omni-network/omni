// Copyright © 2022-2023 Obol Labs Inc. Licensed under the terms of a Business Source License 1.1

// Command verifypr provides a tool to verify omni PRs against our specific conventional commit template.
package main

import (
	"log"
	"os"
)

func main() {
	err := run()
	if err != nil {
		log.Printf("❌ Verification failed: %+v\n", err)
		os.Exit(1)
	}

	log.Println("✅ Verification Success")
}
