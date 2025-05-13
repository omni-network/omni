package fundthresh_test

import (
	"fmt"
	"testing"

	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/solver/fundthresh"
	"github.com/omni-network/omni/solver/rebalance"
)

//go:generate go test . -golden -clean

func TestThresholdReference(t *testing.T) {
	t.Parallel()

	golden := make(map[string]map[string]string)

	for _, token := range rebalance.Tokens() {
		thresh := fundthresh.Get(token)

		key := fmt.Sprintf("%s:%s", evmchain.Name(token.ChainID), token.Symbol)

		golden[key] = map[string]string{
			"min":     token.FormatAmt(thresh.Min()),
			"target":  token.FormatAmt(thresh.Target()),
			"minSwap": token.FormatAmt(thresh.MinSwap()),
			"maxSwap": token.FormatAmt(thresh.MaxSwap()),
		}

		if thresh.NeverSurplus() {
			golden[key]["surplus"] = "inf"
		} else {
			golden[key]["surplus"] = token.FormatAmt(thresh.Surplus())
		}
	}

	tutil.RequireGoldenJSON(t, golden, tutil.WithFilename("thresholds.json"))
}
