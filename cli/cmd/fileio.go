package cmd

import (
	"context"
	"io"
	"net/http"
	"os"

	"github.com/omni-network/omni/lib/errors"
)

func copyFile(src string, dest string) error {
	// Open the source file.
	srcFile, err := os.Open(src)
	if err != nil {
		return errors.Wrap(err, "open source file")
	}
	defer srcFile.Close()

	// Create the destination file.
	destFile, err := os.Create(dest)
	if err != nil {
		return errors.Wrap(err, "create destination file")
	}
	defer destFile.Close()

	// Copy the file contents.
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return errors.Wrap(err, "copy file")
	}

	// Set the same permissions as the source file.
	srcInfo, err := os.Stat(src)
	if err != nil {
		return errors.Wrap(err, "get source file info")
	}

	if err := os.Chmod(dest, srcInfo.Mode()); err != nil {
		return errors.Wrap(err, "set file permissions")
	}

	return nil
}

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
