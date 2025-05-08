package anvil

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// FundAccounts funds the anvil account via the anvil_setBalance RPC method.
func FundAccounts(ctx context.Context, ethCl ethclient.Client, amount *big.Int, accounts ...common.Address) error {
	for _, account := range accounts {
		result := make(map[string]any)
		err := ethCl.CallContext(ctx, &result, "anvil_setBalance", account, hexutil.EncodeBig(amount))
		if err != nil {
			return errors.Wrap(err, "set balance", "account", account)
		}
	}

	return nil
}

type FundERC20Option func(*fundERC20Options)

type fundERC20Options struct {
	slotIdx *big.Int
}

func WithSlotIdx(slotIdx uint64) FundERC20Option {
	return func(o *fundERC20Options) {
		o.slotIdx = bi.N(slotIdx)
	}
}

func defaultFundERC20Options() fundERC20Options {
	return fundERC20Options{
		slotIdx: bi.Zero(), // Standard ERC20s have _balances mapping at slot 0.
	}
}

// FundERC20 funds the account with an ERC20 token balance.
// This only works on standard ERC20 tokens with _balances mapping at slot 0.
func FundERC20(ctx context.Context, client ethclient.Client, token common.Address, amount *big.Int, account common.Address, opts ...FundERC20Option) error {
	// storage value for _balances[address]
	svalue := common.BigToHash(amount).Hex()

	o := defaultFundERC20Options()
	for _, opt := range opts {
		opt(&o)
	}

	err := client.CallContext(ctx, nil, "anvil_setStorageAt", token, balanceSlot(account, o.slotIdx), svalue)
	if err != nil {
		return errors.Wrap(err, "fund erc20", "token", token, "account", account)
	}

	return nil
}

func FundUSDC(ctx context.Context, client ethclient.Client, token common.Address, amount *big.Int, account common.Address) error {
	// USDC mapping `balanceAndBlacklistStates` at slot 9.
	return FundERC20(ctx, client, token, amount, account, WithSlotIdx(9))
}

func FundL1USDT(ctx context.Context, client ethclient.Client, token common.Address, amount *big.Int, account common.Address) error {
	// L1 USDT mapping `_balances` at slot 2.
	return FundERC20(ctx, client, token, amount, account, WithSlotIdx(2))
}

func FundArbUSDT(ctx context.Context, client ethclient.Client, token common.Address, amount *big.Int, account common.Address) error {
	// Arb USDT mapping `_balances` at slot 51.
	return FundERC20(ctx, client, token, amount, account, WithSlotIdx(51))
}

func FundOPUSDT(ctx context.Context, client ethclient.Client, token common.Address, amount *big.Int, account common.Address) error {
	// OP USDT mapping `_balances` at slot 0.
	return FundERC20(ctx, client, token, amount, account, WithSlotIdx(0))
}

var (
	// _balances[account] storage slot == keccak256(abi.encode(account, slot index)).
	slotABI = abi.Arguments{
		{Type: abi.Type{T: abi.AddressTy}},
		{Type: abi.Type{T: abi.UintTy, Size: 256}},
	}
)

// balanceSlot returns the storage slot _balances[account] for a standard ERC20 token (slot 0).
func balanceSlot(account common.Address, slotIdx *big.Int) string {
	slot, err := slotABI.Pack(account, slotIdx)
	if err != nil {
		// known args, this should never happen
		panic(err)
	}

	return hexutil.Encode(crypto.Keccak256(slot))
}
