package cmd

import (
	"fmt"
)

// reportError reports an error to the user. If it's a known error, it will
// print a helpful message.
func reportError(err error) {
	msg := err.Error()

	printErr := func() {
		fmt.Printf("❌ registration failed: \033[1m%s\033[0m\n", msg)
	}

	switch msg {
	case "already registered":
		fmt.Println("You're already registered.")
	case "not an operator":
		printErr()
		fmt.Println("🤔 Have you registered as an operator with Eigen-Layer?")
	case "not in allowlist":
		fmt.Println("Your operator address is not in the allowlist.")
	case "max operators reached":
		fmt.Println("The maximum number of operators has been reached.")
	case "min stake not met":
		fmt.Println("You do not meet the minimum stake requirement.")
	case "invalid delegation manager address":
		printErr()
		fmt.Println("🤔 Is el_delegation_manager set correctly in your operator.yaml?")
	case "no contract code at given address":
		printErr()
		fmt.Println("🤔 Is eth_rpc_url set correctly in your operator.yaml?")
	default:
		printErr()
	}
}
