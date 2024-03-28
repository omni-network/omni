package key

import (
	"bytes"
	"context"
	"os/exec"

	"github.com/omni-network/omni/lib/errors"
)

// createGCPSecret creates a new GCP Secret Manager secret.
func createGCPSecret(ctx context.Context, name string, value []byte, labels map[string]string) error {
	// Init the gcloud command to create a secret
	cmd := exec.CommandContext(ctx, "gcloud", "secrets", "create", name, "--data-file=-")

	// Add labels to the command
	for k, v := range labels {
		cmd.Args = append(cmd.Args, "--labels="+k+"="+v)
	}

	// Set the secret value as the input for the command
	cmd.Stdin = bytes.NewReader(value)

	// Execute the command
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Wrap(err, "gcloud create secret", "out", string(out))
	}

	return nil
}

// getGCPSecret fetches the latest version of the specified secret.
func getGCPSecret(ctx context.Context, name string) ([]byte, error) {
	// Init the gcloud command to access the secret
	cmd := exec.CommandContext(ctx, "gcloud", "secrets", "versions", "access", "latest", "--secret", name)

	var stdOut, stdErr bytes.Buffer
	cmd.Stderr = &stdErr
	cmd.Stdout = &stdOut

	// Execute the command and capture the output
	if err := cmd.Run(); err != nil {
		return nil, errors.Wrap(err, "gcloud fetch secret", "out", stdErr.String())
	}

	return stdOut.Bytes(), nil
}
