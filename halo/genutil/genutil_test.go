package genutil_test

import (
	"testing"
	"time"

	uluwatu1 "github.com/omni-network/omni/halo/app/upgrades/uluwatu"
	"github.com/omni-network/omni/halo/genutil"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"

	k1 "github.com/cometbft/cometbft/crypto/secp256k1"

	"github.com/ethereum/go-ethereum/common"

	_ "github.com/omni-network/omni/halo/sdk" // To init SDK config.
)

//go:generate go test . -golden -clean

func TestMakeGenesis(t *testing.T) {
	t.Parallel()
	timestamp := time.Unix(1, 0)

	val1 := k1.GenPrivKeySecp256k1([]byte("secret1")).PubKey()
	val2 := k1.GenPrivKeySecp256k1([]byte("secret2")).PubKey()

	executionBlockHash := common.BytesToHash([]byte("blockhash"))

	resp, err := genutil.MakeGenesis(t.Context(), netconf.Simnet, timestamp, executionBlockHash, uluwatu1.UpgradeName, val1, val2)
	tutil.RequireNoError(t, err)

	tutil.RequireGoldenJSON(t, resp)
}
