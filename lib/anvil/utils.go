package anvil

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
