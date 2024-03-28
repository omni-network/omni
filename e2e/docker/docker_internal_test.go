package docker

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestGethVersion checks if the geth version is up to date.
func TestGethVersion(t *testing.T) {
	t.Parallel()

	resp, err := http.Get("https://api.github.com/repos/ethereum/go-ethereum/releases/latest") //nolint:noctx // Test is ok
	require.NoError(t, err)
	defer resp.Body.Close()

	bz, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var v version
	require.NoError(t, json.Unmarshal(bz, &v))

	require.Equal(t, gethTag, v.TagName, "A new version of geth has been released, update `gethTag`")
}

type version struct {
	TagName string `json:"tag_name"`
}
