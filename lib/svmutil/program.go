package svmutil

import "github.com/gagliardetto/solana-go"

// Program represents a Solana executable program (smart contract).
type Program struct {
	Name         string
	SharedObject []byte // Compiled BPF shared object
	KeyPairJSON  []byte
}

func (p Program) SOFile() string {
	return p.Name + ".so"
}

func (p Program) KeyPairFile() string {
	return p.Name + "-keypair.json"
}

func (p Program) MustPrivateKey() solana.PrivateKey {
	pk, err := solana.PrivateKeyFromSolanaKeygenFileBytes(p.KeyPairJSON)
	if err != nil {
		panic(err)
	}

	return pk
}

func (p Program) MustPublicKey() solana.PublicKey {
	return p.MustPrivateKey().PublicKey()
}
