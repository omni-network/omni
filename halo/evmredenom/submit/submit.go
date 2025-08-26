package submit

import (
	"context"
	"strings"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/ethp2p"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/protocols/snap"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/trie/trienode"

	"github.com/holiman/uint256"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

// Do submits the account range batches to the redenom contract.
// It connects to the given peer, retrieves account ranges from the state root,
// and submits them in batches of the given size until all accounts are processed.
func Do(
	ctx context.Context,
	from common.Address,
	backend *ethbackend.Backend,
	archive ethclient.Client,
	peer *enode.Node,
	stateRoot common.Hash,
	batchSize uint64,
	concurrency int64,
	preimages map[common.Hash]common.Address,
) error {
	if concurrency <= 0 || batchSize == 0 {
		return errors.New("invalid concurrency or batch size", "concurrency", concurrency, "batch_size", batchSize)
	}

	nodeKey, err := crypto.GenerateKey() // Random node key
	if err != nil {
		return errors.Wrap(err, "generate node key")
	}

	cl, _, err := ethp2p.Dial(ctx, nodeKey, peer)
	if err != nil {
		return errors.Wrap(err, "dial peer")
	}

	contract, err := bindings.NewRedenom(common.HexToAddress(predeploys.Redenom), backend)
	if err != nil {
		return errors.Wrap(err, "new redenom contract")
	}

	// Ensure the contract owner matches the from address.
	if addr, err := contract.Owner(&bind.CallOpts{Context: ctx}); err != nil {
		return errors.Wrap(err, "get contract owner")
	} else if addr != from {
		return errors.New("contract owner mismatch", "expected", from, "got", addr)
	}

	var next common.Hash
	var prevNonce uint64
	eg, ctx := errgroup.WithContext(ctx)
	sema := semaphore.NewWeighted(concurrency)
	for i := 0; ; i++ {
		resp, err := cl.AccountRange(ctx, stateRoot, next, batchSize)
		if err != nil {
			return errors.Wrap(err, "get account range")
		} else if len(resp.Accounts) == 0 {
			return errors.New("empty account range response")
		}

		done, err := verifyBatch(stateRoot, next, resp)
		if err != nil {
			return err
		}

		next = incHash(resp.Accounts[len(resp.Accounts)-1].Hash)

		if err := sema.Acquire(ctx, 1); err != nil {
			return errors.Wrap(err, "acquire semaphore")
		}

		nonceCtx, nonce, err := backend.WithReservedNonce(ctx, from)
		if err != nil {
			return errors.Wrap(err, "get reserved nonce")
		} else if i > 0 && nonce <= prevNonce {
			return errors.New("nonce not incremented", "prev", prevNonce, "got", nonce)
		}

		log.Debug(ctx, "Submitting redenomination account range batch", "batch", i, "len", len(resp.Accounts), "from_map", len(preimages), "nonce", nonce)

		eg.Go(func() error {
			defer sema.Release(1)
			return submitBatch(nonceCtx, from, contract, backend, archive, preimages, resp)()
		})

		if !done {
			continue
		}

		if err := eg.Wait(); err != nil {
			return errors.Wrap(err, "wait submit batch")
		}

		log.Info(ctx, "All redenomination account ranges submitted", "total", i+1)

		return nil
	}
}

//nolint:nestif // Slightly complex nested logic.
func submitBatch(
	ctx context.Context,
	from common.Address,
	contract *bindings.Redenom,
	backend *ethbackend.Backend,
	archive ethclient.Client,
	preimages map[common.Hash]common.Address,
	resp *snap.AccountRangePacket,
) func() error {
	return func() error {
		var addrs []common.Address
		var bodies [][]byte
		var fromMap int
		for _, acc := range resp.Accounts {
			addr, ok := preimages[acc.Hash]
			if !ok {
				preimage, err := archive.Preimage(ctx, acc.Hash)
				if err != nil {
					if strings.Contains(err.Error(), "unknown preimage") {
						// Check all preimages first for better logs/stats.
						continue
					}

					return errors.Wrap(err, "get preimage", "hash", acc.Hash)
				}
				addr, err = cast.EthAddress(preimage)
				if err != nil {
					return errors.Wrap(err, "decode address from preimage", "preimage", preimage)
				}
			} else {
				fromMap++
			}

			addrs = append(addrs, addr)
			bodies = append(bodies, acc.Body)
		}

		if len(addrs) != len(resp.Accounts) {
			return errors.New("missing preimages", "got_total", len(addrs), "got_from_map", from, "expected", len(resp.Accounts), "missing", len(resp.Accounts)-len(addrs))
		}

		txOpts, err := backend.BindOpts(ctx, from)
		if err != nil {
			return errors.Wrap(err, "bind opts")
		}

		batch := bindings.RedenomAccountRange{
			Addresses: addrs,
			Accounts:  bodies,
			Proof:     resp.Proof,
		}

		tx, err := contract.Submit(txOpts, batch)
		if err != nil {
			return errors.Wrap(err, "submit batch")
		}

		rec, err := backend.WaitMined(ctx, tx)
		if err != nil {
			return errors.Wrap(err, "wait submit mined")
		}

		log.Info(ctx, "Submitted redenomination account range batch", "tx", tx.Hash(), "len", len(resp.Accounts), "from_map", fromMap, "block", rec.BlockNumber, "block_hash", rec.BlockHash.Hex())

		return nil
	}
}

// It returns true if this is the last batch (no more accounts to process),.
func verifyBatch(root, origin common.Hash, resp *snap.AccountRangePacket) (bool, error) {
	if len(resp.Accounts) == 0 {
		return false, errors.New("empty account range response")
	}

	hashes := make([][]byte, 0, len(resp.Accounts))
	bodies := make([][]byte, 0, len(resp.Accounts))
	for _, acc := range resp.Accounts {
		body, err := etypes.FullAccountRLP(acc.Body)
		if err != nil {
			return false, errors.Wrap(err, "decode account body")
		}
		bodies = append(bodies, body)
		hashes = append(hashes, acc.Hash[:])
	}

	proof := make(trienode.ProofList, 0, len(resp.Proof))
	for _, node := range resp.Proof {
		proof = append(proof, node)
	}

	more, err := trie.VerifyRangeProof(root, origin[:], hashes, bodies, proof.Set())
	if err != nil {
		return false, errors.Wrap(err, "verify range proof")
	}

	return !more, nil
}

// incHash returns the next hash, in lexicographical order (a.k.a plus one).
// Note it rolls over for common.MaxHash, returning zero hash.
func incHash(h common.Hash) common.Hash {
	return new(uint256.Int).AddUint64(
		new(uint256.Int).SetBytes32(h[:]),
		1,
	).Bytes32()
}
