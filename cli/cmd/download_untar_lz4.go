package cmd

import (
	"archive/tar"
	"context"
	"io"
	"net/http"
	"os"

	"github.com/omni-network/omni/lib/errors"

	"github.com/pierrec/lz4/v4"
)

func downloadUntarLz4(ctx context.Context, srcURL string, outputDir string) error {
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

	if err := untarLz4(resp.Body, outputDir); err != nil {
		return errors.Wrap(err, "decompress tar lz4", "url", srcURL)
	}

	return nil
}

func untarLz4(stream io.Reader, outputDir string) error {
	// Open the .tar.lz4 archive.
	tarLZ4Reader := tar.NewReader(lz4.NewReader(stream))

	// Extract files from the tar archive.
	for {
		header, err := tarLZ4Reader.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return errors.Wrap(err, "read tar archive")
		}

		// Determine the output path.
		outputPath := outputDir + "/" + header.Name

		switch header.Typeflag {
		case tar.TypeDir:
			// Create a directory.
			// #nosec G115: ignoring potential integer overflow in int64 to uint32 conversion
			if err := os.MkdirAll(outputPath, os.FileMode(header.Mode)); err != nil {
				return errors.Wrap(err, "create directory")
			}
		case tar.TypeReg:
			// Create a file.
			outFile, err := os.Create(outputPath)
			if err != nil {
				return errors.Wrap(err, "create file")
			}

			// Copy the file contents.
			if _, err := io.CopyN(outFile, tarLZ4Reader, header.Size); err != nil {
				return errors.Wrap(err, "write file")
			}

			if err := outFile.Close(); err != nil {
				return errors.Wrap(err, "close file")
			}

			// Set permissions.
			// #nosec G115: ignoring potential integer overflow in int64 to uint32 conversion
			if err := os.Chmod(outputPath, os.FileMode(header.Mode)); err != nil {
				return errors.Wrap(err, "set file permissions")
			}
		default:
			return errors.New("unsupported type", "type", header.Typeflag)
		}
	}

	return nil
}
