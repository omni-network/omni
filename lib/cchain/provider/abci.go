package provider

import (
	"context"
	"sync"
	"time"

	atypes "github.com/omni-network/omni/halo/attest/types"
	ptypes "github.com/omni-network/omni/halo/portal/types"
	rtypes "github.com/omni-network/omni/halo/registry/types"
	vtypes "github.com/omni-network/omni/halo/valsync/types"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tracer"
	"github.com/omni-network/omni/lib/xchain"

	rpcclient "github.com/cometbft/cometbft/rpc/client"

	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gogogrpc "github.com/cosmos/gogoproto/grpc"
	"github.com/cosmos/gogoproto/proto"
	"google.golang.org/grpc"
)

// ABCIClient abstracts the cometBFT RPC client consisting of only the required methods.
type ABCIClient interface {
	rpcclient.ABCIClient
	rpcclient.SignClient
}

func NewABCIProvider(abci ABCIClient, network netconf.ID, chainNamer func(xchain.ChainVersion) string) Provider {
	// Stream backoff for 1s, querying new attestations after 1 consensus block
	backoffFunc := func(ctx context.Context) func() {
		return expbackoff.New(ctx, expbackoff.WithPeriodicConfig(time.Second))
	}

	acl := atypes.NewQueryClient(rpcAdaptor{abci: abci})
	vcl := vtypes.NewQueryClient(rpcAdaptor{abci: abci})
	pcl := ptypes.NewQueryClient(rpcAdaptor{abci: abci})
	rcl := rtypes.NewQueryClient(rpcAdaptor{abci: abci})

	return Provider{
		fetch:       newABCIFetchFunc(acl),
		latest:      newABCILatestFunc(acl),
		window:      newABCIWindowFunc(acl),
		valset:      newABCIValsetFunc(vcl),
		portalBlock: newABCIPortalBlockFunc(pcl),
		networkFunc: newABCINetworkFunc(rcl),
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

func newABCIFetchFunc(cl atypes.QueryClient) fetchFunc {
	return func(ctx context.Context, chainVer xchain.ChainVersion, fromOffset uint64) ([]xchain.Attestation, error) {
		const endpoint = "attestations_from"
		defer latency(endpoint)()

		ctx, span := tracer.Start(ctx, spanName(endpoint))
		defer span.End()

		resp, err := cl.AttestationsFrom(ctx, &atypes.AttestationsFromRequest{
			ChainId:    chainVer.ID,
			ConfLevel:  uint32(chainVer.ConfLevel),
			FromOffset: fromOffset,
		})
		if err != nil {
			incQueryErr(endpoint)
			return nil, errors.Wrap(err, "abci query approved-from")
		}

		atts, err := atypes.AttestationsFromProto(resp.Attestations)
		if err != nil {
			return nil, errors.Wrap(err, "attestations from proto")
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

		resp, err := cl.LatestAttestation(ctx, &atypes.LatestAttestationRequest{
			ChainId:   chainVer.ID,
			ConfLevel: uint32(chainVer.ConfLevel),
		})
		if errors.Is(err, sdkerrors.ErrKeyNotFound) {
			return xchain.Attestation{}, false, nil
		} else if err != nil {
			incQueryErr(endpoint)
			return xchain.Attestation{}, false, errors.Wrap(err, "abci query latest attestation")
		}

		att, err := atypes.AttestationFromProto(resp.Attestation)
		if err != nil {
			return xchain.Attestation{}, false, errors.Wrap(err, "attestations from proto")
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

// rpcAdaptor adapts the cometBFT query client to the gRPC client interface.
// Note it only supports the Invoke method, not the NewStream method.
type rpcAdaptor struct {
	gogogrpc.ClientConn
	abci rpcclient.ABCIClient
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

	r, err := a.abci.ABCIQueryWithOptions(ctx, method, bz, rpcclient.ABCIQueryOptions{})
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
