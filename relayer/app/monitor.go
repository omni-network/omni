package relayer

import (
	"context"
	"net/http"
	"time"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// startMonitoring starts the monitoring goroutines.
func startMonitoring(ctx context.Context, network netconf.Network, xprovider xchain.Provider,
	cprovider cchain.Provider, addr common.Address, rpcClients map[uint64]ethclient.Client) {
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
		go monitorAttestedForever(ctx, srcChain, cprovider, xprovider, network)
	}

	// Monitors below only apply to EVM chains.
	for _, srcChain := range network.EVMChains() {
		go monitorAccountForever(ctx, addr, srcChain.Name, rpcClients[srcChain.ID])
		for _, dstChain := range network.EVMChains() {
			if srcChain.ID == dstChain.ID {
				continue
			}

			go monitorOffsetsForever(ctx, xprovider, network, srcChain, dstChain)
		}
	}

	go monitorConsOffsetForever(ctx, network, xprovider)
}

// monitorConsOffsetsForever blocks and periodically monitors the emitted
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

	// Consensus chain messages are broadacst, so query for each EVM chain.
	for _, stream := range network.StreamsFrom(cChain.ID) {
		ref := xchain.EmitRef{ConfLevel: ptr(stream.ConfLevel())}
		emitted, ok, err := xprovider.GetEmittedCursor(ctx, ref, stream)
		if err != nil {
			return err
		} else if !ok {
			continue
		}

		streamName := network.StreamName(stream)
		emitMsgOffset.WithLabelValues(streamName).Set(float64(emitted.MsgOffset))

		submitted, ok, err := xprovider.GetSubmittedCursor(ctx, stream)
		if err != nil {
			return err
		} else if !ok {
			continue
		}

		submitMsgOffset.WithLabelValues(streamName).Set(float64(submitted.MsgOffset))
		submitBlockOffset.WithLabelValues(streamName).Set(float64(submitted.BlockOffset))
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

// monitorAttestedForever blocks and periodically monitors the halo attested height and offsets of the given chain.
func monitorAttestedForever(
	ctx context.Context,
	srcChain netconf.Chain,
	cprovider cchain.Provider,
	xprovider xchain.Provider,
	network netconf.Network,
) {
	chainVer := srcChain.ChainVersions()[0]

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			// First get attested head (so it can't be lower than heads).
			att, err := getAttested(ctx, chainVer, cprovider)
			// Then populate gauges "at the same time" so they update "atomically".
			if err != nil {
				log.Error(ctx, "Monitoring attested failed (will retry)", err, "chain", srcChain.Name)
				continue
			}

			attestedHeight.WithLabelValues(srcChain.Name).Set(float64(att.BlockHeight))
			attestedBlockOffset.WithLabelValues(srcChain.Name).Set(float64(att.BlockOffset))

			// Query stream emit offsets for the original xblock from the chain itself.
			for _, stream := range network.StreamsFrom(srcChain.ID) {
				ref := xchain.EmitRef{
					Height: &att.BlockHeight,
				}

				name := network.StreamName(stream)

				cursor, _, err := xprovider.GetEmittedCursor(ctx, ref, stream)
				if err != nil {
					log.Error(ctx, "Fetching stream offsets for attestation failed (will retry)", err, "stream", name)
					continue
				}

				attestedMsgOffset.WithLabelValues(name).Set(float64(cursor.MsgOffset))
			}
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
		"xfinal": xblock.BlockOffset,
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

// monitorAttestedOnce monitors of the latest attested height by chain.
func getAttested(ctx context.Context, chainVer xchain.ChainVersion, cprovider cchain.Provider) (xchain.Attestation, error) {
	att, ok, err := cprovider.LatestAttestation(ctx, chainVer)
	if err != nil {
		return xchain.Attestation{}, errors.Wrap(err, "latest attestation")
	} else if !ok {
		return xchain.Attestation{}, nil
	}

	return att, nil
}

// monitorAccountsForever blocks and periodically monitors the relayer accounts
// for the given chain.
func monitorAccountForever(ctx context.Context, addr common.Address, chainName string, client ethclient.Client) {
	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := monitorAccountOnce(ctx, addr, chainName, client)
			if ctx.Err() != nil {
				return
			} else if err != nil {
				log.Error(ctx, "Monitoring account failed (will retry)", err,
					"chain", chainName)

				continue
			}
		}
	}
}

// monitorAccountOnce monitors the relayer account for the given chain.
func monitorAccountOnce(ctx context.Context, addr common.Address, chainName string, client ethclient.Client) error {
	balance, err := client.EtherBalanceAt(ctx, addr)
	if err != nil {
		return errors.Wrap(err, "balance at")
	}

	nonce, err := client.NonceAt(ctx, addr, nil)
	if err != nil {
		return errors.Wrap(err, "nonce at")
	}

	accountBalance.WithLabelValues(chainName).Set(balance)
	accountNonce.WithLabelValues(chainName).Set(float64(nonce))

	return nil
}

// monitorOffsetsForever blocks and periodically monitors the emitted and submitted
// offsets for a given source and destination chain.
func monitorOffsetsForever(ctx context.Context, xprovider xchain.Provider, network netconf.Network, src, dst netconf.Chain,
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
func monitorOffsetsOnce(ctx context.Context, xprovider xchain.Provider, network netconf.Network, src, dst netconf.Chain,
) error {
	var lastErr error
	for _, stream := range network.StreamsBetween(src.ID, dst.ID) {
		ref := xchain.EmitRef{ConfLevel: ptr(stream.ConfLevel())}
		emitted, _, err := xprovider.GetEmittedCursor(ctx, ref, stream)
		if err != nil {
			lastErr = errors.Wrap(err, "get emitted cursor", "stream", network.StreamName(stream))
			continue
		}

		submitted, _, err := xprovider.GetSubmittedCursor(ctx, stream)
		if err != nil {
			lastErr = errors.Wrap(err, "get submit cursor", "stream", network.StreamName(stream))
			continue
		}

		name := network.StreamName(stream)
		emitMsgOffset.WithLabelValues(name).Set(float64(emitted.MsgOffset))
		submitMsgOffset.WithLabelValues(name).Set(float64(submitted.MsgOffset))
		submitBlockOffset.WithLabelValues(name).Set(float64(submitted.BlockOffset))
	}

	return lastErr
}

// serveMonitoring starts a goroutine that serves the monitoring API. It
// returns a channel that will receive an error if the server fails to start.
func serveMonitoring(address string) <-chan error {
	errChan := make(chan error)
	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.Handler())

		srv := &http.Server{
			Addr:              address,
			ReadHeaderTimeout: 5 * time.Second,
			IdleTimeout:       5 * time.Second,
			WriteTimeout:      5 * time.Second,
			Handler:           mux,
		}
		errChan <- errors.Wrap(srv.ListenAndServe(), "serve monitoring")
	}()

	return errChan
}

func ptr[T any](t T) *T {
	return &t
}
