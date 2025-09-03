package xmonitor

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
)

// Start starts the xchain monitoring goroutines.
func Start(
	ctx context.Context,
	network netconf.Network,
	xprovider xchain.Provider,
	cprovider cchain.Provider,
	rpcClients map[uint64]ethclient.Client,
) error {
	// Monitor the head of all chains, including consensus.
	for _, srcChain := range network.Chains {
		headsFunc := func(ctx context.Context) map[ethclient.HeadType]uint64 {
			return getEVMHeads(ctx, rpcClients[srcChain.ID])
		}
		if netconf.IsOmniConsensus(network.ID, srcChain.ID) {
			headsFunc = func(ctx context.Context) map[ethclient.HeadType]uint64 {
				return getConsXHead(ctx, cprovider)
			}
		}

		go monitorHeadsForever(ctx, srcChain, headsFunc)
		for _, chainVer := range srcChain.ChainVersions() {
			go monitorAttestedForever(ctx, chainVer, cprovider, network, xprovider)
		}
	}

	// Monitors below only apply to EVM chains.
	for _, srcChain := range network.EVMChains() {
		for _, dstChain := range network.EVMChains() {
			if srcChain.ID == dstChain.ID {
				continue
			}

			go monitorOffsetsForever(ctx, xprovider, network, srcChain, dstChain)
		}
	}

	go monitorConsOffsetForever(ctx, network, xprovider)

	return nil
}

// monitorConsOffsetForever blocks and periodically monitors the emitted
// offsets for a given consensus chain.
// Note that submitted offsets are not monitored as the consensus chain doesn't support submissions.
func monitorConsOffsetForever(ctx context.Context, network netconf.Network, xprovider xchain.Provider) {
	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := monitorConsOffsetOnce(ctx, network, xprovider)
			if ctx.Err() != nil {
				return
			} else if err != nil {
				log.Error(ctx, "Monitoring consensus stream offsets failed (will retry)", err)

				continue
			}
		}
	}
}

func monitorConsOffsetOnce(ctx context.Context, network netconf.Network, xprovider xchain.Provider) error {
	cChain, ok := network.OmniConsensusChain()
	if !ok {
		return nil
	}

	// Consensus chain messages are broadcast, so query for each EVM chain.
	for _, stream := range network.StreamsFrom(cChain.ID) {
		ref := xchain.ConfRef(stream.ConfLevel())
		emitted, ok, err := xprovider.GetEmittedCursor(ctx, ref, stream)
		if err != nil {
			return errors.Wrap(err, "get emit cursor", "stream", network.StreamName(stream))
		} else if !ok {
			continue
		}

		streamName := network.StreamName(stream)
		emitMsgOffset.WithLabelValues(streamName).Set(float64(emitted.MsgOffset))

		submitted, ok, err := xprovider.GetSubmittedCursor(ctx, xchain.LatestRef, stream)
		if err != nil {
			return errors.Wrap(err, "get submit cursor", "stream", network.StreamName(stream))
		} else if !ok {
			continue
		}

		submitMsgOffset.WithLabelValues(streamName).Set(float64(submitted.MsgOffset))
		submitAttestOffset.WithLabelValues(streamName).Set(float64(submitted.AttestOffset))
	}

	return nil
}

// monitorHeadsForever blocks and periodically monitors the latest/safe/final heads of the given chain.
func monitorHeadsForever(
	ctx context.Context,
	chain netconf.Chain,
	headsFunc func(context.Context) map[ethclient.HeadType]uint64,
) {
	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			heads := headsFunc(ctx)
			for typ, head := range heads {
				headHeight.WithLabelValues(chain.Name, typ.String()).Set(float64(head))
			}
		}
	}
}

// monitorAttestedForever blocks and periodically monitors the halo attested height and offsets of the given chain version.
func monitorAttestedForever(
	ctx context.Context,
	chainVer xchain.ChainVersion,
	cprovider cchain.Provider,
	network netconf.Network,
	xprovider xchain.Provider,
) {
	chainVerName := network.ChainName(chainVer.ID)
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	var lastAttestOffset uint64

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			att, ok, err := cprovider.LatestAttestation(ctx, chainVer)
			if err != nil {
				log.Warn(ctx, "Attest offset monitor failed getting latest attestation (will retry)", err, "chain_version", chainVerName)
				continue
			} else if !ok {
				continue
			} else if att.AttestOffset == lastAttestOffset {
				continue
			}

			attestedHeight.WithLabelValues(chainVerName).Set(float64(att.BlockHeight))
			attestedOffset.WithLabelValues(chainVerName).Set(float64(att.AttestOffset))

			// Query stream offsets of the original xblock from the chain itself.
			for _, stream := range network.StreamsFrom(chainVer.ID) {
				if stream.ConfLevel() != chainVer.ConfLevel {
					continue
				}

				streamName := network.StreamName(stream)

				cursor, _, err := xprovider.GetEmittedCursor(ctx, xchain.HeightRef(att.BlockHeight), stream)
				if err != nil {
					log.Warn(ctx, "Attest offset monitor failed getting emit cursor", err, "stream", streamName)
					continue
				}

				attestedMsgOffset.WithLabelValues(streamName).Set(float64(cursor.MsgOffset))
			}

			lastAttestOffset = att.AttestOffset
		}
	}
}

// getConsXHead returns the latest XBlock height for the consensus chain.
// This is equivalent to the latest validator set id.
func getConsXHead(ctx context.Context, cprovider cchain.Provider) map[ethclient.HeadType]uint64 {
	xblock, ok, err := cprovider.XBlock(ctx, 0, true)
	if err != nil || !ok {
		return nil
	}

	return map[ethclient.HeadType]uint64{
		"xfinal": xblock.BlockHeight,
	}
}

func getEVMHeads(ctx context.Context, client ethclient.Client) map[ethclient.HeadType]uint64 {
	heads := []ethclient.HeadType{
		ethclient.HeadLatest,
		ethclient.HeadSafe,
		ethclient.HeadFinalized,
	}

	resp := make(map[ethclient.HeadType]uint64)
	for _, typ := range heads {
		head, err := client.HeaderByType(ctx, typ)
		if err != nil {
			// Not all chains support all types, so just swallow the errors, this is best effort monitoring.
			continue
		}
		resp[typ] = head.Number.Uint64()
	}

	return resp
}

// monitorOffsetsForever blocks and periodically monitors the emitted and submitted
// offsets for a given source and destination chain.
func monitorOffsetsForever(
	ctx context.Context,
	xprovider xchain.Provider,
	network netconf.Network,
	src, dst netconf.Chain,
) {
	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := monitorOffsetsOnce(ctx, xprovider, network, src, dst)
			if ctx.Err() != nil {
				return
			} else if err != nil {
				log.Error(ctx, "Monitoring stream offsets failed (will retry)", err,
					"src_chain", src.Name, "dst_chain", dst.Name)
			}
		}
	}
}

// monitorOffsetsOnce monitors the emitted and submitted offsets for a given source and
// destination chain.
func monitorOffsetsOnce(
	ctx context.Context,
	xprovider xchain.Provider,
	network netconf.Network,
	src, dst netconf.Chain,
) error {
	var lastErr error
	for _, stream := range network.StreamsBetween(src.ID, dst.ID) {
		// First attempt to get the emit cursor from the cache.
		srcChainVer := xchain.ChainVersion{ID: src.ID, ConfLevel: stream.ConfLevel()}
		height, err := xprovider.ChainVersionHeight(ctx, srcChainVer)
		if err != nil {
			lastErr = errors.Wrap(err, "latest height", "stream", network.StreamName(stream))
			continue
		} else if height < src.DeployHeight {
			continue // Don't monitor chains before finalized.
		}

		confLevel := stream.ConfLevel()
		emitted, _, err := xprovider.GetEmittedCursor(ctx, xchain.ConfRef(confLevel), stream)
		if err != nil {
			lastErr = errors.Wrap(err, "get emit cursor", "stream", stream)
			continue
		}

		submitted, _, err := xprovider.GetSubmittedCursor(ctx, xchain.LatestRef, stream)
		if err != nil {
			lastErr = errors.Wrap(err, "get submit cursor", "stream", network.StreamName(stream), "height", height)
			continue
		}

		name := network.StreamName(stream)
		emitMsgOffset.WithLabelValues(name).Set(float64(emitted.MsgOffset))
		submitMsgOffset.WithLabelValues(name).Set(float64(submitted.MsgOffset))
		submitAttestOffset.WithLabelValues(name).Set(float64(submitted.AttestOffset))
	}

	return lastErr
}
