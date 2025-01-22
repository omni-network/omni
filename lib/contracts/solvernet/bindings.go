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

func ParseFillOriginData(data []byte) (bindings.ISolverNetFillOriginData, error) {
	unpacked, err := inputsFillOriginData.Unpack(data)
	if err != nil {
		return bindings.ISolverNetFillOriginData{}, errors.Wrap(err, "unpack fill data")
	}

	wrap := struct {
		Data bindings.ISolverNetFillOriginData
	}{}
	if err := inputsFillOriginData.Copy(&wrap, unpacked); err != nil {
		return bindings.ISolverNetFillOriginData{}, errors.Wrap(err, "copy fill data")
	}

	return wrap.Data, nil
}

func ParseOrderData(data []byte) (bindings.ISolverNetOrderData, error) {
	unpacked, err := inputsOrderData.Unpack(data)
	if err != nil {
		return bindings.ISolverNetOrderData{}, errors.Wrap(err, "unpack fill data")
	}

	wrap := struct {
		Data bindings.ISolverNetOrderData
	}{}
	if err := inputsOrderData.Copy(&wrap, unpacked); err != nil {
		return bindings.ISolverNetOrderData{}, errors.Wrap(err, "copy fill data")
	}

	return wrap.Data, nil
}

func PackOrderData(data bindings.ISolverNetOrderData) ([]byte, error) {
	packed, err := inputsOrderData.Pack(data)
	if err != nil {
		return nil, errors.Wrap(err, "pack fill data")
	}

	return packed, nil
}

func PackFillOriginData(data bindings.ISolverNetFillOriginData) ([]byte, error) {
	packed, err := inputsFillOriginData.Pack(data)
	if err != nil {
		return nil, errors.Wrap(err, "pack fill data")
	}

	return packed, nil
}
