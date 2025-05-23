package anchorinbox

import (
	"sync"

	"github.com/omni-network/omni/anchor/localnet"
	"github.com/omni-network/omni/lib/svmutil"
)

var initOnce sync.Once

// Program returns the program instance, it also ensures that the program ID is set only once.
func Program() svmutil.Program {
	program := svmutil.Program{
		Name:         "solver_inbox",
		SharedObject: localnet.InboxSO,
		KeyPairJSON:  localnet.InboxKeyPairJSON,
	}

	initOnce.Do(func() {
		SetProgramID(program.MustPublicKey())
	})

	return program
}
