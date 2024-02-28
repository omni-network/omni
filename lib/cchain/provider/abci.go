package provider

import (
	"context"

	atypes "github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/xchain"

	rpcclient "github.com/cometbft/cometbft/rpc/client"

	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gogogrpc "github.com/cosmos/gogoproto/grpc"
	"github.com/cosmos/gogoproto/proto"
	"google.golang.org/grpc"
)

func NewABCIProvider(abci rpcclient.ABCIClient, chains map[uint64]string) Provider {
	backoffFunc := func(ctx context.Context) (func(), func()) {
		return expbackoff.NewWithReset(ctx, expbackoff.WithFastConfig())
	}

	cl := atypes.NewQueryClient(rpcAdaptor{abci: abci})

	return Provider{
		fetch:       newABCIFetchFunc(cl),
		latest:      newABCILatestFunc(cl),
		window:      newABCIWindowFunc(cl),
		backoffFunc: backoffFunc,
		chainNames:  chains,
	}
}

func newABCIFetchFunc(cl atypes.QueryClient) func(ctx context.Context, chainID uint64, fromHeight uint64,
) ([]xchain.Attestation, error) {
	return func(ctx context.Context, chainID uint64, fromHeight uint64) ([]xchain.Attestation, error) {
		resp, err := cl.AttestationsFrom(ctx, &atypes.AttestationsFromRequest{
			ChainId:    chainID,
			FromHeight: fromHeight,
		})
		if err != nil {
			return nil, errors.Wrap(err, "abci query approved-from")
		}

		atts, err := atypes.AttestationsFromProto(resp.Attestations)
		if err != nil {
			return nil, errors.Wrap(err, "attestations from proto")
		}

		return atts, nil
	}
}

func newABCIWindowFunc(cl atypes.QueryClient) func(ctx context.Context, chainID uint64, height uint64,
) (int, error) {
	return func(ctx context.Context, chainID uint64, height uint64) (int, error) {
		resp, err := cl.WindowCompare(ctx, &atypes.WindowCompareRequest{
			ChainId: chainID,
			Height:  height,
		})
		if err != nil {
			return 0, errors.Wrap(err, "abci query window compare")
		}

		return int(resp.Cmp), nil
	}
}

func newABCILatestFunc(cl atypes.QueryClient) func(ctx context.Context, chainID uint64,
) (xchain.Attestation, bool, error) {
	return func(ctx context.Context, chainID uint64) (xchain.Attestation, bool, error) {
		resp, err := cl.LatestAttestation(ctx, &atypes.LatestAttestationRequest{
			ChainId: chainID,
		})
		if errors.Is(err, sdkerrors.ErrKeyNotFound) {
			return xchain.Attestation{}, false, nil
		} else if err != nil {
			return xchain.Attestation{}, false, errors.Wrap(err, "abci query latest attestation")
		}

		att, err := atypes.AttestationFromProto(resp.Attestation)
		if err != nil {
			return xchain.Attestation{}, false, errors.Wrap(err, "attestations from proto")
		}

		return att, true, nil
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
