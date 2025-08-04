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
	"github.com/omni-network/omni/e2e/manifests"
	e2etypes "github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/halo/genutil"
	"github.com/omni-network/omni/halo/genutil/evm"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts/omnitoken"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var (
	// forgeScriptABI is the ABI for the AllocPredeploys script.
	forgeScriptABI = mustGetABI(bindings.AllocPredeploysMetaData)

	// genValAlloc is the genesis validator allocation.
	genValAlloc = bi.Ether(genutil.ValidatorPower)
)

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
		// no allocs needed for simnet
		if network == netconf.Simnet {
			continue
		}

		if allocs.IsLocked(network) {
			continue
		}

		nativeBridgeBalance, err := getNativeBridgeBalance(network)
		if err != nil {
			return errors.Wrap(err, "get native bridge balance")
		}

		cfg := bindings.AllocPredeploysConfig{
			Manager:                eoa.MustAddress(network, eoa.RoleManager),
			Upgrader:               eoa.MustAddress(network, eoa.RoleUpgrader),
			Redenomizer:            eoa.MustAddress(network, eoa.RoleRedenomizer),
			ChainId:                bi.N(network.Static().OmniExecutionChainID),
			EnableStakingAllowlist: network.IsProtected(),
			NativeBridgeBalance:    nativeBridgeBalance,
			Output:                 "allocs/" + network.String() + ".json",
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

func getNativeBridgeBalance(network netconf.ID) (*big.Int, error) {
	resp := omnitoken.TotalSupply()

	// if not mainnet, return total supply
	if network != netconf.Mainnet {
		return resp, nil
	}

	// subtract prefunds
	prefunds, err := evm.PrefundAlloc(network)
	if err != nil {
		return nil, errors.Wrap(err, "prefund alloc")
	}

	for _, prefund := range prefunds {
		resp = bi.Sub(resp, prefund.Balance)
	}

	// subtract genesis validator allocations
	manifest, err := manifests.Mainnet()
	if err != nil {
		return nil, errors.Wrap(err, "mainnet manifest")
	}

	for _, node := range manifest.Nodes {
		// empty mode is validator (defauly)
		if node.Mode == string(e2etypes.ModeValidator) || node.Mode == "" {
			resp = bi.Sub(resp, genValAlloc)
		}
	}

	return resp, nil
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
