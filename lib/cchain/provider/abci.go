package provider

import (
	"context"
	"strconv"
	"sync"
	"time"

	atypes "github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/halo/genutil/genserve"
	ptypes "github.com/omni-network/omni/halo/portal/types"
	rtypes "github.com/omni-network/omni/halo/registry/types"
	vtypes "github.com/omni-network/omni/halo/valsync/types"
	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tracer"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/lib/xchain"
	evmengtypes "github.com/omni-network/omni/octane/evmengine/types"

	rpcclient "github.com/cometbft/cometbft/rpc/client"
	"github.com/cometbft/cometbft/rpc/client/http"

	"github.com/ethereum/go-ethereum/common"

	errorsmod "cosmossdk.io/errors"
	utypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	"github.com/cosmos/cosmos-sdk/client/grpc/node"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	btypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	dtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	mtypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	sltypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	gogogrpc "github.com/cosmos/gogoproto/grpc"
	"github.com/cosmos/gogoproto/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Dial returns a ABCI provider to the provided network connecting to well-known public RPCs.
func Dial(network netconf.ID) (Provider, error) {
	consRPC := network.Static().ConsensusRPC()
	if consRPC == "" {
		return Provider{}, errors.New("consensus rpc not configured for network")
	}

	cl, err := http.New(consRPC, "/websocket")
	if err != nil {
		return Provider{}, errors.Wrap(err, "new tendermint client")
	}

	return NewABCI(cl, network), nil
}

// NewABCI returns a new provider using the provided cometBFT ABCI client.
func NewABCI(cmtCl rpcclient.Client, network netconf.ID, opts ...func(*Provider)) Provider {
	return newProvider(rpcAdaptor{abci: cmtCl}, network, opts...)
}

// NewGRPC returns a new provider using the provided gRPC server address.
// This is preferred to NewABCI as it bypasses CometBFT so is much faster
// and doesn't affect chain performance.
func NewGRPC(target string, network netconf.ID, ir codectypes.InterfaceRegistry, opts ...func(*Provider)) (Provider, error) {
	grpcClient, err := grpc.NewClient(target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(ir).GRPCCodec())),
	)
	if err != nil {
		return Provider{}, errors.Wrap(err, "new grpc client")
	}

	return newProvider(grpcClient, network, opts...), nil
}

func newProvider(cc gogogrpc.ClientConn, network netconf.ID, opts ...func(*Provider)) Provider {
	// Stream backoff for 1s, querying new attestations after 1 consensus block
	backoffFunc := func(ctx context.Context) func() {
		return expbackoff.New(ctx, expbackoff.WithPeriodicConfig(time.Second))
	}

	namer := netconf.ChainVersionNamer(network)

	acl := atypes.NewQueryClient(cc)
	vcl := vtypes.NewQueryClient(cc)
	pcl := ptypes.NewQueryClient(cc)
	rcl := rtypes.NewQueryClient(cc)
	gcl := genserve.NewQueryClient(cc)
	ucl := utypes.NewQueryClient(cc)
	scl := stypes.NewQueryClient(cc)
	dcl := dtypes.NewQueryClient(cc)
	slcl := sltypes.NewQueryClient(cc)
	cmtcl := cmtservice.NewServiceClient(cc)
	mcl := mtypes.NewQueryClient(cc)
	evmengcl := evmengtypes.NewQueryClient(cc)
	bcl := btypes.NewQueryClient(cc)
	ncl := node.NewServiceClient(cc)

	p := Provider{
		fetch:         newABCIFetchFunc(acl, cmtcl, namer),
		allAtts:       newABCIAllAttsFunc(acl),
		latest:        newABCILatestFunc(acl),
		window:        newABCIWindowFunc(acl),
		valset:        newABCIValsetFunc(vcl),
		val:           newABCIValFunc(scl),
		vals:          newABCIValsFunc(scl),
		signing:       newABCISigningFunc(slcl),
		rewards:       newABCIRewards(dcl),
		portalBlock:   newABCIPortalBlockFunc(pcl),
		networkFunc:   newABCINetworkFunc(rcl),
		genesisFunc:   newABCIGenesisFunc(gcl),
		plannedFunc:   newABCIPlannedUpgradeFunc(ucl),
		appliedFunc:   newABCIAppliedUpgradeFunc(ucl),
		executionHead: newABCIExecutionHeadFunc(evmengcl),
		chainID:       newChainIDFunc(cmtcl),
		backoffFunc:   backoffFunc,
		chainNamer:    namer,
		network:       network,
		queryClients: cchain.QueryClients{
			Attest:       acl,
			Portal:       pcl,
			Registry:     rcl,
			ValSync:      vcl,
			Staking:      scl,
			Slashing:     slcl,
			Upgrade:      ucl,
			Distribution: dcl,
			Mint:         mcl,
			EvmEngine:    evmengcl,
			Bank:         bcl,
			Node:         ncl,
		},
	}

	for _, opt := range opts {
		opt(&p)
	}

	return p
}

func newABCISigningFunc(cl sltypes.QueryClient) signingFunc {
	return func(ctx context.Context) ([]cchain.SDKSigningInfo, error) {
		const endpoint = "signing_info"
		defer latency(endpoint)()

		ctx, span := tracer.Start(ctx, spanName(endpoint))
		defer span.End()

		req := &sltypes.QuerySigningInfosRequest{}
		resp, err := cl.SigningInfos(ctx, req)
		if err != nil {
			incQueryErr(endpoint)
			return nil, errors.Wrap(err, "abci query signing info")
		}

		params, err := cl.Params(ctx, &sltypes.QueryParamsRequest{})
		if err != nil {
			incQueryErr(endpoint)
			return nil, errors.Wrap(err, "abci query params")
		} else if params.Params.SignedBlocksWindow == 0 {
			return nil, errors.New("signed blocks window is zero")
		}

		var infos []cchain.SDKSigningInfo
		for _, info := range resp.Info {
			// uptime over the past <SignedBlocksWindow> blocks.
			uptime := 1.0 - (float64(info.MissedBlocksCounter) / float64(params.Params.SignedBlocksWindow))

			infos = append(infos, cchain.SDKSigningInfo{
				ValidatorSigningInfo: info,
				Uptime:               uptime,
			})
		}

		return infos, nil
	}
}

func newABCIAppliedUpgradeFunc(ucl utypes.QueryClient) appliedUpgradeFunc {
	return func(ctx context.Context, name string) (utypes.Plan, bool, error) {
		resp, err := ucl.AppliedPlan(ctx, &utypes.QueryAppliedPlanRequest{
			Name: name,
		})
		if err != nil {
			return utypes.Plan{}, false, errors.Wrap(err, "abci query applied plan")
		} else if resp.Height == 0 {
			return utypes.Plan{}, false, nil
		}

		return utypes.Plan{
			Name:   name,
			Height: resp.Height,
		}, true, nil
	}
}

func newABCIPlannedUpgradeFunc(ucl utypes.QueryClient) planedUpgradeFunc {
	return func(ctx context.Context) (utypes.Plan, bool, error) {
		resp, err := ucl.CurrentPlan(ctx, &utypes.QueryCurrentPlanRequest{})
		if err != nil {
			return utypes.Plan{}, false, errors.Wrap(err, "abci query current plan")
		} else if resp.Plan == nil {
			return utypes.Plan{}, false, nil
		}

		return *resp.Plan, true, nil
	}
}

// newChainIDFunc returns a function that returns the consensus chain ID. It caches the result.
func newChainIDFunc(cmtCl cmtservice.ServiceClient) chainIDFunc {
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

		resp, err := cmtCl.GetLatestBlock(ctx, &cmtservice.GetLatestBlockRequest{})
		if err != nil {
			return 0, errors.Wrap(err, "abci header")
		}

		chainID, err = netconf.ConsensusChainIDStr2Uint64(resp.SdkBlock.Header.ChainID)
		if err != nil {
			return 0, errors.Wrap(err, "parse chain ID")
		}

		return chainID, nil
	}
}

func newABCIRewards(cl dtypes.QueryClient) rewardsFunc {
	return func(ctx context.Context, operator common.Address) (float64, bool, error) {
		const endpoint = "rewards"
		defer latency(endpoint)()

		ctx, span := tracer.Start(ctx, spanName(endpoint))
		defer span.End()

		req := &dtypes.QueryValidatorOutstandingRewardsRequest{
			ValidatorAddress: sdk.ValAddress(operator.Bytes()).String(),
		}

		resp, err := cl.ValidatorOutstandingRewards(ctx, req)
		if errors.Is(err, sdkerrors.ErrKeyNotFound) || status.Code(err) == codes.NotFound {
			return 0, false, nil
		} else if err != nil {
			incQueryErr(endpoint)
			return 0, false, errors.Wrap(err, "abci query rewards")
		}

		return resp.Rewards.Rewards.AmountOf(sdk.DefaultBondDenom).MustFloat64(), true, nil
	}
}

func newABCIValFunc(cl stypes.QueryClient) valFunc {
	return func(ctx context.Context, operatorAddr common.Address) (cchain.SDKValidator, bool, error) {
		const endpoint = "validator"
		defer latency(endpoint)()

		ctx, span := tracer.Start(ctx, spanName(endpoint))
		defer span.End()

		valAddr := sdk.ValAddress(operatorAddr.Bytes())
		resp, err := cl.Validator(ctx, &stypes.QueryValidatorRequest{ValidatorAddr: valAddr.String()})
		if errors.Is(err, sdkerrors.ErrKeyNotFound) || status.Code(err) == codes.NotFound {
			return cchain.SDKValidator{}, false, nil
		} else if err != nil {
			incQueryErr(endpoint)
			return cchain.SDKValidator{}, false, errors.Wrap(err, "abci query validator")
		}

		return cchain.SDKValidator{Validator: resp.Validator}, true, nil
	}
}

func newABCIValsFunc(cl stypes.QueryClient) valsFunc {
	return func(ctx context.Context) ([]cchain.SDKValidator, error) {
		const endpoint = "vals"
		defer latency(endpoint)()

		ctx, span := tracer.Start(ctx, spanName(endpoint))
		defer span.End()

		resp, err := cl.Validators(ctx, &stypes.QueryValidatorsRequest{})
		if err != nil {
			incQueryErr(endpoint)
			return nil, errors.Wrap(err, "abci query validators")
		}

		vals := make([]cchain.SDKValidator, 0, len(resp.Validators))
		for _, val := range resp.Validators {
			vals = append(vals, cchain.SDKValidator{Validator: val})
		}

		return vals, nil
	}
}

func newABCIValsetFunc(cl vtypes.QueryClient) valsetFunc {
	return func(ctx context.Context, valSetID uint64, latest bool) (valSetResponse, bool, error) {
		const endpoint = "valset"
		defer latency(endpoint)()

		ctx, span := tracer.Start(ctx, spanName(endpoint))
		defer span.End()

		resp, err := cl.ValidatorSet(ctx, &vtypes.ValidatorSetRequest{Id: valSetID, Latest: latest})
		if errors.Is(err, sdkerrors.ErrKeyNotFound) || status.Code(err) == codes.NotFound {
			return valSetResponse{}, false, nil
		} else if err != nil {
			incQueryErr(endpoint)
			return valSetResponse{}, false, errors.Wrap(err, "abci query valset")
		}

		vals := make([]cchain.PortalValidator, 0, len(resp.Validators))
		for _, v := range resp.Validators {
			ethAddr, err := v.EthereumAddress()
			if err != nil {
				return valSetResponse{}, false, err
			}
			vals = append(vals, cchain.PortalValidator{
				Address: ethAddr,
				Power:   v.Power,
			})
		}

		return valSetResponse{
			ValSetID:        resp.Id,
			Validators:      vals,
			CreatedHeight:   resp.CreatedHeight,
			activatedHeight: resp.ActivatedHeight,
		}, true, nil
	}
}

func newABCIAllAttsFunc(cl atypes.QueryClient) allAttsFunc {
	return func(ctx context.Context, chainVer xchain.ChainVersion, fromOffset uint64) ([]xchain.Attestation, error) {
		var atts []xchain.Attestation
		for _, status := range []uint32{atypes.StatusPending, atypes.StatusApproved} {
			req := &atypes.ListAllAttestationsRequest{
				ChainId:    chainVer.ID,
				ConfLevel:  uint32(chainVer.ConfLevel),
				Status:     status,
				FromOffset: fromOffset,
			}
			resp, err := cl.ListAllAttestations(ctx, req)
			if err != nil {
				return nil, errors.Wrap(err, "abci query all attestations")
			}

			for _, att := range resp.Attestations {
				attX, err := att.ToXChain()
				if err != nil {
					return nil, err
				}

				atts = append(atts, attX)
			}
		}

		return atts, nil
	}
}

func newABCIFetchFunc(attCl atypes.QueryClient, cmtCl cmtservice.ServiceClient, chainNamer func(xchain.ChainVersion) string) fetchFunc {
	return func(ctx context.Context, chainVer xchain.ChainVersion, fromOffset uint64, cursor uint64) ([]xchain.Attestation, uint64, error) {
		const endpoint = "fetch_attestations"
		defer latency(endpoint)()

		ctx, span := tracer.Start(ctx, spanName(endpoint))
		defer span.End()

		chainName := chainNamer(chainVer)

		// try fetching from latest height
		atts, ok, err := attsFromAtHeight(ctx, attCl, chainVer, fromOffset, 0)
		if err != nil {
			incQueryErr(endpoint)
			return nil, 0, errors.Wrap(err, "abci query attestations-from")
		} else if ok {
			binarySearchStepsMetric(chainName, 0)
			lookbackStepsMetric(chainName, 0)

			return atts, cursor, nil
		}

		earliestAttestationAtLatestHeight, ok, err := queryEarliestAttestation(ctx, attCl, chainVer, 0)
		if err != nil {
			incQueryErr(endpoint)
			return nil, 0, errors.Wrap(err, "abci query earliest-attestation-in-state")
		}

		// Either no attestations have happened yet, or the queried fromOffset is in the "future"
		// Caller has to wait and retry in both cases
		if !ok || earliestAttestationAtLatestHeight.AttestOffset < fromOffset {
			// First attestation hasn't happened yet, return empty and set cursor to latest height
			return nil, cursor, nil
		}

		latestBlockResp, err := cmtCl.GetLatestBlock(ctx, &cmtservice.GetLatestBlockRequest{})
		if err != nil {
			return nil, 0, errors.Wrap(err, "query latest block")
		}

		latestHeight, err := umath.ToUint64(latestBlockResp.SdkBlock.Header.Height)
		if err != nil {
			return nil, 0, err
		}

		// Binary search range from cached to latest
		searchStart, searchEnd := cursor, latestHeight
		if cursor == 0 {
			// Unless no cached height provided, then perform lookback search
			searchStart, searchEnd, err = lookbackRange(ctx, attCl, chainVer, chainName, fromOffset, latestHeight)
			if err != nil {
				incQueryErr(endpoint)
				return nil, 0, errors.Wrap(err, "lookback search")
			}
		}

		offsetHeight, err := binarySearch(ctx, attCl, chainVer, chainName, fromOffset, searchStart, searchEnd)
		if err != nil {
			incQueryErr(endpoint)
			return nil, 0, errors.Wrap(err, "binary search")
		}

		atts, ok, err = attsFromAtHeight(ctx, attCl, chainVer, fromOffset, offsetHeight)
		if err != nil {
			incQueryErr(endpoint)
			return nil, 0, errors.Wrap(err, "abci query attestations-from")
		} else if !ok {
			return nil, 0, errors.New("expected to find attestations [BUG]")
		}

		return atts, offsetHeight, nil
	}
}

func newABCIWindowFunc(cl atypes.QueryClient) windowFunc {
	return func(ctx context.Context, chainVer xchain.ChainVersion, attestOffset uint64) (int, error) {
		const endpoint = "window_compare"
		defer latency(endpoint)()

		ctx, span := tracer.Start(ctx, spanName(endpoint))
		defer span.End()

		resp, err := cl.WindowCompare(ctx, &atypes.WindowCompareRequest{
			ChainId:      chainVer.ID,
			ConfLevel:    uint32(chainVer.ConfLevel),
			AttestOffset: attestOffset,
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
	return func(ctx context.Context, attestOffset uint64, latest bool) (*ptypes.BlockResponse, bool, error) {
		const endpoint = "portal_block"
		defer latency(endpoint)()

		ctx, span := tracer.Start(ctx, spanName(endpoint))
		defer span.End()

		resp, err := pcl.Block(ctx, &ptypes.BlockRequest{Id: attestOffset, Latest: latest})
		if errors.Is(err, sdkerrors.ErrKeyNotFound) || status.Code(err) == codes.NotFound {
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
		if errors.Is(err, sdkerrors.ErrKeyNotFound) || status.Code(err) == codes.NotFound {
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

func newABCIExecutionHeadFunc(cl evmengtypes.QueryClient) executionHeadFunc {
	return func(ctx context.Context) (cchain.ExecutionHead, error) {
		const endpoint = "execution_head"
		defer latency(endpoint)()

		ctx, span := tracer.Start(ctx, spanName(endpoint))
		defer span.End()

		resp, err := cl.ExecutionHead(ctx, &evmengtypes.ExecutionHeadRequest{})
		if err != nil {
			incQueryErr(endpoint)
			return cchain.ExecutionHead{}, errors.Wrap(err, "abci query execution head")
		}

		blockHash, err := cast.EthHash(resp.BlockHash)
		if err != nil {
			return cchain.ExecutionHead{}, errors.Wrap(err, "cast block hash")
		}

		blockTime, err := umath.ToInt64(resp.BlockTime)
		if err != nil {
			return cchain.ExecutionHead{}, errors.Wrap(err, "cast block time")
		}

		return cchain.ExecutionHead{
			CreatedHeight: resp.CreatedHeight,
			BlockNumber:   resp.BlockNumber,
			BlockHash:     blockHash,
			BlockTime:     time.Unix(blockTime, 0),
		}, nil
	}
}

// rpcAdaptor adapts the cometBFT query client to the gRPC client interface.
// Note it only supports the Invoke method, not the NewStream method.
type rpcAdaptor struct {
	gogogrpc.ClientConn
	abci rpcclient.ABCIClient
}

// WithCtxHeight returns a copy of the context with the `x-cosmos-block-height` grpc metadata header
// set to the provided height.
//
// This height will be supplied in ABCIQueryOptions when issuing queries in the rpcAdaptor.
// It will also be added to grpc queries automatically.
func WithCtxHeight(ctx context.Context, height uint64) (context.Context, error) {
	if _, ok, err := heightFromCtx(ctx); err != nil {
		return nil, err
	} else if ok {
		return nil, errors.New("height already set in context")
	}

	return metadata.AppendToOutgoingContext(ctx, grpctypes.GRPCBlockHeightHeader, strconv.FormatUint(height, 10)), nil
}

// heightFromCtx returns the `x-cosmos-block-height` grpc metadata header value from the supplied context, if found.
func heightFromCtx(ctx context.Context) (uint64, bool, error) {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return 0, false, nil
	}

	heightHeaders := md.Get(grpctypes.GRPCBlockHeightHeader)
	if len(heightHeaders) == 0 {
		return 0, false, nil
	} else if len(heightHeaders) > 1 {
		return 0, false, errors.New("multiple height headers")
	}

	height, err := strconv.ParseUint(heightHeaders[0], 10, 64)
	if err != nil {
		return 0, false, errors.Wrap(err, "parse height header")
	}

	return height, true, nil
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

	// Zero value for queryHeight will result the latest height
	queryHeight, _, err := heightFromCtx(ctx)
	if err != nil {
		return err
	}

	queryHeightInt64, err := umath.ToInt64(queryHeight)
	if err != nil {
		return err
	}

	r, err := a.abci.ABCIQueryWithOptions(ctx, method, bz, rpcclient.ABCIQueryOptions{Height: queryHeightInt64})
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
	ctx, err := WithCtxHeight(ctx, height)
	if err != nil {
		return xchain.Attestation{}, false, err
	}

	resp, err := cl.EarliestAttestation(ctx,
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

// queryLatestAttestation returns the latest approved attestation for the provided chain version
// at the provided consensus block height, or the latest block height if height is 0.
func queryLatestAttestation(ctx context.Context, cl atypes.QueryClient, chainVer xchain.ChainVersion, height uint64) (xchain.Attestation, bool, error) {
	ctx, err := WithCtxHeight(ctx, height)
	if err != nil {
		return xchain.Attestation{}, false, err
	}

	resp, err := cl.LatestAttestation(ctx,
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

// attsFromAtHeight returns approved attestations for the provided chain version
// at the provided consensus block height, or the latest block height if height is 0.
func attsFromAtHeight(ctx context.Context, cl atypes.QueryClient, chainVer xchain.ChainVersion, fromOffset, height uint64) ([]xchain.Attestation, bool, error) {
	ctx, err := WithCtxHeight(ctx, height)
	if err != nil {
		return nil, false, err
	}

	resp, err := cl.AttestationsFrom(ctx, &atypes.AttestationsFromRequest{
		ChainId:    chainVer.ID,
		ConfLevel:  uint32(chainVer.ConfLevel),
		FromOffset: fromOffset,
	})
	if err != nil {
		return nil, false, errors.Wrap(err, "abci query attestations-from")
	} else if len(resp.Attestations) == 0 {
		return nil, false, nil
	}

	atts, err := atypes.AttestationsFromProto(resp.Attestations)
	if err != nil {
		return nil, false, errors.Wrap(err, "attestations from proto")
	}

	return atts, true, nil
}

// binarySearch uses a binary search between defined start and end consensus block heights to
// find the attestations with the provided fromOffset. It returns the consensus block height
// that contains the attestation offset.
func binarySearch(
	ctx context.Context,
	attCl atypes.QueryClient,
	chainVer xchain.ChainVersion,
	chainName string,
	fromOffset uint64,
	startHeight uint64,
	endHeight uint64,
) (uint64, error) {
	const endpoint = "offset_binary_search"
	defer latency(endpoint)()

	for steps := 0; startHeight <= endHeight; steps++ {
		midHeight := startHeight + umath.SubtractOrZero(endHeight, startHeight)/2

		earliestAtt, ok, err := queryEarliestAttestation(ctx, attCl, chainVer, midHeight)
		if err != nil {
			incQueryErr(endpoint)
			return 0, errors.Wrap(err, "abci query earliest-attestation-in-state")
		} else if !ok {
			// If we're so far back that there's no attestation at all, move forward
			startHeight = midHeight + 1
			continue
		}

		if fromOffset == earliestAtt.AttestOffset {
			binarySearchStepsMetric(chainName, steps)
			return midHeight, nil
		}

		latestAtt, ok, err := queryLatestAttestation(ctx, attCl, chainVer, midHeight)
		if err != nil {
			incQueryErr(endpoint)
			return 0, errors.Wrap(err, "abci query latest-attestation")
		} else if !ok {
			return 0, errors.New("no latest attestation found despite earlier check [BUG]")
		}

		if fromOffset >= earliestAtt.AttestOffset && fromOffset <= latestAtt.AttestOffset {
			binarySearchStepsMetric(chainName, steps)

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
// until attestation found has lower or equal offset than the provided from offset, this
// guarantees the fromOffset attestation will be found in the range. Start and
// end height defining the range are returned.
//
//nolint:nonamedreturns // named returned for clarity
func lookbackRange(
	ctx context.Context,
	attCl atypes.QueryClient,
	chainVer xchain.ChainVersion,
	chainName string,
	fromOffset uint64,
	latestHeight uint64,
) (startHeight uint64, endHeight uint64, err error) {
	const endpoint = "offset_lookback"
	defer latency(endpoint)()

	endHeight = latestHeight
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
		earliestAtt, ok, err := queryEarliestAttestation(ctx, attCl, chainVer, queryHeight)
		if IsErrHistoryPruned(err) {
			// We've jumped to before the prune height, but _might_ still have the requested offset
			earliestStoreHeight, err := getEarliestStoreHeight(ctx, attCl, chainVer, queryHeight+1)
			if err != nil {
				return 0, 0, errors.Wrap(err, "failed to get earliest store height")
			}

			earliestAtt, ok, err = queryEarliestAttestation(ctx, attCl, chainVer, earliestStoreHeight)
			if err != nil {
				incQueryErr(endpoint)
				return 0, 0, errors.Wrap(err, "abci query earliest-attestation-in-state")
			}

			// If we're so far back that no attestation is found, or that we're before fromOffset,
			// that's a good breaking point for binary search
			if !ok || earliestAtt.AttestOffset <= fromOffset {
				lookbackStepsMetric(chainName, steps)
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
			lookbackStepsMetric(chainName, steps)
			startHeight = queryHeight

			return startHeight, endHeight, nil
		}

		// Otherwise, keep moving back
		endHeight = queryHeight
		lookback *= 2
	}
}
