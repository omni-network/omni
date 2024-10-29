package provider

import (
	"context"

	atypes "github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tracer"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type attestations struct {
	query      atypes.QueryClient
	service    cmtservice.ServiceClient
	chainNamer func(xchain.ChainVersion) string
}

func (a attestations) Find(
	ctx context.Context,
	chainVer xchain.ChainVersion,
	fromOffset uint64,
	cachedHeight uint64,
) ([]xchain.Attestation, uint64, error) {
	log.Debug(ctx, "[atts] find", 33, "cached", cachedHeight, "from", fromOffset)

	const endpoint = "fetch_attestations"
	defer latency(endpoint)()

	ctx, span := tracer.Start(ctx, spanName(endpoint))
	defer span.End()

	chainName := a.chainNamer(chainVer)

	// try fetching from latest height
	atts, ok, err := a.fromAtHeight(ctx, chainVer, fromOffset, 0)
	if err != nil {
		incQueryErr(endpoint)
		return nil, 0, errors.Wrap(err, "abci query attestations-from")
	} else if ok {
		debugLogAtts(ctx, "[atts] find [49]", atts)
		fetchStepsMetrics(chainName, 0, 0, 0)

		return atts, 0, nil
	}

	earliestAttestationAtLatestHeight, ok, err := a.earliestAtHeight(ctx, chainVer, 0)
	if err != nil {
		incQueryErr(endpoint)

		return nil, 0, errors.Wrap(err, "abci query earliest-attestation-in-state")
	}

	// Either no attestations have happened yet, or the queried fromOffset is in the "future"
	// Caller has to wait and retry in both cases
	if !ok || earliestAttestationAtLatestHeight.AttestOffset < fromOffset {
		// First attestation hasn't happened yet, return empty
		return []xchain.Attestation{}, 0, nil
	}

	latestBlockResp, err := a.service.GetLatestBlock(ctx, &cmtservice.GetLatestBlockRequest{})
	if err != nil {
		return []xchain.Attestation{}, 0, errors.Wrap(err, "query latest block")
	}
	latestHeight := uint64(latestBlockResp.SdkBlock.Header.Height)

	var startHeight, endHeight uint64
	// if we are performing search for the first time and no previous height was cached do a full search
	if cachedHeight == 0 {
		startHeight, endHeight, err = a.lookbackRange(ctx, chainVer, fromOffset, latestHeight)
		if err != nil {
			incQueryErr(endpoint)
			return nil, 0, errors.Wrap(err, "lookback search")
		}
	} else {
		startHeight = cachedHeight + 1 // we move to the next height
		endHeight = latestHeight

		// try first doing a simple forward search on next height before binary search
		atts, ok, err := a.fromAtHeight(ctx, chainVer, fromOffset, startHeight)
		if err != nil {
			return nil, 0, err
		}
		if ok {
			return atts, startHeight, nil
		}
	}

	offsetHeight, err := a.binarySearch(ctx, chainVer, fromOffset, startHeight, endHeight)
	if err != nil {
		incQueryErr(endpoint)
		return nil, 0, errors.Wrap(err, "binary search")
	}

	atts, ok, err = a.fromAtHeight(ctx, chainVer, fromOffset, offsetHeight)
	if err != nil {
		incQueryErr(endpoint)
		return nil, 0, errors.Wrap(err, "abci query attestations-from")
	} else if !ok {
		return nil, 0, errors.New("expected to find attestations [BUG]")
	}

	return atts, offsetHeight, nil
}

// fromAtHeight searches the consensus state history and
// returns a historical consensus block height that contains an approved attestation
// for the provided chain version at consensus block height.
func (a attestations) fromAtHeight(
	ctx context.Context,
	chainVer xchain.ChainVersion,
	fromOffset uint64,
	height uint64,
) ([]xchain.Attestation, bool, error) {
	resp, err := a.query.AttestationsFrom(withCtxHeight(ctx, height), &atypes.AttestationsFromRequest{
		ChainId:    chainVer.ID,
		ConfLevel:  uint32(chainVer.ConfLevel),
		FromOffset: fromOffset,
	})
	if err != nil {
		return []xchain.Attestation{}, false, errors.Wrap(err, "abci query attestations-from")
	} else if len(resp.Attestations) == 0 {
		return []xchain.Attestation{}, false, nil
	}

	atts, err := atypes.AttestationsFromProto(resp.Attestations)
	if err != nil {
		return []xchain.Attestation{}, false, errors.Wrap(err, "attestations from proto")
	}

	return atts, true, nil
}

// earliestAtHeight returns the earliest approved attestation for the provided chain version
// at the provided consensus block height, or the latest block height if height is 0.
func (a attestations) earliestAtHeight(
	ctx context.Context,
	chainVer xchain.ChainVersion,
	height uint64,
) (xchain.Attestation, bool, error) {
	resp, err := a.query.EarliestAttestation(withCtxHeight(ctx, height),
		&atypes.EarliestAttestationRequest{
			ChainId:   chainVer.ID,
			ConfLevel: uint32(chainVer.ConfLevel),
		})
	if errors.Is(err, sdkerrors.ErrKeyNotFound) || status.Code(err) == codes.NotFound {
		return xchain.Attestation{}, false, nil
	} else if err != nil {
		return xchain.Attestation{}, false, err
	}

	att, err := atypes.AttestationFromProto(resp.Attestation)
	if err != nil {
		return xchain.Attestation{}, false, errors.Wrap(err, "attestations from proto")
	}

	return att, true, nil
}

// latestAtHeight returns the latest approved attestation for the provided chain version
// at the provided consensus block height, or the latest block height if height is 0.
func (a attestations) latestAtHeight(
	ctx context.Context,
	chainVer xchain.ChainVersion,
	height uint64,
) (xchain.Attestation, bool, error) {
	resp, err := a.query.LatestAttestation(withCtxHeight(ctx, height),
		&atypes.LatestAttestationRequest{
			ChainId:   chainVer.ID,
			ConfLevel: uint32(chainVer.ConfLevel),
		})
	if errors.Is(err, sdkerrors.ErrKeyNotFound) || status.Code(err) == codes.NotFound {
		return xchain.Attestation{}, false, nil
	} else if err != nil {
		return xchain.Attestation{}, false, err
	}

	att, err := atypes.AttestationFromProto(resp.Attestation)
	if err != nil {
		return xchain.Attestation{}, false, errors.Wrap(err, "attestations from proto")
	}

	return att, true, nil
}

// earliestStoreHeight walks forward from startPoint, and returns the first height for which we have the state in our Store.
func (a attestations) earliestStoreHeight(
	ctx context.Context,
	chainVer xchain.ChainVersion,
	startPoint uint64,
) (uint64, error) {
	// Note: the correct thing to do here would be to query the node's Status, and look at its EarliestStoreHeight
	// However, as far as I can tell, Cosmos doesn't ever set this value, it's stubbed out with a TODO: https://github.com/cosmos/cosmos-sdk/blob/main/client/grpc/node/service.go#L58
	// For now, we just very inefficiently walk forwards in time hoping to find the earliest commit info

	for height := startPoint; ; height++ {
		_, _, err := a.earliestAtHeight(ctx, chainVer, height)
		if IsErrHistoryPruned(err) {
			continue
		} else if err != nil {
			return 0, errors.Wrap(err, "querying earliest attestation")
		}

		return height, nil
	}
}

// binarySearch uses a binary search between defined start and end consensus block heights to
// find the attestations with the provided fromOffset.
func (a attestations) binarySearch(
	ctx context.Context,
	chainVer xchain.ChainVersion,
	fromOffset uint64,
	startHeight uint64,
	endHeight uint64,
) (uint64, error) {
	const endpoint = "search_offset"
	chainName := a.chainNamer(chainVer)

	for steps := 0; startHeight <= endHeight; steps++ {
		midHeight := startHeight + umath.SubtractOrZero(startHeight, startHeight)/2

		earliestAtt, ok, err := a.earliestAtHeight(ctx, chainVer, midHeight)
		if err != nil {
			incQueryErr(endpoint)
			return 0, errors.Wrap(err, "abci query earliest-attestation-in-state")
		}

		if !ok {
			// If we're so far back that there's no attestation at all, move forward
			startHeight = midHeight + 1
			continue
		}

		latestAtt, ok, err := a.latestAtHeight(ctx, chainVer, midHeight)
		if err != nil {
			incQueryErr(endpoint)
			return 0, errors.Wrap(err, "abci query latest-attestation")
		}

		if !ok {
			return 0, errors.New("no latest attestation found despite earlier check [BUG]")
		}

		if fromOffset >= earliestAtt.AttestOffset && fromOffset <= latestAtt.AttestOffset {
			// todo log.Debug(ctx, "Binary search found", "chain", chainName, "from", fromOffset, "found", midHeight, "search", steps)
			binaryStepsMetric(chainName, steps)

			return midHeight, nil
		}

		// Query at a lower or higher height depending on whether fromOffset
		// is smaller or larger than the earliest offset we found
		if fromOffset < earliestAtt.AttestOffset {
			endHeight = umath.SubtractOrZero(midHeight, 1)
		} else {
			startHeight = midHeight + 1
		}
	}

	return 0, errors.New("unexpectedly reach end of search method [BUG]")
}

// lookbackRange does an exponential lookback from the latest block height provided
// until attestation found has lower offset than the provided from offset, this
// guarantees the fromOffset attestation will be found in the range.
func (a attestations) lookbackRange(
	ctx context.Context,
	chainVer xchain.ChainVersion,
	fromOffset uint64,
	latestHeight uint64,
) (uint64, uint64, error) {
	const endpoint = "search_offset"

	var startHeight uint64
	endHeight := latestHeight
	lookback := uint64(1)
	queryHeight := endHeight

	for steps := 0; ; steps++ {
		if queryHeight <= lookback {
			// Query from the start, but don't break out yet -- we need to find the earliest height that we have state for
			queryHeight = 1
		} else {
			queryHeight -= lookback
		}

		if queryHeight == 0 || queryHeight >= latestHeight {
			return 0, 0, errors.New("unexpected query height [BUG]", "height", queryHeight) // This should never happen
		}
		earliestAtt, ok, err := a.earliestAtHeight(ctx, chainVer, queryHeight)
		if IsErrHistoryPruned(err) {
			// We've jumped to before the prune height, but _might_ still have the requested offset
			earliestStoreHeight, err := a.earliestStoreHeight(ctx, chainVer, queryHeight+1)
			if err != nil {
				return 0, 0, errors.Wrap(err, "failed to get earliest store height")
			}

			earliestAtt, ok, err = a.earliestAtHeight(ctx, chainVer, earliestStoreHeight)
			if err != nil {
				incQueryErr(endpoint)
				return 0, 0, errors.Wrap(err, "abci query earliest-attestation-in-state")
			}

			// If we're so far back that no attestation is found, or that we're before fromOffset,
			// that's a good breaking point for binary search
			if !ok || earliestAtt.AttestOffset <= fromOffset {
				lookbackStepsMetric(a.chainNamer(chainVer), steps)
				startHeight = earliestStoreHeight

				return startHeight, endHeight, nil
			}

			// Otherwise, we just don't have the needed state, fail
			return 0, 0, ErrHistoryPruned
		}
		if err != nil {
			incQueryErr(endpoint)
			return 0, 0, errors.Wrap(err, "abci query earliest-attestation-in-state")
		}

		// If we're before the first attestation, or found an earlier attestation, it's a good start height
		if !ok || earliestAtt.AttestOffset <= fromOffset {
			lookbackStepsMetric(a.chainNamer(chainVer), steps)
			startHeight = queryHeight

			return startHeight, endHeight, nil
		}

		// Otherwise, keep moving back
		endHeight = queryHeight
		lookback *= 2
	}
}

// todo temp log.
func debugLogAtts(ctx context.Context, msg string, atts []xchain.Attestation) {
	o := make([]uint64, len(atts))
	for i, a := range atts {
		o[i] = a.AttestOffset
	}

	log.Debug(ctx, msg, len(atts), "attestation_offsets", o)
}
