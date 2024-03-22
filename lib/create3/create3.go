package create3

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// HashSalt returns the [32]byte hash of the salt string.
func HashSalt(s string) [32]byte {
	var h [32]byte
	copy(h[:], crypto.Keccak256([]byte(s)))

	return h
}

// Omni's Create3 factory namespaces salts by deployer.
func namespacedSalt(deployer common.Address, salt string) [32]byte {
	var h [32]byte
	copy(h[:], crypto.Keccak256(deployer.Bytes(), crypto.Keccak256([]byte(salt))))

	return h
}

//nolint:gochecknoglobals // static hash
var (
	// proxyBytecodeHash is the hash bytecode of the intermediate proxy contract the Create3 factory deploys.
	proxyBytecodeHash = crypto.Keccak256(hexutil.MustDecode("0x67363d3d37363d34f03d5260086018f3"))
)

// Address returns the Create3 address for the given factory, salt, and deployer.
func Address(
	factory common.Address,
	salt string,
	deployer common.Address,
) common.Address {
	// First, get the CREATE2 intermediate proxy address
	proxyAddr := crypto.CreateAddress2(factory, namespacedSalt(deployer, salt), proxyBytecodeHash)

	// Return the CREATE address the proxy deploys to
	return crypto.CreateAddress(proxyAddr, 1)
}
