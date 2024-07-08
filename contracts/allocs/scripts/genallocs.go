// genallocs generates predeploy allocs for each omni network id
// Usage: go run genallocs.go

package main

import (
	"fmt"
	"math/big"
	"os"
	"os/exec"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
		if network == netconf.Simnet {
			continue
		}

		cfg := bindings.AllocPredeploysConfig{
			Admin:                  eoa.MustAddress(network, eoa.RoleAdmin),
			ChainId:                new(big.Int).SetUint64(network.Static().OmniExecutionChainID),
			EnableStakingAllowlist: network.IsProtected(),
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
