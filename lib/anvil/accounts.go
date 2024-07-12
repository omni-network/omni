package anvil

import (
	"crypto/ecdsa"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

//nolint:gochecknoglobals // Static keys and addresses
var (
	acc0 = addr("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	acc1 = addr("0x70997970C51812dc3A010C7d01b50e0d17dc79C8")
	acc2 = addr("0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC")
	acc3 = addr("0x90F79bf6EB2c4f870365E785982E1f101E93b906")
	acc4 = addr("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65")
	acc5 = addr("0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc")
	acc6 = addr("0x976EA74026E726554dB657fA54763abd0C3a0aa9")
	acc7 = addr("0x14dC79964da2C08b23698B3D3cc7Ca32193d9955")
	acc8 = addr("0x23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f")
	acc9 = addr("0xa0Ee7A142d267C1f36714E4a8F75612F20a79720")

	pk0 = mustHexToKey("0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	pk1 = mustHexToKey("0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d")
	pk2 = mustHexToKey("0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a")
	pk3 = mustHexToKey("0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6")
	pk4 = mustHexToKey("0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a")
	pk5 = mustHexToKey("0x8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba")
	pk6 = mustHexToKey("0x92db14e403b83dfe3df233f83dfa3a0d7096f21ca9b0d6d6b8d88b2b4ec1564e")
	pk7 = mustHexToKey("0x4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356")
	pk8 = mustHexToKey("0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97")
	pk9 = mustHexToKey("0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6")
)

//
// Accounts.
//

func DevAccount0() common.Address {
	return acc0
}

func DevAccount1() common.Address {
	return acc1
}

func DevAccount2() common.Address {
	return acc2
}

func DevAccount3() common.Address {
	return acc3
}

func DevAccount4() common.Address {
	return acc4
}

func DevAccount5() common.Address {
	return acc5
}

func DevAccount6() common.Address {
	return acc6
}

func DevAccount7() common.Address {
	return acc7
}

func DevAccount8() common.Address {
	return acc8
}

func DevAccount9() common.Address {
	return acc9
}

func IsDevAccount(addr common.Address) bool {
	switch addr {
	case acc0, acc1, acc2, acc3, acc4, acc5, acc6, acc7, acc8, acc9:
		return true
	}

	return false
}

//
// Private keys.
//

func DevPrivateKey0() *ecdsa.PrivateKey {
	return pk0
}

func DevPrivateKey1() *ecdsa.PrivateKey {
	return pk1
}

func DevPrivateKey2() *ecdsa.PrivateKey {
	return pk2
}

func DevPrivateKey3() *ecdsa.PrivateKey {
	return pk3
}

func DevPrivateKey4() *ecdsa.PrivateKey {
	return pk4
}

func DevPrivateKey5() *ecdsa.PrivateKey {
	return pk5
}

func DevPrivateKey6() *ecdsa.PrivateKey {
	return pk6
}

func DevPrivateKey7() *ecdsa.PrivateKey {
	return pk7
}

func DevPrivateKey8() *ecdsa.PrivateKey {
	return pk8
}

func DevPrivateKey9() *ecdsa.PrivateKey {
	return pk9
}

func DevPrivateKeys() []*ecdsa.PrivateKey {
	return []*ecdsa.PrivateKey{pk0, pk1, pk2, pk3, pk4, pk5, pk6, pk7, pk8, pk9}
}

func PrivateKey(account common.Address) (*ecdsa.PrivateKey, bool) {
	switch account {
	case acc0:
		return pk0, true
	case acc1:
		return pk1, true
	case acc2:
		return pk2, true
	case acc3:
		return pk3, true
	case acc4:
		return pk4, true
	case acc5:
		return pk5, true
	case acc6:
		return pk6, true
	case acc7:
		return pk7, true
	case acc8:
		return pk8, true
	case acc9:
		return pk9, true
	default:
		return nil, false
	}
}

//
// Utils.
//

func addr(hex string) common.Address {
	return common.HexToAddress(hex)
}

func mustHexToKey(privKeyHex string) *ecdsa.PrivateKey {
	privKey, err := crypto.HexToECDSA(strings.TrimPrefix(privKeyHex, "0x"))
	if err != nil {
		panic(err)
	}

	return privKey
}
