package contract

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type BoundFeeOracleV1 struct {
	owner   common.Address        // eoa owner of the FeeOracleV1 dcontract
	addr    common.Address        // address of the FeeOracle oracle contract addrss
	backend *ethbackend.Backend   // ethbackend initialized with owner pk
	bound   *bindings.FeeOracleV1 // bound FeeOracleV1 contract
}

var _ FeeOracleV1 = BoundFeeOracleV1{}

// New creates a new bound FeeOracleV1 contract.
func New(ctx context.Context, chain netconf.Chain, ethCl ethclient.Client, pk *ecdsa.PrivateKey) (BoundFeeOracleV1, error) {
	backend, err := ethbackend.NewBackend(chain.Name, chain.ID, chain.BlockPeriod, ethCl, pk)
	if err != nil {
		return BoundFeeOracleV1{}, errors.Wrap(err, "new backend")
	}

	portal, err := bindings.NewOmniPortal(chain.PortalAddress, ethCl)
	if err != nil {
		return BoundFeeOracleV1{}, errors.Wrap(err, "new omni portal")
	}

	addr, err := portal.FeeOracle(&bind.CallOpts{Context: ctx})
	if err != nil {
		return BoundFeeOracleV1{}, errors.Wrap(err, "fee oracle addr")
	}

	contract, err := bindings.NewFeeOracleV1(addr, backend)
	if err != nil {
		return BoundFeeOracleV1{}, errors.Wrap(err, "new fee oracle")
	}

	owner := crypto.PubkeyToAddress(pk.PublicKey)

	return BoundFeeOracleV1{
		owner:   owner,
		addr:    addr,
		backend: backend,
		bound:   contract,
	}, nil
}

// GasPriceOn returns the gas price on the FeeOracleV1 contract for the destination chain.
func (c BoundFeeOracleV1) GasPriceOn(ctx context.Context, destChainID uint64) (*big.Int, error) {
	return c.bound.GasPriceOn(callOpts(ctx), destChainID)
}

// SetGasPriceOn sets the gas price on the FeeOracleV1 contract for the destination chain.
func (c BoundFeeOracleV1) SetGasPriceOn(ctx context.Context, destChainID uint64, gasPrice *big.Int) error {
	txOpts, err := c.txOptsWithCtx(ctx)
	if err != nil {
		return errors.Wrap(err, "tx opts")
	}

	tx, err := c.bound.SetGasPrice(txOpts, destChainID, gasPrice)
	if err != nil {
		return errors.Wrap(err, "set gas price")
	}

	_, err = c.backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined", "tx", tx.Hash().Hex())
	}

	return nil
}

// SetToNativeRate sets the conversion rate on the FeeOracleV1 contract for the destination chain.
func (c BoundFeeOracleV1) SetToNativeRate(ctx context.Context, destChainID uint64, rate *big.Int) error {
	txOpts, err := c.txOptsWithCtx(ctx)
	if err != nil {
		return errors.Wrap(err, "tx opts")
	}

	tx, err := c.bound.SetToNativeRate(txOpts, destChainID, rate)
	if err != nil {
		return errors.Wrap(err, "set conversion rate")
	}

	_, err = c.backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined", "tx", tx.Hash().Hex())
	}

	return nil
}

// ToNativeRate returns the "to native" conversion rate on the FeeOracleV1 contract for the destination chaink.
func (c BoundFeeOracleV1) ToNativeRate(ctx context.Context, destChainID uint64) (*big.Int, error) {
	return c.bound.ToNativeRate(callOpts(ctx), destChainID)
}

// txOpts returns a new transact opts with the given context.
func (c BoundFeeOracleV1) txOptsWithCtx(ctx context.Context) (*bind.TransactOpts, error) {
	return c.backend.BindOpts(ctx, c.owner)
}

// PostsTo returns the postsTo value on the FeeOracleV1 contract for the destination chain.
func (c BoundFeeOracleV1) PostsTo(ctx context.Context, destChainID uint64) (uint64, error) {
	return c.bound.PostsTo(callOpts(ctx), destChainID)
}

func (c BoundFeeOracleV1) BulkSetFeeParams(ctx context.Context, params []bindings.IFeeOracleV1ChainFeeParams) error {
	txOpts, err := c.txOptsWithCtx(ctx)
	if err != nil {
		return errors.Wrap(err, "tx opts")
	}

	tx, err := c.bound.BulkSetFeeParams(txOpts, params)
	if err != nil {
		return errors.Wrap(err, "bulk set fee params")
	}

	_, err = c.backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined", "tx", tx.Hash().Hex())
	}

	return nil
}

// callOpts returns a new call opts with the given context.
func callOpts(ctx context.Context) *bind.CallOpts {
	return &bind.CallOpts{Context: ctx}
}
