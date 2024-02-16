package provider

import (
	"context"

	halopbv1 "github.com/omni-network/omni/halo/halopb/v1"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/xchain"

	rpcclient "github.com/cometbft/cometbft/rpc/client"

	"google.golang.org/protobuf/proto"
)

var _ FetchFunc = ABCIFetcher{}.ApprovedFrom

func NewABCIProvider(abci rpcclient.ABCIClient, chains map[uint64]string) Provider {
	backoffFunc := func(ctx context.Context) (func(), func()) {
		return expbackoff.NewWithReset(ctx, expbackoff.WithFastConfig())
	}

	return Provider{
		fetch:       NewABCIFetcher(abci).ApprovedFrom,
		backoffFunc: backoffFunc,
		chainNames:  chains,
	}
}

type ABCIFetcher struct {
	abci rpcclient.ABCIClient
}

func NewABCIFetcher(abci rpcclient.ABCIClient) ABCIFetcher {
	return ABCIFetcher{
		abci: abci,
	}
}

func (f ABCIFetcher) ApprovedFrom(ctx context.Context, chainID uint64, fromHeight uint64,
) ([]xchain.AggAttestation, error) {
	req := &halopbv1.ApprovedFromRequest{
		ChainId:    chainID,
		FromHeight: fromHeight,
	}
	bz, err := proto.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "marshal request")
	}

	resp, err := f.abci.ABCIQuery(ctx, halopbv1.HaloService_ApprovedFrom_FullMethodName, bz)
	if err != nil {
		return nil, errors.Wrap(err, "abci query approved-from",
			"chain_id", chainID,
			"from_height", fromHeight,
			"len", len(bz),
		)
	} else if !resp.Response.IsOK() {
		return nil, errors.New("abci query approved-from failed",
			"code", resp.Response.Code,
			"log", resp.Response.Log,
			"info", resp.Response.Info,
		)
	}

	response := new(halopbv1.ApprovedFromResponse)
	err = proto.Unmarshal(resp.Response.Value, response)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal approved-from response")
	}

	aggs, err := halopbv1.AggregatesFromProto(response.GetAggregates())
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal approved-from aggregates")
	}

	return aggs, nil
}
