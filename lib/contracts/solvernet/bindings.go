package solvernet

import (
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var (
	bindingsABI          = mustGetABI(bindings.ISolverNetBindingsMetaData)
	inputsOrderData      = mustGetInputs(bindingsABI, "orderData")
	inputsFillOriginData = mustGetInputs(bindingsABI, "fillOriginData")
)

func mustGetABI(metadata *bind.MetaData) *abi.ABI {
	abi, err := metadata.GetAbi()
	if err != nil {
		panic(err)
	}

	return abi
}

func mustGetInputs(abi *abi.ABI, name string) abi.Arguments {
	method, ok := abi.Methods[name]
	if !ok {
		panic("method not found")
	}

	return method.Inputs
}

func ParseFillOriginData(data []byte) (bindings.SolverNetFillOriginData, error) {
	unpacked, err := inputsFillOriginData.Unpack(data)
	if err != nil {
		return bindings.SolverNetFillOriginData{}, errors.Wrap(err, "unpack fill data")
	}

	wrap := struct {
		Data bindings.SolverNetFillOriginData
	}{}
	if err := inputsFillOriginData.Copy(&wrap, unpacked); err != nil {
		return bindings.SolverNetFillOriginData{}, errors.Wrap(err, "copy fill data")
	}

	return wrap.Data, nil
}

func ParseOrderData(data []byte) (bindings.SolverNetOrderData, error) {
	unpacked, err := inputsOrderData.Unpack(data)
	if err != nil {
		return bindings.SolverNetOrderData{}, errors.Wrap(err, "unpack fill data")
	}

	wrap := struct {
		Data bindings.SolverNetOrderData
	}{}
	if err := inputsOrderData.Copy(&wrap, unpacked); err != nil {
		return bindings.SolverNetOrderData{}, errors.Wrap(err, "copy fill data")
	}

	return wrap.Data, nil
}

func PackOrderData(data bindings.SolverNetOrderData) ([]byte, error) {
	packed, err := inputsOrderData.Pack(data)
	if err != nil {
		return nil, errors.Wrap(err, "pack fill data")
	}

	return packed, nil
}

func PackFillOriginData(data bindings.SolverNetFillOriginData) ([]byte, error) {
	packed, err := inputsFillOriginData.Pack(data)
	if err != nil {
		return nil, errors.Wrap(err, "pack fill data")
	}

	return packed, nil
}
