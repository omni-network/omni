package svmutil

import (
	"context"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"

	"github.com/gagliardetto/solana-go"
	associatedtokenaccount "github.com/gagliardetto/solana-go/programs/associated-token-account"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/programs/tokenregistry"
	"github.com/gagliardetto/solana-go/rpc"
)

type CreateMintResp struct {
	Symbol      string
	MintAccount solana.PublicKey
	AuthATA     solana.PublicKey
	Authority   solana.PrivateKey
}

func CreateMint(ctx context.Context, cl *rpc.Client, authority solana.PrivateKey, symbol string, decimals uint8) (CreateMintResp, error) {
	// Create new random mint and authority ata
	mintPrivKey, err := solana.NewRandomPrivateKey()
	if err != nil {
		return CreateMintResp{}, errors.Wrap(err, "create random mint private key")
	}

	authATA, _, err := solana.FindAssociatedTokenAddress(authority.PublicKey(), mintPrivKey.PublicKey())
	if err != nil {
		return CreateMintResp{}, errors.Wrap(err, "find associated token address")
	}

	mintMeta, _, err := solana.FindTokenMetadataAddress(mintPrivKey.PublicKey())
	if err != nil {
		return CreateMintResp{}, errors.Wrap(err, "find token metadata address")
	}

	name32, err := tokenregistry.NameFromString(symbol)
	if err != nil {
		return CreateMintResp{}, errors.Wrap(err, "convert name")
	}
	symbol32, err := tokenregistry.SymbolFromString(symbol)
	if err != nil {
		return CreateMintResp{}, errors.Wrap(err, "convert name to symbol")
	}

	const mintAmount uint64 = 1e9 // 1G tokens
	var logo tokenregistry.Logo
	var website tokenregistry.Website

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
	createAccount := system.NewCreateAccountInstruction(rent, token.MINT_SIZE, solana.TokenProgramID, authority.PublicKey(), mintPrivKey.PublicKey())
	initMint := token.NewInitializeMint2Instruction(decimals, authority.PublicKey(), authority.PublicKey(), mintPrivKey.PublicKey())
	_ = tokenregistry.NewRegisterTokenInstruction(logo, name32, symbol32, website, mintMeta, authority.PublicKey(), mintPrivKey.PublicKey())
	createATA := associatedtokenaccount.NewCreateInstruction(authority.PublicKey(), authority.PublicKey(), mintPrivKey.PublicKey())
	mintTo := token.NewMintToInstruction(amount.Uint64(), mintPrivKey.PublicKey(), authATA, authority.PublicKey(), nil)

	txSig, err := Send(ctx, cl,
		WithInstructions(
			createAccount.Build(),
			initMint.Build(),
			//  registerMint,
			createATA.Build(),
			mintTo.Build(),
		),
		WithPrivateKeys(authority, mintPrivKey),
	)
	if err != nil {
		return CreateMintResp{}, errors.Wrap(err, "send create mint transaction")
	}

	_, err = AwaitConfirmedTransaction(ctx, cl, txSig)
	if err != nil {
		return CreateMintResp{}, errors.Wrap(err, "await confirmed transaction")
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
		Symbol:      symbol,
		MintAccount: mintPrivKey.PublicKey(),
		AuthATA:     authATA,
		Authority:   authority,
	}, nil
}
