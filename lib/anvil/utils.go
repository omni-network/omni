package anvil

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
)

// FundAccounts funds the anvil account via the anvil_setBalance RPC method.
func FundAccounts(ctx context.Context, rpc string, amount *big.Int, accounts ...common.Address) error {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return err
	}
	defer client.Close()

	for _, account := range accounts {
		result := make(map[string]any)
		err = client.Client().CallContext(ctx, &result, "anvil_setBalance", account, hexutil.EncodeBig(amount))
		if err != nil {
			return errors.Wrap(err, "set balance", "account", account)
		}
	}

	return nil
}
