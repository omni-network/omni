// This file rewrites some of Hyperliquid's signing utils tests in golang.
// ref: https://github.com/hyperliquid-dex/hyperliquid-python-sdk/blob/a8edca1/tests/signing_test.py
package hypercore

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/stretchr/testify/require"
)

// ref: https://github.com/hyperliquid-dex/hyperliquid-python-sdk/blob/a8edca1/tests/signing_test.py#L37
func TestL1ActionSigningMatches(t *testing.T) {
	t.Parallel()

	privateKey, err := crypto.HexToECDSA("0123456789012345678901234567890123456789012345678901234567890123")
	require.NoError(t, err)
	signer := NewPrivateKeySigner(privateKey)

	// Struct required to enforce msgpack encoding order
	type DummyAction struct {
		Type string `json:"type"`
		Num  uint64 `json:"num"`
	}

	action := DummyAction{
		Type: "dummy",
		Num:  100000000000,
	}

	ctx := t.Context()

	signatureMainnet, err := signL1Action(ctx, signer, action, emptyAddress, 0, 0, true)
	require.NoError(t, err)

	// nOTE: Hyperliquid python sdk trims leading zeros - so we do too.
	require.Equal(t, "0x53749d5b30552aeb2fca34b530185976545bb22d0b3ce6f62e31be961a59298", signatureMainnet.R)
	require.Equal(t, "0x755c40ba9bf05223521753995abb2f73ab3229be8ec921f350cb447e384d8ed8", signatureMainnet.S)
	require.Equal(t, uint8(27), signatureMainnet.V)

	signatureTestnet, err := signL1Action(ctx, signer, action, emptyAddress, 0, 0, false)
	require.NoError(t, err)

	require.Equal(t, "0x542af61ef1f429707e3c76c5293c80d01f74ef853e34b76efffcb57e574f9510", signatureTestnet.R)
	require.Equal(t, "0x17b8b32f086e8cdede991f1e2c529f5dd5297cbe8128500e00cbaf766204a613", signatureTestnet.S)
	require.Equal(t, uint8(28), signatureTestnet.V)
}

type OrderAction struct {
	Type     string      `json:"type"`
	Orders   []OrderWire `json:"orders"`
	Grouping string      `json:"grouping"`
}

type OrderWire struct {
	A int           `json:"a"`
	B bool          `json:"b"`
	P string        `json:"p"`
	S string        `json:"s"`
	R bool          `json:"r"`
	T OrderTypeWire `json:"t"`
	C string        `json:"c,omitempty"`
}

type OrderTypeWire struct {
	Limit   LimitOrderType       `json:"limit,omitempty"`
	Trigger TriggerOrderTypeWire `json:"trigger,omitempty"`
}

type TriggerOrderTypeWire struct {
	TriggerPx string `json:"triggerPx"`
	IsMarket  bool   `json:"isMarket"`
	TPSL      string `json:"tpsl"` // "tp" | "sl"
}

type LimitOrderType struct {
	TIF string `json:"tif"` // "Alo" | "Ioc" | "Gtc"
}

// NOTE: order wiring was omitted, preferring to inline order action (wiring code has not been copied over).
func TestPhantomAgentCreationMatchesProductino(t *testing.T) {
	t.Parallel()

	timestamp := uint64(1677777606040)
	action := OrderAction{
		Type: "order",
		Orders: []OrderWire{
			{
				A: 4,
				B: true,
				P: "1670.1",
				S: "0.0147",
				R: false,
				T: OrderTypeWire{
					Limit: LimitOrderType{
						TIF: "Ioc",
					},
				},
			},
		},
		Grouping: "na",
	}

	hash, err := actionHash(action, common.Address{}, timestamp, 0)
	require.NoError(t, err)

	agent := phantomAgent(hash, true)
	require.Equal(t, "0x0fcbeda5ae3c4950a548021552a4fea2226858c4453571bf3f24ba017eac2908", hexutil.Encode(agent.ConnectionID[:]))
}
