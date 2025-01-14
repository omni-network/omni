package cmd

import (
	"context"
	"io"
	"net/http"
	"os"

	"github.com/omni-network/omni/lib/errors"
)

func downloadFile(ctx context.Context, srcURL string, destFilePath string) error {
	// Build an HTTP GET request with an injected context.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, srcURL, nil)
	if err != nil {
		return errors.Wrap(err, "build GET request")
	}

	// Send the request.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "send request")
	}
	defer resp.Body.Close()

	// Check if the HTTP status is OK.
	if resp.StatusCode != http.StatusOK {
		return errors.Wrap(err, "download file", "status_code", resp.StatusCode)
	}

	// Create the destination file.
	outFile, err := os.Create(destFilePath)
	if err != nil {
		return errors.Wrap(err, "create file")
	}
	defer outFile.Close()

	// Copy the response body to the file.
	if _, err = io.Copy(outFile, resp.Body); err != nil {
		return errors.Wrap(err, "copy content to file")
	}

	return nil
}
