package keeper

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"testing"

	"github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/stretchr/testify/require"
)

// AttestTable returns the attestations ORM table.
func (k *Keeper) AttestTable() AttestationTable {
	return k.attTable
}

// SignatureTable returns the attestations ORM table.
func (k *Keeper) SignatureTable() SignatureTable {
	return k.sigTable
}

func TestWindowCompose(t *testing.T) {
	t.Parallel()
	const window = 64
	const mid = 32
	const bigMid = 256

	const in = 0
	const above = 1
	const below = -1

	tests := []struct {
		Mid      uint64
		Target   uint64
		Expected int
	}{
		{mid, mid, in},
		{mid, mid + 1, in},
		{mid, mid - 1, in},
		{mid, mid + window, in},
		{mid, 0, in},
		{mid, mid + window + 1, above},
		{mid, mid + window + window, above},
		{bigMid, bigMid - window - 1, below},
		{bigMid, bigMid - window - window, below},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("mid_%d_target_%d", tt.Mid, tt.Target), func(t *testing.T) {
			t.Parallel()
			got := windowCompare(window, tt.Mid, tt.Target)
			if got != tt.Expected {
				t.Errorf("Test %d: Expected %d, got %d", i, tt.Expected, got)
			}
		})
	}
}

func TestVerifyAggVotes(t *testing.T) {
	t.Parallel()
	const (
		voteExtLimit = 4
		cChainID     = 999
		windowOffset = 64
		power        = 10
		srcChain1    = 1
	)
	val1 := genPrivKey(t)
	val2 := genPrivKey(t)
	val3 := genPrivKey(t)
	valX := genPrivKey(t) // Not in set
	vals := []*ecdsa.PrivateKey{val1, val2, val3, valX}
	portalReg := testPortalRegistry{srcChain1: []xchain.ConfLevel{xchain.ConfFinalized}}

	valset := ValSet{Vals: map[common.Address]int64{
		addr(val1): power,
		addr(val2): power,
		addr(val3): power,
	}}

	att1 := &types.AttestHeader{
		ConsensusChainId: cChainID,
		SourceChainId:    srcChain1,
		ConfLevel:        uint32(xchain.ConfFinalized),
		AttestOffset:     1,
	}

	block1A := &types.BlockHeader{
		ChainId:     srcChain1,
		BlockHeight: 1,
		BlockHash:   tutil.RandomHash().Bytes(),
	}

	block1B := &types.BlockHeader{
		ChainId:     srcChain1,
		BlockHeight: 1,
		BlockHash:   tutil.RandomHash().Bytes(),
	}
	msgRootA := tutil.RandomHash().Bytes()
	msgRootB := tutil.RandomHash().Bytes()

	windowCompareFunc := func(ctx context.Context, chainVersion xchain.ChainVersion, attestOffset uint64) (int, error) {
		if attestOffset > windowOffset {
			return 1, nil
		}

		return 0, nil
	}

	tests := []struct {
		name   string
		aggs   []*types.AggVote
		errStr string
	}{
		{
			name: "no votes",
		},
		{
			name: "same att header, different blocks",
			aggs: []*types.AggVote{
				{
					AttestHeader: att1,
					BlockHeader:  block1A,
					MsgRoot:      msgRootA,
					Signatures:   toSign(val1, val2),
				},
				{
					AttestHeader: att1,
					BlockHeader:  block1B,
					MsgRoot:      msgRootB,
					Signatures:   toSign(val3),
				},
			},
		},
		{
			name:   "same att header, same blocks",
			errStr: "invalid duplicate aggregate vote",
			aggs: []*types.AggVote{
				{
					AttestHeader: att1,
					BlockHeader:  block1A,
					MsgRoot:      msgRootA,
					Signatures:   toSign(val1, val2),
				},
				{
					AttestHeader: att1,
					BlockHeader:  block1A,
					MsgRoot:      msgRootA,
					Signatures:   toSign(val3),
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			keeper := Keeper{
				portalRegistry: portalReg,
				namer:          netconf.SimnetNetwork().ChainVersionName,
				voter:          nil,
				voteWindow:     0,
				voteExtLimit:   voteExtLimit,
			}

			for i := 0; i < len(test.aggs); i++ {
				test.aggs[i].Signatures = sign(t, test.aggs[i], vals)
			}

			err := keeper.verifyAggVotes(ctx, cChainID, valset, test.aggs, windowCompareFunc)
			if test.errStr == "" {
				require.NoError(t, err)
			} else {
				require.ErrorContains(t, err, test.errStr)
			}
		})
	}
}

func sign(t *testing.T, vote *types.AggVote, vals []*ecdsa.PrivateKey) []*types.SigTuple {
	t.Helper()
	attRoot, err := vote.AttestationRoot()
	require.NoError(t, err)

	for i := 0; i < len(vote.Signatures); i++ {
		sig := vote.Signatures[i]
		var found bool
		for _, val := range vals {
			if bytes.Equal(sig.ValidatorAddress, addr(val).Bytes()) {
				pk, err := k1util.StdPrivKeyToComet(val)
				require.NoError(t, err)

				s, err := k1util.Sign(pk, attRoot)
				require.NoError(t, err)

				vote.Signatures[i].Signature = s[:]
				found = true

				break
			}
		}

		require.True(t, found, "signature not found")
	}

	return vote.Signatures
}

func addr(val *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(val.PublicKey)
}

func toSign(vals ...*ecdsa.PrivateKey) []*types.SigTuple {
	var resp []*types.SigTuple
	for _, val := range vals {
		resp = append(resp, &types.SigTuple{
			ValidatorAddress: addr(val).Bytes(),
		})
	}

	return resp
}

func genPrivKey(t *testing.T) *ecdsa.PrivateKey {
	t.Helper()

	privKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	return privKey
}

type testPortalRegistry map[uint64][]xchain.ConfLevel

func (t testPortalRegistry) ConfLevels(context.Context) (map[uint64][]xchain.ConfLevel, error) {
	return t, nil
}
