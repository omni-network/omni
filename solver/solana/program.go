package solana

import (
	"github.com/omni-network/omni/solver/solana/events"

	"github.com/gagliardetto/solana-go"

	_ "embed"
)

var (
	//go:embed events/events.so
	soEvents []byte

	//go:embed events/events-keypair.json
	keyPairEvents []byte

	ProgramEvents = Program{
		Name:         "events",
		SharedObject: soEvents,
		KeyPairJSON:  keyPairEvents,
	}

	Programs = []Program{
		ProgramEvents,
	}
)

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

func (p Program) Init() {
	events.SetProgramID(p.MustPublicKey())
}

//nolint:gochecknoinits // init isn't cool, but the generated code requires it.
func init() {
	for _, program := range Programs {
		program.Init()
	}
}
