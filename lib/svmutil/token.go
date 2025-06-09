package svmutil

import (
	"context"
	"crypto/ed25519"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"

	"github.com/gagliardetto/solana-go"
	associatedtokenaccount "github.com/gagliardetto/solana-go/programs/associated-token-account"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
)

var (
	// DevnetUSDCMint is the USDC mint private key on local devnet.
	DevnetUSDCMint = solana.PrivateKey(ed25519.NewKeyFromSeed([]byte(usdcSeed)))
	usdcSeed       = "usdc_devnet_seed________________"
)

type CreateMintResp struct {
	MintAccount solana.PublicKey
	Authority   solana.PrivateKey
	ATAs        map[solana.PublicKey]solana.PublicKey // Map of funded associated token accounts
}

func (r CreateMintResp) AuthATA() solana.PublicKey {
	return r.ATAs[r.Authority.PublicKey()]
}

func CreateMint(
	ctx context.Context,
	cl *rpc.Client,
	authority solana.PrivateKey,
	mint solana.PrivateKey,
	decimals uint8,
	toFund ...solana.PublicKey,
) (CreateMintResp, error) {
	atas := make(map[solana.PublicKey]solana.PublicKey)
	authATA, _, err := solana.FindAssociatedTokenAddress(authority.PublicKey(), mint.PublicKey())
	if err != nil {
		return CreateMintResp{}, errors.Wrap(err, "find associated token address")
	}
	atas[authority.PublicKey()] = authATA

	const mintAmount uint64 = 1e9 // 1G tokens
	amount := bi.N(mintAmount)
	for range decimals {
		amount = bi.MulRaw(amount, 10)
	}

	// Calculate rent
	rent, err := cl.GetMinimumBalanceForRentExemption(ctx, token.MINT_SIZE, rpc.CommitmentConfirmed)
	if err != nil {
		return CreateMintResp{}, errors.Wrap(err, "get minimum balance for rent exemption")
	}

	// Create mint account, and initialize it, create associated token account, and mint 1G tokens
	createAccount := system.NewCreateAccountInstruction(rent, token.MINT_SIZE, solana.TokenProgramID, authority.PublicKey(), mint.PublicKey())
	initMint := token.NewInitializeMint2Instruction(decimals, authority.PublicKey(), authority.PublicKey(), mint.PublicKey())
	createATA := associatedtokenaccount.NewCreateInstruction(authority.PublicKey(), authority.PublicKey(), mint.PublicKey())
	mintTo := token.NewMintToInstruction(amount.Uint64(), mint.PublicKey(), authATA, authority.PublicKey(), nil)

	txSig, err := Send(ctx, cl,
		WithInstructions(
			createAccount.Build(),
			initMint.Build(),
			createATA.Build(),
			mintTo.Build(),
		),
		WithPrivateKeys(authority, mint),
	)
	if err != nil {
		return CreateMintResp{}, errors.Wrap(err, "send create mint transaction")
	}

	_, err = AwaitConfirmedTransaction(ctx, cl, txSig)
	if err != nil {
		return CreateMintResp{}, errors.Wrap(err, "await initial confirmed transaction")
	}

	// Add additional mintTos for all fund accounts.
	var instrs []solana.Instruction
	for _, fund := range toFund {
		if _, ok := atas[fund]; ok {
			continue // If the ATA already exists, skip creating it.
		}

		fundATA, _, err := solana.FindAssociatedTokenAddress(fund, mint.PublicKey())
		if err != nil {
			return CreateMintResp{}, errors.Wrap(err, "find associated token address for fund account")
		}

		createATA := associatedtokenaccount.NewCreateInstruction(authority.PublicKey(), fund, mint.PublicKey())
		fundTo := token.NewMintToInstruction(amount.Uint64(), mint.PublicKey(), fundATA, authority.PublicKey(), nil)
		instrs = append(instrs, createATA.Build(), fundTo.Build())
		atas[fund] = fundATA // TODO(corver): Split if too large to submit
	}
	if len(instrs) > 0 {
		txSig, err = Send(ctx, cl,
			WithInstructions(instrs...),
			WithPrivateKeys(authority),
		)
		if err != nil {
			return CreateMintResp{}, errors.Wrap(err, "send mint to fund accounts transaction")
		}

		_, err = AwaitConfirmedTransaction(ctx, cl, txSig)
		if err != nil {
			return CreateMintResp{}, errors.Wrap(err, "await second confirmed transaction")
		}
	}

	bal, err := cl.GetTokenAccountBalance(ctx, authATA, rpc.CommitmentConfirmed)
	if err != nil {
		return CreateMintResp{}, errors.Wrap(err, "get token account balance")
	} else if bal.Value == nil || bal.Value.UiAmount == nil {
		return CreateMintResp{}, errors.New("token account balance is nil")
	} else if *bal.Value.UiAmount != float64(mintAmount) {
		return CreateMintResp{}, errors.New("unexpected token account balance", "got", *bal.Value.UiAmount, "want", float64(mintAmount))
	}

	return CreateMintResp{
		MintAccount: mint.PublicKey(),
		ATAs:        atas,
		Authority:   authority,
	}, nil
}

// EnsureATA returns (and possibly creates) the owner's associated token account (ATA) for the given mint.
func EnsureATA(ctx context.Context, cl *rpc.Client, mint solana.PublicKey, owner solana.PrivateKey) (solana.PublicKey, error) {
	resp, err := EnsureATAs(ctx, cl, mint, owner)
	if err != nil {
		return solana.PublicKey{}, errors.Wrap(err, "ensure associated token account")
	}

	return resp[owner.PublicKey()], nil
}

// EnsureATAs finds or creates the owners' associated token account (ATA) for the given mint.
func EnsureATAs(ctx context.Context, cl *rpc.Client, mint solana.PublicKey, owners ...solana.PrivateKey) (map[solana.PublicKey]solana.PublicKey, error) {
	resp := make(map[solana.PublicKey]solana.PublicKey)
	var instrs []solana.Instruction
	for _, owner := range owners {
		ata, _, err := solana.FindAssociatedTokenAddress(owner.PublicKey(), mint)
		if err != nil {
			return nil, errors.Wrap(err, "find associated token address")
		}

		// Check if the ATA already exists
		info, err := cl.GetAccountInfo(ctx, ata)
		if err != nil {
			return nil, errors.Wrap(err, "get account info")
		}
		if info.Value == nil {
			// Create the ATA instruction if it does not exist
			instr := associatedtokenaccount.NewCreateInstruction(owner.PublicKey(), owner.PublicKey(), mint)
			instrs = append(instrs, instr.Build())
			resp[owner.PublicKey()] = ata
		}
		resp[owner.PublicKey()] = ata
	}

	if len(instrs) == 0 {
		// All ATAs already exist, nothing to do
		return resp, nil
	}

	txSig, err := Send(ctx, cl, WithInstructions(instrs...), WithPrivateKeys(owners...))
	if err != nil {
		return nil, errors.Wrap(err, "send create ATA transaction")
	}

	_, err = AwaitConfirmedTransaction(ctx, cl, txSig)
	if err != nil {
		return nil, errors.Wrap(err, "await confirmed transaction")
	}

	return resp, nil
}
