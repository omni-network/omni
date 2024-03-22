package contracts

import (
	"bytes"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// PackInitCode packs the init code for a contract deployment.
func PackInitCode(abi *abi.ABI, bytecodeHex string, params ...any) ([]byte, error) {
	input, err := abi.Pack("", params...)
	if err != nil {
		return nil, errors.Wrap(err, "pack init code")
	}

	bytecode, err := hexutil.Decode(bytecodeHex)
	if err != nil {
		return nil, errors.Wrap(err, "decode bytecode")
	}

	return append(bytecode, input...), nil
}

//nolint:gochecknoglobals // static hash
var (
	// proxyBytecodeHash is the hash bytecode of the intermediate proxy
	// contract the Create3 factory deploys.
	proxyBytecodeHash = crypto.Keccak256(hexutil.MustDecode("0x67363d3d37363d34f03d5260086018f3"))
)

// Create3Address returns the Create3 address for the given factory, salt, and deployer.
func Create3Address(
	factory common.Address,
	salt string,
	deployer common.Address,
) common.Address {
	// Omni's Create3 factory namespaces salt with the deployer address.
	namespacedSalt := crypto.Keccak256(deployer.Bytes(), crypto.Keccak256([]byte(salt)))

	// First, get the CREATE2 intermediate proxy address
	proxyAddress := common.BytesToAddress(
		crypto.Keccak256(
			bytes.Join(
				[][]byte{
					{0xff},
					factory.Bytes(),
					namespacedSalt,
					proxyBytecodeHash,
				},
				nil,
			),
		),
	)

	// Return the CREATE address the proxy deploys to
	return common.BytesToAddress(
		crypto.Keccak256(
			bytes.Join(
				[][]byte{
					// 0xd6 = 0xc0 (short RLP prefix) + 0x16 (length of: 0x94 ++ proxy ++ 0x01)
					// 0x94 = 0x80 + 0x14 (0x14 = the length of an address, 20 bytes, in hex)
					{0xd6, 0x94},
					proxyAddress.Bytes(),
					// 0x01 = nonce of the proxy
					{0x01},
				},
				nil,
			),
		),
	)
}
