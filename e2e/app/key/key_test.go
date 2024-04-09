package key_test

import (
	"context"
	"flag"
	"testing"

	"github.com/omni-network/omni/e2e/app/key"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/stretchr/testify/require"
)

func TestKeys(t *testing.T) {
	t.Parallel()
	for _, typ := range []key.Type{key.Validator, key.P2PConsensus, key.P2PExecution, key.EOA} {
		t.Run(typ.String(), func(t *testing.T) {
			t.Parallel()

			key1 := key.Generate(typ)

			addrA, err := key1.Addr()
			require.NoError(t, err)

			key2, err := key.FromBytes(typ, key1.Bytes())
			require.NoError(t, err)

			addrB, err := key2.Addr()
			require.NoError(t, err)
			require.Equal(t, addrA, addrB)

			require.True(t, key1.Equals(key2.PrivKey))

			ecdsaKey, err := key1.ECDSA()
			switch typ {
			case key.Validator, key.P2PExecution, key.EOA:
				require.NoError(t, err)
				addrC := crypto.PubkeyToAddress(ecdsaKey.PublicKey)
				require.Equal(t, addrA, addrC.Hex())
			case key.P2PConsensus:
				require.Error(t, err)
			}
		})
	}
}

var integration = flag.Bool("integration", false, "run integration tests")

//go:generate go test . -integration -run=TestIntegration -v

func TestIntegration(t *testing.T) {
	t.Parallel()
	if !*integration {
		t.Skip("skipping integration tests")
	}

	for _, typ := range []key.Type{key.Validator, key.P2PConsensus, key.P2PExecution} {
		t.Run(typ.String(), func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			network := netconf.Simnet
			name := "deleteme"

			k, err := key.UploadNew(ctx, key.UploadConfig{
				Network: network,
				Name:    name,
				Type:    typ,
			})
			require.NoError(t, err)

			addr, err := k.Addr()
			require.NoError(t, err)

			k2, err := key.Download(ctx, network, name, typ, addr)
			tutil.RequireNoError(t, err)

			require.True(t, k.Equals(k2.PrivKey))

			key.DeleteSecretForT(ctx, t, network, name, typ, addr)
		})
	}
}
