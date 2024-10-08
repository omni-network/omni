// genallocs generates predeploy allocs for each omni network id
// Usage: go run genallocs.go

package main

import (
	"fmt"
	"math/big"
	"os"
	"os/exec"

	"github.com/omni-network/omni/contracts/allocs"
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/halo/genutil/evm"
	"github.com/omni-network/omni/lib/contracts/omnitoken"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

var forgeScriptABI = mustGetABI(bindings.AllocPredeploysMetaData)

func main() {
	err := genallocs()

	if err != nil {
		fmt.Printf("genallocs failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("genallocs succeeded")
}

func genallocs() error {
	for _, network := range netconf.All() {
		// always skip simnet. skip mainnet until it is required
		if network == netconf.Simnet || network == netconf.Mainnet {
			continue
		}

		if allocs.IsLocked(network) {
			continue
		}

		prefunds, err := evm.PrefundAlloc(network)
		if err != nil {
			return errors.Wrap(err, "prefund alloc")
		}

		cfg, err := allocConfig(network, prefunds)
		if err != nil {
			return errors.Wrap(err, "alloc config")
		}

		calldata, err := forgeScriptABI.Pack("run", cfg)
		if err != nil {
			return errors.Wrap(err, "pack calldata")
		}

		_, err = execCmd(".", "forge", "script", "AllocPredeploys", "--root", "core", "--sig", hexutil.Encode(calldata))
		if err != nil {
			return errors.Wrap(err, "alloc predeploys")
		}

		// format and sort output

		formatted, err := execCmd(".", "jq", "-S", ".", cfg.Output)
		if err != nil {
			return errors.Wrap(err, "format output")
		}

		err = os.WriteFile(cfg.Output, []byte(formatted), 0644)
		if err != nil {
			return errors.Wrap(err, "write output")
		}
	}

	return nil
}

func allocConfig(network netconf.ID, prefunds types.GenesisAlloc) (bindings.AllocPredeploysConfig, error) {
	nativeBridgeBalance := new(big.Int).Set(omnitoken.TotalSupply)

	// for mainnet (when supported), we subtract the prefunds from the native bridge balance
	if network == netconf.Mainnet {
		for _, prefund := range prefunds {
			nativeBridgeBalance.Sub(nativeBridgeBalance, prefund.Balance)
		}

		// sanity check - require we do not subtract more than 100 OMNI
		minSaneBalance := new(big.Int).Sub(
			omnitoken.TotalSupply,
			new(big.Int).Mul(big.NewInt(100), big.NewInt(params.Ether)),
		)

		if nativeBridgeBalance.Cmp(minSaneBalance) < 0 {
			return bindings.AllocPredeploysConfig{}, errors.New("native bridge below sane min", "balance", nativeBridgeBalance)
		}
	}

	return bindings.AllocPredeploysConfig{
		Manager:                eoa.MustAddress(network, eoa.RoleManager),
		Upgrader:               eoa.MustAddress(network, eoa.RoleUpgrader),
		ChainId:                new(big.Int).SetUint64(network.Static().OmniExecutionChainID),
		EnableStakingAllowlist: network.IsProtected(),
		NativeBridgeBalance:    nativeBridgeBalance,
		Output:                 "allocs/" + network.String() + ".json",
	}, nil
}

func execCmd(dir string, cmd string, args ...string) (string, error) {
	c := exec.Command(cmd, args...)
	c.Dir = dir

	out, err := c.CombinedOutput()
	if err != nil {
		return string(out), errors.Wrap(err, "exec", "out", string(out))
	}

	return string(out), nil
}

func mustGetABI(metadata *bind.MetaData) *abi.ABI {
	abi, err := metadata.GetAbi()
	if err != nil {
		panic(err)
	}

	return abi
}
