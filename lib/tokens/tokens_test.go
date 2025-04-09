package tokens_test

import (
	e2e "github.com/omni-network/omni/e2e/solve"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tutil"
	"testing"
)

//go:generate go test . -golden

// TestGenMockTokens generates mock_tokens.json
// See func mocks() for details.
func TestGenMockTokens(t *testing.T) {
	var tkns []tokens.Token
	for _, mock := range e2e.MockTokens() {
		chainClass := tokens.ClassDevent
		if mock.NetworkID == netconf.Staging {
			chainClass = tokens.ClassTestnet
		}

		tkns = append(tkns, tokens.Token{
			Asset:      mock.Asset,
			Address:    mock.Address(),
			ChainID:    mock.ChainID,
			ChainClass: chainClass,
			IsMock:     true,
		})
	}

	tutil.RequireGoldenJSON(t, tkns, tutil.WithFilename("../mock_tokens.json"))
}
