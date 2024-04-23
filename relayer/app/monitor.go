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
	for _, srcChain := range network.Chains {
		go monitorHeightsForever(ctx, srcChain, cprovider, rpcClients[srcChain.ID])
		if srcChain.IsOmniConsensus {
			// Below monitors only apply to EVM chains.
			continue
		}

		go monitorAccountForever(ctx, addr, srcChain.Name, rpcClients[srcChain.ID])
		for _, dstChain := range network.EVMChains() {
			if srcChain.ID == dstChain.ID {
				continue
			}

			go monitorOffsetsForever(ctx, srcChain.ID, dstChain.ID, srcChain.Name, dstChain.Name, xprovider)
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
	for _, destChain := range network.EVMChains() {
		emitted, ok, err := xprovider.GetEmittedCursor(ctx, cChain.ID, destChain.ID)
		if err != nil {
			return err
		} else if !ok {
			return nil
		}

		emitCursor.WithLabelValues(cChain.Name, destChain.Name).Set(float64(emitted.Offset))
	}

	return nil
}

// monitorHeightsForever blocks and periodically monitors the latest/safe/final heads
// and halo attested height of the given chain.
func monitorHeightsForever(ctx context.Context, chain netconf.Chain,
	cprovider cchain.Provider, client ethclient.Client,
) {
	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			// First get attested head (so it can't be lower than heads).
			attested, err := getAttested(ctx, chain.ID, chain.DeployHeight, cprovider)

			// then get chain heads (so it is always higher than attested).
			var heads map[ethclient.HeadType]uint64
			if chain.IsOmniConsensus {
				heads = getConsXHead(ctx, cprovider)
			} else {
				heads = getEVMHeads(ctx, client)
			}

			// Then populate gauges "at the same time" so they update "atomically".
			if err != nil {
				log.Error(ctx, "Monitoring attested failed (will retry)", err,
					"chain", chain.Name)
			} else {
				attestedHeight.WithLabelValues(chain.Name).Set(float64(attested))
			}

			for typ, head := range heads {
				headHeight.WithLabelValues(chain.Name, typ.String()).Set(float64(head))
			}
			stratHeight, ok := heads[ethclient.HeadType(chain.FinalizationStrat)]
			if ok {
				headHeight.WithLabelValues(chain.Name, "xfinal").Set(float64(stratHeight))
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

// monitorAttestedOnce monitors of the latest attested height by chain.
func getAttested(ctx context.Context, chainID uint64, deployHeight uint64, cprovider cchain.Provider) (uint64, error) {
	att, ok, err := cprovider.LatestAttestation(ctx, chainID)
	if err != nil {
		return 0, errors.Wrap(err, "latest attestation")
	} else if !ok {
		return deployHeight - 1, nil
	}

	return att.BlockHeader.BlockHeight, nil
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
func monitorOffsetsForever(ctx context.Context, src, dst uint64, srcChain, dstChain string,
	xprovider xchain.Provider) {
	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := monitorOffsetsOnce(ctx, src, dst, srcChain, dstChain, xprovider)
			if ctx.Err() != nil {
				return
			} else if err != nil {
				log.Error(ctx, "Monitoring stream offsets failed (will retry)", err,
					"src_chain", srcChain, "dst_chain", dstChain)

				continue
			}
		}
	}
}

// monitorOffsetsOnce monitors the emitted and submitted offsets for a given source and
// destination chain.
func monitorOffsetsOnce(ctx context.Context, src, dst uint64, srcChain, dstChain string,
	xprovider xchain.Provider) error {
	emitted, ok, err := xprovider.GetEmittedCursor(ctx, src, dst)
	if err != nil {
		return err
	} else if !ok {
		return nil
	}

	submitted, _, err := xprovider.GetSubmittedCursor(ctx, dst, src)
	if err != nil {
		return err
	}

	emitCursor.WithLabelValues(srcChain, dstChain).Set(float64(emitted.Offset))
	submitCursor.WithLabelValues(srcChain, dstChain).Set(float64(submitted.Offset))

	return nil
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
