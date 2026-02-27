package app

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/txmgr"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const drainRecipient = "0x79Ef4d1224a055Ad4Ee5e2226d0cb3720d929AE7"

// ethTransferGas is the gas limit for a simple ETH transfer.
const ethTransferGas uint64 = 21000

// DrainRelayerMonitor transfers all relayer and monitor ETH balances to the drain recipient on all chains.
func DrainRelayerMonitor(ctx context.Context, def Definition, dryRun bool) error {
	network := networkFromDef(def)
	endpoints := ExternalEndpoints(def)
	roles := []eoa.Role{eoa.RoleRelayer, eoa.RoleMonitor}
	recipient := common.HexToAddress(drainRecipient)

	// Load keys
	var keys []*ecdsa.PrivateKey
	for _, role := range roles {
		key, err := eoa.PrivateKey(ctx, network.ID, role)
		if err != nil {
			return errors.Wrap(err, "private key", "role", role)
		}

		keys = append(keys, key)
	}

	backends, err := ethbackend.BackendsFromNetwork(ctx, network, endpoints, keys...)
	if err != nil {
		return errors.Wrap(err, "backends from network")
	}

	// Drain each chain
	for _, chain := range network.EVMChains() {
		if evmchain.IsDisabled(chain.ID) {
			log.Warn(ctx, "Ignoring disabled chain", nil, "chain", chain.Name)
			continue
		}

		backend, err := backends.Backend(chain.ID)
		if err != nil {
			return errors.Wrap(err, "backend", "chain", chain.Name)
		}

		for _, role := range roles {
			addr, ok := eoa.Address(network.ID, role)
			if !ok {
				log.Warn(ctx, "No address for role, skipping", nil, "role", role, "chain", chain.Name)
				continue
			}

			if err := drainAccount(ctx, backend, chain.Name, role, addr, recipient, dryRun); err != nil {
				return errors.Wrap(err, "drain account", "chain", chain.Name, "role", role)
			}
		}
	}

	return nil
}

func drainAccount(
	ctx context.Context,
	backend *ethbackend.Backend,
	chainName string,
	role eoa.Role,
	from common.Address,
	to common.Address,
	dryRun bool,
) error {
	ctx = log.WithCtx(ctx, "chain", chainName, "role", role, "from", from)

	balance, err := backend.BalanceAt(ctx, from, nil)
	if err != nil {
		log.Warn(ctx, "Failed fetching balance, skipping", err)
		return nil
	}

	if balance.Sign() == 0 {
		log.Info(ctx, "Zero balance, skipping")
		return nil
	}

	log.Info(ctx, "Draining account",
		"balance", etherStr(balance),
		"to", to,
	)

	if dryRun {
		log.Warn(ctx, "Skipping actual drain tx due to dry-run", nil)
		return nil
	}

	tx, rec, err := transferNativeMax(ctx, backend, from, to, balance)
	if err != nil {
		return errors.Wrap(err, "transfer native max")
	} else if rec.Status != ethtypes.ReceiptStatusSuccessful {
		return errors.New("drain tx failed", "tx", tx.Hash())
	}

	remaining, err := backend.BalanceAt(ctx, from, nil)
	if err != nil {
		return errors.Wrap(err, "get remaining balance")
	}

	log.Info(ctx, "Account drained",
		"amount", etherStr(tx.Value()),
		"remaining", etherStr(remaining),
		"tx", tx.Hash().Hex(),
	)

	return nil
}

// maxDrainRetries is the number of times to retry with increasing gas reserve.
const maxDrainRetries = 5

// transferNativeMax attempts to transfer the maximum native balance from `from` to `to`.
// It estimates gas cost as gasPrice * 21000 * 2, and sends balance - gasCost.
// On L2s, actual fees can exceed the initial estimate, so it retries with 2x more gas
// reserve each attempt.
func transferNativeMax(
	ctx context.Context,
	backend *ethbackend.Backend,
	from common.Address,
	to common.Address,
	balance *big.Int,
) (*ethtypes.Transaction, *ethclient.Receipt, error) {
	gasPrice, err := backend.SuggestGasPrice(ctx)
	if err != nil {
		return nil, nil, errors.Wrap(err, "suggest gas price")
	}

	// Start with 2x gas reserve, double each retry.
	gasMultiplier := uint64(2)

	for range maxDrainRetries {
		gasCost := new(big.Int).Mul(gasPrice, new(big.Int).SetUint64(ethTransferGas*gasMultiplier))

		amount := bi.Sub(balance, gasCost)
		if amount.Sign() <= 0 {
			return nil, nil, errors.New("balance too low to cover gas",
				"balance", etherStr(balance),
				"gas_cost", etherStr(gasCost),
				"multiplier", gasMultiplier,
			)
		}

		log.Info(ctx, "Attempting transfer",
			"amount", etherStr(amount),
			"gas_reserve", etherStr(gasCost),
			"multiplier", gasMultiplier,
		)

		tx, rec, err := backend.Send(ctx, from, txmgr.TxCandidate{
			To:       &to,
			GasLimit: ethTransferGas,
			Value:    amount,
		})
		if err != nil {
			log.Warn(ctx, "Transfer failed, retrying with higher gas reserve", err, "multiplier", gasMultiplier)
			gasMultiplier *= 2

			continue
		}

		return tx, rec, nil
	}

	return nil, nil, errors.New("transfer failed after retries", "retries", maxDrainRetries)
}

// DrainAllowed returns an error if the drain command is not allowed for the given network.
func DrainAllowed(network netconf.ID) error {
	if network == netconf.Simnet || network == netconf.Devnet {
		return errors.New("cannot drain on simnet or devnet")
	}

	return nil
}
