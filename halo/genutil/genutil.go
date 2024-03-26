package genutil

import (
	"encoding/json"
	"os"
	"time"

	attesttypes "github.com/omni-network/omni/halo/attest/types"
	etypes "github.com/omni-network/omni/halo/evmengine/types"
	vtypes "github.com/omni-network/omni/halo/valsync/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/cometbft/cometbft/crypto"

	"cosmossdk.io/math"
	"cosmossdk.io/x/tx/signing"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmosstd "github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authcodec "github.com/cosmos/cosmos-sdk/x/auth/codec"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	atypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	btypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	dtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	gtypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/gogoproto/proto"
)

func MakeGenesis(network netconf.ID, genesisTime time.Time, valPubkeys ...crypto.PubKey) (*gtypes.AppGenesis, error) {
	cdc := getCodec()
	txConfig := authtx.NewTxConfig(cdc, nil)

	// Step 1: Create the default genesis app state for all modules.
	appStateBz, err := json.MarshalIndent(defaultAppState(cdc.MustMarshalJSON), "", " ")
	if err != nil {
		return nil, errors.Wrap(err, "marshal app state")
	}

	// Step 2: Create the app genesis object and store it to disk.

	appGen := &gtypes.AppGenesis{
		AppName:       "halo",
		AppVersion:    "TBD",
		GenesisTime:   genesisTime.UTC(),
		ChainID:       network.Static().OmniConsensusChainIDStr(),
		InitialHeight: 1,
		Consensus:     defaultConsensusGenesis(),
		AppState:      appStateBz,
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
	valTxs := make([]sdk.Tx, 0, len(valPubkeys))
	for _, pubkey := range valPubkeys {
		tx, err := addValidator(txConfig, pubkey, cdc, tempFile.Name())
		if err != nil {
			return nil, errors.Wrap(err, "add validator")
		}
		valTxs = append(valTxs, tx)
	}

	// Step 4: Collect the MsgCreateValidator txs and update the app state (again).
	appState, err := collectGenTxs(cdc, txConfig, tempFile.Name(), valTxs)
	if err != nil {
		return nil, errors.Wrap(err, "collect genesis transactions")
	}

	appStateBz, err = json.MarshalIndent(appState, "", " ")
	if err != nil {
		return nil, errors.Wrap(err, "marshal app state")
	}

	appGen.AppState = appStateBz

	// Step 5: Validate
	if err := appGen.ValidateAndComplete(); err != nil {
		return nil, errors.Wrap(err, "validate and complete genesis")
	}

	return appGen, validateGenesis(cdc, appState)
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
	sstate := stypes.GetGenesisStateFromAppState(cdc, appState)
	if err := staking.ValidateGenesis(sstate); err != nil {
		return errors.Wrap(err, "validate staking genesis")
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
	amount := sdk.NewCoin(sdk.DefaultBondDenom, sdk.DefaultPowerReduction)

	err = genutil.AddGenesisAccount(cdc, addr.Bytes(), false, genFile, amount.String(), "", 0, 0, "")
	if err != nil {
		return nil, errors.Wrap(err, "add genesis account")
	}

	pub, err := k1util.PubKeyToCosmos(pubkey)
	if err != nil {
		return nil, err
	}

	var zero = math.LegacyZeroDec()

	msg, err := stypes.NewMsgCreateValidator(
		sdk.ValAddress(addr.Bytes()).String(),
		pub,
		amount,
		stypes.Description{Moniker: addr.Hex()},
		stypes.NewCommissionRates(zero, zero, zero),
		sdk.DefaultPowerReduction)
	if err != nil {
		return nil, errors.Wrap(err, "create validator message")
	}

	builder := txConfig.NewTxBuilder()

	if err := builder.SetMsgs(msg); err != nil {
		return nil, errors.Wrap(err, "set message")
	}

	return builder.GetTx(), nil
}

// defaultAppState returns the default genesis application state.
func defaultAppState(marshal func(proto.Message) []byte) map[string]json.RawMessage {
	return map[string]json.RawMessage{
		stypes.ModuleName: marshal(stypes.DefaultGenesisState()),
		atypes.ModuleName: marshal(atypes.DefaultGenesisState()),
		btypes.ModuleName: marshal(btypes.DefaultGenesisState()),
		dtypes.ModuleName: marshal(dtypes.DefaultGenesisState()),
		vtypes.ModuleName: marshal(vtypes.DefaultGenesisState()),
	}
}

func getCodec() *codec.ProtoCodec {
	// TODO(corver): Use depinject to get all of this.
	sdkConfig := sdk.GetConfig()
	reg, err := codectypes.NewInterfaceRegistryWithOptions(codectypes.InterfaceRegistryOptions{
		ProtoFiles: proto.HybridResolver,
		SigningOptions: signing.Options{
			AddressCodec:          authcodec.NewBech32Codec(sdkConfig.GetBech32AccountAddrPrefix()),
			ValidatorAddressCodec: authcodec.NewBech32Codec(sdkConfig.GetBech32ValidatorAddrPrefix()),
		},
	})
	if err != nil {
		panic(err)
	}

	cosmosstd.RegisterInterfaces(reg)
	atypes.RegisterInterfaces(reg)
	stypes.RegisterInterfaces(reg)
	btypes.RegisterInterfaces(reg)
	dtypes.RegisterInterfaces(reg)
	etypes.RegisterInterfaces(reg)
	attesttypes.RegisterInterfaces(reg)

	return codec.NewProtoCodec(reg)
}
