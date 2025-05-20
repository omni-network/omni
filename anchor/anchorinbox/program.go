package anchorinbox

import (
	"sync"

	"github.com/omni-network/omni/anchor/localnet"
	"github.com/omni-network/omni/lib/solutil"
)

var initOnce sync.Once

// Program returns the program instance, it also ensures that the program ID is set only once.
func Program() solutil.Program {
	program := solutil.Program{
		Name:         "solver_inbox",
		SharedObject: localnet.InboxSO,
		KeyPairJSON:  localnet.InboxKeyPairJSON,
	}

	initOnce.Do(func() {
		SetProgramID(program.MustPublicKey())
	})

	return program
}
