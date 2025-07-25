package submit

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/ethp2p"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/protocols/snap"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/trie/trienode"

	"github.com/holiman/uint256"
)

// Do submits the account range batches to the redenom contract.
// It connects to the given peer, retrieves account ranges from the state root,
// and submits them in batches of the given size until all accounts are processed.
func Do(
	ctx context.Context,
	from common.Address,
	backend *ethbackend.Backend,
	peer *enode.Node,
	stateRoot common.Hash,
	batchSize uint64,
) error {
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
	txOpts, err := backend.BindOpts(ctx, from)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	var next common.Hash
	for i := 0; ; i++ {
		resp, err := cl.AccountRange(ctx, stateRoot, next, batchSize)
		if err != nil {
			return errors.Wrap(err, "get account range")
		}

		done, err := verifyBatch(stateRoot, next, resp)
		if err != nil {
			return err
		}

		var addrs []common.Address
		var bodies [][]byte
		for _, acc := range resp.Accounts {
			preimage, err := backend.Preimage(ctx, acc.Hash)
			if err != nil {
				return errors.Wrap(err, "get preimage")
			}
			addr, err := cast.EthAddress(preimage)
			if err != nil {
				return errors.Wrap(err, "decode address from preimage")
			}

			addrs = append(addrs, addr)
			bodies = append(bodies, acc.Body)
			next = incHash(acc.Hash)
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

		log.Info(ctx, "Submitted account range batch", "tx", tx.Hash(), "index", i, "done", done, "next", next)

		if done {
			return nil
		}
	}
}

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

	done, err := trie.VerifyRangeProof(root, origin[:], hashes, bodies, proof.Set())
	if err != nil {
		return false, errors.Wrap(err, "verify range proof")
	}

	return done, nil
}

// incHash returns the next hash, in lexicographical order (a.k.a plus one).
// Note it rolls over for common.MaxHash, returning zero hash.
func incHash(h common.Hash) common.Hash {
	return new(uint256.Int).AddUint64(
		new(uint256.Int).SetBytes32(h[:]),
		1,
	).Bytes32()
}
