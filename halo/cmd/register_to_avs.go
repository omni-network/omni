package cmd

import (
	"context"
	"math/big"
	"os"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/cometbft/cometbft/privval"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/spf13/cobra"
)

func RegisterOperatorToOmniAVS(cfg *OperatorConfig) *cobra.Command {
	registerToAVSCmd := &cobra.Command{
		Use:   "register",
		Short: "register validator to omni avs",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg.HaloConfig.HomeDir = cfg.HomeDir

			// read the comet config based on the home directory
			cometCfg, err := parseCometConfig(cmd.Context(), cfg.HaloConfig.HomeDir)
			if err != nil {
				return err
			}
			cfg.CometConfig = cometCfg

			return register(cmd.Context(), cfg)
		},
	}

	bindOperatorFlags(registerToAVSCmd.Flags(), cfg)

	return registerToAVSCmd
}

func register(ctx context.Context, cfg *OperatorConfig) error {
	privVal, client, chain, err := loadKeysAndChain(ctx, cfg)
	if err != nil {
		return err
	}

	err = validateContractAddresses(ctx, cfg, client)
	if err != nil {
		return err
	}

	// load contract bindings
	omniAvs, err := bindings.NewOmniAVS(common.HexToAddress(cfg.OmniAVSAddr), client)
	if err != nil {
		return err
	}
	avsDirectoryAddr, err := omniAvs.AvsDirectory(&bind.CallOpts{})
	if err != nil {
		return err
	}
	avsDirectory, err := bindings.NewAVSDirectory(avsDirectoryAddr, client)
	if err != nil {
		return err
	}

	// get the operator private key and operators ethereum address
	operPK, err := crypto.ToECDSA(privVal.Key.PrivKey.Bytes())
	if err != nil {
		return errors.Wrap(err, "could not convert pk to ecdsa")
	}
	pubKey, err := privVal.GetPubKey()
	if err != nil {
		return errors.Wrap(err, "get pubkey")
	}
	operAddr, err := k1util.PubKeyToAddress(pubKey)
	if err != nil {
		return errors.Wrap(err, "could not convert to ethereum address")
	}

	// calculate operator signature and digest hash
	blockNumber, err := client.BlockNumber(ctx)
	if err != nil {
		return errors.Wrap(err, "getting blockNumber ")
	}
	block, err := client.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		return errors.Wrap(err, "getting blockByNumber ")
	}

	operatorSignatureWithSaltAndExpiry := bindings.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: []byte{0},
		Salt:      crypto.Keccak256Hash(operAddr.Bytes()),
		Expiry:    big.NewInt(int64(block.Time()) + int64(24*time.Hour)),
	}
	digestHash, err := avsDirectory.CalculateOperatorAVSRegistrationDigestHash(&bind.CallOpts{},
		common.HexToAddress(privVal.GetAddress().String()),
		common.HexToAddress(cfg.OmniAVSAddr),
		operatorSignatureWithSaltAndExpiry.Salt,
		operatorSignatureWithSaltAndExpiry.Expiry)
	if err != nil {
		return err
	}

	sig, err := k1util.Sign(privVal.Key.PrivKey, [32]byte(digestHash[:32]))
	if err != nil {
		return errors.Wrap(err, "error signing)")
	}

	operatorSignatureWithSaltAndExpiry.Signature = sig[:]
	if len(operatorSignatureWithSaltAndExpiry.Signature) != 65 {
		return errors.New("invalid signature length")
	}
	operatorSignatureWithSaltAndExpiry.Signature[64] += 27
	txOpts, err := bind.NewKeyedTransactorWithChainID(operPK, big.NewInt(int64(chain.ID)))
	if err != nil {
		return errors.Wrap(err, "error getting txopts")
	}

	txOpts.Context = ctx
	tx, err := omniAvs.RegisterOperatorToAVS(txOpts, operAddr, operatorSignatureWithSaltAndExpiry)
	if err != nil {
		return err
	}
	log.Info(ctx, "Submitted registration to AVS")

	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		return errors.Wrap(err, "error waiting for mining tx")
	}
	log.Info(ctx, "Operator registered with AVS", "address", operAddr.String(), "txHash", receipt.TxHash)

	return nil
}

func loadKeysAndChain(ctx context.Context, cfg *OperatorConfig) (*privval.FilePV, *ethclient.Wrapper, *netconf.Chain, error) {
	// check for home directory where the config files exist
	if !directoryExists(cfg.HaloConfig.HomeDir) {
		log.Info(ctx, "Make sure to run \"init\" command before running operator commands")
		err := errors.New("directory does not exists", "home", cfg.HaloConfig.HomeDir)

		return nil, nil, nil, err
	}

	// load network config for the layer1 c
	c, err := getChainConfig(cfg)
	if err != nil {
		return nil, nil, nil, err
	}

	// load a private validator key and state from disk (this hard exits on any error).
	privVal := privval.LoadFilePVEmptyState(cfg.CometConfig.PrivValidatorKeyFile(), cfg.CometConfig.PrivValidatorStateFile())

	// connect to the rpc endpoint
	client, err := ethclient.Dial(c.Name, c.RPCURL)
	if err != nil {
		return nil, nil, nil, err
	}

	return privVal, &client, c, nil
}

func getChainConfig(cfg *OperatorConfig) (*netconf.Chain, error) {
	network, err := netconf.Load(cfg.HaloConfig.NetworkFile())
	if err != nil {
		return nil, errors.Wrap(err, "load network config")
	} else if err := network.Validate(); err != nil {
		return nil, errors.Wrap(err, "validate network config")
	}

	chain, present := network.EthereumChain()
	if present {
		return &chain, nil
	}

	return nil, errors.New("chain not found")
}

func directoryExists(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}

	return true
}

func validateContractAddresses(ctx context.Context, cfg *OperatorConfig, client *ethclient.Wrapper) error {
	// check if contracts are deployed in respective addresses
	blockNum, err := client.BlockNumber(ctx)
	if err != nil {
		return err
	}
	_, err = client.CodeAt(ctx, common.HexToAddress(cfg.OmniAVSAddr), big.NewInt(int64(blockNum)))
	if err != nil {
		return err
	}

	return nil
}
