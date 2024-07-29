package provider

import (
	"context"
	"sync"
	"time"

	atypes "github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/halo/genutil/genserve"
	ptypes "github.com/omni-network/omni/halo/portal/types"
	rtypes "github.com/omni-network/omni/halo/registry/types"
	vtypes "github.com/omni-network/omni/halo/valsync/types"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tracer"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/lib/xchain"

	rpcclient "github.com/cometbft/cometbft/rpc/client"

	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gogogrpc "github.com/cosmos/gogoproto/grpc"
	"github.com/cosmos/gogoproto/proto"
	"google.golang.org/grpc"
)

func NewABCIProvider(abci rpcclient.Client, network netconf.ID, chainNamer func(xchain.ChainVersion) string) Provider {
	// Stream backoff for 1s, querying new attestations after 1 consensus block
	backoffFunc := func(ctx context.Context) func() {
		return expbackoff.New(ctx, expbackoff.WithPeriodicConfig(time.Second))
	}

	acl := atypes.NewQueryClient(rpcAdaptor{abci: abci})
	vcl := vtypes.NewQueryClient(rpcAdaptor{abci: abci})
	pcl := ptypes.NewQueryClient(rpcAdaptor{abci: abci})
	rcl := rtypes.NewQueryClient(rpcAdaptor{abci: abci})
	gcl := genserve.NewQueryClient(rpcAdaptor{abci: abci})

	return Provider{
		fetch:       newABCIFetchFunc(acl, abci, chainNamer),
		latest:      newABCILatestFunc(acl),
		window:      newABCIWindowFunc(acl),
		valset:      newABCIValsetFunc(vcl),
		portalBlock: newABCIPortalBlockFunc(pcl),
		networkFunc: newABCINetworkFunc(rcl),
		genesisFunc: newABCIGenesisFunc(gcl),
		chainID:     newChainIDFunc(abci),
		header:      abci.Header,
		backoffFunc: backoffFunc,
		chainNamer:  chainNamer,
		network:     network,
	}
}

// newChainIDFunc returns a function that returns the consensus chain ID. It caches the result.
func newChainIDFunc(abci rpcclient.SignClient) chainIDFunc {
	var mu sync.Mutex
	var chainID uint64

	return func(ctx context.Context) (uint64, error) {
		mu.Lock()
		defer mu.Unlock()
		if chainID != 0 {
			return chainID, nil
		}

		ctx, span := tracer.Start(ctx, spanName("chain_id"))
		defer span.End()

		resp, err := abci.Header(ctx, nil)
		if err != nil {
			return 0, errors.Wrap(err, "abci header")
		}

		chainID, err = netconf.ConsensusChainIDStr2Uint64(resp.Header.ChainID)
		if err != nil {
			return 0, errors.Wrap(err, "parse chain ID")
		}

		return chainID, nil
	}
}

func newABCIValsetFunc(cl vtypes.QueryClient) valsetFunc {
	return func(ctx context.Context, valSetID uint64, latest bool) (valSetResponse, bool, error) {
		const endpoint = "valset"
		defer latency(endpoint)()

		ctx, span := tracer.Start(ctx, spanName(endpoint))
		defer span.End()

		resp, err := cl.ValidatorSet(ctx, &vtypes.ValidatorSetRequest{Id: valSetID, Latest: latest})
		if errors.Is(err, sdkerrors.ErrKeyNotFound) {
			return valSetResponse{}, false, nil
		} else if err != nil {
			incQueryErr(endpoint)
			return valSetResponse{}, false, errors.Wrap(err, "abci query valset")
		}

		vals := make([]cchain.Validator, 0, len(resp.Validators))
		for _, v := range resp.Validators {
			ethAddr, err := v.EthereumAddress()
			if err != nil {
				return valSetResponse{}, false, err
			}
			vals = append(vals, cchain.Validator{
				Address: ethAddr,
				Power:   v.Power,
			})
		}

		return valSetResponse{
			ValSetID:      resp.Id,
			Validators:    vals,
			CreatedHeight: resp.CreatedHeight,
			activedHeight: resp.ActivatedHeight,
		}, true, nil
	}
}

func newABCIFetchFunc(cl atypes.QueryClient, client rpcclient.Client, chainNamer func(xchain.ChainVersion) string) fetchFunc {
	return func(ctx context.Context, chainVer xchain.ChainVersion, fromOffset uint64) ([]xchain.Attestation, error) {
		const endpoint = "fetch_attestations"
		defer latency(endpoint)()

		ctx, span := tracer.Start(ctx, spanName(endpoint))
		defer span.End()

		chainName := chainNamer(chainVer)

		// try fetching from latest height
		atts, ok, err := attsFromAtHeight(ctx, cl, chainVer, fromOffset, 0)
		if err != nil {
			incQueryErr(endpoint)
			return nil, errors.Wrap(err, "abci query attestations-from")
		}

		if ok {
			fetchStepsMetrics(chainName, 0, 0)
			return atts, nil
		}

		earliestAttestationAtLatestHeight, ok, err := queryEarliestAttestation(ctx, cl, chainVer, 0)
		if err != nil {
			incQueryErr(endpoint)
			return nil, errors.Wrap(err, "abci query earliest-attestation-in-state")
		}

		// Either no attestations have happened yet, or the queried fromOffset is in the "future"
		// Caller has to wait and retry in both cases
		if !ok || earliestAttestationAtLatestHeight.AttestOffset < fromOffset {
			// First attestation hasn't happened yet, return empty
			return []xchain.Attestation{}, nil
		}

		offsetHeight, err := searchOffsetInHistory(ctx, client, cl, chainVer, chainName, fromOffset)
		if err != nil {
			incQueryErr(endpoint)
			return nil, errors.Wrap(err, "searching offset in history")
		}

		atts, attsFromOk, err := attsFromAtHeight(ctx, cl, chainVer, fromOffset, offsetHeight)
		if err != nil {
			incQueryErr(endpoint)
			return nil, errors.Wrap(err, "abci query attestations-from")
		}

		if !attsFromOk {
			return nil, errors.New("expected to find attestations [BUG]")
		}

		return atts, nil
	}
}

func newABCIWindowFunc(cl atypes.QueryClient) windowFunc {
	return func(ctx context.Context, chainVer xchain.ChainVersion, xBlockOffset uint64) (int, error) {
		const endpoint = "window_compare"
		defer latency(endpoint)()

		ctx, span := tracer.Start(ctx, spanName(endpoint))
		defer span.End()

		resp, err := cl.WindowCompare(ctx, &atypes.WindowCompareRequest{
			ChainId:     chainVer.ID,
			ConfLevel:   uint32(chainVer.ConfLevel),
			BlockOffset: xBlockOffset,
		})
		if err != nil {
			incQueryErr(endpoint)
			return 0, errors.Wrap(err, "abci query window compare")
		}

		return int(resp.Cmp), nil
	}
}

func newABCILatestFunc(cl atypes.QueryClient) latestFunc {
	return func(ctx context.Context, chainVer xchain.ChainVersion) (xchain.Attestation, bool, error) {
		const endpoint = "latest_attestation"
		defer latency(endpoint)()

		ctx, span := tracer.Start(ctx, spanName(endpoint))
		defer span.End()

		att, ok, err := queryLatestAttestation(ctx, cl, chainVer, 0)
		if err != nil {
			incQueryErr(endpoint)
			return xchain.Attestation{}, false, errors.Wrap(err, "abci query latest attestation")
		}
		if !ok {
			return xchain.Attestation{}, false, nil
		}

		return att, true, nil
	}
}

func newABCIPortalBlockFunc(pcl ptypes.QueryClient) portalBlockFunc {
	return func(ctx context.Context, blockOffset uint64, latest bool) (*ptypes.BlockResponse, bool, error) {
		const endpoint = "portal_block"
		defer latency(endpoint)()

		ctx, span := tracer.Start(ctx, spanName(endpoint))
		defer span.End()

		resp, err := pcl.Block(ctx, &ptypes.BlockRequest{Id: blockOffset, Latest: latest})
		if errors.Is(err, sdkerrors.ErrKeyNotFound) {
			return nil, false, nil
		} else if err != nil {
			incQueryErr(endpoint)
			return nil, false, errors.Wrap(err, "abci query portal block")
		}

		return resp, true, nil
	}
}

func newABCINetworkFunc(pcl rtypes.QueryClient) networkFunc {
	return func(ctx context.Context, networkID uint64, latest bool) (*rtypes.NetworkResponse, bool, error) {
		const endpoint = "registry_network"
		defer latency(endpoint)()

		ctx, span := tracer.Start(ctx, spanName(endpoint))
		defer span.End()

		resp, err := pcl.Network(ctx, &rtypes.NetworkRequest{Id: networkID, Latest: latest})
		if errors.Is(err, sdkerrors.ErrKeyNotFound) {
			return nil, false, nil
		} else if err != nil {
			incQueryErr(endpoint)
			return nil, false, errors.Wrap(err, "abci query network")
		}

		return resp, true, nil
	}
}

func newABCIGenesisFunc(gcl genserve.QueryClient) genesisFunc {
	return func(ctx context.Context) (execution []byte, consensus []byte, err error) { //nolint:nonamedreturns // Disambiguate identical return types
		const endpoint = "genesis"
		defer latency(endpoint)()

		ctx, span := tracer.Start(ctx, spanName(endpoint))
		defer span.End()

		resp, err := gcl.Genesis(ctx, &genserve.GenesisRequest{})
		if err != nil {
			incQueryErr(endpoint)
			return nil, nil, errors.Wrap(err, "abci genesis")
		}

		return resp.ExecutionGenesisJson, resp.ConsensusGenesisJson, nil
	}
}

// rpcAdaptor adapts the cometBFT query client to the gRPC client interface.
// Note it only supports the Invoke method, not the NewStream method.
type rpcAdaptor struct {
	gogogrpc.ClientConn
	abci rpcclient.ABCIClient
}

type heightKey struct{}

// withCtxHeight returns a copy of the context with the `heightKey` set to the provided height.
// This height will be supplied in ABCIQueryOptions when issuing queries in the rpcAdaptor.
func withCtxHeight(ctx context.Context, height uint64) context.Context {
	return context.WithValue(ctx, heightKey{}, height)
}

// heightFromCtx returns the `heightKey` from the supplied context, if found.
func heightFromCtx(ctx context.Context) (uint64, bool) {
	v, ok := ctx.Value(heightKey{}).(uint64)
	return v, ok
}

func (a rpcAdaptor) Invoke(ctx context.Context, method string, req, resp any, _ ...grpc.CallOption) error {
	reqpb, ok := req.(proto.Message)
	if !ok {
		return errors.New("args not proto.Message")
	}
	resppb, ok := resp.(proto.Message)
	if !ok {
		return errors.New("args not proto.Message")
	}

	bz, err := proto.Marshal(reqpb)
	if err != nil {
		return errors.Wrap(err, "marshal approved-from request")
	}

	// If left as zero value, the latest height will be used
	var queryHeight uint64
	if v, ok := heightFromCtx(ctx); ok {
		queryHeight = v
	}

	r, err := a.abci.ABCIQueryWithOptions(ctx, method, bz, rpcclient.ABCIQueryOptions{Height: int64(queryHeight)})
	if err != nil {
		return errors.Wrap(err, "abci query")
	} else if !r.Response.IsOK() {
		err := errorsmod.ABCIError(r.Response.Codespace, r.Response.Code, r.Response.Log)
		return errors.Wrap(err, "abci query failed", "code", r.Response.Code, "info", r.Response.Info, "log", r.Response.Log)
	}

	err = proto.Unmarshal(r.Response.Value, resppb)
	if err != nil {
		return errors.Wrap(err, "unmarshal approved-from response")
	}

	return nil
}

func spanName(endpoint string) string {
	return "cprovider/" + endpoint
}

// searchOffsetInHistory searches the consensus state history and
// returns a historical consensus block height that contains an approved attestation
// for the provided chain version and fromOffset.
func searchOffsetInHistory(ctx context.Context, client rpcclient.Client, cl atypes.QueryClient, chainVer xchain.ChainVersion, chainName string, fromOffset uint64) (uint64, error) {
	const endpoint = "search_offset"
	defer latency(endpoint)

	// Exponentially backoff to find a good start point for binary search, this prefers more recent queries
	info, err := client.ABCIInfo(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "abci query info")
	}

	var startHeightIndex uint64
	endHeightIndex := uint64(info.Response.LastBlockHeight)
	lookback := uint64(1)
	var lookbackStepsCounter uint64 // For metrics only
	queryHeight := endHeightIndex
	for {
		lookbackStepsCounter++
		if queryHeight <= lookback {
			// Query from the start, but don't break out yet -- we need to find the earliest height that we have state for
			queryHeight = 1
		} else {
			queryHeight -= lookback
		}

		if queryHeight == 0 || queryHeight >= uint64(info.Response.LastBlockHeight) {
			return 0, errors.New("unexpected query height [BUG]", "queryHeight", queryHeight) // This should never happen
		}
		earliestAtt, ok, err := queryEarliestAttestation(ctx, cl, chainVer, queryHeight)
		if IsErrHistoryPruned(err) {
			// We've jumped to before the prune height, but _might_ still have the requested offset
			earliestStoreHeight, err := getEarliestStoreHeight(ctx, cl, chainVer, queryHeight+1)
			if err != nil {
				return 0, errors.Wrap(err, "failed to get earliest store height")
			}
			earliestAtt, ok, err = queryEarliestAttestation(ctx, cl, chainVer, earliestStoreHeight)
			if err != nil {
				incQueryErr(endpoint)
				return 0, errors.Wrap(err, "abci query earliest-attestation-in-state")
			}

			// If we're so far back that no attestation is found, or that we're before fromOffset,
			// that's a good breaking point for binary search
			if !ok || earliestAtt.AttestOffset <= fromOffset {
				startHeightIndex = earliestStoreHeight
				break
			}

			// Otherwise, we just don't have the needed state, fail
			return 0, ErrHistoryPruned
		}
		if err != nil {
			incQueryErr(endpoint)
			return 0, errors.Wrap(err, "abci query earliest-attestation-in-state")
		}

		// If we're before the first attestation, or found an earlier attestation, it's a good start height
		if !ok || earliestAtt.AttestOffset <= fromOffset {
			startHeightIndex = queryHeight
			break
		}

		// Otherwise, keep moving back
		endHeightIndex = queryHeight
		lookback *= 2
	}

	// We now have reasonable start and end indices for binary search
	var binarySearchStepsCounter uint64 // For metrics only
	for startHeightIndex <= endHeightIndex {
		binarySearchStepsCounter++
		midHeightIndex := startHeightIndex + umath.SubtractOrZero(endHeightIndex, startHeightIndex)/2

		earliestAtt, ok, err := queryEarliestAttestation(ctx, cl, chainVer, midHeightIndex)
		if err != nil {
			incQueryErr(endpoint)
			return 0, errors.Wrap(err, "abci query earliest-attestation-in-state")
		}

		if !ok {
			// If we're so far back that there's no attestation at all, move forward
			startHeightIndex = midHeightIndex + 1
			continue
		}

		latestAtt, ok, err := queryLatestAttestation(ctx, cl, chainVer, midHeightIndex)
		if err != nil {
			incQueryErr(endpoint)
			return 0, errors.Wrap(err, "abci query latest-attestation")
		}

		if !ok {
			return 0, errors.New("no latest attestation found despite earlier check [BUG]")
		}

		if fromOffset >= earliestAtt.AttestOffset && fromOffset <= latestAtt.AttestOffset {
			fetchStepsMetrics(chainName, lookbackStepsCounter, binarySearchStepsCounter)
			return midHeightIndex, nil
		}

		// Query at a lower or higher height depending on whether fromOffset
		// is smaller or larger than the earliest offset we found
		if fromOffset < earliestAtt.AttestOffset {
			endHeightIndex = umath.SubtractOrZero(midHeightIndex, 1)
		} else {
			startHeightIndex = midHeightIndex + 1
		}
	}

	return 0, errors.New("unexpectedly reach end of search method [BUG]")
}

// getEarliestStoreHeight walks forward from startPoint, and returns the first height for which we have the state in our Store.
func getEarliestStoreHeight(ctx context.Context, cl atypes.QueryClient, chainVer xchain.ChainVersion, startPoint uint64) (uint64, error) {
	// Note: the correct thing to do here would be to query the node's Status, and look at its EarliestStoreHeight
	// However, as far as I can tell, Cosmos doesn't ever set this value, it's stubbed out with a TODO: https://github.com/cosmos/cosmos-sdk/blob/main/client/grpc/node/service.go#L58
	// For now, we just very inefficiently walk forwards in time hoping to find the earliest commit info

	for height := startPoint; ; height++ {
		_, _, err := queryEarliestAttestation(ctx, cl, chainVer, height)
		if IsErrHistoryPruned(err) {
			continue
		} else if err != nil {
			return 0, errors.Wrap(err, "querying earliest attestation")
		}

		return height, nil
	}
}

// queryEarliestAttestation returns the earliest approved attestation for the provided chain version
// at the provided consensus block height, or the latest block height if height is 0.
func queryEarliestAttestation(ctx context.Context, cl atypes.QueryClient, chainVer xchain.ChainVersion, height uint64) (xchain.Attestation, bool, error) {
	resp, err := cl.EarliestAttestation(withCtxHeight(ctx, height),
		&atypes.EarliestAttestationRequest{
			ChainId:   chainVer.ID,
			ConfLevel: uint32(chainVer.ConfLevel),
		})
	if errors.Is(err, sdkerrors.ErrKeyNotFound) {
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

// queryLatestAttestation returns the latest approved attestation for the provided chain version
// at the provided consensus block height, or the latest block height if height is 0.
func queryLatestAttestation(ctx context.Context, cl atypes.QueryClient, chainVer xchain.ChainVersion, height uint64) (xchain.Attestation, bool, error) {
	resp, err := cl.LatestAttestation(withCtxHeight(ctx, height),
		&atypes.LatestAttestationRequest{
			ChainId:   chainVer.ID,
			ConfLevel: uint32(chainVer.ConfLevel),
		})
	if errors.Is(err, sdkerrors.ErrKeyNotFound) {
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

// attsFromAtHeight returns approved attestations for the provided chain version
// at the provided consensus block height, or the latest block height if height is 0.
func attsFromAtHeight(ctx context.Context, cl atypes.QueryClient, chainVer xchain.ChainVersion, fromOffset, height uint64) ([]xchain.Attestation, bool, error) {
	resp, err := cl.AttestationsFrom(withCtxHeight(ctx, height), &atypes.AttestationsFromRequest{
		ChainId:    chainVer.ID,
		ConfLevel:  uint32(chainVer.ConfLevel),
		FromOffset: fromOffset,
	})
	if err != nil {
		return []xchain.Attestation{}, false, errors.Wrap(err, "abci query attestations-from")
	}

	if len(resp.Attestations) == 0 {
		return []xchain.Attestation{}, false, nil
	}

	atts, err := atypes.AttestationsFromProto(resp.Attestations)
	if err != nil {
		return []xchain.Attestation{}, false, errors.Wrap(err, "attestations from proto")
	}

	return atts, len(resp.Attestations) != 0, nil
}
