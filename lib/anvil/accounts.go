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

	pk0 = privkey("0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	pk1 = privkey("0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d")
	pk2 = privkey("0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a")
	pk3 = privkey("0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6")
	pk4 = privkey("0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a")
	pk5 = privkey("0x8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba")
	pk6 = privkey("0x92db14e403b83dfe3df233f83dfa3a0d7096f21ca9b0d6d6b8d88b2b4ec1564e")
	pk7 = privkey("0x4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356")
	pk8 = privkey("0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97")
	pk9 = privkey("0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6")
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

// ExternalAccounts defines a list of well-known genesis-funded (1M OMNI) EOAs for use by external apps (non-golang).
var ExternalAccounts = map[common.Address]*ecdsa.PrivateKey{
	addr("0xf2b7b8aBe9FA4F2c4Aea113ab396165f20E9fdEC"): privkey("0x55bd495fdeb8db80fa102995c5a8de3aa14306a03b4bcf94722723d9c4c680f6"),
	addr("0xecC80DF24e41FcEaF49FE9C27FC962E540c22178"): privkey("0x5e8288546360010e8d03eee564dbdd5dcb049c3548ddd3b163eb88f2577a0682"),
	addr("0x6d03bbf81759C48cD9f2AE73cFb8F354f21972ab"): privkey("0x41e79ac587f190d64fee4e2a11bfb06e162feb230737d9518310f3174314aaaf"),
	addr("0xb8426BE760A5c5b852dD3a6f09566bE334502e23"): privkey("0x8602a3d7650c6f3c1fb43b1f3c9e995b647cc5b45442bcd26c2b6329ee6aebc0"),
	addr("0xa3b4067cFbc474d1f75808807B603279AE394831"): privkey("0x360e96f85bb1ac794e0018f70c997e2768396a3f78d8177574e80184bc13704a"),
	addr("0x1f86Fb8aa281033C60688EF7e967Aa7310B057A2"): privkey("0x30ce7e482a2ca15d068ee41638cfaa8e5fda02a81de7008075a3f2294845abf5"),
	addr("0x749c0bfDfe0BB271b6e65b2359F0CefEC9a60c6D"): privkey("0x3be0fd30e6410c6a7b3c5ad60035f57c763a370a194fd109388f6f1c4ae5e93a"),
	addr("0xccC3D46dB2dADdF35be1eD9e76f4a2596272d494"): privkey("0x5d61e5cd876189526b51a694eb8ec47ef029524d7e3c63c46db5e5f5ae0faa71"),
	addr("0xe668Df3744D6Bd761456154430b6ACD2f1B6B59b"): privkey("0x84bf5c761eefa3b0ff2ef224cd19de2ea13390b3345c3bfbd877c6dc5c35b0ec"),
	addr("0xCC7fE83163f9F1d1901aeA7810953B0a8655E8F0"): privkey("0xe6f881b604af7ea169b290009d5e258cee7af6945194e459de5ce438bd3fa3f2"),
	addr("0xAaAF516cAF63eb8E4D7961160832edAcefEA6C06"): privkey("0x026f94500df4998bca119ed04e1c757be6fb7150e26f42dc4598d7d196af77e0"),
	addr("0x83e08a21b0bab74C78B24d2a6FA7F869cdE3d65C"): privkey("0x4e517daaf3204459b1d3c3e682df9aecbeeb1fbd833894d52c0e2271f34829bc"),
	addr("0x9F239f03020AcF4B6dCBCE18Cb343CD68663cf23"): privkey("0x7b1070a4cd5577abe1fad5d44c0491675a001ff8585f71dcd3cf72bfd519fb64"),
	addr("0xeB2fcdb2F48d112EdCDBd90227D07a633a3E6DC6"): privkey("0xe3aaa280e0d2086e7a2a84a29940fc89b7844d4832f919a161e91f92679409e8"),
	addr("0xe81911953fba2d0A84680454E3d3e41Cc1894034"): privkey("0x2d5174493ed3142674a0d89af5ac993e75391d02f1eede59299b511545c59a65"),
	addr("0x9884754FaDAad8C1B2CbfFaFD8292f50A880c987"): privkey("0x1b8c4fed0a4b8bd9bbe922ad527f5a7864443fff8d2e3f01a47603ce6518b770"),
	addr("0x7ba1d63591A464612d8f49BDD675E2b5A1dfAe18"): privkey("0x1b87064a9d7e84970d1555932f2e22f19e6050be3cdc7c0da62c3eda8df955b9"),
	addr("0xdEF0f650cB7571fdD83EAf253FC6cF141b1d8721"): privkey("0x12f9136c4cd59932511d46fbd7fd0fea71072c51623feb32cb92cf3a62fb74be"),
	addr("0xe0461cE0Fb767906DdffDa2177139922c4d1acF5"): privkey("0x7f0aa4716a0d8273ae9b26ffe4ec1bd98ff5a3dc5535a5160787a439b1ced42e"),
	addr("0xfcf7DDF31642C3c572813d86FAA6aCd942aD92D9"): privkey("0x6b86c7b78f9eeeef39a13747af93a7c9f3e5d46ea16d4062e558b975ba010b3b"),
	addr("0x7b1260687b3a8E09440B5BddD1B8c13414DaFA2B"): privkey("0xa4c6d9eea3e1dbe48fe669df6885823cf5998ccb7e337e62d0f7555cf9c73ce8"),
	addr("0xcf05d0CdbDFfe0aF5D3B01d857970EE22c1C983E"): privkey("0x2fd5ff7436a531d0f7bf702f8a3a566f4843e02c4b877bf0cbd8189dbbea86a7"),
	addr("0x6Da3e1eec0D7937a27D6D27A6b130ECA153A9A74"): privkey("0x8d591b5904c30afcb77f97c849dcf58e69ab07e7daae624941bd4b525525ab7b"),
	addr("0xD231205ca62a542594990DF78F75fcf49d2DDcA5"): privkey("0x8203d53d31a23f40f366447bc3f84e1bf00216aea0fae54e1e84b804f4bf5caa"),
	addr("0x80ecc9DA9667e862C0bBbCC2c24233CbBD4CBc90"): privkey("0x4d927d7cadd8fc19787e3f04b69ebccc80578e9f5a1919b00c20b886de1a98c0"),
	addr("0x180F59BB0AA7215AC69F7916774EF4b78095f220"): privkey("0x09e56ad74e86661e08d887e1402b8db87195184573a4c4cc5d6fb67cc2b68925"),
	addr("0x8B450783c31e516AD77440feC0E200C1851b32c9"): privkey("0xb8c2933fa8b492c1dd8f3f60eee9be583b127368a57e5fea29eed07067e47cd3"),
	addr("0x7e3af6da131c64C212E56002D64407E3Cba4580C"): privkey("0x05398c7c989654018c122ad96d2372c2fc448bc4fa716234f6c02d431e6671fd"),
	addr("0x1ca1Fd568bb67cBe516043F16b46CBd2185Fd4c7"): privkey("0xcac019956cf59ce7b3e9fec55495df08789a1b0907d340ea12ff00ee4282170d"),
	addr("0xDfc3b08739c5fC3744E7b98E9b3A0584d08f6A9f"): privkey("0x88aed8816ec0e891c5eb34deb4a29f3655b36f1d6c6c87d22e37e3b084c8aac9"),
	addr("0x7fe20ae030362Db3A5dFF06ee5e556C2088D579D"): privkey("0xeda58be6fb77ff82caee900a1ae455d1d8bd9968f1e3f0931eaa5e1216f88982"),
	addr("0x7Ad20D3BFD931a49897dC6CBAbefdE4AcE25c525"): privkey("0xe4abc5daef2fce7e24ebda2ca217fed3079bbd078d0e8212b342971507e41ed1"),
	addr("0xEd0dDbCbA145aa7f6B8b765CF3dB3F6F6dB04884"): privkey("0x65144e873b4c865d4047cb9be04d16eb20e65d56bae1b383da674ba1de0b08f3"),
	addr("0x2fb06161957ac3C2fdf951F460d8685F5EF55DAc"): privkey("0x297edfbd5cfe720bb5d96a5e2c5a3334a03b7d4a4671f1cc44464b38f0630b6d"),
	addr("0x80A37dd40d265Da50192B449Fbec9F53Fb146CFb"): privkey("0x3d75dcb1e3e3d5a17200f9a6064f9dc7ad8db5d2e5387f64ceee4a19bb9e4457"),
	addr("0xc56d16B0F862dBe86B67C36825A4Cd17cbb0b5c1"): privkey("0x8f1a85bc142e4d400baf631da419b3fbe411aff5502473087deec560d5ffb03a"),
	addr("0x2519409aeb0BAf036CA61b28e8BB86B25d5d8400"): privkey("0xe6200e5bcbf9a4c4b80f80d1eb6148d7940f514fe8c59a12e079bfe0425645fc"),
	addr("0xd99648FF74e1825a4455D773480Ce5d66cF47260"): privkey("0xc6c3cc159f8c150136bb68754e319757d83ed3aac1e95f2667603691fc156b97"),
	addr("0xe46D8269632288b00D54900Db2aE122853eeA33c"): privkey("0xa658080579738b7abb72cd43f6b30c9f0a58fff4fb8963a346414395e38d70b0"),
	addr("0x803156801039392DF33d59f24d721D2c95b23f75"): privkey("0xaedcf756189631aee6e99cae57f88347d0859bba9d8fcd4e8a5aecaf87ca0026"),
	addr("0x03eddb70F1b5ceBf05B4Ac60e6520CBA3efdDbC9"): privkey("0x69100d2216af2ed8a2906808bbf9e053572633fe278edc53cd9e67fa3085906d"),
	addr("0xB9CD86719B236E681C4723E637f28706EFDA475D"): privkey("0x8d565889959795b99d20cd1c7ee4385462f48b41dced673db0c48a734ce89ff6"),
	addr("0xc9f1efC3b4DB437D10c7238De5CCe55ED42c39De"): privkey("0xc84ceb6d2d1ecd2e6223936df873377f917b0e664938fd151243983598f482dd"),
	addr("0xE039Faae15979ff2e07E60027283e20C70104D11"): privkey("0x73895da81e23504403fcfed8453550ac0b4c4b5ad8d87c5aeeb3496b9fc4527f"),
	addr("0xd2d6283C6dd35Fccb5152cf1901fcF297C95B2F0"): privkey("0x5d12c538b541edf1bc2408e976f210c1a8c98badc86712904c3196eaa09382ac"),
	addr("0x3667FCEA29E6A6A180844C979aAC7cB0e4280A4A"): privkey("0x4838cf7d06bcf4ad6376073327efc810b0a2472cc9e56e083efb26ed5b5754d1"),
	addr("0xFbDa4078B4aB093572633B3Bfe4Cb0e0AAbAFB5D"): privkey("0x6a32db912a372b0d142de95327bda0b00b2c3596c0cd37d73349f290992bf64c"),
	addr("0x82f0F3F7c8adEF29b6E2c540B9fA9E51d0B485C7"): privkey("0x65af5f173a51fad5c3e41260334f48dfd177e5774a4bcabb63979a50e01d5f3e"),
	addr("0xaC8DE55aB9d081130679c17B3c9072f5A3Ed940E"): privkey("0x0e12c5229b7b784e554451d226b66bd80621c8351c9158ff9da17373b3ff8b89"),
	addr("0xEe035D5153c472Ce5490fC9e493e3f7CA5b3bfcb"): privkey("0x125a9973c23100aa9120f84496c83a9c717698b0031c19c52c0c4d325e2ccca8"),
}

//
// Utils.
//

func addr(hex string) common.Address {
	return common.HexToAddress(hex)
}

func privkey(privKeyHex string) *ecdsa.PrivateKey {
	privKey, err := crypto.HexToECDSA(strings.TrimPrefix(privKeyHex, "0x"))
	if err != nil {
		panic(err)
	}

	return privKey
}
