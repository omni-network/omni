package provider

import (
	"context"

	halopb "github.com/omni-network/omni/halo/halopb/v1"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"

	rpcclient "github.com/cometbft/cometbft/rpc/client"

	"google.golang.org/protobuf/proto"
)

var _ FetchFunc = ABCIFetcher{}.ApprovedFrom

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
	req := &halopb.ApprovedFromRequest{
		ChainId:    chainID,
		FromHeight: fromHeight,
	}
	bz, err := proto.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "marshal request")
	}

	resp, err := f.abci.ABCIQuery(ctx, halopb.HaloService_ApprovedFrom_FullMethodName, bz)
	if err != nil {
		return nil, errors.Wrap(err, "abci query approved-from")
	} else if !resp.Response.IsOK() {
		return nil, errors.New("abci query approved-from failed",
			"code", resp.Response.Code,
			"log", resp.Response.Log,
			"info", resp.Response.Info,
		)
	}

	response := new(halopb.ApprovedFromResponse)
	err = proto.Unmarshal(resp.Response.Value, response)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal approved-from response")
	}

	aggs, err := halopb.AggregatesFromProto(response.GetAggregates())
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal approved-from aggregates")
	}

	return aggs, nil
}
