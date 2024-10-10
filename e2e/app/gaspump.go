package app

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/gaspump"
	"github.com/omni-network/omni/lib/contracts/gasstation"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/txmgr"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

// DeployGasApp deploys OmniGasPump and OmniGasStation contracts.
func DeployGasApp(ctx context.Context, def Definition) error {
	if err := deployGasPumps(ctx, def); err != nil {
		return errors.Wrap(err, "deploy gas pumps")
	}

	if err := deployGasStation(ctx, def); err != nil {
		return errors.Wrap(err, "deploy gas station")
	}

	if err := fundGasStation(ctx, def); err != nil {
		return errors.Wrap(err, "fund gas station")
	}

	return nil
}

// deployGasPumps deploys OmniGasPump contracts to all chains except Omni's EVM.
func deployGasPumps(ctx context.Context, def Definition) error {
	network := NetworkFromDef(def)
	omniEVM, ok := network.OmniEVMChain()
	if !ok {
		return errors.New("no omni evm chain")
	}

	for _, chain := range network.EVMChains() {
		// GasPump not deployed on OmniEVM
		if chain.ID == omniEVM.ID {
			continue
		}

		backend, err := def.Backends().Backend(chain.ID)
		if err != nil {
			return errors.Wrap(err, "backend", "chain", chain.Name)
		}

		addr, receipt, err := gaspump.DeployIfNeeded(ctx, def.Testnet.Network, backend)
		if err != nil {
			return errors.Wrap(err, "deploy", "chain", chain.Name, "tx", maybeTxHash(receipt))
		}

		log.Info(ctx, "Gas pump deployed", "chain", chain.Name, "address", addr.Hex(), "tx", maybeTxHash(receipt))
	}

	return nil
}

// deployGasStation deploys OmniGasStation contract to Omni's EVM.
func deployGasStation(ctx context.Context, def Definition) error {
	network := NetworkFromDef(def)
	omniEVM, ok := network.OmniEVMChain()
	if !ok {
		return errors.New("no omni evm chain")
	}

	backend, err := def.Backends().Backend(omniEVM.ID)
	if err != nil {
		return errors.Wrap(err, "backend")
	}

	addrs, err := contracts.GetAddresses(ctx, network.ID)
	if err != nil {
		return errors.Wrap(err, "get addrs")
	}

	gasPumps := make([]bindings.OmniGasStationGasPump, 0, len(network.EVMChains())-1)
	for _, chain := range network.EVMChains() {
		if chain.ID == omniEVM.ID {
			continue
		}

		gasPumps = append(gasPumps, bindings.OmniGasStationGasPump{
			ChainID: chain.ID,
			Addr:    addrs.GasPump,
		})
	}

	addr, receipt, err := gasstation.DeployIfNeeded(ctx, def.Testnet.Network, backend, gasPumps)
	if err != nil {
		return errors.Wrap(err, "deploy", "tx", maybeTxHash(receipt))
	}

	log.Info(ctx, "Gas station deployed", "address", addr.Hex(), "tx", maybeTxHash(receipt))

	return nil
}

// fundGasStation funds a network's OmniGasStation contract on Omni's EVM.
//
// TODO: handle funding / monitoring properly.
// consider joining with e2e/app/eoa, or introduce something similar for contracts.
func fundGasStation(ctx context.Context, def Definition) error {
	network := NetworkFromDef(def)
	omniEVM, ok := network.OmniEVMChain()
	if !ok {
		return errors.New("no omni evm chain")
	}

	backend, err := def.Backends().Backend(omniEVM.ID)
	if err != nil {
		return errors.Wrap(err, "backend")
	}

	funder := eoa.MustAddress(network.ID, eoa.RoleFunder)

	addrs, err := contracts.GetAddresses(ctx, network.ID)
	if err != nil {
		return errors.Wrap(err, "get addrs")
	}

	// 1000 OMNI
	amt := new(big.Int).Mul(big.NewInt(1000), big.NewInt(params.Ether))

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

func maybeTxHash(receipt *ethtypes.Receipt) string {
	if receipt != nil {
		return receipt.TxHash.Hex()
	}

	return "nil"
}

type GasPumpTest struct {
	Recipient  common.Address
	TargetOMNI *big.Int
}

var (
	GasPumpTests = []GasPumpTest{
		{
			Recipient:  common.HexToAddress("0x0000000000000000000000000000000000001111"),
			TargetOMNI: big.NewInt(5000000000000000), // 0.005 OMNI
		},
		{
			Recipient:  common.HexToAddress("0x0000000000000000000000000000000000002222"),
			TargetOMNI: big.NewInt(10000000000000000), // 0.01 OMNI
		},
		{
			Recipient:  common.HexToAddress("0x0000000000000000000000000000000000003333"),
			TargetOMNI: big.NewInt(15000000000000000), // 0.015 OMNI
		},
	}
)

func testGasPumps(ctx context.Context, def Definition) error {
	networkID := def.Testnet.Network
	network := NetworkFromDef(def)

	addrs, err := contracts.GetAddresses(ctx, networkID)
	if err != nil {
		return errors.Wrap(err, "get addrs")
	}

	omniEVM, ok := network.OmniEVMChain()
	if !ok {
		return errors.New("no omni evm chain")
	}

	if !networkID.IsEphemeral() {
		log.Warn(ctx, "Skipping bridge test", errors.New("only ephemeral networks"))
		return nil
	}

	// just need an account with funds on ephemeral network chains
	payor := anvil.DevAccount7()

	for _, chain := range network.EVMChains() {
		if chain.ID == omniEVM.ID {
			continue
		}

		backend, err := def.Backends().Backend(chain.ID)
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

		for _, test := range GasPumpTests {
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

			if actualOMNI.Cmp(test.TargetOMNI) != 0 {
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
