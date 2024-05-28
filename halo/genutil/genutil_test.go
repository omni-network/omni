package genutil_test

import (
	"testing"
	"time"

	"github.com/omni-network/omni/halo/genutil"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/cometbft/cometbft/crypto"
	k1 "github.com/cometbft/cometbft/crypto/secp256k1"

	_ "github.com/omni-network/omni/halo/app" // To init SDK config.
)

func insecureValKeyFromConsKey(consKey crypto.PrivKey) crypto.PrivKey {
	// reorg some bytes
	bz := consKey.Bytes()
	bz[0], bz[1] = bz[1], bz[0]
	return k1.PrivKey(bz)
}

//go:generate go test . -golden -clean

func TestMakeGenesis(t *testing.T) {
	t.Parallel()
	timestamp := time.Unix(1, 0)

	val1ConsKey := k1.GenPrivKeySecp256k1([]byte("secret1"))
	val2ConsKey := k1.GenPrivKeySecp256k1([]byte("secret2"))

	val1ValKey := insecureValKeyFromConsKey(val1ConsKey)
	val2ValKey := insecureValKeyFromConsKey(val2ConsKey)

	val1Addr, err := k1util.PubKeyToAddress(val1ValKey.PubKey())
	tutil.RequireNoError(t, err)

	val2Addr, err := k1util.PubKeyToAddress(val2ValKey.PubKey())
	tutil.RequireNoError(t, err)

	val1 := genutil.Validator{
		ConsPubKey: val1ConsKey.PubKey(),
		Addr:       val1Addr,
	}

	val2 := genutil.Validator{
		ConsPubKey: val2ConsKey.PubKey(),
		Addr:       val2Addr,
	}

	resp, err := genutil.MakeGenesis(netconf.Simnet, timestamp, val1, val2)
	tutil.RequireNoError(t, err)

	tutil.RequireGoldenJSON(t, resp)
}
