package e2e_test

import (
	"context"
	"crypto/sha256"
	"math/big"
	"testing"
	"time"

	"github.com/omni-network/omni/e2e/app/geth"
	"github.com/omni-network/omni/halo/genutil/evm"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/umath"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/misc/eip4844"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/ethereum/go-ethereum/params"

	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
)

// TestGethConfig ensure that the geth config is setup correctly.
func TestGethConfig(t *testing.T) {
	t.Parallel()
	testOmniEVM(t, func(t *testing.T, client ethclient.Client) {
		t.Helper()
		ctx := context.Background()

		cfg := geth.MakeGethConfig(geth.Config{})

		block, err := client.BlockByNumber(ctx, big.NewInt(1))
		require.NoError(t, err)

		require.EqualValues(t, int(cfg.Eth.Miner.GasCeil), int(block.GasLimit()))
		require.Equal(t, big.NewInt(0), block.Difficulty())

		require.NotNil(t, block.BeaconRoot())
		require.NotEqual(t, common.Hash{}, *block.BeaconRoot())
	})
}

func TestBlobTx(t *testing.T) {
	t.Parallel()
	testOmniEVM(t, func(t *testing.T, client ethclient.Client) {
		t.Helper()
		err := sendBlobTx(context.Background(), client, evm.DefaultChainConfig(netconf.Devnet))
		require.NoError(t, err)
	})
}

func sendBlobTx(ctx context.Context, client ethclient.Client, config *params.ChainConfig) error {
	privKey := anvil.DevPrivateKey1()
	addr := crypto.PubkeyToAddress(privKey.PublicKey)

	// Create a blob tx
	blobTx, err := makeUnsignedBlobTx(ctx, client, config, addr)
	if err != nil {
		return err
	}
	tx, err := ethtypes.SignNewTx(privKey,
		ethtypes.NewCancunSigner(umath.NewBigInt(netconf.Devnet.Static().OmniExecutionChainID)),
		blobTx)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	// Submit it and confirm status
	err = client.SendTransaction(ctx, tx)
	if err != nil {
		return err
	}
	rec, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		return err
	} else if ethtypes.ReceiptStatusSuccessful != rec.Status {
		return errors.New("receipt status not successful")
	}

	return nil
}

// makeUnsignedBlobTx is a utility method to construct a random blob transaction
// without signing it.
// Reference: github.com/ethereum/go-ethereum@v1.14.11/core/txpool/blobpool/blobpool_test.go:184.
func makeUnsignedBlobTx(ctx context.Context, client ethclient.Client, config *params.ChainConfig, from common.Address) (*ethtypes.BlobTx, error) {
	nonce, err := client.NonceAt(ctx, from, nil)
	if err != nil {
		return nil, errors.Wrap(err, "nonce")
	}

	tipCap, baseFee, blobFee, err := estimateGasPrice(ctx, client, config)
	if err != nil {
		return nil, errors.Wrap(err, "estimate gas price")
	}

	emptyBlob := new(kzg4844.Blob)
	emptyBlobCommit, err := kzg4844.BlobToCommitment(emptyBlob)
	if err != nil {
		return nil, errors.Wrap(err, "blob commitment")
	}
	emptyBlobProof, err := kzg4844.ComputeBlobProof(emptyBlob, emptyBlobCommit)
	if err != nil {
		return nil, errors.Wrap(err, "blob proof")
	}
	emptyBlobVHash := kzg4844.CalcBlobHashV1(sha256.New(), &emptyBlobCommit)

	return &ethtypes.BlobTx{
		ChainID:    uint256.NewInt(netconf.Devnet.Static().OmniExecutionChainID),
		Nonce:      nonce,
		GasTipCap:  uint256.NewInt(tipCap.Uint64()),
		GasFeeCap:  uint256.NewInt(baseFee.Uint64()),
		BlobFeeCap: uint256.NewInt(blobFee.Uint64()),
		Gas:        21000,
		BlobHashes: []common.Hash{emptyBlobVHash},
		Sidecar: &ethtypes.BlobTxSidecar{
			Blobs:       []kzg4844.Blob{*emptyBlob},
			Commitments: []kzg4844.Commitment{emptyBlobCommit},
			Proofs:      []kzg4844.Proof{emptyBlobProof},
		},
	}, nil
}

func estimateGasPrice(ctx context.Context, backend ethclient.Client, config *params.ChainConfig) (*big.Int, *big.Int, *big.Int, error) {
	tip, err := backend.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, nil, nil, err
	}

	head, err := backend.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, nil, nil, err
	}
	if head.BaseFee == nil {
		return nil, nil, nil, errors.New("txmgr does not support pre-london blocks that do not have a base fee")
	}

	baseFee := head.BaseFee
	minBase := big.NewInt(params.GWei) // Minimum base fee is 1 GWei.
	if baseFee.Cmp(minBase) < 0 {
		baseFee = minBase
	}

	var blobFee *big.Int
	if head.ExcessBlobGas != nil {
		blobFee = eip4844.CalcBlobFee(config, head)
	}

	// The tip must be at most the base fee.
	if tip.Cmp(baseFee) > 0 {
		tip = baseFee
	}

	return tip, baseFee, blobFee, nil
}
