package netman

import (
	"context"
	"fmt"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common"
)

func logBalance(ctx context.Context, backend *ethbackend.Backend, chain string, addr common.Address, name string,
) error {
	balance, err := backend.EtherBalanceAt(ctx, addr)
	if err != nil {
		return errors.Wrap(err, "get balance")
	}

	log.Info(ctx, "Provided public chain key balance",
		"chain", chain,
		"address", addr.Hex(),
		"balance", fmt.Sprintf("%.2f", balance),
		"key_name", name,
	)

	return nil
}
