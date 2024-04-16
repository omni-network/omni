package monitor

import (
	"context"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

// startAVSSync starts a forever-loop that calls `OmniAVS.SyncWithOmni` once per day.
// This results in a xmsg from the AVS contract to the OmniRestaking contract with the latest Eigen delegations.
func startAVSSync(ctx context.Context, cfg Config, network netconf.Network) error {
	privateKey, err := ethcrypto.LoadECDSA(cfg.PrivateKey)
	if err != nil {
		return errors.Wrap(err, "load private key")
	}

	ethL1, ok := network.EthereumChain()
	if !ok {
		log.Warn(ctx, "Not syncing avs since no ethereum chain defined", nil)
		return nil
	} else if ethL1.AVSContractAddr == (common.Address{}) {
		log.Warn(ctx, "Not syncing avs since AVSContractAddr empty", nil)
		return nil
	} else if ethL1.PortalAddress == (common.Address{}) {
		log.Warn(ctx, "Not syncing avs since no l1 portal defined", nil)
		return nil
	} else if omniEVM, ok := network.OmniEVMChain(); !ok || omniEVM.PortalAddress == (common.Address{}) {
		log.Warn(ctx, "Not syncing avs since no omni evm portal defined", nil)
		return nil
	}

	ethCl, err := ethclient.Dial(ethL1.Name, ethL1.RPCURL)
	if err != nil {
		return errors.Wrap(err, "dial ethereum chain")
	}

	backend, err := ethbackend.NewBackend(ethL1.Name, ethL1.ID, ethL1.BlockPeriod, ethCl)
	if err != nil {
		return err
	}

	from, err := backend.AddAccount(privateKey)
	if err != nil {
		return err
	}

	omniAVS, err := bindings.NewOmniAVS(ethL1.AVSContractAddr, backend)
	if err != nil {
		return errors.Wrap(err, "new omni portal")
	}

	go syncAVSForever(ctx, omniAVS, backend, from)

	return nil
}

func syncAVSForever(ctx context.Context, omniAVS *bindings.OmniAVS, backend *ethbackend.Backend, from common.Address) {
	timer := time.NewTimer(time.Hour)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
		case <-timer.C:
			err := syncAVSOnce(ctx, omniAVS, backend, from)
			if err != nil {
				log.Warn(ctx, "Syncing avs failed (will retry)", err)
				timer.Reset(time.Hour)
			} else {
				timer.Reset(time.Hour * 24)
			}
		}
	}
}

func syncAVSOnce(ctx context.Context, omniAVS *bindings.OmniAVS, backend *ethbackend.Backend, from common.Address) error {
	txOpts, err := backend.BindOpts(ctx, from)
	if err != nil {
		return err
	}

	tx, err := omniAVS.SyncWithOmni(txOpts)
	if err != nil {
		return errors.Wrap(err, "sync with omni")
	}

	if _, err := backend.WaitMined(ctx, tx); err != nil {
		return errors.Wrap(err, "wait mined")
	}

	log.Info(ctx, "Successfully synced AVS with omni")

	return nil
}
