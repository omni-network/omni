package genutil

import (
	"context"
	"encoding/json"
	"os"
	"sort"
	"time"

	haloapp "github.com/omni-network/omni/halo/app"
	magellan2 "github.com/omni-network/omni/halo/app/upgrades/magellan"
	uluwatu1 "github.com/omni-network/omni/halo/app/upgrades/uluwatu"
	"github.com/omni-network/omni/halo/evmupgrade"
	vtypes "github.com/omni-network/omni/halo/valsync/types"
	"github.com/omni-network/omni/lib/buildinfo"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/netconf"
	evmtypes "github.com/omni-network/omni/octane/evmengine/types"

	"github.com/cometbft/cometbft/crypto"

	"github.com/ethereum/go-ethereum/common"

	"cosmossdk.io/math"
	etypes "cosmossdk.io/x/evidence/types"
	utypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	atypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	btypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	dtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	gtypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	sltypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	sttypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/gogoproto/proto"
)

// slashingWindows overrides the default slashing signed_blocks_window from 100 to 1000
// since Omni block period (+-1s) is very fast, roughly 10x normal period of 10s.
const slashingBlocksWindow = 1000

// ValidatorPower is the default power assigned to ephemeral genesis validators.
const ValidatorPower = 1000

func MakeGenesis(
	ctx context.Context,
	network netconf.ID,
	genesisTime time.Time,
	executionBlockHash common.Hash,
	upgradeName string,
	valPubkeys ...crypto.PubKey,
) (*gtypes.AppGenesis, error) {
	if upgradeName != "" && upgradeName != uluwatu1.UpgradeName && upgradeName != magellan2.UpgradeName {
		// TODO(corver): Add support for chain genesis (block 0) directly from subsequent upgrades,
		//  ie, no actual network upgrade required since block 0 is already upgraded.
		return nil, errors.New("unsupported genesis network upgrade", "upgrade", upgradeName)
	}

	encConf, err := haloapp.ClientEncodingConfig(ctx, network)
	if err != nil {
		return nil, errors.Wrap(err, "marshal app state")
	}

	// Step 1: Create the default genesis app state for all modules.
	appState1 := defaultAppState(network.Static().MaxValidators, executionBlockHash, encConf.Codec.MustMarshalJSON)
	appState1Bz, err := json.MarshalIndent(appState1, "", " ")
	if err != nil {
		return nil, errors.Wrap(err, "marshal app state")
	}

	// Step 2: Create the app genesis object and store it to disk.
	appGen := &gtypes.AppGenesis{
		AppName:       "halo",
		AppVersion:    buildinfo.Version(),
		GenesisTime:   genesisTime.UTC(),
		ChainID:       network.Static().OmniConsensusChainIDStr(),
		InitialHeight: 1,
		Consensus:     defaultConsensusGenesis(),
		AppState:      appState1Bz,
	}

	// Use this temp file as "disk cache", since the genutil functions require a file path
	tempFile, err := os.CreateTemp("", "")
	if err != nil {
		return nil, errors.Wrap(err, "create temp file")
	}
	if err := genutil.ExportGenesisFile(appGen, tempFile.Name()); err != nil {
		return nil, errors.Wrap(err, "export genesis file")
	}

	// Step 3: Create the genesis validators; genesis account and a MsgCreateValidator.
	var genTxs []sdk.Tx
	for _, pubkey := range sortByAddress(valPubkeys) {
		tx, err := addValidator(encConf.TxConfig, pubkey, encConf.Codec, tempFile.Name())
		if err != nil {
			return nil, errors.Wrap(err, "add validator")
		}
		genTxs = append(genTxs, tx)
	}

	if upgradeName != "" {
		tx, err := addUpgrade(encConf.TxConfig, upgradeName)
		if err != nil {
			return nil, errors.Wrap(err, "add upgrade")
		}
		genTxs = append(genTxs, tx)
	}

	// Step 4: Collect the MsgCreateValidator txs and update the app state (again).
	appState2, err := collectGenTxs(encConf.Codec, encConf.TxConfig, tempFile.Name(), genTxs)
	if err != nil {
		return nil, errors.Wrap(err, "collect genesis transactions")
	}
	appGen.AppState, err = json.MarshalIndent(appState2, "", " ")
	if err != nil {
		return nil, errors.Wrap(err, "marshal app state")
	}

	// Step 5: Validate
	if err := appGen.ValidateAndComplete(); err != nil {
		return nil, errors.Wrap(err, "validate and complete genesis")
	}

	return appGen, validateGenesis(encConf.Codec, appState2)
}

func defaultConsensusGenesis() *gtypes.ConsensusGenesis {
	pb := DefaultConsensusParams().ToProto()
	resp := gtypes.NewConsensusGenesis(pb, nil)
	// NewConsensusGenesis has a bug, it doesn't set VoteExtensionsEnableHeight
	resp.Params.ABCI.VoteExtensionsEnableHeight = pb.Abci.VoteExtensionsEnableHeight

	return resp
}

func validateGenesis(cdc codec.Codec, appState map[string]json.RawMessage) error {
	// Staking module
	ststate := sttypes.GetGenesisStateFromAppState(cdc, appState)
	if err := staking.ValidateGenesis(ststate); err != nil {
		return errors.Wrap(err, "validate staking genesis")
	}

	// Slashing module
	var slstate sltypes.GenesisState
	if err := cdc.UnmarshalJSON(appState[sltypes.ModuleName], &slstate); err != nil {
		return errors.Wrap(err, "unmarshal slashing genesis")
	}
	if err := sltypes.ValidateGenesis(slstate); err != nil {
		return errors.Wrap(err, "validate slashing genesis")
	}

	// Bank module
	bstate := btypes.GetGenesisStateFromAppState(cdc, appState)
	if err := bstate.Validate(); err != nil {
		return errors.Wrap(err, "validate bank genesis")
	}

	// Distribution module
	dstate := new(dtypes.GenesisState)
	if err := cdc.UnmarshalJSON(appState[dtypes.ModuleName], dstate); err != nil {
		return errors.Wrap(err, "unmarshal distribution genesis")
	}
	if err := dtypes.ValidateGenesis(dstate); err != nil {
		return errors.Wrap(err, "validate distribution genesis")
	}

	// Auth module
	astate := atypes.GetGenesisStateFromAppState(cdc, appState)
	if err := atypes.ValidateGenesis(astate); err != nil {
		return errors.Wrap(err, "validate auth genesis")
	}

	// Evidence module
	estate := new(etypes.GenesisState)
	if err := cdc.UnmarshalJSON(appState[etypes.ModuleName], estate); err != nil {
		return errors.Wrap(err, "unmarshal evidence genesis")
	} else if err := estate.Validate(); err != nil {
		return errors.Wrap(err, "validate evidence genesis")
	}

	return nil
}

func collectGenTxs(cdc codec.Codec, txConfig client.TxConfig, genFile string, genTXs []sdk.Tx,
) (map[string]json.RawMessage, error) {
	appState, _, err := gtypes.GenesisStateFromGenFile(genFile)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal genesis state")
	}

	appState, err = genutil.SetGenTxsInAppGenesisState(cdc, txConfig.TxJSONEncoder(), appState, genTXs)
	if err != nil {
		return nil, errors.Wrap(err, "set genesis transactions")
	}

	return appState, nil
}

func addValidator(txConfig client.TxConfig, pubkey crypto.PubKey, cdc codec.Codec, genFile string) (sdk.Tx, error) {
	// We use the validator pubkey as the account address
	addr, err := k1util.PubKeyToAddress(pubkey)
	if err != nil {
		return nil, err
	}

	// Add validator with 1 power (1e18 $STAKE ~= 1 ether $STAKE)
	amount := sdk.NewCoin(sdk.DefaultBondDenom, sdk.DefaultPowerReduction.MulRaw(ValidatorPower))

	err = genutil.AddGenesisAccount(cdc, addr.Bytes(), false, genFile, amount.String(), "", 0, 0, "")
	if err != nil {
		return nil, errors.Wrap(err, "add genesis account")
	}

	pub, err := k1util.PubKeyToCosmos(pubkey)
	if err != nil {
		return nil, err
	}

	var zero = math.LegacyZeroDec()

	msg, err := sttypes.NewMsgCreateValidator(
		sdk.ValAddress(addr.Bytes()).String(),
		pub,
		amount,
		sttypes.Description{Moniker: addr.Hex()},
		sttypes.NewCommissionRates(zero, zero, zero),
		sdk.DefaultPowerReduction,
	)
	if err != nil {
		return nil, errors.Wrap(err, "create validator message")
	}

	builder := txConfig.NewTxBuilder()

	if err := builder.SetMsgs(msg); err != nil {
		return nil, errors.Wrap(err, "set message")
	}

	return builder.GetTx(), nil
}

func addUpgrade(txConfig client.TxConfig, name string) (sdk.Tx, error) {
	msg := &utypes.MsgSoftwareUpgrade{
		Authority: sdk.AccAddress(address.Module(evmupgrade.ModuleName)).String(),
		Plan: utypes.Plan{
			Name:   name,
			Height: 1,
			Info:   "genesis upgrade",
		},
	}

	builder := txConfig.NewTxBuilder()

	if err := builder.SetMsgs(msg); err != nil {
		return nil, errors.Wrap(err, "set message")
	}

	return builder.GetTx(), nil
}

// defaultAppState returns the default genesis application state.
func defaultAppState(
	maxVals uint32,
	executionBlockHash common.Hash,
	marshal func(proto.Message) []byte,
) map[string]json.RawMessage {
	stakingGenesis := sttypes.DefaultGenesisState()
	stakingGenesis.Params.MaxValidators = maxVals

	slashingGenesis := sltypes.DefaultGenesisState()
	slashingGenesis.Params.SignedBlocksWindow = slashingBlocksWindow

	evmengGenesis := evmtypes.NewGenesisState(executionBlockHash)

	return map[string]json.RawMessage{
		sttypes.ModuleName:  marshal(stakingGenesis),
		sltypes.ModuleName:  marshal(slashingGenesis),
		atypes.ModuleName:   marshal(atypes.DefaultGenesisState()),
		btypes.ModuleName:   marshal(btypes.DefaultGenesisState()),
		dtypes.ModuleName:   marshal(dtypes.DefaultGenesisState()),
		etypes.ModuleName:   marshal(etypes.DefaultGenesisState()),
		vtypes.ModuleName:   marshal(vtypes.DefaultGenesisState()),
		gtypes.ModuleName:   marshal(gtypes.DefaultGenesisState()),
		evmtypes.ModuleName: marshal(evmengGenesis),
		utypes.ModuleName:   []byte("{}"), // See cosmossdk.io/x/upgrade@v0.1.4/module.go#DefaultGenesis
	}
}

func sortByAddress(pubkeys []crypto.PubKey) []crypto.PubKey {
	sort.Slice(pubkeys, func(i, j int) bool {
		return pubkeys[i].Address().String() < pubkeys[j].Address().String()
	})

	return pubkeys
}
