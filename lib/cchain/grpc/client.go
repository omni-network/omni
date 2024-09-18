package grpc

import (
	atypes "github.com/omni-network/omni/halo/attest/types"
	ptypes "github.com/omni-network/omni/halo/portal/types"
	"github.com/omni-network/omni/lib/errors"

	sltypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	sttypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	Staking  sttypes.QueryClient
	Slashing sltypes.QueryClient
	Attest   atypes.QueryClient
	Portal   ptypes.QueryClient
}

func Dial(target string) (Client, error) {
	grpcClient, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return Client{}, errors.Wrap(err, "new grpc client")
	}

	return Client{
		Staking:  sttypes.NewQueryClient(grpcClient),
		Slashing: sltypes.NewQueryClient(grpcClient),
		Attest:   atypes.NewQueryClient(grpcClient),
		Portal:   ptypes.NewQueryClient(grpcClient),
	}, nil
}
