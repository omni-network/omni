package key

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
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
	cmd := exec.CommandContext(ctx, "gcloud", "secrets", "versions", "access", "latest", "--secret", name, "--format=json")

	var stdOut, stdErr bytes.Buffer
	cmd.Stderr = &stdErr
	cmd.Stdout = &stdOut

	// Execute the command and capture the output
	if err := cmd.Run(); err != nil {
		return nil, errors.Wrap(err, "gcloud fetch secret", "out", stdErr.String())
	}

	// Unmarshal the json response
	var resp response
	if err := json.Unmarshal(stdOut.Bytes(), &resp); err != nil {
		return nil, errors.Wrap(err, "unmarshal secret response")
	}

	// Decode the base64 encoded secret data
	bz, err := base64.URLEncoding.DecodeString(resp.Payload.DataBase64)
	if err != nil {
		return nil, errors.Wrap(err, "decode secret data")
	}

	return bz, nil
}

// response from gloud secret manager --format=json.
type response struct {
	Name    string `json:"name"`
	Payload struct {
		DataBase64 string `json:"data"`
		DataCRC32c string `json:"dataCrc32c"`
	} `json:"payload"`
}
