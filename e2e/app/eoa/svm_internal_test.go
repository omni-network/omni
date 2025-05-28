package eoa_test

import (
	"bytes"
	"flag"
	"fmt"
	"testing"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/svmutil"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/stretchr/testify/require"
)

var integration = flag.Bool("integration", false, "run integration tests")

// TestGenSVMAddrs generates the evm address to svm address mapping in svmaddrs.toml.
func TestGenSVMAddrs(t *testing.T) {
	t.Parallel()
	if !*integration {
		t.Skip("integration tests not enabled")
	}

	var b bytes.Buffer
	b.WriteString("[addrs]\n")

	unique := make(map[common.Address]bool)
	for _, network := range netconf.All() {
		if network == netconf.Simnet || network.IsProtected() {
			continue
		}

		for _, role := range eoa.AllRoles() {
			pk, err := eoa.PrivateKey(t.Context(), network, role)
			if err != nil {
				t.Logf("Skipping %s %s: %v", network, role, err)
				continue
			}

			svmAddr := svmutil.MapEVMKey(pk).PublicKey()
			ethAddr := crypto.PubkeyToAddress(pk.PublicKey)
			if unique[ethAddr] {
				continue
			}
			unique[ethAddr] = true

			b.WriteString(fmt.Sprintf(`  "%s" = "%s" # %s-%s`, ethAddr, svmAddr, network, role))
			b.WriteString("\n")
		}
	}

	tutil.RequireGoldenBytes(t, b.Bytes(), tutil.WithFilename("../svmaddrs.toml"))
}

func TestSVMAddrs(t *testing.T) {
	t.Parallel()

	account, ok := eoa.AccountForRole(netconf.Devnet, eoa.RoleCold)
	require.True(t, ok)

	svmAddr, err := account.SVMAddress()
	require.NoError(t, err)
	require.Equal(t, "CgtM4BPSVy8LbLZ4dTHsPr7p6d91aoiGJdxokWEkpqh3", svmAddr.String())
}
