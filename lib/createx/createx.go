package createx

import (
	"github.com/omni-network/omni/lib/cast"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	TrueFlag        = "True"
	FalseFlag       = "False"
	UnspecifiedFlag = "Unspecified"

	MsgSender   = "MsgSender"
	ZeroAddress = "ZeroAddress"
	Random      = "Random"
)

var (
	// CreateXAddress is the canonical CreateX factory address.
	CreateXAddress = common.HexToAddress("0xba5Ed099633D3B313e4D5F7bdc1305d3c28ba5Ed")
)

// Create2Address returns the CREATE2 address for a contract deployed via CreateX factory.
func Create2Address(salt string, initCodeHash common.Hash, deployer common.Address) common.Address {
	saltBytes := []byte(salt)
	var salt32 [32]byte

	// If salt is exactly 32 bytes, use it directly
	if len(saltBytes) == 32 {
		copy(salt32[:], saltBytes)
	} else {
		// For string salts, hash them to get a 32-byte salt
		hashedSalt := crypto.Keccak256Hash(saltBytes)
		copy(salt32[:], hashedSalt[:])
	}

	// Apply CreateX guarding logic
	guardedSalt := GuardSalt(salt32, deployer)
	if guardedSalt == [32]byte{} {
		return common.Address{}
	}

	// Compute CREATE2 address
	return crypto.CreateAddress2(CreateXAddress, guardedSalt, initCodeHash[:])
}

// GuardSalt applies CreateX's salt guarding logic.
func GuardSalt(salt [32]byte, deployer common.Address) [32]byte {
	// Parse the salt to determine its format
	senderBytes, redeployFlag := ParseSalt(salt, deployer)

	switch {
	case redeployFlag == TrueFlag:
		// Redeploy protection is not supported
		return [32]byte{}

	case senderBytes == MsgSender && redeployFlag == UnspecifiedFlag:
		// Unspecified flag for MsgSender
		return [32]byte{}

	case senderBytes == ZeroAddress && redeployFlag == UnspecifiedFlag:
		// Unspecified flag for ZeroAddress
		return [32]byte{}

	case senderBytes == MsgSender:
		// Configures permissioned deploy protection
		return EfficientHash(common.LeftPadBytes(deployer.Bytes(), 32), salt[:])

	default:
		// For random cases (any flag) or ZeroAddress with False flag, hash the salt
		return crypto.Keccak256Hash(salt[:])
	}
}

// ParseSalt parses the salt format according to CreateX's logic.
func ParseSalt(salt [32]byte, deployer common.Address) (string, string) {
	saltAddr := cast.MustEthAddress(salt[:20])
	flag := salt[20]

	if saltAddr == deployer {
		if flag == 0x01 {
			return MsgSender, TrueFlag
		} else if flag == 0x00 {
			return MsgSender, FalseFlag
		}

		return MsgSender, UnspecifiedFlag
	}

	if saltAddr == (common.Address{}) {
		if flag == 0x01 {
			return ZeroAddress, TrueFlag
		} else if flag == 0x00 {
			return ZeroAddress, FalseFlag
		}

		return ZeroAddress, UnspecifiedFlag
	}

	if flag == 0x01 {
		return Random, TrueFlag
	} else if flag == 0x00 {
		return Random, FalseFlag
	}

	return Random, UnspecifiedFlag
}

// EfficientHash mimics CreateX's _efficientHash function.
func EfficientHash(a, b []byte) [32]byte {
	var hash [32]byte
	data := make([]byte, 0, len(a)+len(b))
	data = append(data, a...)
	data = append(data, b...)
	copy(hash[:], crypto.Keccak256(data))

	return hash
}
