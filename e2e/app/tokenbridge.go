package app

import (
	"context"
	"math/big"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/l1bridge"
	"github.com/omni-network/omni/lib/contracts/omnitoken"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// setupTokenBridge deploys the OmniBridgeL1 & OmniToken contracts (if necessary), and configures the OmniBridgeNative predeploy.
func setupTokenBridge(ctx context.Context, def Definition) error {
	networkID := def.Testnet.Network
	network := networkFromDef(def)

	l1, ok := network.EthereumChain()
	if !ok {
		log.Warn(ctx, "Skipping token bridge setup", errors.New("no ethereum L1 chain"))
		return nil
	}

	omniEVM, ok := network.OmniEVMChain()
	if !ok {
		return errors.New("no omni evm chain")
	}

	admin := eoa.MustAddress(networkID, eoa.RoleAdmin)

	portalAddr := contracts.Portal(networkID)

	l1Backend, err := def.Backends().Backend(l1.ID)
	if err != nil {
		return errors.Wrap(err, "backend")
	}

	omniBackend, err := def.Backends().Backend(omniEVM.ID)
	if err != nil {
		return errors.Wrap(err, "omni backend")
	}

	// Deploy the token

	tokenAddr, receipt, err := omnitoken.DeployIfNeeded(ctx, networkID, l1Backend)
	if err != nil {
		return errors.Wrap(err, "deploy omni token")
	}

	if receipt != nil {
		log.Info(ctx, "Deployed Omni Token", "chain", l1.Name, "addr", tokenAddr.Hex(), "block", receipt.BlockNumber)
	} else if tokenAddr != networkID.Static().TokenAddress {
		log.Warn(ctx, "Omni token already deployed, but not in network static", errors.New("missing static token addr"), "addr", tokenAddr.Hex())
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

	nativeBridge, err := bindings.NewOmniBridgeNative(common.HexToAddress(predeploys.OmniBridgeNative), omniBackend)
	if err != nil {
		return errors.Wrap(err, "bridge native")
	}

	txOpts, err := l1Backend.BindOpts(ctx, admin)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	tx, err := nativeBridge.Setup(txOpts, l1.ID, portalAddr, l1BridgeAddr)
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

type BridgeTest struct {
	To     common.Address
	Amount *big.Int
}

var ToNativeBridgeTests = []BridgeTest{
	{
		To:     common.HexToAddress("0x1111111111111111111111111111111111111111"),
		Amount: ether(1000),
	},
	{
		To:     common.HexToAddress("0x2222222222222222222222222222222222222222"),
		Amount: ether(1000),
	},
	{
		To:     common.HexToAddress("0x3333333333333333333333333333333333333333"),
		Amount: ether(1000),
	},
}

var ToL1BridgeTests = []BridgeTest{
	{
		To:     common.HexToAddress("0x1111111111111111111111111111111111111111"),
		Amount: ether(100),
	},
	{
		To:     common.HexToAddress("0x2222222222222222222222222222222222222222"),
		Amount: ether(100),
	},
	{
		To:     common.HexToAddress("0x3333333333333333333333333333333333333333"),
		Amount: ether(100),
	},
}

// testBridge bridges some tokens from L1 to OmniEVM, and some from OmniEVM to L1.
// Tokens must be bridged to OmniEVM first, before the native bridge contract will allow bridging back to L1.
func testBridge(ctx context.Context, def Definition) error {
	networkID := def.Testnet.Network
	network := networkFromDef(def)

	if !networkID.IsEphemeral() {
		log.Warn(ctx, "Skipping bridge test", errors.New("only ephemeral networks"))
		return nil
	}

	if _, ok := network.EthereumChain(); !ok {
		log.Warn(ctx, "Skipping bridge test ", errors.New("no ethereum L1 chain"))
		return nil
	}

	if err := bridgeToNative(ctx, def, ToNativeBridgeTests); err != nil {
		return errors.Wrap(err, "bridge to native")
	}

	if err := waitNativeBridges(ctx, def, ToNativeBridgeTests); err != nil {
		return errors.Wrap(err, "wait native bridges")
	}

	if err := bridgeToL1(ctx, def, ToL1BridgeTests); err != nil {
		return errors.Wrap(err, "bridge to L1")
	}

	return nil
}

// bridgeToNative bridges tokens from L1 to OmniEVM.
func bridgeToNative(ctx context.Context, def Definition, toBridge []BridgeTest) error {
	networkID := def.Testnet.Network
	network := networkFromDef(def)

	l1, ok := network.EthereumChain()
	if !ok {
		return errors.New("no ethereum L1 chain")
	}

	// payor is initial supply recipient, the only account with OMNI on L1
	payor := omnitoken.InitialSupplyRecipient(networkID)
	l1BridgeAddr := contracts.L1Bridge(networkID)
	tokenAddr := contracts.Token(networkID)

	txOpts, backend, err := def.Backends().BindOpts(ctx, l1.ID, payor)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	token, err := bindings.NewOmni(tokenAddr, backend)
	if err != nil {
		return errors.Wrap(err, "token")
	}

	tx, err := token.IncreaseAllowance(txOpts, l1BridgeAddr, omnitoken.TotalSupply)
	if err != nil {
		return errors.Wrap(err, "increase allowance")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	bridge, err := bindings.NewOmniBridgeL1(l1BridgeAddr, backend)
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
func waitNativeBridges(ctx context.Context, def Definition, bridges []BridgeTest) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	network := networkFromDef(def)
	omniEVM, ok := network.OmniEVMChain()
	if !ok {
		return errors.New("no omni evm")
	}

	backend, err := def.Backends().Backend(omniEVM.ID)
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

				if balance.Cmp(test.Amount) == 0 {
					bridged++
				}
			}

			if bridged == len(bridges) {
				log.Debug(ctx, "All native bridges complete")
				return nil
			}
		}
	}
}

// bridgeToL1 bridges tokens from OmniEVM to L1.
func bridgeToL1(ctx context.Context, def Definition, toBridge []BridgeTest) error {
	network := networkFromDef(def)

	omniEVM, ok := network.OmniEVMChain()
	if !ok {
		return errors.New("no omni evm chain")
	}

	backend, err := def.Backends().Backend(omniEVM.ID)
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

		txOpts.Value = new(big.Int).Add(test.Amount, fee)

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

func ether(n int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(n), big.NewInt(1e18))
}
