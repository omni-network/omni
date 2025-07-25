package ethclient

import (
	"context"
	"net/http"
	"time"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/tracer"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

//go:generate go run genwrap/genwrap.go

var _ Client = wrapper{}

type HeadType string

func (h HeadType) String() string {
	return string(h)
}

func (h HeadType) Verify() error {
	if !allHeadTypes[h] {
		return errors.New("invalid head type", "head", h)
	}

	return nil
}

//nolint:gochecknoglobals // Static mappings
var allHeadTypes = map[HeadType]bool{
	HeadLatest:    true,
	HeadEarliest:  true,
	HeadPending:   true,
	HeadSafe:      true,
	HeadFinalized: true,
}

const (
	HeadLatest    HeadType = "latest"
	HeadEarliest  HeadType = "earliest"
	HeadPending   HeadType = "pending"
	HeadSafe      HeadType = "safe"
	HeadFinalized HeadType = "finalized"
)

// wrapper wraps an ethclient.Client adding metrics and wrapped errors.
type wrapper struct {
	cl      *ethclient.Client
	httpCl  *http.Client
	name    string
	address string
}

// NewClient wraps an *rpc.Client adding metrics and wrapped errors and a header cache.
func NewClient(cl *rpc.Client, name, address string) (Client, error) {
	return newHeaderCache(wrapper{
		cl:      ethclient.NewClient(cl),
		name:    name,
		address: address,
	})
}

// Dial calls DialContext with a background context.
// See DialContext for more details.
func Dial(chainName string, url string) (Client, error) {
	return DialContext(context.Background(), chainName, url)
}

// DialContext connects a client to the given URL. It returns a wrapped client adding metrics and wrapped errors and a header cache.
//
// Note if the URL is http(s), it doesn't return an error if it cannot connect to the URL.
// It will retry connecting on every call to a wrapped method. In this case, the context is ignored.
// It will only return an error if the url is invalid.
func DialContext(ctx context.Context, chainName string, url string) (Client, error) {
	client := &http.Client{Timeout: defaultRPCHTTPTimeout}

	rpcClient, err := rpc.DialOptions(ctx, url, rpc.WithHTTPClient(client))
	if err != nil {
		return engineClient{}, errors.Wrap(err, "dial", "chain", chainName, "url", url)
	}

	return newHeaderCache(wrapper{
		cl:      ethclient.NewClient(rpcClient),
		name:    chainName,
		address: url,
		httpCl:  client,
	})
}

// CloseIdleConnectionsForever blocks and closes idle connections periodically.
// It returns when the context is canceled.
// This is useful to close TCP-keep-alive connections to load-balanced RPC servers
// which could sometimes remain connected to stalled (but alive) servers.
// Note this is a noop if the client wasn't created by Dial or DialContext.
func (w wrapper) CloseIdleConnectionsForever(ctx context.Context) {
	if w.httpCl == nil {
		return
	}

	const period = time.Minute * 5

	ticker := time.NewTicker(period)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			w.httpCl.CloseIdleConnections()
		}
	}
}

// Close closes the underlying RPC connection.
func (w wrapper) Close() {
	w.cl.Close()
}

// Address returns the underlying RPC address.
func (w wrapper) Address() string {
	return w.address
}

// Name returns the client or chain name.
func (w wrapper) Name() string {
	return w.name
}

// ProgressIfSyncing returns the sync progress (and true) if the node is syncing else false if node is
// not syncing or an error.
//
// This wrap SyncProgress which returns nil-nil if the node is not syncing which results in panics.
func (w wrapper) ProgressIfSyncing(ctx context.Context) (*ethereum.SyncProgress, bool, error) {
	const endpoint = "sync_progress"
	defer latency(w.name, endpoint)()

	resp, err := w.cl.SyncProgress(ctx)
	if err != nil {
		incError(w.name, endpoint)
		return nil, false, errors.Wrap(err, "json-rpc", "endpoint", endpoint)
	} else if resp == nil {
		return nil, false, nil
	}

	return resp, true, nil
}

// HeaderByType returns the block header for the given head type.
func (w wrapper) HeaderByType(ctx context.Context, typ HeadType) (*types.Header, error) {
	const endpoint = "header_by_type"
	defer latency(w.name, endpoint)()

	var bn rpc.BlockNumber
	if err := bn.UnmarshalJSON([]byte(typ.String())); err != nil {
		return nil, errors.Wrap(err, "unmarshal head type")
	}

	header, err := w.cl.HeaderByNumber(ctx, bi.N(bn))
	if err != nil {
		incError(w.name, endpoint)
		err = errors.Wrap(err, "json-rpc", "endpoint", endpoint)
	}

	return header, err
}

// SetHead sets the current head of the local chain by block number.
// Note, this is a destructive action and may severely damage your chain.
// Use with extreme caution.
func (w wrapper) SetHead(ctx context.Context, height uint64) error {
	const endpoint = "set_head"
	defer latency(w.name, endpoint)()

	err := w.cl.Client().CallContext(
		ctx,
		nil,
		"debug_setHead",
		height,
	)
	if err != nil {
		incError(w.name, endpoint)
		return errors.Wrap(err, "json-rpc", "endpoint", endpoint)
	}

	return nil
}

// PeerCount returns the number of p2p peers as reported by the net_peerCount method.
func (w wrapper) PeerCount(ctx context.Context) (uint64, error) {
	const endpoint = "peer_count"
	defer latency(w.name, endpoint)()

	resp, err := w.cl.PeerCount(ctx)
	if err != nil {
		incError(w.name, endpoint)
		return 0, errors.Wrap(err, "json-rpc", "endpoint", endpoint)
	}

	return resp, nil
}

// EtherBalanceAt returns the current balance in ether of the provided account.
// Note this converts big.Int to float64 so IS NOT accurate.
// Only use if accuracy is not required, i.e., for display/metrics purposes.
func (w wrapper) EtherBalanceAt(ctx context.Context, addr common.Address) (float64, error) {
	b, err := w.BalanceAt(ctx, addr, nil)
	if err != nil {
		return 0, err
	}

	return bi.ToEtherF64(b), nil
}

// TxReceipt returns the transaction receipt for the given transaction hash.
// It includes additional fields not present in the geth ethclient, such as OP L1 fee info.
func (w wrapper) TxReceipt(ctx context.Context, txHash common.Hash) (*Receipt, error) {
	const endpoint = "tx_receipt"
	defer latency(w.name, endpoint)()

	ctx, span := tracer.Start(ctx, spanName(endpoint))
	defer span.End()

	var r *Receipt
	err := w.cl.Client().CallContext(ctx, &r, "eth_getTransactionReceipt", txHash)

	// mirror geth ethclient behavior
	if err == nil && r == nil {
		err = ethereum.NotFound
	}

	if err != nil {
		incError(w.name, endpoint)
		err = errors.Wrap(err, "json-rpc", "endpoint", endpoint)
	}

	return r, err
}

func (w wrapper) Preimage(ctx context.Context, hash common.Hash) (hexutil.Bytes, error) {
	const endpoint = "preimage"
	defer latency(w.name, endpoint)()

	ctx, span := tracer.Start(ctx, spanName(endpoint))
	defer span.End()

	var preimage hexutil.Bytes
	err := w.cl.Client().CallContext(ctx, &preimage, "debug_preimage", hash)
	if err != nil {
		incError(w.name, endpoint)
		err = errors.Wrap(err, "json-rpc", "endpoint", endpoint)
	}

	return preimage, err
}

//nolint:revive // interface{} required by upstream.
func (w wrapper) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	const endpoint = "raw_call"
	defer latency(w.name, endpoint)()

	err := w.cl.Client().CallContext(ctx, result, method, args...)
	if err != nil {
		incError(w.name, endpoint)
		err = errors.Wrap(err, "json-rpc", "endpoint", endpoint)
	}

	return err
}
