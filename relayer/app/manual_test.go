package relayer_test

import (
	"context"
	"errors"
	"flag"
	"os"
	"testing"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	"github.com/omni-network/omni/lib/xchain/connect"

	"github.com/stretchr/testify/require"
)

// infuraSecretEnv defines Infura secret environment variable name.
const infuraSecretEnv = "INFURA_SECRET"

var (
	flagIntegration = flag.Bool("integration", false, "run integration tests")
	flagNetwork     = flag.String("network", "", "network")
	flagSrcChain    = flag.String("src-chain", "", "source chain name")
	flagDestChain   = flag.String("dest-chain", "", "destination chain name")
	flagShard       = flag.String("shard", "", "latest or finalized shard")
	flagOverride    = flag.Bool("override", false, "use finalized override")
)

//go:generate go test . -v -run=TestManualSubmission -integration -network=staging -src-chain=holesky -dest-chain=omni_evm -shard=latest

// TestManualSubmission is a script to manually do the next relayer submission
// for the provided network+stream.
//
// Note that an infura secret is required as env var INFURA_SECRET.
//
//nolint:paralleltest // Integration test
func TestManualSubmission(t *testing.T) {
	if !*flagIntegration {
		t.Skip("skipping integration test")
		return
	}
	_, ok := os.LookupEnv(infuraSecretEnv)
	require.True(t, ok, "missing "+infuraSecretEnv+" env var")

	ctx := t.Context()
	network := netconf.ID(*flagNetwork)
	streamName := *flagSrcChain + "|" + *flagShard + "|" + *flagDestChain

	srcChain, err := getChainID(*flagSrcChain, network)
	require.NoError(t, err, "src-chain unknown")
	destChain, err := getChainID(*flagDestChain, network)
	require.NoError(t, err, "dest-chain unknown")

	var shard xchain.ShardID
	switch *flagShard {
	case "finalized":
		shard = xchain.ShardFinalized0
	case "latest":
		shard = xchain.ShardLatest0
	case "broadcast":
		shard = xchain.ShardBroadcast0
	default:
		require.Fail(t, "unknown shard")
	}

	log.Info(ctx, "Attempting manual submission", "network", network, "stream", streamName)

	err = manualSubmission(
		ctx,
		network,
		xchain.StreamID{
			SourceChainID: srcChain,
			DestChainID:   destChain,
			ShardID:       shard,
		},
		*flagOverride,
		eoa.RoleXCaller,
	)
	require.NoError(t, err)
}

func getChainID(name string, network netconf.ID) (uint64, error) {
	if name == "omni_evm" {
		return network.Static().OmniExecutionChainID, nil
	}

	meta, ok := evmchain.MetadataByName(name)
	if !ok {
		return 0, errors.New("chain not found")
	}

	return meta.ChainID, nil
}

func manualSubmission(ctx context.Context, network netconf.ID, stream xchain.StreamID, override bool, role eoa.Role) error {
	privkey, err := eoa.PrivateKey(ctx, network, role)
	if err != nil {
		return err
	}
	from, ok := eoa.Address(network, role)
	if !ok {
		return errors.New("address not found")
	}

	conn, err := connect.New(ctx, network,
		connect.WithPrivKey(privkey),
		connect.WithInfuraENV(infuraSecretEnv),
	)
	if err != nil {
		return err
	}

	portal, err := conn.Portal(ctx, stream.DestChainID)
	if err != nil {
		return err
	}

	backend, err := conn.Backend(stream.DestChainID)
	if err != nil {
		return err
	}

	txOpts, err := backend.BindOpts(ctx, from)
	if err != nil {
		return err
	}
	// TODO(corver): if gas-estimation or gas-price is an issue, provide optional flags to override them.

	// Get the latest submit cursor (we need to submit the next msg offset)
	cursor, ok, err := conn.XProvider.GetSubmittedCursor(ctx, xchain.LatestRef, stream)
	if err != nil {
		return err
	} else if !ok {
		log.Warn(ctx, "No submitted cursor found, trying to submit first msg", nil)
		cursor = xchain.SubmitCursor{
			StreamID:     stream,
			AttestOffset: 1,
			MsgOffset:    0,
		}
	}

	log.Info(ctx, "Using submitted cursor", "attest_offset", cursor.AttestOffset, "msg_offset", cursor.MsgOffset)

	// Fetch the attestation and block with the next msg offset
	att, block, err := fetchNextAttBlock(ctx, conn, cursor, override)
	if err != nil {
		return err
	}

	// Make submission
	sub, err := makeSubmission(att, block, cursor)
	if err != nil {
		return err
	}

	// Submit it
	tx, err := portal.Xsubmit(txOpts, sub)
	if err != nil {
		return err
	}

	rec, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	log.Info(ctx, "🎉 Submission sent and included on-chain",
		"link", conn.Network.ID.Static().OmniScanTXURL(tx.Hash()),
		"block", rec.BlockNumber.Uint64(),
	)

	return nil
}

func makeSubmission(att xchain.Attestation, block xchain.Block, cursor xchain.SubmitCursor) (bindings.XSubmission, error) {
	tree, err := xchain.NewMsgTree(block.Msgs)
	if err != nil {
		return bindings.XSubmission{}, err
	} else if att.MsgRoot != tree.MsgRoot() {
		return bindings.XSubmission{}, errors.New("msg root mismatch")
	}

	var msgs []xchain.Msg
	for _, msg := range block.Msgs {
		if msg.StreamOffset == cursor.MsgOffset+1 {
			msgs = append(msgs, msg)
		}
	}

	proof, err := tree.Proof(msgs)
	if err != nil {
		return bindings.XSubmission{}, err
	}

	attRoot, err := xchain.AttestationRoot(att.AttestHeader, att.BlockHeader, tree.MsgRoot())
	if err != nil {
		return bindings.XSubmission{}, err
	} else if root, err := att.AttestationRoot(); err != nil {
		return bindings.XSubmission{}, err
	} else if root != attRoot {
		return bindings.XSubmission{}, errors.New("attestation root mismatch")
	}

	sub := xchain.Submission{
		AttestationRoot: attRoot,
		ValidatorSetID:  att.ValidatorSetID,
		AttHeader:       att.AttestHeader,
		BlockHeader:     block.BlockHeader,
		Msgs:            msgs,
		Proof:           proof.Proof,
		ProofFlags:      proof.ProofFlags,
		Signatures:      att.Signatures,
		DestChainID:     cursor.DestChainID,
	}

	return xchain.SubmissionToBinding(sub), nil
}

func fetchNextAttBlock(ctx context.Context, conn connect.Connector, cursor xchain.SubmitCursor, override bool) (xchain.Attestation, xchain.Block, error) {
	chainVer := cursor.ChainVersion()
	if override {
		chainVer.ConfLevel = xchain.ConfFinalized
	}

	// Stop once we reach latest
	latest, ok, err := conn.CProvider.LatestAttestation(ctx, chainVer)
	if err != nil {
		return xchain.Attestation{}, xchain.Block{}, err
	} else if !ok {
		return xchain.Attestation{}, xchain.Block{}, errors.New("no attestation found")
	}

	errDone := errors.New("done")

	var respAtt xchain.Attestation
	var respBlock xchain.Block
	err = conn.CProvider.StreamAttestations(ctx, chainVer, cursor.AttestOffset, "",
		func(ctx context.Context, att xchain.Attestation) error {
			log.Debug(ctx, "Checking attestation", "attest_offset", att.AttestOffset)
			block, ok, err := conn.XProvider.GetBlock(ctx, xchain.ProviderRequest{
				ChainID:   cursor.SourceChainID,
				Height:    att.BlockHeight,
				ConfLevel: cursor.ShardID.ConfLevel(),
			})
			if err != nil {
				return err
			} else if !ok {
				return errors.New("block not found")
			} else if block.BlockHash != att.BlockHash {
				return errors.New("block hash mismatch")
			}

			for _, msg := range block.Msgs {
				log.Debug(ctx, "Checking block message", "height", block.BlockHeight, "msg_dest", msg.DestChainID, "msg_shard", msg.ShardID, "msg_offset", msg.StreamOffset)
				offsetOk := msg.StreamOffset == cursor.MsgOffset+1
				destOk := msg.DestChainID == cursor.DestChainID
				shardOk := msg.ShardID == cursor.ShardID
				if offsetOk && destOk && shardOk {
					respAtt = att
					respBlock = block
					log.Info(ctx, "Found attestation and msg and block",
						"attest_offset", att.AttestOffset,
						"msg_offset", msg.StreamOffset,
						"block_height", block.BlockHeight,
					)

					return errDone
				}
			}

			if att.AttestOffset == latest.AttestOffset {
				return errors.New("reached latest attestation")
			}

			return nil
		})
	if errors.Is(err, errDone) {
		return respAtt, respBlock, nil
	}

	return xchain.Attestation{}, xchain.Block{}, err
}
