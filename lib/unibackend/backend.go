package unibackend

import (
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/uni"

	"github.com/gagliardetto/solana-go/rpc"
)

type Backends map[uint64]Backend

func (b Backends) Backend(chainID uint64) (Backend, error) {
	resp, ok := b[chainID]
	if !ok {
		return Backend{}, errors.New("unknown backend chain id", "chain_id", chainID)
	}

	return resp, nil
}

func EthBackends(backends ethbackend.Backends) Backends {
	resp := make(Backends)
	for _, backend := range backends.All() {
		_, id := backend.Chain()
		resp[id] = EthBackend(backend)
	}

	return resp
}

type Backend struct {
	kind    uni.Kind
	eth     *ethbackend.Backend
	sol     *rpc.Client
	chainID uint64
}

func EthBackend(backend *ethbackend.Backend) Backend {
	_, id := backend.Chain()

	return Backend{
		kind:    uni.KindEth,
		eth:     backend,
		sol:     nil,
		chainID: id,
	}
}

func SolBackend(cl *rpc.Client, chainID uint64) Backend {
	return Backend{
		kind:    uni.KindSol,
		sol:     cl,
		chainID: chainID,
	}
}

func (b Backend) ChainID() uint64 {
	return b.chainID
}

func (b Backend) IsSol() bool {
	return b.kind == uni.KindSol
}

func (b Backend) IsEth() bool {
	return b.kind == uni.KindEth
}

func (b Backend) EthClient() ethclient.Client {
	return b.eth
}

func (b Backend) EthBackend() *ethbackend.Backend {
	return b.eth
}

func (b Backend) SolClient() *rpc.Client {
	return b.sol
}
