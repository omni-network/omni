package cmd

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"strings"

	"github.com/omni-network/omni/lib/errors"

	"github.com/spf13/cobra"
)

func newDeveloperCmds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "developer",
		Short: "XApp development commands",
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(
		newForgeProjectCmd(),
	)

	return cmd
}

func newForgeProjectCmd() *cobra.Command {
	var cfg developerForgeProjectConfig

	cmd := &cobra.Command{
		Use:   "new",
		Short: "Scaffold a Forge project",
		Long:  `Scaffold a new Forge project accompanied by simple mocked testing and a multi-chain deployment script.`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return newForgeProjectTemplate(cmd.Context(), cfg)
		},
	}

	bindDeveloperForgeProjectConfig(cmd, &cfg)

	return cmd
}

type developerForgeProjectConfig struct {
	templateURL string
}

// newForgeProjectTemplate creates a new project using the forge template.
func newForgeProjectTemplate(_ context.Context, cfg developerForgeProjectConfig) error {
	// Check if forge is installed
	if !checkForgeInstalled() {
		// Forge is not installed, return an error with a suggestion.
		return &cliError{
			Msg:     "forge is not installed.",
			Suggest: "You can install foundry by visiting https://github.com/foundry-rs/foundry",
		}
	}

	sanitizedInput, err := sanitizeInput(cfg.templateURL)
	if err != nil {
		return &cliError{
			Msg:     "failed to sanitize input",
			Suggest: "Please provide a valid URL",
		}
	}

	// Execute the `forge init` command.
	cmd := exec.Command("forge", "init", "--template", sanitizedInput)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return errors.Wrap(err, "failed to run forge init")
	}

	return nil
}

// checkForgeInstalled checks if forge is installed by attempting to run 'forge --version'.
func checkForgeInstalled() bool {
	cmd := exec.Command("forge", "--version")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	return err == nil // If there is no error, forge is installed.
}

// sanitizeInput validates and sanitizes the input to prevent command injection.
func sanitizeInput(input string) (string, error) {
	// Replace any characters that are not alphanumeric or allowed special characters.
	// You might need to adjust this based on what characters are allowed in your URLs.
	// This example allows alphanumeric characters, slashes, dots, and hyphens.
	var allowedChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-./"
	var sanitized strings.Builder
	for _, char := range input {
		if strings.ContainsAny(string(char), allowedChars) {
			if _, err := sanitized.WriteRune(char); err != nil {
				return "", errors.Wrap(err, "failed to write sanitized input")
			}
		}
	}

	return sanitized.String(), nil
}
