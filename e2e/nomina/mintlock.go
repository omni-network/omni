package nomina

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

// MintLock is the deployed address of the NominaMintLock contract.
var MintLock = common.HexToAddress("0xF9046e60f10000c97316D76Ba0DbAB399C3D8752")

// LockNomMint calls Nomina.setMintAuthority to queue the NominaMintLock contract as the
// pending mint authority, permanently locking the ability to mint NOM once accepted.
func LockNomMint(ctx context.Context, backend *ethbackend.Backend) error {
	nomAddr := contracts.NomAddr(netconf.Mainnet)

	nomina, err := bindings.NewNomina(nomAddr, backend)
	if err != nil {
		return errors.Wrap(err, "new nomina")
	}

	mintAuthority := eoa.MustAddress(netconf.Mainnet, eoa.RoleNomAuthority)

	txOpts, err := backend.BindOpts(ctx, mintAuthority)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	tx, err := nomina.SetMintAuthority(txOpts, MintLock)
	if err != nil {
		return errors.Wrap(err, "set mint authority")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	log.Info(ctx, "NOM mint lock queued", "nom", nomAddr, "mint_lock", MintLock)

	return nil
}
