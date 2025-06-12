package svmutil_test

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/omni-network/omni/lib/svmutil"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"
)

func TestSavePrivateKey(t *testing.T) {
	t.Parallel()

	k, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	file := filepath.Join(t.TempDir(), "key.json")
	err = svmutil.SavePrivateKey(k, file)
	require.NoError(t, err)

	k2, err := solana.PrivateKeyFromSolanaKeygenFile(file)
	require.NoError(t, err)
	require.Equal(t, k, k2)
}

func TestMapEVMKey(t *testing.T) {
	t.Parallel()

	const privHex = "289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032"
	evm, err := crypto.HexToECDSA(privHex)
	require.NoError(t, err)

	svm := svmutil.MapEVMKey(evm)

	evmAddr := crypto.PubkeyToAddress(evm.PublicKey).Hex()
	svmAddr := svm.PublicKey().String()

	require.Equal(t, "0x970E8128AB834E8EAC17Ab8E3812F010678CF791", evmAddr)
	require.Equal(t, "p6K5Ff9XkxxkUbriA9Kzrnn49H3cuGXEw1AiD4faNJyPQwW1Kxgqw8oN8nwQpAQYwfpF7G72tqyuXXXnvZ4dYYu", svm.String())
	require.Equal(t, "G8n9W2tGpcazbyrdQzniSKV6p4uFJArWCENhapDmHTx9", svmAddr)
}

func TestFilterDataLogs(t *testing.T) {
	t.Parallel()
	program := solana.MustPublicKeyFromBase58("whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc")

	expected := [][]string{
		{ // len(logs1)==1
			"Program data: 4cpJr5MroJY4JTBvYKAWFg8AIX2IKcUTlfhh44l86K51YxwETLy2jgBGVr/24UJIeAQBAAAAAAAAtzspUdMcSDAGAQAAAAAAAKNS1475OgAAfWsHOAAAAAAAAAAAAAAAAAAAAAAAAAAAVfdHWYMAAACKoHegEwAAAA==",
		},
		{ // len(logs2)==4
			"Program data: 4cpJr5MroJYpaAqbwx37a8ztWUYMupTp2tsZ/N8dkljA5maKm7tkWgD0lSG1lpAAQgAAAAAAAAAAojlOk/fjAEIAAAAAAAAAAH4JEQMAAAAAVa4gLgAAAAAAAAAAAAAAAAAAAAAAAAAAfREAAAAAAACcAgAAAAAAAA==",
			"Program data: 4cpJr5MroJYadt5TZ+XOA3IuI67qXNX56EwbmbA3icHAnZSH+TTBDQFN5Ow6wT9NkAAAAAAAAAAAfDWSdjVLSpAAAAAAAAAAAFWuIC4AAAAACNelDgAAAAAAAAAAAAAAAAAAAAAAAAAABCMFAAAAAAB/xAAAAAAAAA==",
			"Program data: 4cpJr5MroJZxeaOgaqQiDwwrJfgWEkHBny6g4OKYgjZki9Mpys2zBwC7UuCihSxdCQEAAAAAAAAAYtNAzYAtXQkBAAAAAAAAAAjXpQ4AAAAAXoShDQAAAAAAAAAAAAAAAAAAAAAAAAAAhVMAAAAAAAB6DAAAAAAAAA==",
			"Program data: 4cpJr5MroJaSGF30VKHk6aHXzz99JeTyA05hiqSVtbeRIIcOyEefqQEX6AmOjWl2eQAAAAAAAAAA/OPoQ2XLdXkAAAAAAAAAAF6EoQ0AAAAAjyIRAwAAAAAAAAAAAAAAAAAAAAAAAAAAmYQBAAAAAAAQOgAAAAAAAA==",
		},
		nil, // len(logs3)==0
	}

	for i, logs := range []string{logs1, logs2, logs3} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			ll := strings.Split(strings.TrimSpace(logs), "\n")
			filtered, ok, err := svmutil.FilterDataLogs(ll, program)
			require.NoError(t, err)
			require.Equal(t, expected[i], filtered, "filtered logs should match expected for test %d", i)
			if i < 2 {
				require.True(t, ok, "logs should not be truncated for test %d", i)
			} else {
				require.False(t, ok, "logs should be truncated for test %d", i)
			}
		})
	}
}

var (
	logs1 = `
Program ComputeBudget111111111111111111111111111111 invoke [1]
Program ComputeBudget111111111111111111111111111111 success
Program ComputeBudget111111111111111111111111111111 invoke [1]
Program ComputeBudget111111111111111111111111111111 success
Program Evo1veLzNhWcEXFc4xv9nwoDg4MBoyo2N2j6nWWHQF69 invoke [1]
Program 9W959DqEETiGZocYWCQPaJ6sBmUzgfxXfqGeTEdp3aQP invoke [2]
Program log: Instruction: Swap
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [3]
Program log: Instruction: Transfer
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 4736 of 140794 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [3]
Program log: Instruction: MintTo
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 4492 of 111445 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [3]
Program log: Instruction: Transfer
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 4645 of 104015 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program 9W959DqEETiGZocYWCQPaJ6sBmUzgfxXfqGeTEdp3aQP consumed 58078 of 156925 compute units
Program 9W959DqEETiGZocYWCQPaJ6sBmUzgfxXfqGeTEdp3aQP success
Program whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc invoke [2]
Program log: Instruction: Swap
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [3]
Program log: Instruction: Transfer
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 4645 of 54084 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [3]
Program log: Instruction: Transfer
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 4736 of 46379 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program data: 4cpJr5MroJY4JTBvYKAWFg8AIX2IKcUTlfhh44l86K51YxwETLy2jgBGVr/24UJIeAQBAAAAAAAAtzspUdMcSDAGAQAAAAAAAKNS1475OgAAfWsHOAAAAAAAAAAAAAAAAAAAAAAAAAAAVfdHWYMAAACKoHegEwAAAA==
Program whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc consumed 57843 of 96057 compute units
Program whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc success
Program Evo1veLzNhWcEXFc4xv9nwoDg4MBoyo2N2j6nWWHQF69 consumed 122008 of 160104 compute units
Program Evo1veLzNhWcEXFc4xv9nwoDg4MBoyo2N2j6nWWHQF69 failed: custom program error: 0xdeadbeef`

	logs2 = `
Program ComputeBudget111111111111111111111111111111 invoke [1]
Program ComputeBudget111111111111111111111111111111 success
Program NA365bsPdvZ8sP58qJ5QFg7eXygCe8aPRRxR9oeMbR5 invoke [1]
Program whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc invoke [2]
Program log: Instruction: Swap
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [3]
Program log: Instruction: Transfer
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 4645 of 561687 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [3]
Program log: Instruction: Transfer
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 4645 of 553982 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program data: 4cpJr5MroJYpaAqbwx37a8ztWUYMupTp2tsZ/N8dkljA5maKm7tkWgD0lSG1lpAAQgAAAAAAAAAAojlOk/fjAEIAAAAAAAAAAH4JEQMAAAAAVa4gLgAAAAAAAAAAAAAAAAAAAAAAAAAAfREAAAAAAACcAgAAAAAAAA==
Program whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc consumed 50745 of 596561 compute units
Program whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc success
Program whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc invoke [2]
Program log: Instruction: Swap
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [3]
Program log: Instruction: Transfer
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 4645 of 506556 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [3]
Program log: Instruction: Transfer
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 4645 of 498854 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program data: 4cpJr5MroJYadt5TZ+XOA3IuI67qXNX56EwbmbA3icHAnZSH+TTBDQFN5Ow6wT9NkAAAAAAAAAAAfDWSdjVLSpAAAAAAAAAAAFWuIC4AAAAACNelDgAAAAAAAAAAAAAAAAAAAAAAAAAABCMFAAAAAAB/xAAAAAAAAA==
Program whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc consumed 52764 of 543498 compute units
Program whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc success
Program whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc invoke [2]
Program log: Instruction: Swap
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [3]
Program log: Instruction: Transfer
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 4645 of 456308 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [3]
Program log: Instruction: Transfer
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 4645 of 448603 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program data: 4cpJr5MroJZxeaOgaqQiDwwrJfgWEkHBny6g4OKYgjZki9Mpys2zBwC7UuCihSxdCQEAAAAAAAAAYtNAzYAtXQkBAAAAAAAAAAjXpQ4AAAAAXoShDQAAAAAAAAAAAAAAAAAAAAAAAAAAhVMAAAAAAAB6DAAAAAAAAA==
Program whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc consumed 47927 of 488364 compute units
Program whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc success
Program whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc invoke [2]
Program log: Instruction: Swap
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [3]
Program log: Instruction: Transfer
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 4645 of 402393 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [3]
Program log: Instruction: Transfer
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 4645 of 394691 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program data: 4cpJr5MroJaSGF30VKHk6aHXzz99JeTyA05hiqSVtbeRIIcOyEefqQEX6AmOjWl2eQAAAAAAAAAA/OPoQ2XLdXkAAAAAAAAAAF6EoQ0AAAAAjyIRAwAAAAAAAAAAAAAAAAAAAAAAAAAAmYQBAAAAAAAQOgAAAAAAAA==
Program whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc consumed 51541 of 438158 compute units
Program whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc success
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [2]
Program log: Instruction: Transfer
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 4644 of 384630 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program NA365bsPdvZ8sP58qJ5QFg7eXygCe8aPRRxR9oeMbR5 consumed 219874 of 599850 compute units
Program NA365bsPdvZ8sP58qJ5QFg7eXygCe8aPRRxR9oeMbR5 success
Program 11111111111111111111111111111111 invoke [1]
Program 11111111111111111111111111111111 success
`

	logs3 = `
Program ComputeBudget111111111111111111111111111111 invoke [1]
Program ComputeBudget111111111111111111111111111111 success
Program ComputeBudget111111111111111111111111111111 invoke [1]
Program ComputeBudget111111111111111111111111111111 success
Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL invoke [1]
Program log: CreateIdempotent
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [2]
Program log: Instruction: GetAccountDataSize
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 1569 of 923675 compute units
Program return: TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA pQAAAAAAAAA=
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program 11111111111111111111111111111111 invoke [2]
Program 11111111111111111111111111111111 success
Program log: Initialize the associated token account
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [2]
Program log: Instruction: InitializeImmutableOwner
Program log: Please upgrade to SPL Token 2022 for immutable owner support
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 1405 of 917088 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [2]
Program log: Instruction: InitializeAccount3
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 3158 of 913206 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL consumed 19315 of 929080 compute units
Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL success
Program 11111111111111111111111111111111 invoke [1]
Program 11111111111111111111111111111111 success
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [1]
Program log: Instruction: SyncNative
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 3045 of 909615 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL invoke [1]
Program log: CreateIdempotent
Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL consumed 4338 of 906570 compute units
Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL success
Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL invoke [1]
Program log: CreateIdempotent
Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL consumed 7438 of 902232 compute units
Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL success
Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL invoke [1]
Program log: CreateIdempotent
Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL consumed 4437 of 894794 compute units
Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL success
Program tuna4uSQZncNeeiAMKbstuxA9CUkHH6HmC64wgmnogD invoke [1]
Program log: Instruction: OpenPositionWithLiquidityOrca
Program 11111111111111111111111111111111 invoke [2]
Program 11111111111111111111111111111111 success
Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL invoke [2]
Program log: Create
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [3]
Program log: Instruction: GetAccountDataSize
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 1569 of 836764 compute units
Program return: TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA pQAAAAAAAAA=
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program 11111111111111111111111111111111 invoke [3]
Program 11111111111111111111111111111111 success
Program log: Initialize the associated token account
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [3]
Program log: Instruction: InitializeImmutableOwner
Program log: Please upgrade to SPL Token 2022 for immutable owner support
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 1405 of 830177 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [3]
Program log: Instruction: InitializeAccount3
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 3158 of 826293 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL consumed 25408 of 848239 compute units
Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL success
Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL invoke [2]
Program log: Create
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [3]
Program log: Instruction: GetAccountDataSize
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 1622 of 801219 compute units
Program return: TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA pQAAAAAAAAA=
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program 11111111111111111111111111111111 invoke [3]
Program 11111111111111111111111111111111 success
Program log: Initialize the associated token account
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [3]
Program log: Instruction: InitializeImmutableOwner
Program log: Please upgrade to SPL Token 2022 for immutable owner support
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 1405 of 794579 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [3]
Program log: Instruction: InitializeAccount3
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 4241 of 790695 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL consumed 25044 of 811194 compute units
Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL success
Program log: tick_lower=-17900, tick_upper=-17032
Program whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc invoke [2]
Program log: Instruction: OpenPositionWithTokenExtensions
Program 11111111111111111111111111111111 invoke [3]
Program 11111111111111111111111111111111 success
Program 11111111111111111111111111111111 invoke [3]
Program 11111111111111111111111111111111 success
Program TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb invoke [3]
Program log: Instruction: InitializeMintCloseAuthority
Program TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb consumed 1080 of 733251 compute units
Program TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb success
Program TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb invoke [3]
Program log: Instruction: InitializeMint2
Program TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb consumed 1963 of 729902 compute units
Program TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb success
Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL invoke [3]
Program log: Create
Program TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb invoke [4]
Program log: Instruction: GetAccountDataSize
Program TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb consumed 1370 of 710947 compute units
Program return: TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb qgAAAAAAAAA=
Program TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb success
Program 11111111111111111111111111111111 invoke [4]
Program 11111111111111111111111111111111 success
Program log: Initialize the associated token account
Program TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb invoke [4]
Program log: Instruction: InitializeImmutableOwner
Program TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb consumed 527 of 704647 compute units
Program TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb success
Program TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb invoke [4]
Program log: Instruction: InitializeAccount3
Program TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb consumed 1993 of 701731 compute units
Program TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb success
Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL consumed 19900 of 719334 compute units
Program ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL success
Program TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb invoke [3]
Program log: Instruction: MintTo
Program TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb consumed 1576 of 696269 compute units
Program TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb success
Program TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb invoke [3]
Program log: Instruction: SetAuthority
Program TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb consumed 1351 of 691934 compute units
Program TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb success
Program whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc consumed 59043 of 748149 compute units
Program whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc success
Program log: Collateral=[6100000000; 0], Borrow=[4559750000; 2374263843]
Program log: sqrt_price: 7692956142750524896; tick: -17493
Program log: Oracle/Spot price: 0.1739470777146977 / 0.17391901540414914; Deviation: 0.00016132671452284697
Program log: Protocol fees: [2584875; 1187131]
Program log: Transferring 6100000000 A-tokens from user to position (transfer fee included)
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [2]
Program log: Instruction: TransferChecked
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 6238 of 670898 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program log: Vault: So11111111111111111111111111111111111111112, total deposited f/s: 41821307599088 / 38892083958144, total borrowed f/s: 20555316727628 / 18065069261113, added interest: 34737683
Program log: Vault: EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v, total deposited f/s: 5441859025857 / 5108551313549, total borrowed f/s: 2094794823247 / 1854016125011, added interest: 3063504
Program log: Vault: So11111111111111111111111111111111111111112, total borrowed f/s: 20555351465311 / 18065069261113, borrow f/s: 4559750000 / 4007335983
Program log: Vault: EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v, total borrowed f/s: 2094797886751 / 1854016125011, borrow f/s: 2374263843 / 2101359505
Program log: Loan shares after borrowing: [4007335983; 2101359505]
Program log: Transferring 4559750000 A-tokens from vault to position (transfer fee included)
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA invoke [2]
Program log: Instruction: TransferChecked
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA consumed 6238 of 582176 compute units
Program TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA success
Program log: Transferring 2374263843 B-tokens from vault to position (transfer fee included)
Log truncated
Program 11111111111111111111111111111111 success
`
)
