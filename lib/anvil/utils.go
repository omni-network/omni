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

// FundERC20 funds the account with an ERC20 token balance.
// This only works on standard ERC20 tokens with _balances mapping at slot 0.
func FundERC20(ctx context.Context, client ethclient.Client,
	token common.Address, amount *big.Int, accounts ...common.Address) error {
	// storage value for _balances[address]
	svalue := common.BigToHash(amount).Hex()

	for _, account := range accounts {
		err := client.CallContext(ctx, nil, "anvil_setStorageAt", token, balanceSlot(account), svalue)
		if err != nil {
			return errors.Wrap(err, "fund erc20", "token", token, "account", account)
		}
	}

	return nil
}

var (
	// _balances[account] storage slot == keccak256(abi.encode(account, idx)).
	slotIdx = bi.Zero()
	slotABI = abi.Arguments{
		{Type: abi.Type{T: abi.AddressTy}},
		{Type: abi.Type{T: abi.UintTy, Size: 256}},
	}
)

// balanceSlot returns the storage slot _balances[account] for a standard ERC20 token (slot 0).
func balanceSlot(account common.Address) string {
	slot, err := slotABI.Pack(account, slotIdx)
	if err != nil {
		// known args, this should never happen
		panic(err)
	}

	return hexutil.Encode(crypto.Keccak256(slot))
}
