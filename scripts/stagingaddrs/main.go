// Command stagingaddrs prints the contract addresses on the current live staging network.
// This is useful because staging create3 salts are derived from block 1 hash.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	ctx := context.Background()
	err := run(ctx)
	if err != nil {
		log.Error(ctx, "❌ Failed", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	addrs, err := contracts.GetAddresses(ctx, netconf.Staging)
	if err != nil {
		return errors.Wrap(err, "get addrs")
	}

	addrsJSON := map[string]common.Address{
		"create3":         addrs.Create3Factory,
		"avs":             addrs.AVS,
		"portal":          addrs.Portal,
		"l1bridge":        addrs.L1Bridge,
		"token":           addrs.Token,
		"nom":             addrs.NomToken,
		"gaspump":         addrs.GasPump,
		"gasstation":      addrs.GasStation,
		"solvernetinbox":  addrs.SolverNetInbox,
		"solvernetoutbox": addrs.SolverNetOutbox,
		"feeoraclev2":     addrs.FeeOracleV2,
	}

	prettyJSON, err := json.Marshal(addrsJSON)
	if err != nil {
		return errors.Wrap(err, "marshal json")
	}

	fmt.Println(string(prettyJSON))

	return nil
}
