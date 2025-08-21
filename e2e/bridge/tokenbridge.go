package bridge

import (
	"context"
	"math/big"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/l1bridge"
	"github.com/omni-network/omni/lib/contracts/omnitoken"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// DeployBridge deploys the OmniBridgeL1 & OmniToken contracts (if necessary), and configures the OmniBridgeNative predeploy.
func DeployBridge(ctx context.Context, testnet types.Testnet, backends ethbackend.Backends) error {
	networkID := testnet.Network

	if networkID.IsProtected() {
		return errors.New("cannot deploy bridge on protected networks")
	}

	// Amount deposited to L1 bridge on network creation.
	// We only deploy bridge on ephemeral networks now, so l1Deposits is always zero.
	l1Deposits := bi.Zero()

	l1, ok := testnet.EthereumChain()
	if !ok {
		log.Warn(ctx, "Skipping token bridge setup", errors.New("no ethereum L1 chain"))
		return nil
	}

	omniEVM, ok := testnet.OmniEVMChain()
	if !ok {
		return errors.New("no omni evm chain")
	}

	manager := eoa.MustAddress(networkID, eoa.RoleManager)

	addrs, err := contracts.GetAddresses(ctx, networkID)
	if err != nil {
		return errors.Wrap(err, "get addrs")
	}

	l1Backend, err := backends.Backend(l1.ChainID)
	if err != nil {
		return errors.Wrap(err, "backend")
	}

	omniBackend, err := backends.Backend(omniEVM.ChainID)
	if err != nil {
		return errors.Wrap(err, "omni backend")
	}

	// Deploy the bridge

	l1BridgeAddr, receipt, err := l1bridge.DeployIfNeeded(ctx, networkID, l1Backend)
	if err != nil {
		return errors.Wrap(err, "deploy l1 bridge")
	}

	if receipt != nil {
		log.Info(ctx, "Deployed L1 Bridge", "chain", l1.Name, "addr", l1BridgeAddr.Hex(), "block", receipt.BlockNumber)
	} else if l1BridgeAddr != networkID.Static().L1BridgeAddress {
		log.Warn(ctx, "L1 bridge already deployed, but not in network static", errors.New("missing static bridge addr"), "addr", l1BridgeAddr.Hex())
	}

	// Configure the OmniBridge native predeploy

	nativeBridgeAddr := common.HexToAddress(predeploys.OmniBridgeNative)
	nativeBridge, err := bindings.NewOmniBridgeNative(nativeBridgeAddr, omniBackend)
	if err != nil {
		return errors.Wrap(err, "bridge native")
	}

	txOpts, err := l1Backend.BindOpts(ctx, manager)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	tx, err := nativeBridge.Setup(txOpts, l1.ChainID, addrs.Portal, l1BridgeAddr, l1Deposits)
	if err != nil {
		return errors.Wrap(err, "setup bridge native")
	}

	_, err = omniBackend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	log.Info(ctx, "Configured OmniBridge native predeploy", "chain", l1.Name, "addr", predeploys.OmniBridgeNative)

	return nil
}

type ToAmount struct {
	To     common.Address
	Amount *big.Int
}

// Test bridges some tokens from L1 to OmniEVM, and some from OmniEVM to L1.
// Tokens must be bridged to OmniEVM first, before the native bridge contract will allow bridging back to L1.
func Test(ctx context.Context, testnet types.Testnet, backends ethbackend.Backends, toNatives, toL1s []ToAmount) error {
	networkID := testnet.Network
	if !networkID.IsEphemeral() {
		log.Warn(ctx, "Skipping bridge test", errors.New("only ephemeral networks"))
		return nil
	}

	if _, ok := testnet.EthereumChain(); !ok {
		log.Warn(ctx, "Skipping bridge test ", errors.New("no ethereum L1 chain"))
		return nil
	}

	if err := bridgeToNative(ctx, testnet, backends, toNatives); err != nil {
		return errors.Wrap(err, "bridge to native")
	}

	if err := waitNativeBridges(ctx, testnet, backends, toNatives); err != nil {
		return errors.Wrap(err, "wait native bridges")
	}

	if err := bridgeToL1(ctx, testnet, backends, toL1s); err != nil {
		return errors.Wrap(err, "bridge to L1")
	}

	return nil
}

// bridgeToNative bridges tokens from L1 to OmniEVM.
func bridgeToNative(ctx context.Context, testnet types.Testnet, backends ethbackend.Backends, toBridge []ToAmount) error {
	networkID := testnet.Network

	l1, ok := testnet.EthereumChain()
	if !ok {
		return errors.New("no ethereum L1 chain")
	}

	// payor is initial supply recipient, the only account with OMNI on L1 (for non-mainnet networks)
	payor, err := omnitoken.InitialSupplyRecipient(networkID)
	if err != nil {
		return err
	}

	addrs, err := contracts.GetAddresses(ctx, networkID)
	if err != nil {
		return errors.Wrap(err, "get addrs")
	}

	txOpts, backend, err := backends.BindOpts(ctx, l1.ChainID, payor)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	token, err := bindings.NewIERC20(addrs.Token, backend)
	if err != nil {
		return errors.Wrap(err, "token")
	}

	tx, err := token.Approve(txOpts, addrs.L1Bridge, omnitoken.TotalSupply())
	if err != nil {
		return errors.Wrap(err, "increase allowance")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	bridge, err := bindings.NewOmniBridgeL1(addrs.L1Bridge, backend)
	if err != nil {
		return errors.Wrap(err, "l1 bridge")
	}

	txns := make([]*ethtypes.Transaction, len(toBridge))

	for i, test := range toBridge {
		fee, err := bridge.BridgeFee(&bind.CallOpts{Context: ctx}, txOpts.From, test.To, test.Amount)
		if err != nil {
			return errors.Wrap(err, "bridge fee")
		}

		txOpts.Value = fee

		log.Debug(ctx, "Bridging to native", "to", test.To.Hex(), "amount", test.Amount, "fee", fee)

		tx, err := bridge.Bridge(txOpts, test.To, test.Amount)
		if err != nil {
			return errors.Wrap(err, "bridge")
		}

		txns[i] = tx
	}

	for _, tx := range txns {
		_, err := backend.WaitMined(ctx, tx)
		if err != nil {
			return errors.Wrap(err, "wait mined")
		}
	}

	return nil
}

// waitNativeBridges waits for all native bridge test cases to complete.
// This is required before bridging back to L1, because the native bridge must be informed that L1 tokens are available.
func waitNativeBridges(ctx context.Context, testnet types.Testnet, backends ethbackend.Backends, bridges []ToAmount) error {
	omniEVM, ok := testnet.OmniEVMChain()
	if !ok {
		return errors.New("no omni evm")
	}

	backend, err := backends.Backend(omniEVM.ChainID)
	if err != nil {
		return errors.Wrap(err, "backend")
	}

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return errors.Wrap(ctx.Err(), "timeout")
		case <-ticker.C:
			bridged := 0

			for _, test := range bridges {
				balance, err := backend.BalanceAt(ctx, test.To, nil)
				if err != nil {
					return errors.Wrap(err, "balance of")
				}

				if bi.EQ(balance, test.Amount) {
					bridged++
				}
			}

			if bridged == len(bridges) {
				log.Debug(ctx, "All native bridges complete")
				return nil
			}

			log.Debug(ctx, "Waiting for native bridges", "bridged", bridged, "total", len(bridges))
		}
	}
}

// bridgeToL1 bridges tokens from OmniEVM to L1.
func bridgeToL1(ctx context.Context, testnet types.Testnet, backends ethbackend.Backends, toBridge []ToAmount) error {
	omniEVM, ok := testnet.OmniEVMChain()
	if !ok {
		return errors.New("no omni evm chain")
	}

	backend, err := backends.Backend(omniEVM.ChainID)
	if err != nil {
		return errors.Wrap(err, "omni backend")
	}

	// payor is an anvil dev account, which is prefunded with native OMNI on ephemeral networks
	// note that in production, the only way to get native OMNI is to bridge it from L1
	payor := anvil.DevAccount7()

	bridge, err := bindings.NewOmniBridgeNative(common.HexToAddress(predeploys.OmniBridgeNative), backend)
	if err != nil {
		return errors.Wrap(err, "l1 bridge")
	}

	txOpts, err := backend.BindOpts(ctx, payor)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	txns := make([]*ethtypes.Transaction, len(toBridge))

	for i, test := range toBridge {
		fee, err := bridge.BridgeFee(&bind.CallOpts{Context: ctx}, test.To, test.Amount)
		if err != nil {
			return errors.Wrap(err, "bridge fee")
		}

		txOpts.Value = bi.Add(test.Amount, fee)

		log.Debug(ctx, "Bridging to L1", "to", test.To.Hex(), "amount", test.Amount, "fee", fee)

		tx, err := bridge.Bridge(txOpts, test.To, test.Amount)
		if err != nil {
			return errors.Wrap(err, "bridge")
		}

		txns[i] = tx
	}

	for _, tx := range txns {
		_, err := backend.WaitMined(ctx, tx)
		if err != nil {
			return errors.Wrap(err, "wait mined")
		}
	}

	return nil
}
