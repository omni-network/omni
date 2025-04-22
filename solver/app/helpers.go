package app

import (
	"bytes"
	"context"

	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// detectContractChains returns the chains on which the contract is deployed at the provided address.
func detectContractChains(ctx context.Context, network netconf.Network, backends ethbackend.Backends, address common.Address) ([]uint64, error) {
	var resp []uint64
	for _, chain := range network.EVMChains() {
		backend, err := backends.Backend(chain.ID)
		if err != nil {
			return nil, err
		}

		code, err := backend.CodeAt(ctx, address, nil)
		if err != nil {
			return nil, errors.Wrap(err, "get code", "chain", chain.Name)
		} else if len(code) == 0 {
			continue
		}

		resp = append(resp, chain.ID)
	}

	return resp, nil
}

// toEthAddr converts a 32-byte address to an Ethereum address.
func toEthAddr(bz [32]byte) (common.Address, error) {
	addr := cast.MustEthAddress(bz[12:])
	if !cmpAddrs(addr, bz) {
		return common.Address{}, errors.New("invalid eth address", "address", hexutil.Encode(bz[:]))
	}

	return addr, nil
}

// cmpAddrs returns true if the eth address is equal to the 32-byte address.
func cmpAddrs(addr common.Address, bz [32]byte) bool {
	addrBz := addr.Bytes()
	var addrBz32 [32]byte
	copy(addrBz32[12:], addrBz)

	return bytes.Equal(addrBz32[:], bz[:])
}

// toBz32 converts an Ethereum address to a 32-byte address.
func toBz32(addr common.Address) [32]byte {
	var bz [32]byte
	copy(bz[12:], addr.Bytes())

	return bz
}

func maybeDebugRevert(ctx context.Context, cl ethclient.Client, from common.Address, tx *ethtypes.Transaction, rec *ethclient.Receipt) (bool, error) {
	if rec == nil || rec.Status == ethtypes.ReceiptStatusSuccessful {
		return false, nil
	}

	// Assume 95%+ gas usage was due to out of gas
	if rec.GasUsed >= tx.Gas()*95/100 {
		return true, errors.New("tx probably reverted due to out of gas",
			"dest_chain", cl.Name(),
			"tx", tx.Hash(),
			"receipt_height", rec.BlockNumber,
			"gas_used", rec.GasUsed,
			"gas_limit", tx.Gas(),
		)
	}

	// Try and get debug information of the reverted transaction
	resp, err := cl.CallContract(ctx, callFromTx(from, tx), rec.BlockNumber)
	if err == nil {
		return false, nil // It didn't revert again
	}

	return true, errors.Wrap(err, "tx reverted",
		"dest_chain", cl.Name(),
		"tx", tx.Hash(),
		"receipt_height", rec.BlockNumber,
		"call_resp", hexutil.Encode(resp),
		"custom", solvernet.DetectCustomError(err),
	)
}

func callFromTx(from common.Address, tx *ethtypes.Transaction) ethereum.CallMsg {
	resp := ethereum.CallMsg{
		From:          from,
		To:            tx.To(),
		Gas:           tx.Gas(),
		Value:         tx.Value(),
		Data:          tx.Data(),
		AccessList:    tx.AccessList(),
		BlobGasFeeCap: tx.BlobGasFeeCap(),
		BlobHashes:    tx.BlobHashes(),
	}

	// Either populate gas price or gas caps (not both).
	if tx.GasPrice() != nil && tx.GasPrice().Sign() != 0 {
		resp.GasPrice = tx.GasPrice()
	} else {
		resp.GasFeeCap = tx.GasFeeCap()
		resp.GasTipCap = tx.GasTipCap()
	}

	return resp
}
