package ethclient_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/omni-network/omni/lib/ethclient"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"

	"github.com/golang-jwt/jwt/v5"
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

func TestGetPayloadV2(t *testing.T) {
	t.Parallel()
	fuzzer := fuzz.New().NilChance(0)

	var param1 engine.PayloadID
	fuzzer.Fuzz(&param1)

	var resp engine.ExecutionPayloadEnvelope
	fuzzer.Fuzz(&resp)

	call := func(ctx context.Context, engineCl ethclient.EngineClient) (any, error) {
		return engineCl.GetPayloadV2(ctx, param1)
	}

	testEndpoint(t, call, resp, param1)
}

func TestGetPayloadV3(t *testing.T) {
	t.Parallel()
	fuzzer := fuzz.New().NilChance(0)

	var param1 engine.PayloadID
	fuzzer.Fuzz(&param1)

	var resp engine.ExecutionPayloadEnvelope
	fuzzer.Fuzz(&resp)

	call := func(ctx context.Context, engineCl ethclient.EngineClient) (any, error) {
		return engineCl.GetPayloadV3(ctx, param1)
	}

	testEndpoint(t, call, resp, param1)
}

func TestNewPayloadV2(t *testing.T) {
	t.Parallel()
	fuzzer := fuzz.New().NilChance(0)

	var param1 engine.ExecutableData
	fuzzer.Fuzz(&param1)

	var resp engine.PayloadStatusV1
	fuzzer.Fuzz(&resp)

	call := func(ctx context.Context, engineCl ethclient.EngineClient) (any, error) {
		return engineCl.NewPayloadV2(ctx, param1)
	}

	testEndpoint(t, call, resp, param1)
}

func TestNewPayloadV3(t *testing.T) {
	t.Parallel()
	fuzzer := fuzz.New().NilChance(0)

	var param1 engine.ExecutableData
	fuzzer.Fuzz(&param1)

	var param2 []common.Hash
	fuzzer.Fuzz(&param2)

	var param3 common.Hash
	fuzzer.Fuzz(&param3)

	var resp engine.PayloadStatusV1
	fuzzer.Fuzz(&resp)

	call := func(ctx context.Context, engineCl ethclient.EngineClient) (any, error) {
		return engineCl.NewPayloadV3(ctx, param1, param2, &param3)
	}

	testEndpoint(t, call, resp, param1, param2, param3)
}

func TestForkchoiceUpdatedV2(t *testing.T) {
	t.Parallel()
	fuzzer := fuzz.New().NilChance(0)

	var param1 engine.ForkchoiceStateV1
	fuzzer.Fuzz(&param1)

	var param2 engine.PayloadAttributes
	fuzzer.Fuzz(&param2)

	var resp engine.ForkChoiceResponse
	fuzzer.Fuzz(&resp)

	call := func(ctx context.Context, engineCl ethclient.EngineClient) (any, error) {
		return engineCl.ForkchoiceUpdatedV2(ctx, param1, &param2)
	}

	testEndpoint(t, call, resp, param1, param2)
}

func TestForkchoiceUpdatedV3(t *testing.T) {
	t.Parallel()
	fuzzer := fuzz.New().NilChance(0)

	var param1 engine.ForkchoiceStateV1
	fuzzer.Fuzz(&param1)

	var param2 engine.PayloadAttributes
	fuzzer.Fuzz(&param2)

	var resp engine.ForkChoiceResponse
	fuzzer.Fuzz(&resp)

	call := func(ctx context.Context, engineCl ethclient.EngineClient) (any, error) {
		return engineCl.ForkchoiceUpdatedV3(ctx, param1, &param2)
	}

	testEndpoint(t, call, resp, param1, param2)
}

func testEndpoint(t *testing.T, callback func(context.Context, ethclient.EngineClient) (any, error),
	resp any, params ...any,
) {
	t.Helper()

	const jwtSecret = "secret"

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse JWT from Authorization Bearer header.
		tokenString := r.Header.Get("Authorization")
		require.NotEmpty(t, tokenString)
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})
		require.NoError(t, err)

		require.Equal(t, "/", r.URL.Path)

		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)

		var rpcReq jsonRPCRequest
		err = json.Unmarshal(body, &rpcReq)
		require.NoError(t, err)

		for i, actualJSON := range rpcReq.Params {
			expectJSON, err := json.Marshal(params[i])
			require.NoError(t, err)

			require.Equal(t, string(expectJSON), string(actualJSON))
		}

		rpcResp := jsonRPCResponse{
			JSONRPC: "2.0",
			ID:      rpcReq.ID,
			Result:  resp,
		}
		buf, err := json.Marshal(rpcResp)
		require.NoError(t, err)

		_, _ = w.Write(buf)
	}))
	defer srv.Close()

	ctx := context.Background()

	api, err := ethclient.NewAuthClient(ctx, srv.URL, []byte(jwtSecret))
	require.NoError(t, err)

	got, err := callback(ctx, api)
	require.NoError(t, err)

	equalJSON(t, resp, got)
}

type jsonRPCRequest struct {
	JSONRPC string            `json:"jsonrpc"`
	ID      any               `json:"id"`
	Method  string            `json:"method"`
	Params  []json.RawMessage `json:"params"`
}

type jsonRPCResponse struct {
	JSONRPC string `json:"jsonrpc"`
	ID      any    `json:"id"`
	Result  any    `json:"result"`
}

// equalJSON asserts that two values are equal after marshaling to JSON.
func equalJSON(t *testing.T, a, b any) {
	t.Helper()

	aa, err := json.Marshal(a)
	require.NoError(t, err)

	bb, err := json.Marshal(b)
	require.NoError(t, err)

	require.Equal(t, string(aa), string(bb))
}
