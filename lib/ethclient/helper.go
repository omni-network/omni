package ethclient

import "strings"

// IsErrMethodNotAvailable returns true if the error indicates that the
// RPC method is not available on the server.
// This is useful to handle some endpoints that are not always enabled/configured.
func IsErrMethodNotAvailable(err error) bool {
	if err == nil {
		return false
	}

	// See go-ethereum/rpc/errors.go: return fmt.Sprintf("the method %s does not exist/is not available", e.method)

	return strings.Contains(err.Error(), "the method") &&
		strings.Contains(err.Error(), "not available")
}
