package balancesnap

import (
	"context"
	"math/big"
	"strings"
	"sync/atomic"

	"github.com/omni-network/omni/halo/evmredenom"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethp2p"
	"github.com/omni-network/omni/lib/forkjoin"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/protocols/snap"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/trie/trienode"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	dtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/holiman/uint256"
)

type EVM struct {
	Address       common.Address `json:"address"`
	Balance       *big.Int       `json:"balance"`
	BalancePretty string         `json:"balance_pretty"`
}

type Stake struct {
	Address     common.Address `json:"address"`
	Total       *big.Int       `json:"total"` // Total: delegation + unbonding + rewards (all * evmredenom.Factor)
	TotalPretty string         `json:"total_pretty"`
	Delegation  *big.Int       `json:"delegation"` // Active delegated amount * evmredenom.Factor
	Unbonding   *big.Int       `json:"unbonding"`  // Unbonding amount * evmredenom.Factor
	Rewards     *big.Int       `json:"rewards"`    // Accrued rewards * evmredenom.Factor
}

// GetEVMBalances returns the balances of all EVM addresses at the specified state root.
// It uses the Ethereum snap protocol to fetch account ranges from a peer node.
//
// Parameters:
//   - peer: The ENR of the EVM node to fetch accounts from (must support snap protocol)
//   - stateRoot: The state root hash to query balances from (e.g., at halt height)
//   - archive: Archive node client for preimage lookups
//   - preimages: Map of account hash to address from genesis allocs
//   - batchSize: Number of accounts to fetch per batch (default: 5000 if 0)
func GetEVMBalances(
	ctx context.Context,
	peer *enode.Node,
	stateRoot common.Hash,
	archive ethclient.Client,
	preimages map[common.Hash]common.Address,
	batchSize uint64,
) ([]EVM, error) {
	if batchSize == 0 {
		batchSize = 5000 // Default batch size
	}

	// Generate a random node key for P2P connection
	nodeKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, errors.Wrap(err, "generate node key")
	}

	// Establish P2P connection to the EVM node
	cl, _, err := ethp2p.Dial(ctx, nodeKey, peer)
	if err != nil {
		return nil, errors.Wrap(err, "dial peer")
	}

	var balances []EVM
	var next common.Hash // Starting hash for account range

	for i := 0; ; i++ {
		// Fetch account range from peer
		resp, err := cl.AccountRange(ctx, stateRoot, next, batchSize)
		if err != nil {
			return nil, errors.Wrap(err, "get account range", "batch", i)
		} else if len(resp.Accounts) == 0 {
			return nil, errors.New("empty account range response", "batch", i)
		}

		// Verify the batch cryptographically
		done, err := verifyAccountBatch(stateRoot, next, resp)
		if err != nil {
			return nil, errors.Wrap(err, "verify batch", "batch", i)
		}

		// Process accounts in this batch
		for _, acc := range resp.Accounts {
			// Resolve account hash to address
			addr, err := resolveAddress(ctx, acc.Hash, preimages, archive)
			if err != nil {
				return nil, errors.Wrap(err, "resolve address", "hash", acc.Hash, "batch", i)
			}

			// Convert slim account body to full RLP format
			fullBody, err := etypes.FullAccountRLP(acc.Body)
			if err != nil {
				return nil, errors.Wrap(err, "convert to full account RLP", "addr", addr, "batch", i)
			}

			// Decode account body to extract balance
			var account etypes.StateAccount
			if err := rlp.DecodeBytes(fullBody, &account); err != nil {
				return nil, errors.Wrap(err, "decode account", "addr", addr, "batch", i)
			}

			// Only include accounts with non-zero balance
			if account.Balance.Sign() > 0 {
				bal := account.Balance.ToBig()
				balances = append(balances, EVM{
					Address:       addr,
					Balance:       bal,
					BalancePretty: FormatBalance(bal),
				})
			}
		}

		log.Info(ctx, "Processed account range batch",
			"batch", i,
			"accounts", len(resp.Accounts),
			"total_balances", len(balances),
		)

		if done {
			log.Info(ctx, "Completed fetching all EVM balances", "total", len(balances))
			return balances, nil
		}

		// Move to next range
		next = incHash(resp.Accounts[len(resp.Accounts)-1].Hash)
	}
}

// resolveAddress resolves an account hash to its address using preimages or archive node.
func resolveAddress(
	ctx context.Context,
	hash common.Hash,
	preimages map[common.Hash]common.Address,
	archive ethclient.Client,
) (common.Address, error) {
	// Check preimages map first
	if addr, ok := preimages[hash]; ok {
		return addr, nil
	}

	// Query archive node for preimage
	preimage, err := archive.Preimage(ctx, hash)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "get preimage from archive")
	}

	addr, err := cast.EthAddress(preimage)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "decode address from preimage")
	}

	return addr, nil
}

// verifyAccountBatch cryptographically verifies an account range batch using Merkle proofs.
// It returns true if this is the last batch (no more accounts to process).
func verifyAccountBatch(root, origin common.Hash, resp *snap.AccountRangePacket) (bool, error) {
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

// incHash returns the next hash in lexicographical order (i.e., plus one).
// Note: it rolls over for common.MaxHash, returning zero hash.
func incHash(h common.Hash) common.Hash {
	return new(uint256.Int).AddUint64(
		new(uint256.Int).SetBytes32(h[:]),
		1,
	).Bytes32()
}

// Example: 1234567890123456789012345678 -> 1_234_567_890.123456789012345678.
func FormatBalance(balance *big.Int) string {
	if balance == nil || balance.Sign() == 0 {
		return "0.000000000000000000"
	}

	// 10^18 for 18 decimal places
	decimals := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)

	// Get whole and fractional parts
	whole := new(big.Int).Div(balance, decimals)
	remainder := new(big.Int).Mod(balance, decimals)

	// Format whole part with underscores
	wholeStr := whole.String()
	wholeFormatted := addUnderscores(wholeStr)

	// Format fractional part (pad to 18 digits)
	fracStr := remainder.String()
	fracStr = strings.Repeat("0", 18-len(fracStr)) + fracStr

	return wholeFormatted + "." + fracStr
}

// Example: 1234567890 -> 1_234_567_890.
func addUnderscores(s string) string {
	if len(s) <= 3 {
		return s
	}

	var result strings.Builder
	for i, digit := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			_, _ = result.WriteRune('_')
		}
		_, _ = result.WriteRune(digit)
	}

	return result.String()
}

// GetStakingBalances returns the staking balances of all delegators on the consensus chain.
// It queries validators, delegations, unbonding delegations, and rewards.
// All amounts are multiplied by evmredenom.Factor to convert from consensus chain units to EVM NOM tokens.
func GetStakingBalances(ctx context.Context, cprov cchain.Provider) ([]Stake, error) {
	qc := cprov.QueryClients()

	// Get all validators (including unbonded/jailed)
	validators, err := cprov.SDKValidators(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get validators")
	}

	log.Info(ctx, "Found validators", "count", len(validators))

	// Collect all delegations and unbonding delegations by querying at validator level
	delegatorMap := make(map[string]*Stake)

	// Query delegations for all validators
	for _, val := range validators {
		// Get delegations
		delRequest := &stypes.QueryValidatorDelegationsRequest{
			ValidatorAddr: val.OperatorAddress,
			Pagination: &query.PageRequest{
				Limit: 10000,
			},
		}

		for {
			delResp, err := qc.Staking.ValidatorDelegations(ctx, delRequest)
			if err != nil {
				return nil, errors.Wrap(err, "query validator delegations", "validator", val.OperatorAddress)
			}

			for _, del := range delResp.DelegationResponses {
				if del.Balance.Denom != "stake" {
					continue
				}

				evmAddr, err := consensusAddrToEVM(del.Delegation.DelegatorAddress)
				if err != nil {
					return nil, errors.Wrap(err, "convert delegator address", "addr", del.Delegation.DelegatorAddress)
				}

				if _, exists := delegatorMap[evmAddr.Hex()]; !exists {
					delegatorMap[evmAddr.Hex()] = &Stake{
						Address:    evmAddr,
						Delegation: bi.Zero(),
						Unbonding:  bi.Zero(),
						Rewards:    bi.Zero(),
					}
				}

				delegatorMap[evmAddr.Hex()].Delegation = bi.Add(
					delegatorMap[evmAddr.Hex()].Delegation,
					del.Balance.Amount.BigInt(),
				)
			}

			if len(delResp.Pagination.NextKey) == 0 {
				break
			}
			delRequest.Pagination.Key = delResp.Pagination.NextKey
		}

		// Get unbonding delegations for this validator
		unbondingRequest := &stypes.QueryValidatorUnbondingDelegationsRequest{
			ValidatorAddr: val.OperatorAddress,
			Pagination: &query.PageRequest{
				Limit: 10000,
			},
		}

		for {
			unbondingResp, err := qc.Staking.ValidatorUnbondingDelegations(ctx, unbondingRequest)
			if err != nil {
				return nil, errors.Wrap(err, "query validator unbonding", "validator", val.OperatorAddress)
			}

			for _, unbonding := range unbondingResp.UnbondingResponses {
				evmAddr, err := consensusAddrToEVM(unbonding.DelegatorAddress)
				if err != nil {
					return nil, errors.Wrap(err, "convert unbonding address", "addr", unbonding.DelegatorAddress)
				}

				if _, exists := delegatorMap[evmAddr.Hex()]; !exists {
					delegatorMap[evmAddr.Hex()] = &Stake{
						Address:    evmAddr,
						Delegation: bi.Zero(),
						Unbonding:  bi.Zero(),
						Rewards:    bi.Zero(),
					}
				}

				for _, entry := range unbonding.Entries {
					delegatorMap[evmAddr.Hex()].Unbonding = bi.Add(
						delegatorMap[evmAddr.Hex()].Unbonding,
						entry.Balance.BigInt(),
					)
				}
			}

			if len(unbondingResp.Pagination.NextKey) == 0 {
				break
			}
			unbondingRequest.Pagination.Key = unbondingResp.Pagination.NextKey
		}
	}

	log.Info(ctx, "Collected delegations and unbonding", "unique_delegators", len(delegatorMap))

	// Convert to slice for forkjoin enrichment (rewards only)
	type delegatorInput struct {
		addr  string
		stake *Stake
	}
	var inputs []delegatorInput
	for addr, stake := range delegatorMap {
		inputs = append(inputs, delegatorInput{addr: addr, stake: stake})
	}

	// Enrich with rewards in parallel
	var i atomic.Int64
	enrichFunc := func(ctx context.Context, input delegatorInput) (Stake, error) {
		// Convert EVM address back to consensus address for rewards query
		consAddr := evmToConsensusAddr(input.stake.Address)

		rewards, err := getDelegatorRewards(ctx, qc.Distribution, consAddr)
		if err != nil {
			return Stake{}, errors.Wrap(err, "get delegator rewards", "addr", consAddr)
		}

		// Calculate total
		delegation := new(big.Int).Set(input.stake.Delegation)
		unbonding := new(big.Int).Set(input.stake.Unbonding)
		total := bi.Add(delegation, unbonding, rewards)

		// Multiply by evmredenom.Factor
		delegation = bi.MulRaw(delegation, evmredenom.Factor)
		unbonding = bi.MulRaw(unbonding, evmredenom.Factor)
		rewards = bi.MulRaw(rewards, evmredenom.Factor)
		total = bi.MulRaw(total, evmredenom.Factor)

		i.Add(1)
		if ii := i.Load(); ii%100 == 0 {
			log.Info(ctx, "Processed delegators", "count", ii)
		}

		return Stake{
			Address:     input.stake.Address,
			Total:       total,
			TotalPretty: FormatBalance(total),
			Delegation:  delegation,
			Unbonding:   unbonding,
			Rewards:     rewards,
		}, nil
	}

	result, cancel := forkjoin.NewWithInputs(ctx, enrichFunc, inputs, forkjoin.WithWorkers(16))
	defer cancel()

	stakes, err := result.Flatten()
	if err != nil {
		return nil, errors.Wrap(err, "enrich delegators with rewards")
	}

	log.Info(ctx, "Collected staking balances", "total_delegators", len(stakes))

	return stakes, nil
}

// consensusAddrToEVM converts a consensus bech32 address to EVM address.
func consensusAddrToEVM(bech32Addr string) (common.Address, error) {
	accAddr, err := sdk.AccAddressFromBech32(bech32Addr)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "parse bech32 address")
	}

	return cast.EthAddress(accAddr)
}

// evmToConsensusAddr converts an EVM address to consensus bech32 address.
func evmToConsensusAddr(evmAddr common.Address) string {
	accAddr := sdk.AccAddress(evmAddr.Bytes())
	return accAddr.String()
}

// Helper functions for staking queries

func getDelegatorRewards(ctx context.Context, qc dtypes.QueryClient, delegatorAddr string) (*big.Int, error) {
	resp, err := qc.DelegationTotalRewards(ctx, &dtypes.QueryDelegationTotalRewardsRequest{
		DelegatorAddress: delegatorAddr,
	})
	if err != nil {
		return nil, errors.Wrap(err, "query delegation rewards")
	}

	total := bi.Zero()
	for _, reward := range resp.Total {
		amount := reward.Amount.TruncateInt()
		total.Add(total, amount.BigInt())
	}

	return total, nil
}
