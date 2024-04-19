package app

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"sort"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/txmgr"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"

	"cosmossdk.io/math"
)

// FundValidatorsForTesting funds validators in ephemeral networks: devnet and staging.
// This is required by load generation for periodic validator self-delegation.
func FundValidatorsForTesting(ctx context.Context, def Definition) error {
	if def.Testnet.Network != netconf.Devnet && def.Testnet.Network != netconf.Staging {
		// Only fund validators in ephemeral networks, devnet and staging.
		return nil
	}

	log.Info(ctx, "Funding validators for testing")

	network := externalNetwork(def)
	omniEVM, _ := network.OmniEVMChain()
	funder := def.Netman().Operator()
	_, fundBackend, err := def.Backends().BindOpts(ctx, omniEVM.ID, funder)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	// Iterate over all nodes, since all maybe become validators.
	for _, node := range def.Testnet.Nodes {
		addr, _ := k1util.PubKeyToAddress(node.PrivvalKey.PubKey())
		tx, _, err := fundBackend.Send(ctx, funder, txmgr.TxCandidate{
			To:       &addr,
			GasLimit: 100_000,
			Value:    math.NewInt(1000).MulRaw(params.Ether).BigInt(),
		})
		if err != nil {
			return errors.Wrap(err, "send")
		} else if _, err := fundBackend.WaitMined(ctx, tx); err != nil {
			return errors.Wrap(err, "wait mined")
		}
	}

	return nil
}

type valUpdate struct {
	Height int64
	Powers map[*e2e.Node]int64
}

func StartValidatorUpdates(ctx context.Context, def Definition) func() error {
	errChan := make(chan error, 1)
	returnErr := func(err error) {
		if err != nil {
			log.Error(ctx, "Validator updates failed", err)
		}
		select {
		case errChan <- err:
		default:
			log.Error(ctx, "Error channel full, dropping error", err)
		}
	}

	go func() {
		// Get all halo private keys
		var privkeys []*ecdsa.PrivateKey
		for _, node := range def.Testnet.Nodes {
			pk, err := k1util.StdPrivKeyFromComet(node.PrivvalKey)
			if err != nil {
				returnErr(err)
				return
			}

			privkeys = append(privkeys, pk)
		}

		// Get a sorted list of validator updates
		var updates []valUpdate
		for height, powers := range def.Testnet.ValidatorUpdates {
			updates = append(updates, valUpdate{
				Height: height,
				Powers: powers,
			})
		}
		sort.Slice(updates, func(i, j int) bool {
			return updates[i].Height < updates[j].Height
		})

		// Create a backend to trigger deposits from
		network := externalNetwork(def)
		omniEVM, _ := network.OmniEVMChain()
		ethCl, err := ethclient.Dial(omniEVM.Name, omniEVM.RPCURL)
		if err != nil {
			returnErr(errors.Wrap(err, "dial"))
			return
		}
		valBackend, err := ethbackend.NewBackend(omniEVM.Name, omniEVM.ID, omniEVM.BlockPeriod, ethCl, privkeys...)
		if err != nil {
			returnErr(errors.Wrap(err, "new backend"))
			return
		}

		// Create the OmniStake contract
		omniStake, err := bindings.NewOmniStake(common.HexToAddress(predeploys.OmniStake), valBackend)
		if err != nil {
			returnErr(errors.Wrap(err, "new omni stake"))
			return
		}

		// Wait for each update, then submit deposit txns.
		for _, update := range updates {
			log.Debug(ctx, "Waiting for next validator update", "wait_for_height", update.Height)
			_, _, err := waitForHeight(ctx, def.Testnet.Testnet, update.Height)
			if err != nil {
				returnErr(errors.Wrap(err, "wait for height"))
				return
			}

			for node, power := range update.Powers {
				addr, err := k1util.PubKeyToAddress(node.PrivvalKey.PubKey())
				if err != nil {
					returnErr(errors.Wrap(err, "pubkey to addr"))
					return
				}

				balance, err := valBackend.EtherBalanceAt(ctx, addr)
				if err != nil {
					returnErr(errors.Wrap(err, "balance att"))
					return
				}

				attrs := []any{"node", node.Name, "balance", balance}

				txOpts, err := valBackend.BindOpts(ctx, addr)
				if err != nil {
					returnErr(errors.Wrap(err, "bind opts"))
					return
				}
				txOpts.Value = math.NewInt(power).MulRaw(params.Ether).BigInt()

				pubkey, err := valBackend.PublicKey(addr)
				if err != nil {
					returnErr(errors.Wrap(err, "public key"))
					return
				}

				tx, err := omniStake.Deposit(txOpts, k1util.PubKeyToBytes64(pubkey))
				if err != nil {
					returnErr(errors.Wrap(err, "deposit", attrs...))
					return
				}
				rec, err := valBackend.WaitMined(ctx, tx)
				if err != nil {
					returnErr(errors.Wrap(err, "wait minded", attrs...))
					return
				}

				log.Info(ctx, "Deposited stake",
					"validator", node.Name,
					"power", power,
					"height", rec.BlockNumber.Uint64(),
					"balance", fmt.Sprintf("%.2f", balance),
				)
			}
		}

		returnErr(nil)
	}()

	return func() error {
		select {
		case err := <-errChan:
			return err
		case <-ctx.Done():
			return errors.Wrap(ctx.Err(), "timeout")
		}
	}
}
