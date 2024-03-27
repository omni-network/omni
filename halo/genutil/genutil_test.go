package genutil_test

import (
	"testing"
	"time"

	"github.com/omni-network/omni/halo/genutil"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"

	k1 "github.com/cometbft/cometbft/crypto/secp256k1"

	_ "github.com/omni-network/omni/halo/app" // To init SDK config.
)

//go:generate go test . -golden -clean

func TestMakeGenesis(t *testing.T) {
	t.Parallel()
	timestamp := time.Unix(1, 0)

	val1 := k1.GenPrivKeySecp256k1([]byte("secret1")).PubKey()
	val2 := k1.GenPrivKeySecp256k1([]byte("secret2")).PubKey()

	resp, err := genutil.MakeGenesis(netconf.Simnet, timestamp, val1, val2)
	tutil.RequireNoError(t, err)

	tutil.RequireGoldenJSON(t, resp)
}
