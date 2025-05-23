package eoa

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/BurntSushi/toml"
	"github.com/gagliardetto/solana-go"

	_ "embed"
)

// svmRemoteAddrs maps remote key evm addresses to svm addresses.
var svmRemoteAddrs = map[common.Address]solana.PublicKey{
	hexToAddr("0x64Bf40F5E6C4DE0dfe8fE6837F6339455657A2F5"): base58ToAddr("5FUCCCq5Bnd6XtTToHg2ZsiB9BvhGPkoTQQgMbE7Noo1"), // shared-cold
	hexToAddr("0xf63316AA39fEc9D2109AB0D9c7B1eE3a6F60AEA4"): base58ToAddr("HnsgSgr78S5hMrKFztxWMfekR6UQqogN4dqMoabT5oiB"), // shared-hot
	// TODO(corver): Add rest of fileblocks accounts here.
}

//go:embed svmaddrs.toml
var svmAddrsTOML []byte
var svmAddrs = func() map[common.Address]solana.PublicKey {
	s := struct {
		Addrs map[string]string `toml:"addrs"`
	}{}

	if _, err := toml.Decode(string(svmAddrsTOML), &s); err != nil {
		panic(err)
	}

	m := make(map[common.Address]solana.PublicKey, len(s.Addrs))
	for k, v := range s.Addrs {
		m[common.HexToAddress(k)] = solana.MustPublicKeyFromBase58(v)
	}

	return m
}()

func hexToAddr(hex string) common.Address {
	return common.HexToAddress(hex)
}

func base58ToAddr(base58 string) solana.PublicKey {
	return solana.MustPublicKeyFromBase58(base58)
}
