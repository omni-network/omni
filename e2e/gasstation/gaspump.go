package gasstation

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/gaspump"
	"github.com/omni-network/omni/lib/contracts/gasstation"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/txmgr"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// DeployEphemeralGasApp deploys OmniGasPump and OmniGasStation contracts to ephemeral networks.
func DeployEphemeralGasApp(ctx context.Context, testnet types.Testnet, backends ethbackend.Backends) error {
	if !testnet.Network.IsEphemeral() {
		return nil
	}

	if err := deployGasPumps(ctx, testnet, backends); err != nil {
		return errors.Wrap(err, "deploy gas pumps")
	}

	if err := deployGasStation(ctx, testnet, backends); err != nil {
		return errors.Wrap(err, "deploy gas station")
	}

	if testnet.Network != netconf.Mainnet {
		if err := fundGasStation(ctx, testnet, backends); err != nil {
			return errors.Wrap(err, "fund gas station")
		}
	}

	return nil
}

// deployGasPumps deploys OmniGasPump contracts to all chains except Omni's EVM.
func deployGasPumps(ctx context.Context, testnet types.Testnet, backends ethbackend.Backends) error {
	omniEVM, ok := testnet.OmniEVMChain()
	if !ok {
		return errors.New("no omni evm chain")
	}

	for _, chain := range testnet.EVMChains() {
		// GasPump not deployed on OmniEVM
		if chain.ChainID == omniEVM.ChainID {
			continue
		}

		backend, err := backends.Backend(chain.ChainID)
		if err != nil {
			return errors.Wrap(err, "backend", "chain", chain.Name)
		}

		addr, receipt, err := gaspump.DeployIfNeeded(ctx, testnet.Network, backend)
		if err != nil {
			return errors.Wrap(err, "deploy", "chain", chain.Name, "tx", maybeTxHash(receipt))
		}

		log.Info(ctx, "Gas pump deployed", "chain", chain.Name, "address", addr.Hex(), "tx", maybeTxHash(receipt))
	}

	return nil
}

// deployGasStation deploys OmniGasStation contract to Omni's EVM.
func deployGasStation(ctx context.Context, testnet types.Testnet, backends ethbackend.Backends) error {
	omniEVM, ok := testnet.OmniEVMChain()
	if !ok {
		return errors.New("no omni evm chain")
	}

	backend, err := backends.Backend(omniEVM.ChainID)
	if err != nil {
		return errors.Wrap(err, "backend")
	}

	addrs, err := contracts.GetAddresses(ctx, testnet.Network)
	if err != nil {
		return errors.Wrap(err, "get addrs")
	}

	gasPumps := make([]bindings.OmniGasStationGasPump, 0, len(testnet.EVMChains())-1)
	for _, chain := range testnet.EVMChains() {
		if chain.ChainID == omniEVM.ChainID {
			continue
		}

		gasPumps = append(gasPumps, bindings.OmniGasStationGasPump{
			ChainID: chain.ChainID,
			Addr:    addrs.GasPump,
		})
	}

	addr, receipt, err := gasstation.DeployIfNeeded(ctx, testnet.Network, backend, gasPumps)
	if err != nil {
		return errors.Wrap(err, "deploy", "tx", maybeTxHash(receipt))
	}

	log.Info(ctx, "Gas station deployed", "address", addr.Hex(), "tx", maybeTxHash(receipt))

	return nil
}

// fundGasStation funds a network's OmniGasStation contract on Omni's EVM.
func fundGasStation(ctx context.Context, testnet types.Testnet, backends ethbackend.Backends) error {
	omniEVM, ok := testnet.OmniEVMChain()
	if !ok {
		return errors.New("no omni evm chain")
	}

	backend, err := backends.Backend(omniEVM.ChainID)
	if err != nil {
		return errors.Wrap(err, "backend")
	}

	funder := eoa.MustAddress(testnet.Network, eoa.RoleHot)

	addrs, err := contracts.GetAddresses(ctx, testnet.Network)
	if err != nil {
		return errors.Wrap(err, "get addrs")
	}

	// 1000 OMNI
	amt := bi.Ether(1_000)

	tx, rec, err := backend.Send(ctx, funder, txmgr.TxCandidate{
		To:       &addrs.GasStation,
		GasLimit: 0,
		Value:    amt,
	})
	if err != nil {
		return errors.Wrap(err, "send tx")
	} else if rec.Status != ethtypes.ReceiptStatusSuccessful {
		return errors.New("fund tx failed", "tx", tx.Hash())
	}

	log.Info(ctx, "Funded gas station", "tx", tx.Hash(), "amount", amt)

	return nil
}

func maybeTxHash(receipt *ethclient.Receipt) string {
	if receipt != nil {
		return receipt.TxHash.Hex()
	}

	return "nil"
}

type GasPumpTest struct {
	Recipient  common.Address
	TargetOMNI *big.Int
}

func TestGasPumps(ctx context.Context, testnet types.Testnet, backends ethbackend.Backends, tests []GasPumpTest) error {
	networkID := testnet.Network

	addrs, err := contracts.GetAddresses(ctx, networkID)
	if err != nil {
		return errors.Wrap(err, "get addrs")
	}

	omniEVM, ok := testnet.OmniEVMChain()
	if !ok {
		return errors.New("no omni evm chain")
	}

	if !networkID.IsEphemeral() {
		log.Warn(ctx, "Skipping bridge test", errors.New("only ephemeral networks"))
		return nil
	}

	// just need an account with funds on ephemeral network chains
	payor := anvil.DevAccount7()

	for _, chain := range testnet.EVMChains() {
		if chain.ChainID == omniEVM.ChainID {
			continue
		}

		backend, err := backends.Backend(chain.ChainID)
		if err != nil {
			return errors.Wrap(err, "backend", "chain", chain.Name)
		}

		gasPump, err := bindings.NewOmniGasPump(addrs.GasPump, backend)
		if err != nil {
			return errors.Wrap(err, "new gas pump")
		}

		txOpts, err := backend.BindOpts(ctx, payor)
		if err != nil {
			return errors.Wrap(err, "bind opts")
		}

		for _, test := range tests {
			neededETH, err := gasPump.Quote(&bind.CallOpts{Context: ctx}, test.TargetOMNI)
			if err != nil {
				return errors.Wrap(err, "quote", "chain", chain.Name)
			}

			actualOMNI, wouldSucceed, revertReason, err := gasPump.DryFillUp(&bind.CallOpts{Context: ctx}, neededETH)
			if err != nil {
				return errors.Wrap(err, "dry fill up", "chain", chain.Name, "needed_eth", neededETH)
			}

			if !wouldSucceed {
				return errors.New("dry fill up failed", "chain", chain.Name, "revert_reason", revertReason)
			}

			if bi.NEQ(actualOMNI, test.TargetOMNI) {
				return errors.New("inaccurate quote", "chain", chain.Name, "actual_omni", actualOMNI, "provided_eth", neededETH, "target_omni", test.TargetOMNI)
			}

			txOpts.Value = neededETH
			tx, err := gasPump.FillUp(txOpts, test.Recipient)
			if err != nil {
				return errors.Wrap(err, "pump", "chain", chain.Name, "recipient", test.Recipient.Hex(), "target_omni", test.TargetOMNI, "needed_eth", neededETH)
			}

			log.Info(ctx, "Pumped gas", "chain", chain.Name, "tx", tx.Hash(), "recipient", test.Recipient.Hex(), "target_omni", test.TargetOMNI, "needed_eth", neededETH)
		}
	}

	return nil
}
