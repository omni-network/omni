package cmd

import (
	"archive/tar"
	"bytes"
	"context"
	"io"
	"os"
	"strconv"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/pierrec/lz4/v4"
)

func decompressTarLz4(ctx context.Context, inputFile string, outputDir string) error {
	// Open the .tar.lz4 file.
	file, err := os.Open(inputFile)
	if err != nil {
		return errors.Wrap(err, "open file")
	}
	defer file.Close()

	// Create an LZ4 reader.
	lz4Reader := lz4.NewReader(file)

	// Decompress into memory or stream directly.
	var buf bytes.Buffer
	if _, err = io.Copy(&buf, lz4Reader); err != nil {
		return errors.Wrap(err, "decompress lz4")
	}

	// Open the .tar archive.
	tarReader := tar.NewReader(&buf)

	// Extract files from the tar archive.
	for {
		header, err := tarReader.Next()
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
			if header.Mode < 0 || header.Mode > 0o777 {
				return errors.New("invalid mode for directory: " + strconv.FormatInt(header.Mode, 10))
			}
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

			// copyFile the file contents.
			if _, err := io.CopyN(outFile, tarReader, header.Size); err != nil {
				if err := outFile.Close(); err != nil {
					return errors.Wrap(err, "close file")
				}

				return errors.Wrap(err, "write file")
			}

			if err := outFile.Close(); err != nil {
				return errors.Wrap(err, "close file")
			}

			if header.Mode < 0 || header.Mode > 0o777 {
				return errors.New("invalid mode for file: " + strconv.FormatInt(header.Mode, 10))
			}
			// Set permissions.
			// #nosec G115: ignoring potential integer overflow in int64 to uint32 conversion
			if err := os.Chmod(outputPath, os.FileMode(header.Mode)); err != nil {
				return errors.Wrap(err, "set file permissions")
			}
		default:
			// Handle other types (symlinks, etc.) if necessary.
			log.Info(ctx, "Ignoring unsupported type", "type", header.Typeflag)
		}
	}

	return nil
}
