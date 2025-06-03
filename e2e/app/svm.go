package app

import (
	"context"
	"path/filepath"
	"time"

	"github.com/omni-network/omni/anchor/anchorinbox"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/svmutil"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

// svmInitAsync initializes the SVM asynchronously.
func svmInitAsync(ctx context.Context, def Definition) <-chan error {
	resp := make(chan error, 1)
	go func() {
		resp <- svmInit(ctx, def)
	}()

	return resp
}

// svmInit initializes the SVM, deploying programs and funding accounts.
func svmInit(ctx context.Context, def Definition) error {
	if len(def.Testnet.SVMChains) == 0 {
		return nil
	} else if len(def.Testnet.SVMChains) > 1 {
		return errors.New("multiple SVM chains")
	} else if def.Testnet.Network != netconf.Devnet {
		return errors.New("svm is only available on Devnet")
	}

	svmChain := def.Testnet.SVMChains[0]
	svmDir := filepath.Join(def.Testnet.Dir, "svm")
	cl := rpc.New(svmChain.ExternalRPC)

	roleKeys, err := svmRoleKeys(ctx, def.Testnet.Network, svmChain.ChainID)
	if err != nil {
		return err
	}

	var roleAccounts []solana.PublicKey
	for _, key := range roleKeys {
		roleAccounts = append(roleAccounts, key.PublicKey())
	}

	log.Debug(ctx, "SVM: Requesting role airdrops")
	// Fund all roles with SOL
	fundAmount := solana.LAMPORTS_PER_SOL * 1e6 // 1M SOL in lamports
	for _, account := range roleAccounts {
		if _, err := cl.RequestAirdrop(ctx, account, fundAmount, ""); err != nil {
			return errors.Wrap(err, "request airdrop for role account", "account", account)
		}
	}

	log.Debug(ctx, "SVM: Creating USDC mint and funding role tokens")
	// Create USDC mint (and fund all role accounts with tokens)
	mintResp, err := svmutil.CreateMint(ctx, cl, roleKeys[eoa.RoleDeployer], svmutil.DevnetUSDCMint, 6, roleAccounts...)
	if err != nil {
		return errors.Wrap(err, "create USDC mint")
	}

	// Deploy the anchorinbox program
	log.Debug(ctx, "SVM: Deploying anchorinbox program")
	_, err = svmutil.Deploy(ctx, cl, svmDir, anchorinbox.Program())
	if err != nil {
		return errors.Wrap(err, "deploy anchorinbox program")
	}

	log.Debug(ctx, "SVM: Initializing anchorinbox program")
	const closeBuffer = time.Second * 0 // Allow immediate closing in devnet
	init, err := anchorinbox.NewInit(svmChain.ChainID, closeBuffer, roleKeys[eoa.RoleSolver].PublicKey())
	if err != nil {
		return errors.Wrap(err, "create anchorinbox init instruction")
	}
	_, err = svmutil.SendSimple(ctx, cl, roleKeys[eoa.RoleSolver], init.Build())
	if err != nil {
		return err
	}

	log.Info(ctx, "SVM initialized", "usdc_mint", mintResp.MintAccount, "anchor_inbox", anchorinbox.ProgramID)

	return nil
}

func svmRoleKeys(ctx context.Context, network netconf.ID, chainID uint64) (map[eoa.Role]solana.PrivateKey, error) {
	if network != netconf.Devnet {
		return nil, errors.New("svm role keys are only available on Devnet")
	}

	keys := make(map[eoa.Role]solana.PrivateKey)
	for _, role := range eoa.AllRoles() {
		if solvernet.SkipRole(chainID, role) {
			continue
		}

		privKey, err := eoa.PrivateKey(ctx, network, role)
		if err != nil {
			return nil, errors.Wrap(err, "get private key for role", "role", role)
		}

		keys[role] = svmutil.MapEVMKey(privKey)
	}

	return keys, nil
}
