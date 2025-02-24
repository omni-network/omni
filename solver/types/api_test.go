package types_test

import (
	"math/big"
	"testing"

	"github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/stretchr/testify/require"
)

func TestParseCalls(t *testing.T) {
	t.Parallel()

	calls := types.JSONCalls{
		{
			// selector + params
			Value:  ether1(),
			Target: common.HexToAddress("0x36e66fbbce51e4cd5bd3c62b637eb411b18949d4"),
			Data:   mustHexDecode("0x70a08231000000000000000000000000e3481474b23f88a8917dbcb4cbc55efcf0f68cc7"),
		},
		{
			// just selector
			Value:  ether1(),
			Target: common.HexToAddress("0x36e66fbbce51e4cd5bd3c62b637eb411b18949d4"),
			Data:   mustHexDecode("0x70a08231"),
		},
		{
			// no calldata
			Value:  ether1(),
			Target: common.HexToAddress("0x36e66fbbce51e4cd5bd3c62b637eb411b18949d4"),
			Data:   nil,
		},
		{
			// nil value
			Value:  nil,
			Target: common.HexToAddress("0x36e66fbbce51e4cd5bd3c62b637eb411b18949d4"),
			Data:   nil,
		},
	}

	parsed := calls.Parse()
	require.Len(t, parsed, 4)

	// full calldata
	require.Equal(t, calls[0].Value.ToInt(), parsed[0].Value)
	require.Equal(t, calls[0].Target, parsed[0].Target)
	require.Equal(t, ([]byte)(*calls[0].Data), parsed[0].Data)

	// just selector
	require.Equal(t, calls[1].Value.ToInt(), parsed[1].Value)
	require.Equal(t, calls[1].Target, parsed[1].Target)
	require.Equal(t, ([]byte)(*calls[1].Data), parsed[1].Data)

	// no calldata
	require.Equal(t, calls[2].Value.ToInt(), parsed[2].Value)
	require.Equal(t, calls[2].Target, parsed[2].Target)
	require.Equal(t, []byte(nil), parsed[2].Data)

	// nil value
	require.Equal(t, big.NewInt(0), parsed[3].Value)
	require.Equal(t, calls[3].Target, parsed[3].Target)
	require.Equal(t, []byte(nil), parsed[3].Data)
}

func TestParseExpenses(t *testing.T) {
	t.Parallel()

	// expenses
	expenses := types.JSONExpenses{
		{
			// full
			Spender: common.HexToAddress("0x36e66fbbce51e4cd5bd3c62b637eb411b18949d4"),
			Token:   common.HexToAddress("0x36e66fbbce51e4cd5bd3c62b637eb411b18949d4"),
			Amount:  ether1(),
		},
		{
			// no spender
			Token:  common.HexToAddress("0x36e66fbbce51e4cd5bd3c62b637eb411b18949d4"),
			Amount: ether1(),
		},
		{
			// no token
			Spender: common.HexToAddress("0x36e66fbbce51e4cd5bd3c62b637eb411b18949d4"),
			Amount:  ether1(),
		},
		{
			// no amount
			Spender: common.HexToAddress("0x36e66fbbce51e4cd5bd3c62b637eb411b18949d4"),
			Token:   common.HexToAddress("0x36e66fbbce51e4cd5bd3c62b637eb411b18949d4"),
			Amount:  nil,
		},
	}

	parsed := expenses.Parse()

	require.Len(t, parsed, 4)

	// full
	require.Equal(t, expenses[0].Spender, parsed[0].Spender)
	require.Equal(t, expenses[0].Token, parsed[0].Token)
	require.Equal(t, expenses[0].Amount.ToInt(), parsed[0].Amount)

	// no spender
	require.Equal(t, common.Address{}, parsed[1].Spender)
	require.Equal(t, expenses[1].Token, parsed[1].Token)
	require.Equal(t, expenses[1].Amount.ToInt(), parsed[1].Amount)

	// no token
	require.Equal(t, expenses[2].Spender, parsed[2].Spender)
	require.Equal(t, common.Address{}, parsed[2].Token)
	require.Equal(t, expenses[2].Amount.ToInt(), parsed[2].Amount)

	// no amount
	require.Equal(t, expenses[3].Spender, parsed[3].Spender)
	require.Equal(t, expenses[3].Token, parsed[3].Token)
	require.Equal(t, big.NewInt(0), parsed[3].Amount)
}

func TestParseDeposit(t *testing.T) {
	t.Parallel()

	// full
	deposit := types.JSONDeposit{
		Token:  common.HexToAddress("0x36e66fbbce51e4cd5bd3c62b637eb411b18949d4"),
		Amount: ether1(),
	}

	parsed := deposit.Parse()
	require.Equal(t, deposit.Token, parsed.Token)
	require.Equal(t, deposit.Amount.ToInt(), parsed.Amount)

	// no amount
	deposit = types.JSONDeposit{
		Token:  common.HexToAddress("0x36e66fbbce51e4cd5bd3c62b637eb411b18949d4"),
		Amount: nil,
	}

	parsed = deposit.Parse()
	require.Equal(t, deposit.Token, parsed.Token)
	require.Equal(t, big.NewInt(0), parsed.Amount)
}

func TestParseQuoteUnit(t *testing.T) {
	t.Parallel()

	// full
	quoteUnit := types.JSONQuoteUnit{
		Token:  common.HexToAddress("0x36e66fbbce51e4cd5bd3c62b637eb411b18949d4"),
		Amount: ether1(),
	}

	parsed := quoteUnit.Parse()
	require.Equal(t, quoteUnit.Token, parsed.Token)
	require.Equal(t, quoteUnit.Amount.ToInt(), parsed.Amount)

	back := parsed.ToJSON()
	require.Equal(t, quoteUnit, back)

	// no amount
	quoteUnit = types.JSONQuoteUnit{
		Token:  common.HexToAddress("0x36e66fbbce51e4cd5bd3c62b637eb411b18949d4"),
		Amount: nil,
	}

	parsed = quoteUnit.Parse()
	require.Equal(t, quoteUnit.Token, parsed.Token)
	require.Equal(t, big.NewInt(0), parsed.Amount)

	back = parsed.ToJSON()
	require.Equal(t, quoteUnit.Token, back.Token)
	require.Equal(t, (*hexutil.Big)(big.NewInt(0)), back.Amount) // amount is set from nil to zero on parse
}

func ether1() *hexutil.Big {
	b := new(big.Int).Mul(big.NewInt(1), big.NewInt(1e18))
	return (*hexutil.Big)(b)
}

func mustHexDecode(s string) *hexutil.Bytes {
	b := hexutil.MustDecode(s)
	return (*hexutil.Bytes)(&b)
}
