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

func (b Backends) EVMBackends() ethbackend.Backends {
	resp := make(map[uint64]*ethbackend.Backend)
	for _, backend := range b {
		if backend.IsEVM() {
			resp[backend.chainID] = backend.EVMBackend()
		}
	}

	return ethbackend.BackendsFrom(resp)
}

func EVMBackends(backends ethbackend.Backends) Backends {
	resp := make(Backends)
	for _, backend := range backends.All() {
		_, id := backend.Chain()
		resp[id] = EVMBackend(backend)
	}

	return resp
}

type Backend struct {
	kind    uni.Kind
	eth     *ethbackend.Backend
	sol     *rpc.Client
	chainID uint64
}

func EVMBackend(backend *ethbackend.Backend) Backend {
	_, id := backend.Chain()

	return Backend{
		kind:    uni.KindEVM,
		eth:     backend,
		sol:     nil,
		chainID: id,
	}
}

func SVMBackend(cl *rpc.Client, chainID uint64) Backend {
	return Backend{
		kind:    uni.KindSVM,
		sol:     cl,
		chainID: chainID,
	}
}

func (b Backend) ChainID() uint64 {
	return b.chainID
}

func (b Backend) IsSVM() bool {
	return b.kind == uni.KindSVM
}

func (b Backend) IsEVM() bool {
	return b.kind == uni.KindEVM
}

func (b Backend) EVMClient() ethclient.Client {
	return b.eth
}

func (b Backend) EVMBackend() *ethbackend.Backend {
	return b.eth
}

func (b Backend) SVMClient() *rpc.Client {
	return b.sol
}
