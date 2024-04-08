package state

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
)

// MemDB is an in stateful wrapper around core.Genesis, that makes it easier to build genesis state.
type MemDB struct {
	rw      sync.RWMutex
	genesis *core.Genesis
}

func NewMemDB(genesis *core.Genesis) *MemDB {
	return &MemDB{
		genesis: genesis,
		rw:      sync.RWMutex{},
	}
}

func (db *MemDB) Genesis() *core.Genesis {
	return db.genesis
}

func (db *MemDB) SetCode(addr common.Address, code []byte) {
	db.rw.Lock()
	defer db.rw.Unlock()

	account := db.getOrCreate(addr)
	account.Code = code
	db.genesis.Alloc[addr] = account
}

func (db *MemDB) SetState(addr common.Address, key, value common.Hash) {
	db.rw.Lock()
	defer db.rw.Unlock()

	account := db.getOrCreate(addr)
	account.Storage[key] = value
	db.genesis.Alloc[addr] = account
}

func (db *MemDB) getOrCreate(addr common.Address) types.Account {
	account, ok := db.genesis.Alloc[addr]

	if !ok {
		return types.Account{
			Code:    []byte{},
			Storage: make(map[common.Hash]common.Hash),
			Balance: big.NewInt(0),
			Nonce:   0,
		}
	}

	return account
}
