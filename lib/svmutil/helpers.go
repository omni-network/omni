package svmutil

import (
	"context"
	"crypto/ecdsa"
	"crypto/ed25519"
	"math/big"
	"regexp"
	"strings"
	"time"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"

	"github.com/ethereum/go-ethereum/crypto"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

// AwaitConfirmedTransaction waits for a transaction to be confirmed.
func AwaitConfirmedTransaction(ctx context.Context, cl *rpc.Client, txSig solana.Signature) (*rpc.GetTransactionResult, error) {
	for {
		tx, err := cl.GetTransaction(ctx, txSig, &rpc.GetTransactionOpts{
			Encoding:                       solana.EncodingBase64,
			Commitment:                     rpc.CommitmentConfirmed,
			MaxSupportedTransactionVersion: ptr(rpc.MaxSupportedTransactionVersion0),
		})
		if errors.Is(err, rpc.ErrNotFound) {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			return nil, errors.Wrap(err, "get confirmed transaction")
		} else if tx.Meta.Err != nil {
			return tx, errors.New("transaction failed", "meta_err", tx.Meta.Err, "signature", txSig)
		}

		return tx, nil
	}
}

// GetAccountDataInto retrieves account data and decodes it into the provided value.
// It uses commitment level of "confirmed".
func GetAccountDataInto(ctx context.Context, cl *rpc.Client, address solana.PublicKey, val any) (*rpc.GetAccountInfoResult, error) {
	info, err := cl.GetAccountInfoWithOpts(ctx, address, &rpc.GetAccountInfoOpts{
		Commitment: rpc.CommitmentConfirmed,
		Encoding:   solana.EncodingBase64,
	})
	if err != nil {
		return nil, errors.Wrap(WrapRPCError(err, "GetAccountDataInto"), "get account info")
	}

	err = bin.NewBorshDecoder(info.Value.Data.GetBinary()).Decode(val)
	if err != nil {
		return nil, errors.Wrap(err, "decode account data")
	}

	return info, nil
}

var chainIDsByHash = map[solana.Hash]uint64{
	solana.MustHashFromBase58("5eykt4UsFv8P8NJdTREpY1vzqKqZKvdpKuc147dw2N9d"): evmchain.IDSolana,
	solana.MustHashFromBase58("EtWTRABZaYq6iMfeYKouRu166VU2xqa1wcaWoxPkrZBG"): evmchain.IDSolanaDevnet,
}

func ChainID(ctx context.Context, cl *rpc.Client) (uint64, error) {
	hash, err := cl.GetGenesisHash(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "get chain ID")
	}

	if chainID, ok := chainIDsByHash[hash]; ok {
		return chainID, nil
	}

	return evmchain.IDSolanaLocal, nil
}

func NativeBalanceAt(ctx context.Context, cl *rpc.Client, addr solana.PublicKey) (*big.Int, error) {
	resp, err := cl.GetBalance(ctx, addr, rpc.CommitmentConfirmed)
	if err != nil {
		return nil, errors.Wrap(WrapRPCError(err, "getBalance"), "get balance")
	}

	return bi.N(resp.Value), nil
}

func TokenBalanceAt(ctx context.Context, cl *rpc.Client, mint, wallet solana.PublicKey) (*big.Int, error) {
	tokenAccount, _, err := solana.FindAssociatedTokenAddress(wallet, mint)
	if err != nil {
		return nil, errors.Wrap(err, "find ata")
	}

	resp, err := cl.GetTokenAccountBalance(ctx, tokenAccount, rpc.CommitmentConfirmed)
	if err != nil {
		return nil, errors.Wrap(WrapRPCError(err, "getTokenAccountBalance"), "get token balance")
	}

	bal, ok := new(big.Int).SetString(resp.Value.Amount, 10)
	if !ok {
		return nil, errors.New("invalid balance", "amount", resp.Value.Amount)
	}

	return bal, nil
}

// MapEVMKey returns a deterministic mapping of an EVM secp256k1 private key to a Solana ed25519 private key.
func MapEVMKey(key *ecdsa.PrivateKey) solana.PrivateKey {
	return solana.PrivateKey(ed25519.NewKeyFromSeed(crypto.FromECDSA(key)))
}

var stackRegex = regexp.MustCompile(`Program (\w+) (invoke|success|failed)(.*)?`)

// FilterDataLogs filters the logs for a specific program, returning only the data logs
// and true if logs were not truncated.
func FilterDataLogs(logs []string, program solana.PublicKey) ([]string, bool, error) {
	var stack []string
	push := func(program string) {
		stack = append(stack, program)
	}
	pop := func() string {
		if len(stack) == 0 {
			return ""
		}
		resp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		return resp
	}
	current := func() bool {
		return len(stack) > 0 && stack[len(stack)-1] == program.String()
	}

	var filtered []string
	for _, log := range logs {
		if isTruncated(log) {
			break
		}

		// If target program is current, filter any data logs
		if current() && strings.HasPrefix(log, "Program data: ") {
			filtered = append(filtered, log)
			continue
		}

		// Check for stack operations
		matches := stackRegex.FindStringSubmatch(log)
		if len(matches) < 3 {
			continue
		}

		programID := matches[1]
		action := matches[2]

		switch action {
		case "invoke":
			push(programID)
		case "success", "failed":
			if pop() != programID {
				return nil, false, errors.New("stack mismatch", "expected", programID, "got", pop())
			}
		}
	}

	return filtered, len(stack) == 0, nil
}

func isTruncated(l string) bool {
	return l == "Log truncated"
}

// GetBlock is a convenience function returning the block for the given slot,
// or false if no block found for the slot,
// or and error.
func GetBlock(ctx context.Context, cl *rpc.Client, slot uint64, details rpc.TransactionDetailsType) (*rpc.GetBlockResult, bool, error) {
	block, err := cl.GetBlockWithOpts(ctx, slot, &rpc.GetBlockOpts{
		TransactionDetails:             details,
		Rewards:                        ptr(false),
		Commitment:                     rpc.CommitmentConfirmed,
		MaxSupportedTransactionVersion: ptr(rpc.MaxSupportedTransactionVersion0),
	})
	if errors.Is(err, rpc.ErrNotConfirmed) {
		return nil, false, nil
	} else if err != nil {
		return nil, false, WrapRPCError(err, "getBlock")
	}

	return block, true, nil
}
