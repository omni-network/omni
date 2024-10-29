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
	"github.com/cometbft/cometbft/rpc/client/http"

	"github.com/ethereum/go-ethereum/common"

	errorsmod "cosmossdk.io/errors"
	utypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	dtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	sltypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	gogogrpc "github.com/cosmos/gogoproto/grpc"
	"github.com/cosmos/gogoproto/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
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

	return NewABCI(cl, network, netconf.ChainVersionNamer(network)), nil
}

// NewABCI returns a new provider using the provided cometBFT ABCI client.
func NewABCI(cmtCl rpcclient.Client, network netconf.ID, chainNamer func(xchain.ChainVersion) string) Provider {
	return newProvider(rpcAdaptor{abci: cmtCl}, network, chainNamer)
}

// NewGRPC returns a new provider using the provided gRPC server address.
// This is preferred to NewABCI as it bypasses CometBFT so is much faster
// and doesn't affect chain performance.
func NewGRPC(target string, network netconf.ID, chainNamer func(xchain.ChainVersion) string) (Provider, error) {
	grpcClient, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return Provider{}, errors.Wrap(err, "new grpc client")
	}

	return newProvider(grpcClient, network, chainNamer), nil
}

func newProvider(cc gogogrpc.ClientConn, network netconf.ID, chainNamer func(xchain.ChainVersion) string) Provider {
	// Stream backoff for 1s, querying new attestations after 1 consensus block
	backoffFunc := func(ctx context.Context) func() {
		return expbackoff.New(ctx, expbackoff.WithPeriodicConfig(time.Second))
	}

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

	return Provider{
		fetch:       newABCIFetchFunc(acl, cmtcl, chainNamer),
		allAtts:     newABCIAllAttsFunc(acl),
		latest:      newABCILatestFunc(acl),
		window:      newABCIWindowFunc(acl),
		valset:      newABCIValsetFunc(vcl),
		val:         newABCIValFunc(scl),
		vals:        newABCIValsFunc(scl),
		signing:     newABCISigningFunc(slcl),
		rewards:     newABCIRewards(dcl),
		portalBlock: newABCIPortalBlockFunc(pcl),
		networkFunc: newABCINetworkFunc(rcl),
		genesisFunc: newABCIGenesisFunc(gcl),
		plannedFunc: newABCIPlannedUpgradeFunc(ucl),
		appliedFunc: newABCIAppliedUpgradeFunc(ucl),
		chainID:     newChainIDFunc(cmtcl),
		backoffFunc: backoffFunc,
		chainNamer:  chainNamer,
		network:     network,
	}
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
			ValSetID:      resp.Id,
			Validators:    vals,
			CreatedHeight: resp.CreatedHeight,
			activedHeight: resp.ActivatedHeight,
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
	return attestations{
		query:      attCl,
		service:    cmtCl,
		chainNamer: chainNamer,
	}.Find
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

// queryLatestAttestation returns the latest approved attestation for the provided chain version
// at the provided consensus block height, or the latest block height if height is 0.
func queryLatestAttestation(ctx context.Context, cl atypes.QueryClient, chainVer xchain.ChainVersion, height uint64) (xchain.Attestation, bool, error) {
	resp, err := cl.LatestAttestation(withCtxHeight(ctx, height),
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
