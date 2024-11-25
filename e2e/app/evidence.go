package app

import (
	"context"
	"time"

	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/crypto/tmhash"
	"github.com/cometbft/cometbft/evidence"
	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
	cmttypes "github.com/cometbft/cometbft/types"
)

// awaitSlashed returns nil when the provided validator is slashed.
func awaitSlashed(ctx context.Context, def Definition, valAddr crypto.Address) error {
	client, err := def.Testnet.BroadcastNode().Client()
	if err != nil {
		return errors.Wrap(err, "broadcast client")
	}

	cprov := provider.NewABCI(client, def.Testnet.Network)

	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	for {
		if err := ctx.Err(); err != nil {
			return errors.Wrap(err, "timeout")
		}

		infos, err := cprov.SDKSigningInfos(ctx)
		if err != nil {
			return errors.Wrap(err, "signing infos")
		}

		for _, info := range infos {
			if addr, err := info.ConsensusCmtAddr(); err != nil {
				return errors.Wrap(err, "consensus address")
			} else if addr.String() != valAddr.String() {
				continue
			}

			// Ensure jailed
			if info.Jailed() {
				log.Info(ctx, "Validator slashed", "address", valAddr)
				return nil
			}
		}
	}
}

// injectEvidence takes a running testnet and generates an
// DuplicateVoteEvidence against the last validator and
// broadcasts it via the broadcast node rpc endpoint `/broadcast_evidence`.
// It returns the address of the validator that was slashed.
//
// This was copied from cometbft/test/e2e/runner/evidence.go.
func injectEvidence(ctx context.Context, testnet types.Testnet) (crypto.Address, error) {
	chainID := testnet.Network.Static().OmniConsensusChainIDStr()

	client, err := testnet.BroadcastNode().Client()
	if err != nil {
		return nil, errors.Wrap(err, "client")
	}

	// request the latest block and validator set from the node
	blockRes, err := client.Block(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "block")
	}
	evidenceHeight := blockRes.Block.Height
	waitHeight := blockRes.Block.Height + 3

	nValidators := 100
	valRes, err := client.Validators(ctx, &evidenceHeight, nil, &nValidators)
	if err != nil {
		return nil, errors.Wrap(err, "validators")
	}

	valSet, err := cmttypes.ValidatorSetFromExistingValidators(valRes.Validators)
	if err != nil {
		return nil, errors.Wrap(err, "valset")
	}

	// Get the private keys of all the validators in the network
	privVals := getPrivateValidatorKeys(testnet.Testnet)

	// Slash the last validator
	valIdx := len(privVals) - 1
	dve, err := generateDuplicateVoteEvidence(privVals, valIdx, evidenceHeight, valSet, chainID, blockRes.Block.Time)
	if err != nil {
		return nil, err
	}

	// Ensure it is valid
	if err := evidence.VerifyDuplicateVote(dve, chainID, valSet); err != nil {
		return nil, errors.Wrap(err, "verify evidence")
	}

	// Wait for the node to reach the height above the forged height so that
	// it is able to validate the evidence
	_, err = waitForNode(ctx, testnet.BroadcastNode(), waitHeight, time.Minute)
	if err != nil {
		return nil, err
	}

	_, err = client.BroadcastEvidence(ctx, dve)
	if err != nil {
		return nil, errors.Wrap(err, "broadcast evidence")
	}

	log.Info(ctx, "Injected double signing evidence", "evidence_height", evidenceHeight, "submit_height", waitHeight)

	return privVals[valIdx].PrivKey.PubKey().Address(), nil
}

func getPrivateValidatorKeys(testnet *e2e.Testnet) []cmttypes.MockPV {
	var privVals []cmttypes.MockPV
	for _, node := range testnet.Nodes {
		if node.Mode == e2e.ModeValidator {
			// Create mock private validators from the validators private key. MockPV is
			// stateless which means we can double vote and do other funky stuff
			privVals = append(privVals, cmttypes.NewMockPVWithParams(node.PrivvalKey, false, false))
		}
	}

	return privVals
}

// generateDuplicateVoteEvidence returns duplicate vote evidence against the valIdx validator.
// This was copied from cometbft/test/e2e/runner/evidence.go.
func generateDuplicateVoteEvidence(
	privVals []cmttypes.MockPV,
	valIdx int,
	height int64,
	vals *cmttypes.ValidatorSet,
	chainID string,
	time time.Time,
) (*cmttypes.DuplicateVoteEvidence, error) {
	voteA, err := cmttypes.MakeVote(privVals[valIdx], chainID, int32(valIdx), height, 0, 2, makeRandomBlockID(), time) //nolint:gosec // Overflow not possible
	if err != nil {
		return nil, errors.Wrap(err, "make vote")
	}
	voteB, err := cmttypes.MakeVote(privVals[valIdx], chainID, int32(valIdx), height, 0, 2, makeRandomBlockID(), time) //nolint:gosec // Overflow not possible
	if err != nil {
		return nil, errors.Wrap(err, "make vote")
	}
	ev, err := cmttypes.NewDuplicateVoteEvidence(voteA, voteB, time, vals)
	if err != nil {
		return nil, errors.Wrap(err, "new evidence")
	}

	return ev, nil
}

// makeRandomBlockID was copied from cometbft/test/e2e/runner/evidence.go.
func makeRandomBlockID() cmttypes.BlockID {
	return makeBlockID(crypto.CRandBytes(tmhash.Size), 100, crypto.CRandBytes(tmhash.Size))
}

// makeBlockID was copied from cometbft/test/e2e/runner/evidence.go.
func makeBlockID(hash []byte, partSetSize uint32, partSetHash []byte) cmttypes.BlockID {
	var (
		h   = make([]byte, tmhash.Size)
		psH = make([]byte, tmhash.Size)
	)
	copy(h, hash)
	copy(psH, partSetHash)

	return cmttypes.BlockID{
		Hash: h,
		PartSetHeader: cmttypes.PartSetHeader{
			Total: partSetSize,
			Hash:  psH,
		},
	}
}
