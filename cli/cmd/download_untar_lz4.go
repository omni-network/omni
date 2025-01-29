package cmd

import (
	"archive/tar"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/omni-network/omni/lib/errors"

	"github.com/pierrec/lz4"
	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
)

func downloadUntarLz4(ctx context.Context, srcURL string, outputDir string, progress *mpb.Progress, clientName string) error {
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

	// Get the content length for the progress bar.
	contentLength := resp.ContentLength
	if contentLength <= 0 {
		return errors.New("unknown content length for progress bar")
	}

	// Create a progress bar.
	bar := progress.New(
		contentLength,
		mpb.BarStyle().Lbound("╢").Filler("▌").Tip("▌").Padding("░").Rbound("╟"),
		mpb.PrependDecorators(
			decor.Name(fmt.Sprintf("Downloading %s snapshot: ", clientName)),
			decor.Counters(decor.SizeB1024(0), "% .1f / % .1f"),
			decor.Name(" Elapsed: "),
			decor.Elapsed(decor.ET_STYLE_GO),
		),
		mpb.AppendDecorators(
			decor.Percentage(decor.WCSyncWidth),
			decor.Name(" ETA: "),
			decor.AverageETA(decor.ET_STYLE_GO),
			decor.Name(" at "),
			decor.EwmaSpeed(decor.SizeB1024(0), "% .1f", 30),
		),
	)

	// Wrap the response body with the progress bar reader.
	progressReader := bar.ProxyReader(resp.Body)
	// Defer cleanup of the progress bar to avoid it being re-rendered.
	defer progressReader.Close()

	if err := untarLz4(progressReader, outputDir); err != nil {
		return errors.Wrap(err, "decompress tar lz4", "url", srcURL)
	}

	// Mark the progress bar as complete.
	bar.SetCurrent(contentLength)

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
