package app

import (
	"os"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"

	"github.com/cometbft/cometbft/crypto"
	cmtjson "github.com/cometbft/cometbft/libs/json"
	"github.com/cometbft/cometbft/privval"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

// loadPrivVal returns a privval.FilePV by loading either a CometBFT priv validator key or an Ethereum keystore file.
func loadPrivVal(cfg Config) (*privval.FilePV, error) {
	cmtFile := cfg.Comet.PrivValidatorKeyFile()
	cmtExists := exists(cmtFile)
	keystoreFile, keystoreExists, err := cfg.KeystoreFile()
	if err != nil {
		return nil, err
	}

	if !cmtExists && !keystoreExists {
		return nil, errors.New("neither a cometBFT priv validator key nor an eigenlayer operator key file exists", "comet_file", cmtFile, "eigen_file", cfg.KeystoreGlob())
	} else if cmtExists && keystoreExists {
		return nil, errors.New("both a cometBFT priv validator key and an eigenlayer operator key file exist", "comet_file", cmtFile, "eigen_file", keystoreFile)
	}

	var key crypto.PrivKey
	if keystoreExists {
		key, err = loadEthKeystore(keystoreFile, cfg.EigenKeyPassword)
	} else {
		key, err = loadCometFilePV(cmtFile)
	}
	if err != nil {
		return nil, err
	}

	state, err := loadCometPVState(cfg.Comet.PrivValidatorStateFile())
	if err != nil {
		return nil, err
	}

	// Create a new privval.FilePV with the loaded key and state.
	// This is a workaround for the fact that there is no other way
	// to set FilePVLastSignState filePath field.
	resp := privval.NewFilePV(key, "", cfg.Comet.PrivValidatorStateFile())
	resp.LastSignState.Step = state.Step
	resp.LastSignState.Round = state.Round
	resp.LastSignState.Height = state.Height
	resp.LastSignState.Signature = state.Signature
	resp.LastSignState.SignBytes = state.SignBytes

	return resp, nil
}

// loadEthKeystore loads an Ethereum keystore file and returns the private key.
func loadEthKeystore(keystoreFile string, password string) (crypto.PrivKey, error) {
	bz, err := os.ReadFile(keystoreFile)
	if err != nil {
		return nil, errors.Wrap(err, "read keystore file", "path", keystoreFile)
	}

	key, err := keystore.DecryptKey(bz, password)
	if err != nil {
		return nil, errors.Wrap(err, "decrypt keystore file", "path", keystoreFile)
	}

	return k1util.StdPrivKeyToComet(key.PrivateKey)
}

// loadCometFilePV loads a CometBFT privval file and returns the private key.
func loadCometFilePV(file string) (crypto.PrivKey, error) {
	bz, err := os.ReadFile(file)
	if err != nil {
		return nil, errors.Wrap(err, "read comet privval", "path", file)
	}

	var pvKey privval.FilePVKey
	err = cmtjson.Unmarshal(bz, &pvKey)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal comet privval", "path", file)
	}

	return pvKey.PrivKey, nil
}

// loadCometPVState loads a CometBFT privval state file.
func loadCometPVState(file string) (privval.FilePVLastSignState, error) {
	bz, err := os.ReadFile(file)
	if err != nil {
		return privval.FilePVLastSignState{}, errors.Wrap(err, "read comet privval state", "path", file)
	}

	var state privval.FilePVLastSignState
	err = cmtjson.Unmarshal(bz, &state)
	if err != nil {
		return privval.FilePVLastSignState{}, errors.Wrap(err, "unmarshal comet privval state", "path", file)
	}

	return state, nil
}

func exists(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}
