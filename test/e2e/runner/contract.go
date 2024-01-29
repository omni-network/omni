package main

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	// privKeyHex0 of pre-funded anvil account 0.
	privKeyHex0 = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	// privKeyHex1 of pre-funded anvil account 1.
	privKeyHex1 = "0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
)

type Portal struct {
	Chain    netconf.Chain
	Client   *ethclient.Client
	Contract *bindings.OmniPortal
}

func DeployContracts(ctx context.Context, network netconf.Network) (map[uint64]Portal, error) {
	log.Info(ctx, "Deploying portal contracts")

	resp := make(map[uint64]Portal)
	for _, chain := range network.Chains {
		ethClient, err := ethclient.Dial(chain.RPCURL)
		if err != nil {
			return nil, errors.Wrap(err, "dial chain")
		}

		txOpts, err := newTxOpts(ctx, privKeyHex0, chain.ID)
		if err != nil {
			return nil, err
		}

		addr, _, _, err := bindings.DeployOmniPortal(txOpts, ethClient)
		if err != nil {
			return nil, errors.Wrap(err, "deploy portal")
		} else if addr.Hex() != chain.PortalAddress {
			return nil, errors.New("portal address mismatch",
				"chain", chain.Name,
				"expect", chain.PortalAddress,
				"actual", addr.Hex(),
			)
		}

		contract, err := bindings.NewOmniPortal(addr, ethClient)
		if err != nil {
			return nil, errors.Wrap(err, "create portal contract")
		}

		resp[chain.ID] = Portal{
			Chain:    chain,
			Client:   ethClient,
			Contract: contract,
		}
	}

	return resp, nil
}

func newTxOpts(ctx context.Context, privKeyHex string, chainID uint64) (*bind.TransactOpts, error) {
	pk, err := crypto.HexToECDSA(strings.TrimPrefix(privKeyHex, "0x"))
	if err != nil {
		return nil, errors.Wrap(err, "parse private key")
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(
		pk,
		big.NewInt(int64(chainID)),
	)
	if err != nil {
		return nil, errors.Wrap(err, "keyed tx ops")
	}

	txOpts.Context = ctx

	return txOpts, nil
}

func StartSendingXMsgs(ctx context.Context, portals map[uint64]Portal) error {
	log.Info(ctx, "Generating cross chain messages async")
	go func() {
		for ctx.Err() == nil {
			err := SendXMsgs(ctx, portals)
			if ctx.Err() == nil {
				return
			} else if err != nil {
				log.Error(ctx, "Failed to send xmsgs, giving up", err)
				return
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return nil
}

// SendXMsgs sends one xmsg from every chain to every other chain.
func SendXMsgs(ctx context.Context, portals map[uint64]Portal) error {
	for _, from := range portals {
		for _, to := range portals {
			if from.Chain.ID == to.Chain.ID {
				continue
			}

			if err := xcall(ctx, from, to.Chain.ID); err != nil {
				return err
			}
		}
	}

	return nil
}

// xcall sends a ethereum transaction to the portal contract, triggering a xcall.
func xcall(ctx context.Context, from Portal, destChainID uint64) error {
	txOpts, err := newTxOpts(ctx, privKeyHex0, from.Chain.ID)
	if err != nil {
		return err
	}

	_, err = from.Contract.Xcall(txOpts, destChainID, common.Address{}, nil)
	if err != nil {
		return errors.Wrap(err, "xcall",
			"sourc_chain", from.Chain.ID,
			"dest_chain", destChainID,
		)
	}

	return nil
}
