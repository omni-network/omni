package createx_test

import (
	"testing"

	"github.com/omni-network/omni/lib/createx"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/stretchr/testify/require"
)

func TestAddress(t *testing.T) {
	t.Parallel()

	// Test data based on CreateX test from Preinstalls.t.sol
	deployer := common.HexToAddress("0x1234567890123456789012345678901234567890")

	// Create a salt in the format: deployer (20 bytes) + flag (1 byte) + random (11 bytes)
	saltRandom := crypto.Keccak256([]byte("SALT"))[:11]
	var saltBytes [32]byte
	copy(saltBytes[:20], deployer.Bytes())
	saltBytes[20] = 0x00 // MsgSender + False flag
	copy(saltBytes[21:], saltRandom)

	saltString := string(saltBytes[:])

	// Create dummy init code hash
	initCodeHash := crypto.Keccak256Hash([]byte("dummy contract bytecode"))

	// Compute address using our function
	addr := createx.Create2Address(saltString, initCodeHash, deployer)

	// Verify it's a valid address (not zero)
	require.NotEqual(t, common.Address{}, addr)

	// Test that different salts produce different addresses
	saltBytes[31] = 0x01 // Change last byte
	saltString2 := string(saltBytes[:])
	addr2 := createx.Create2Address(saltString2, initCodeHash, deployer)
	require.NotEqual(t, addr, addr2)

	// Test with string salt (should be hashed)
	stringSalt := "test-salt-string"
	addr3 := createx.Create2Address(stringSalt, initCodeHash, deployer)
	require.NotEqual(t, common.Address{}, addr3)
	require.NotEqual(t, addr, addr3)
	require.NotEqual(t, addr2, addr3)

	// Test with random salt (should be hashed)
	var randomSalt [32]byte
	randomSalt[0] = 0x99  // Random address
	randomSalt[20] = 0x00 // False flag
	randomSaltString := string(randomSalt[:])
	addr4 := createx.Create2Address(randomSaltString, initCodeHash, deployer)
	require.NotEqual(t, common.Address{}, addr4)
}

func TestGuardSalt(t *testing.T) {
	t.Parallel()

	deployer := common.HexToAddress("0x1234567890123456789012345678901234567890")

	// Test deployer protection (MsgSender + False)
	var salt [32]byte
	copy(salt[:20], deployer.Bytes())
	salt[20] = 0x00 // False flag

	guardedSalt := createx.GuardSalt(salt, deployer)

	// For deployer protection, it should be efficientHash(deployer, salt)
	expected := createx.EfficientHash(common.LeftPadBytes(deployer.Bytes(), 32), salt[:])
	require.Equal(t, expected[:], guardedSalt[:])

	// Test unsupported redeploy flag (MsgSender + True) - should return zero
	var salt2 [32]byte
	copy(salt2[:20], deployer.Bytes())
	salt2[20] = 0x01 // True flag

	guardedSalt2 := createx.GuardSalt(salt2, deployer)
	require.Equal(t, [32]byte{}, guardedSalt2)

	// Test unsupported redeploy flag (ZeroAddress + True) - should return zero
	var salt3 [32]byte
	// Leave first 20 bytes as zero (zero address)
	salt3[20] = 0x01 // True flag

	guardedSalt3 := createx.GuardSalt(salt3, deployer)
	require.Equal(t, [32]byte{}, guardedSalt3)

	// Test MsgSender with unspecified flag - should return zero
	var salt4 [32]byte
	copy(salt4[:20], deployer.Bytes())
	salt4[20] = 0x02 // Unspecified flag

	guardedSalt4 := createx.GuardSalt(salt4, deployer)
	require.Equal(t, [32]byte{}, guardedSalt4)

	// Test ZeroAddress with unspecified flag - should return zero
	var salt5 [32]byte
	// Leave first 20 bytes as zero (zero address)
	salt5[20] = 0x02 // Unspecified flag

	guardedSalt5 := createx.GuardSalt(salt5, deployer)
	require.Equal(t, [32]byte{}, guardedSalt5)

	// Test ZeroAddress with False flag (should hash the salt)
	var salt6 [32]byte
	// Leave first 20 bytes as zero (zero address)
	salt6[20] = 0x00 // False flag

	guardedSalt6 := createx.GuardSalt(salt6, deployer)

	expected6 := crypto.Keccak256Hash(salt6[:])
	require.Equal(t, expected6[:], guardedSalt6[:])

	// Test Random address with False flag (should hash the salt)
	var salt7 [32]byte
	salt7[0] = 0x99  // Random address
	salt7[20] = 0x00 // False flag

	guardedSalt7 := createx.GuardSalt(salt7, deployer)

	expected7 := crypto.Keccak256Hash(salt7[:])
	require.Equal(t, expected7[:], guardedSalt7[:])
}

func TestParseSalt(t *testing.T) {
	t.Parallel()

	deployer := common.HexToAddress("0x1234567890123456789012345678901234567890")

	// Test MsgSender + False
	var salt [32]byte
	copy(salt[:20], deployer.Bytes())
	salt[20] = 0x00

	senderType, redeployFlag := createx.ParseSalt(salt, deployer)
	require.Equal(t, createx.MsgSender, senderType)
	require.Equal(t, createx.FalseFlag, redeployFlag)

	// Test MsgSender + True
	salt[20] = 0x01
	senderType, redeployFlag = createx.ParseSalt(salt, deployer)
	require.Equal(t, createx.MsgSender, senderType)
	require.Equal(t, createx.TrueFlag, redeployFlag)

	// Test ZeroAddress + False
	var salt2 [32]byte
	salt2[20] = 0x00
	senderType, redeployFlag = createx.ParseSalt(salt2, deployer)
	require.Equal(t, createx.ZeroAddress, senderType)
	require.Equal(t, createx.FalseFlag, redeployFlag)

	// Test Random + True
	var salt3 [32]byte
	salt3[0] = 0x99 // Random address
	salt3[20] = 0x01
	senderType, redeployFlag = createx.ParseSalt(salt3, deployer)
	require.Equal(t, createx.Random, senderType)
	require.Equal(t, createx.TrueFlag, redeployFlag)
}

func TestEfficientHash(t *testing.T) {
	t.Parallel()

	a := []byte("hello")
	b := []byte("world")

	result := createx.EfficientHash(a, b)

	// Verify it matches keccak256(a + b)
	expected := crypto.Keccak256Hash(append(a, b...))
	require.Equal(t, expected, common.Hash(result))
}
