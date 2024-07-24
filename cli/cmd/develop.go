package cmd

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/omni-network/omni/lib/errors"

	"github.com/spf13/cobra"
)

const (
	defaultTemplate       = "hello-world-template"
	defaultTemplateCommit = "main" // Commit hash of the template, update as needed
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
		RunE: func(cmd *cobra.Command, _ []string) error {
			return newForgeProjectTemplate(cmd.Context(), cfg)
		},
	}

	bindDeveloperForgeProjectConfig(cmd, &cfg)

	return cmd
}

type developerForgeProjectConfig struct {
	templateName string
}

// newForgeProjectTemplate creates a new project by cloning a repository, checking out a specific commit if necessary, and initializing a new Git repository.
func newForgeProjectTemplate(ctx context.Context, cfg developerForgeProjectConfig) error {
	destinationPath, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "failed to get current working directory")
	}

	// Clone the repository
	//nolint:gosec // This is a command-line tool, not a library
	cmd := exec.CommandContext(ctx, "git", "clone", "https://github.com/omni-network/"+cfg.templateName+".git", destinationPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return errors.Wrap(err, "failed to clone repository")
	}

	// Check out the specific commit if using the default template
	if cfg.templateName == defaultTemplate {
		cmd = exec.CommandContext(ctx, "git", "checkout", defaultTemplateCommit)
		cmd.Dir = destinationPath
		cmd.Stdout = nil // Suppress output
		cmd.Stderr = nil // Suppress output
		if err := cmd.Run(); err != nil {
			return errors.Wrap(err, "failed to checkout commit")
		}
	}

	// Initialize submodules
	cmd = exec.CommandContext(ctx, "git", "submodule", "update", "--init", "--recursive")
	cmd.Dir = destinationPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return errors.Wrap(err, "failed to update submodules")
	}

	// Delete the .git directory to remove history
	err = os.RemoveAll(filepath.Join(destinationPath, ".git"))
	if err != nil {
		return errors.Wrap(err, "failed to remove .git directory")
	}

	// Initialize a new Git repository
	cmd = exec.CommandContext(ctx, "git", "init")
	cmd.Dir = destinationPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return errors.Wrap(err, "failed to initialize new git repository")
	}

	return nil
}
