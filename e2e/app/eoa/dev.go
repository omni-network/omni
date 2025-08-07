package eoa

import (
	"crypto/ecdsa"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// dev nemonic: "lemon prefer stereo virtual exact metal kit genuine journey detail ivory camera"

var (
	acc0 = addr("0xE0cF003AC27FaeC91f107E3834968A601842e9c6")
	acc1 = addr("0x3C298a8fAb961CC155F48557872D37D39015c5bc")
	acc2 = addr("0xD1862026335f1cfD7c5DB13e58bA4C97247A7998")
	acc3 = addr("0xC38Ef10ecC4aD9c24554cACd67cA4896B4fb2F9C")
	acc4 = addr("0x8696e3d6cAD982C32F86320a6f0E1AB8aB3Db3b9")
	acc5 = addr("0xB5775c7e3796822dA059B73218E8a9033dA9bC67")
	acc6 = addr("0x9637Bc245647B4cdD85e3bB092c672e2ddD28539")
	acc7 = addr("0x23a4523A3EE6220fB2CdDc5Ab94A2780D6493230")
	acc8 = addr("0xACb32F1b31F818511139a67b79010fA011960764")
	acc9 = addr("0x3c56a0cDB54D07A91791b698d8B390aB53208E92")

	createXDeployer = addr("0xeD456e05CaAb11d66C4c797dD6c1D6f9A7F352b5")

	pk0 = mustHexToKey("0xbb119deceaff95378015e684292e91a37ef2ae1522f300a2cfdcb5b004bbf00d")
	pk1 = mustHexToKey("0xed3dccb053880be5b681f6f0256fc18410f99bd69fedfc80f6b37e7930d1c526")
	pk2 = mustHexToKey("0xe0cff2136e89d72576e1e8a3af640af8509fa07191429766d89814923b9ddbc2")
	pk3 = mustHexToKey("0x7b63e2b097b37620054a0c3ba9e2c0253751b26f62f866dcbde64508071c0048")
	pk4 = mustHexToKey("0x2b6c35e1914655810e6471aea16358c335995985eb5d1a8e2a49ee5dca6779c1")
	pk5 = mustHexToKey("0x49f5dabfb06f9febf27b3b07dc68f5c3a022edf8ac56631e92ba8879d7d7f44e")
	pk6 = mustHexToKey("0x07fe974baf69d3d4d93438698155adffae10e24ab14eaa468e69a17c2a1295d3")
	pk7 = mustHexToKey("0x8b8638c2c593903c63f096200dad748af3e36ca4dbcfd11f57cfa83e98cfdbc8")
	pk8 = mustHexToKey("0xdf6cd4fcdba1068873acef34a91f41c9af581dafd52fbc2908add6fb213f5e02")
	pk9 = mustHexToKey("0xefebe60ea08dbc53c3e59612380501d4cbe7309548d84692ec978e06e3c61137")
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

func DevAccounts() []common.Address {
	return []common.Address{acc0, acc1, acc2, acc3, acc4, acc5, acc6, acc7, acc8, acc9}
}

func IsDevAccount(addr common.Address) bool {
	switch addr {
	case acc0, acc1, acc2, acc3, acc4, acc5, acc6, acc7, acc8, acc9:
		return true
	}

	return false
}

func CreateXDeployer() common.Address {
	return createXDeployer
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

func DevPrivateKey(account common.Address) (*ecdsa.PrivateKey, bool) {
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
