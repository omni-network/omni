package app

import (
	"context"
	"sort"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/txmgr"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"

	"cosmossdk.io/math"
)

type valUpdate struct {
	Height int64
	Powers map[*e2e.Node]int64
}

func StartValidatorUpdates(ctx context.Context, def Definition) func() error {
	errChan := make(chan error, 1)
	returnErr := func(err error) {
		select {
		case errChan <- err:
		default:
			log.Error(ctx, "Error channel full, dropping error", err)
		}
	}

	network := externalNetwork(def.Testnet, def.Netman.DeployInfo())
	omniEVM, _ := network.OmniChain()
	funder, _, fundBackend, err := def.Backends.BindOpts(ctx, omniEVM.ID)
	if err != nil {
		return func() error { return errors.Wrap(err, "bind opts") }
	}

	go func() {
		// Get a sorted list of validator updates (and a map of total power per validator)
		var updates []valUpdate
		totalPowers := make(map[*e2e.Node]int64)
		for height, powers := range def.Testnet.ValidatorUpdates {
			updates = append(updates, valUpdate{
				Height: height,
				Powers: powers,
			})
			for node, power := range powers {
				totalPowers[node] += power
			}
		}
		sort.Slice(updates, func(i, j int) bool {
			return updates[i].Height < updates[j].Height
		})

		valBackend, err := ethbackend.NewBackend(omniEVM.Name, omniEVM.ID, omniEVM.BlockPeriod, fundBackend.Client)
		if err != nil {
			returnErr(errors.Wrap(err, "new backend"))
			return
		}

		omniStake, err := bindings.NewOmniStake(common.HexToAddress(predeploys.OmniStake), valBackend)
		if err != nil {
			returnErr(errors.Wrap(err, "new omni stake"))
			return
		}

		// Fund each validator with <total_power>+1 $OMNI to stake and pay for gas
		for node, power := range totalPowers {
			addr, _ := k1util.PubKeyToAddress(node.PrivvalKey.PubKey())
			tx, _, err := fundBackend.Send(ctx, funder, txmgr.TxCandidate{
				To:       &addr,
				GasLimit: 100_000,
				Value:    math.NewInt(power + 1).MulRaw(params.Ether).BigInt(),
			})
			if err != nil {
				returnErr(err)
				return
			} else if _, err := fundBackend.WaitMined(ctx, tx); err != nil {
				returnErr(err)
				return
			}

			// Add the validator privkey to the backend so we trigger deposits in its name.
			privKey, err := crypto.ToECDSA(node.PrivvalKey.Bytes())
			if err != nil {
				returnErr(errors.Wrap(err, "privkey to ecdsa"))
				return
			} else if _, err := valBackend.AddAccount(privKey); err != nil {
				returnErr(errors.Wrap(err, "add account"))
				return
			}
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

				txOpts, err := valBackend.BindOpts(ctx, addr)
				if err != nil {
					returnErr(errors.Wrap(err, "bind opts"))
					return
				}
				txOpts.Value = math.NewInt(power).MulRaw(params.GWei).BigInt()

				pubkey, err := valBackend.PublicKey(addr)
				if err != nil {
					returnErr(errors.Wrap(err, "public key"))
					return
				}

				log.Info(ctx, "Depositing stake",
					"validator", node.Name,
					"power", power,
				)

				tx, err := omniStake.Deposit(txOpts, k1util.PubKeyToBytes64(pubkey))
				if err != nil {
					returnErr(errors.Wrap(err, "deposit"))
					return
				} else if _, err := valBackend.WaitMined(ctx, tx); err != nil {
					returnErr(errors.Wrap(err, "wait minded"))
					return
				}
			}
		}

		returnErr(nil)
	}()

	return func() error {
		select {
		case err := <-errChan:
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
