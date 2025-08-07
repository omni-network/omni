package keeper

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/halo/evmredenom"
	"github.com/omni-network/omni/halo/evmredenom/types"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	evmenginetypes "github.com/omni-network/omni/octane/evmengine/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/trie/trienode"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
)

type Keeper struct {
	status    StatusTable
	evmEngine types.EVMEngineKeeper
	contract  *bindings.Redenom
	address   common.Address
}

func New(
	storeService store.KVStoreService,
) (*Keeper, error) {
	schema := &ormv1alpha1.ModuleSchemaDescriptor{SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: File_halo_evmredenom_keeper_evmredenom_proto.Path()},
	}}

	modDB, err := ormdb.NewModuleDB(schema, ormdb.ModuleDBOptions{KVStoreService: storeService})
	if err != nil {
		return nil, errors.Wrap(err, "create module db")
	}

	s, err := NewEvmredenomStore(modDB)
	if err != nil {
		return nil, errors.Wrap(err, "create store")
	}

	address := common.HexToAddress(predeploys.Redenom)
	contract, err := bindings.NewRedenom(address, nil) // Passing nil backend if safe since only Parse functions are used.
	if err != nil {
		return &Keeper{}, errors.Wrap(err, "new redenom contract")
	}

	return &Keeper{
		status:   s.StatusTable(),
		contract: contract,
		address:  address,
	}, nil
}

func (p *Keeper) SetEVMEngineKeeper(keeper types.EVMEngineKeeper) {
	p.evmEngine = keeper
}

func (*Keeper) Name() string {
	return types.ModuleName
}

// InitStatus initializes the redenomination with the provided block state root hash.
// This block defines the accounts and balances that will be redenominated.
func (p *Keeper) InitStatus(ctx context.Context, root common.Hash) error {
	if s, err := p.status.Get(ctx); err != nil {
		return errors.Wrap(err, "get status")
	} else if len(s.GetRoot()) != 0 {
		return errors.New("status already initialized [BUG]")
	}

	var zero common.Hash
	status := Status{
		Root: root[:],
		Done: false,
		Next: zero[:], // First batch always starts from zero.
	}

	if err := p.status.Save(ctx, &status); err != nil {
		return errors.Wrap(err, "set status")
	}

	return nil
}

// FilterParams defines the matching EVM log events, see github.com/ethereum/go-ethereum#FilterQuery.
func (p *Keeper) FilterParams() ([]common.Address, [][]common.Hash) {
	return []common.Address{p.address}, [][]common.Hash{{submittedEvent.ID}}
}

// Deliver processes a omni redenom submitted log event.
func (p *Keeper) Deliver(ctx context.Context, _ common.Hash, elog evmenginetypes.EVMEvent) error {
	ethlog, err := elog.ToEthLog()
	if err != nil {
		return err
	}

	if ethlog.Topics[0] != submittedEvent.ID {
		return errors.New("unknown event")
	}

	submitted, err := p.contract.ParseSubmitted(ethlog)
	if err != nil {
		return errors.Wrap(err, "parse submitted event")
	}

	status, err := p.verifyBatch(ctx, submitted)
	if err != nil {
		return errors.Wrap(err, "verify batch")
	}

	if err := p.status.Save(ctx, status); err != nil {
		return errors.Wrap(err, "save status")
	}

	var nonZero int
	for i, body := range submitted.Accounts {
		addr := submitted.Addresses[i]

		account, err := etypes.FullAccount(body)
		if err != nil {
			return errors.Wrap(err, "decode account body")
		}

		mint, err := calcMint(account.Balance, evmredenom.Factor)
		if err != nil {
			return err
		}

		if eoa.IsDevAccount(addr) {
			log.Info(ctx, "Redenominating dev account", "address", addr, "balance", account.Balance, "mint", mint)
		}

		if bi.IsZero(mint) {
			continue
		}
		nonZero++

		if err := p.evmEngine.InsertWithdrawal(ctx, addr, mint); err != nil {
			return errors.Wrap(err, "insert withdrawal")
		}
	}

	log.Info(ctx, "🌾 Processed redenomination batch",
		"total", len(submitted.Accounts),
		"non_zero", nonZero,
		"next", hexutil.Encode(status.GetNext()),
		"done", status.GetDone(),
	)

	return nil
}

// verifyBatch verifies the batch of accounts and returns the new status.
func (p *Keeper) verifyBatch(ctx context.Context, batch *bindings.RedenomSubmitted) (*Status, error) {
	if batch == nil {
		return nil, errors.New("batch is nil")
	}

	status, err := p.status.Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get status")
	} else if len(status.GetRoot()) == 0 {
		return nil, errors.New("redenomination not initialized")
	} else if status.GetDone() {
		return nil, errors.New("redenomination already done")
	}

	root, err := cast.EthHash(status.GetRoot())
	if err != nil {
		return nil, err
	}

	first, err := cast.EthHash(status.GetNext())
	if err != nil {
		return nil, err
	}

	if len(batch.Addresses) == 0 || len(batch.Accounts) == 0 || len(batch.Proof) == 0 {
		return nil, errors.New("empty batch")
	} else if len(batch.Addresses) != len(batch.Accounts) {
		return nil, errors.New("batch address and account length mismatch", "accounts", len(batch.Accounts), "addresses", len(batch.Addresses))
	}

	// Convert bodies from slim to full RLP format.
	bodies := make([][]byte, 0, len(batch.Accounts))
	for _, body := range batch.Accounts {
		fullBody, err := etypes.FullAccountRLP(body)
		if err != nil {
			return nil, errors.Wrap(err, "decode account body")
		}
		bodies = append(bodies, fullBody)
	}

	// Construct a proof set
	proof := make(trienode.ProofList, 0, len(batch.Proof))
	for _, node := range batch.Proof {
		proof = append(proof, node)
	}

	// Calculate account hashes from addresses.
	hashes := make([][]byte, 0, len(batch.Addresses))
	for _, addr := range batch.Addresses {
		hashes = append(hashes, crypto.Keccak256(addr[:]))
	}

	// Verify the range proof against the trie.
	// More indicates whether there are more batches or whether this is the last batch.
	more, err := trie.VerifyRangeProof(root, first[:], hashes, bodies, proof.Set())
	if err != nil {
		return nil, errors.Wrap(err, "verify batch")
	}

	// Calculate next hash from last address hash.
	last, err := cast.EthHash(hashes[len(hashes)-1])
	if err != nil {
		return nil, errors.Wrap(err, "cast last account hash")
	}
	next := incHash(last)
	if !more {
		next = common.Hash{} // If done, reset next to zero hash.
	}

	return &Status{
		Root: root[:],
		Done: !more,
		Next: next[:],
	}, nil
}
