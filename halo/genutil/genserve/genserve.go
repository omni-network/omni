// Package genserve provides a simple gRPC server that allows querying the consensus and execution genesis files.
package genserve

import (
	"context"

	grpc1 "github.com/cosmos/gogoproto/grpc"
)

// Register constructs a new genesis query server instance and registers it with the provided gRPC server.
func Register(s grpc1.Server, execution, consensus []byte) {
	RegisterQueryServer(s, &server{
		execution: execution,
		consensus: consensus,
	})
}

var _ QueryServer = (*server)(nil)

type server struct {
	execution []byte
	consensus []byte
}

func (s server) Genesis(context.Context, *GenesisRequest) (*GenesisResponse, error) {
	return &GenesisResponse{
		ExecutionGenesisJson: s.execution,
		ConsensusGenesisJson: s.consensus,
	}, nil
}
