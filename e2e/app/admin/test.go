//nolint:gosec // no need for secure randomneness
package admin

import (
	"context"
	"math/rand"
	"slices"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// Test tests all admin commands against an ephemeral network.
func Test(ctx context.Context, def app.Definition) error {
	if !def.Testnet.Network.IsEphemeral() {
		return errors.New("only ephemeral networks")
	}

	log.Info(ctx, "Running contract admin tests.")

	if err := testEnsurePortalSpec(ctx, def); err != nil {
		return err
	}

	if err := testEnsureBridgeSpec(ctx, def); err != nil {
		return err
	}

	if err := testUpgradePortal(ctx, def); err != nil {
		return err
	}

	if err := tesUpgradeFeeOracleV1(ctx, def); err != nil {
		return err
	}

	if err := testUpgradeGasStation(ctx, def); err != nil {
		return err
	}

	if err := testUpgradeGasPump(ctx, def); err != nil {
		return err
	}

	if err := testUpgradeSlashing(ctx, def); err != nil {
		return err
	}

	if err := testUpgradeStaking(ctx, def); err != nil {
		return err
	}

	/*
		if err := testUpgradeBridgeNative(ctx, def); err != nil {
			return err
		}
	*/

	if err := testUpgradeBridgeL1(ctx, def); err != nil {
		return err
	}

	log.Info(ctx, "Done.")

	return nil
}

// noCheck always returns nil. Use for upgrade actions, where only check is if upgrade succeeds.
func noCheck(context.Context, app.Definition, types.EVMChain) error { return nil }

// testUpgradePortal tests UpgradePortal command.
func testUpgradePortal(ctx context.Context, def app.Definition) error {
	err := forOne(ctx, def, randChain(def.Testnet.EVMChains()), UpgradePortal, noCheck)
	if err != nil {
		return errors.Wrap(err, "upgrade portal")
	}

	err = forAll(ctx, def, UpgradePortal, noCheck)
	if err != nil {
		return errors.Wrap(err, "upgrade all portals")
	}

	return nil
}

func tesUpgradeFeeOracleV1(ctx context.Context, def app.Definition) error {
	err := forOne(ctx, def, randChain(def.Testnet.EVMChains()), UpgradeFeeOracleV1, noCheck)
	if err != nil {
		return errors.Wrap(err, "upgrade feeoracle")
	}

	err = forAll(ctx, def, UpgradeFeeOracleV1, noCheck)
	if err != nil {
		return errors.Wrap(err, "upgrade all feeoracles")
	}

	return nil
}

func testUpgradeGasStation(ctx context.Context, def app.Definition) error {
	err := UpgradeGasStation(ctx, def, Config{Broadcast: true})
	if err != nil {
		return errors.Wrap(err, "upgrade gas station")
	}

	return nil
}

func testUpgradeGasPump(ctx context.Context, def app.Definition) error {
	// cannot UpgradeGasPump on omni evm
	c := randChain(def.Testnet.EVMChains())
	for c.Name == omniEVMName {
		c = randChain(def.Testnet.EVMChains())
	}

	err := forOne(ctx, def, c, UpgradeGasPump, noCheck)
	if err != nil {
		return errors.Wrap(err, "upgrade gas pump")
	}

	err = forAll(ctx, def, UpgradeGasPump, noCheck)
	if err != nil {
		return errors.Wrap(err, "upgrade all gas pumps")
	}

	return nil
}

func testUpgradeSlashing(ctx context.Context, def app.Definition) error {
	err := UpgradeSlashing(ctx, def, Config{Broadcast: true})
	if err != nil {
		return errors.Wrap(err, "upgrade slashing")
	}

	return nil
}

func testUpgradeStaking(ctx context.Context, def app.Definition) error {
	err := UpgradeStaking(ctx, def, Config{Broadcast: true})
	if err != nil {
		return errors.Wrap(err, "upgrade staking")
	}

	return nil
}

/*
func testUpgradeBridgeNative(ctx context.Context, def app.Definition) error {
	err := UpgradeBridgeNative(ctx, def, Config{Broadcast: true})
	if err != nil {
		return errors.Wrap(err, "upgrade bridge native")
	}

	return nil
}
*/

func testUpgradeBridgeL1(ctx context.Context, def app.Definition) error {
	err := UpgradeBridgeL1(ctx, def, Config{Broadcast: true})
	if err != nil {
		return errors.Wrap(err, "upgrade bridge l1")
	}

	return nil
}

func testEnsurePortalSpec(ctx context.Context, def app.Definition) error {
	chains := def.Testnet.EVMChains()
	expected := randPortalSpec(chains)

	ensurePortalSpec := func(ctx context.Context, def app.Definition, cfg Config) error {
		return EnsurePortalSpec(ctx, def, cfg, expected)
	}

	err := forOne(ctx, def, randChain(chains), ensurePortalSpec, checkPortalSpec(chains, expected))
	if err != nil {
		return errors.Wrap(err, "ensure portal spec")
	}

	// new random expected values
	*expected = *randPortalSpec(def.Testnet.EVMChains())

	err = forAll(ctx, def, ensurePortalSpec, checkPortalSpec(chains, expected))
	if err != nil {
		return errors.Wrap(err, "ensure all portal specs")
	}

	return nil
}

func testEnsureBridgeSpec(ctx context.Context, def app.Definition) error {
	addrs, err := contracts.GetAddresses(ctx, def.Testnet.Network)
	if err != nil {
		return errors.Wrap(err, "get addrs")
	}

	omniEVM, ok := def.Testnet.OmniEVMChain()
	if !ok {
		return errors.New("no omni evm chain")
	}

	omniBackend, err := def.Backends().Backend(omniEVM.ChainID)
	if err != nil {
		return errors.Wrap(err, "backend", "chain", omniEVM.Name)
	}

	l1, ok := def.Testnet.EthereumChain()
	if !ok {
		return errors.New("no ethereum chain")
	}

	l1Backend, err := def.Backends().Backend(l1.ChainID)
	if err != nil {
		return errors.Wrap(err, "backend", "chain", l1.Name)
	}

	nativebridge, err := bindings.NewOmniBridgeNative(common.HexToAddress(predeploys.OmniBridgeNative), omniBackend)
	if err != nil {
		return errors.Wrap(err, "new omni bridge native")
	}

	l1bridge, err := bindings.NewOmniBridgeL1(addrs.L1Bridge, l1Backend)
	if err != nil {
		return errors.Wrap(err, "new omni bridge l1")
	}

	//  test rand spec

	randSpec := &NetworkBridgeSpec{
		Native: randBridgeSpec(),
		L1:     randBridgeSpec(),
	}

	if err := ensureBridgeSpec(ctx, def, l1bridge, nativebridge, randSpec); err != nil {
		return err
	}

	// reset to defaults

	defaultSpec := DefaultBridgeSpec()

	return ensureBridgeSpec(ctx, def, l1bridge, nativebridge, &defaultSpec)
}

func ensureBridgeSpec(
	ctx context.Context,
	def app.Definition,
	l1bridge *bindings.OmniBridgeL1,
	nativebridge *bindings.OmniBridgeNative,
	spec *NetworkBridgeSpec,
) error {
	if err := EnsureBridgeSpec(ctx, def, Config{Broadcast: true}, spec); err != nil {
		return errors.Wrap(err, "ensure bridge spec")
	}

	l1Spec, err := liveBridgeSpec(ctx, l1bridge)
	if err != nil {
		return errors.Wrap(err, "live l1 bridge spec")
	}

	nativeSpec, err := liveBridgeSpec(ctx, nativebridge)
	if err != nil {
		return errors.Wrap(err, "live native bridge spec")
	}

	if !cmp.Equal(nativeSpec, spec.Native, cmpopts.EquateEmpty()) {
		return errors.New("live native bridge spec mismatch", "live", nativeSpec, "expected", spec.Native)
	}

	if !cmp.Equal(l1Spec, spec.L1, cmpopts.EquateEmpty()) {
		return errors.New("live l1 bridge spec mismatch", "live", l1Spec, "expected", spec.L1)
	}

	return nil
}

// forOne runs an action & check configured for a single chain (Config{Chain: "name"}).
func forOne(
	ctx context.Context,
	def app.Definition,
	chain types.EVMChain,
	action func(context.Context, app.Definition, Config) error,
	check func(context.Context, app.Definition, types.EVMChain) error,
) error {
	if err := action(ctx, def, Config{Broadcast: true, Chain: chain.Name}); err != nil {
		return errors.Wrap(err, "act", "chain", chain.Name)
	}

	if err := check(ctx, def, chain); err != nil {
		return errors.Wrap(err, "check", "chain", chain.Name)
	}

	return nil
}

// forAll runs an action & check configured for all chains (Config{Chain: ""}).
func forAll(
	ctx context.Context,
	def app.Definition,
	action func(context.Context, app.Definition, Config) error,
	check func(context.Context, app.Definition, types.EVMChain) error,
) error {
	if err := action(ctx, def, Config{Broadcast: true}); err != nil {
		return errors.Wrap(err, "act")
	}

	for _, chain := range def.Testnet.EVMChains() {
		if err := check(ctx, def, chain); err != nil {
			return errors.Wrap(err, "check", "chain", chain.Name)
		}
	}

	return nil
}

func checkPortalSpec(chains []types.EVMChain, expected *PortalSpec) func(context.Context, app.Definition, types.EVMChain) error {
	return func(ctx context.Context, def app.Definition, chain types.EVMChain) error {
		backend, err := def.Backends().Backend(chain.ChainID)
		if err != nil {
			return errors.Wrap(err, "backend", "chain", chain.Name)
		}

		addrs, err := contracts.GetAddresses(ctx, def.Testnet.Network)
		if err != nil {
			return errors.Wrap(err, "get addrs")
		}

		live, err := livePortalSpec(ctx, chains, chain, addrs.Portal, backend)
		if err != nil {
			return errors.Wrap(err, "live portal spec", "chain", chain.Name)
		}

		// sort chain IDs
		if len(live.PauseXCallTo) != len(expected.PauseXCallTo) {
			return errors.New("live portal spec mismatch", "chain", chain.Name, "live", live, "expected", *expected)
		}

		if len(live.PauseXSubmitFrom) != len(expected.PauseXSubmitFrom) {
			return errors.New("live portal spec mismatch", "chain", chain.Name, "live", live, "expected", *expected)
		}

		// sort chain IDs, for comparison
		sortUint64(live.PauseXCallTo)
		sortUint64(live.PauseXSubmitFrom)
		sortUint64(expected.PauseXCallTo)
		sortUint64(expected.PauseXSubmitFrom)

		if !cmp.Equal(live, *expected, cmpopts.EquateEmpty()) {
			return errors.New("live portal spec mismatch", "chain", chain.Name, "live", live, "expected", *expected)
		}

		return nil
	}
}

func randPortalSpec(chains []types.EVMChain) *PortalSpec {
	pauseAll := randBool()
	if pauseAll {
		return &PortalSpec{PauseAll: true}
	}

	spec := &PortalSpec{
		PauseXCall:   randBool(),
		PauseXSubmit: randBool(),
	}

	if !spec.PauseXCall {
		spec.PauseXCallTo = randChainIDs(chains)
	}

	if !spec.PauseXSubmit {
		spec.PauseXSubmitFrom = randChainIDs(chains)
	}

	return spec
}

func randBridgeSpec() BridgeSpec {
	pauseAll := randBool()
	if pauseAll {
		return BridgeSpec{PauseAll: true}
	}

	return BridgeSpec{
		PauseWithdraw: randBool(),
		PauseBridge:   randBool(),
	}
}

func sortUint64(ns []uint64) {
	slices.Sort(ns)
}

func randChain(chains []types.EVMChain) types.EVMChain {
	return chains[rand.Intn(len(chains))]
}

func randChains(chains []types.EVMChain) []types.EVMChain {
	n := rand.Intn(len(chains))
	if n == 0 {
		return nil
	}

	rand.Shuffle(len(chains), func(i, j int) {
		chains[i], chains[j] = chains[j], chains[i]
	})

	return chains[:n]
}

func randChainIDs(chains []types.EVMChain) []uint64 {
	chains = randChains(chains)

	chainIDs := make([]uint64, len(chains))
	for i, chain := range chains {
		chainIDs[i] = chain.ChainID
	}

	return chainIDs
}

func randBool() bool {
	return rand.Intn(2) == 0
}
