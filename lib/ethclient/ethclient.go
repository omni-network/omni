package ethclient

import (
	"context"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/tracer"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

//go:generate go run genwrap/genwrap.go

var _ Client = Wrapper{}

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

// Wrapper wraps an ethclient.Client adding metrics and wrapped errors.
type Wrapper struct {
	cl      *ethclient.Client
	chain   string
	address string
}

// NewClient wraps an *rpc.Client adding metrics and wrapped errors.
func NewClient(cl *rpc.Client, chain, address string) Wrapper {
	return Wrapper{
		cl:      ethclient.NewClient(cl),
		chain:   chain,
		address: address,
	}
}

// Dial connects a client to the given URL.
//
// Note if the URL is http(s), it doesn't return an error if it cannot connect to the URL.
// It will retry connecting on every call to a wrapped method. It will only return an error if the
// url is invalid.
func Dial(chainName string, url string) (Wrapper, error) {
	cl, err := ethclient.Dial(url)
	if err != nil {
		return Wrapper{}, errors.Wrap(err, "dial", "chain", chainName, "url", url)
	}

	return Wrapper{
		cl:      cl,
		chain:   chainName,
		address: url,
	}, nil
}

// Close closes the underlying RPC connection.
func (w Wrapper) Close() {
	w.cl.Close()
}

// Address returns the underlying RPC address.
func (w Wrapper) Address() string {
	return w.address
}

// ProgressIfSyncing returns the sync progress (and true) if the node is syncing else false if node is
// not syncing or an error.
//
// This wrap SyncProgress which returns nil-nil if the node is not syncing which results in panics.
func (w Wrapper) ProgressIfSyncing(ctx context.Context) (*ethereum.SyncProgress, bool, error) {
	const endpoint = "sync_progress"
	defer latency(w.chain, endpoint)()

	resp, err := w.cl.SyncProgress(ctx)
	if err != nil {
		incError(w.chain, endpoint)
		return nil, false, errors.Wrap(err, "json-rpc", "endpoint", endpoint)
	} else if resp == nil {
		return nil, false, nil
	}

	return resp, true, nil
}

// HeaderByType returns the block header for the given head type.
func (w Wrapper) HeaderByType(ctx context.Context, typ HeadType) (*types.Header, error) {
	const endpoint = "header_by_type"
	defer latency(w.chain, endpoint)()

	var bn rpc.BlockNumber
	if err := bn.UnmarshalJSON([]byte(typ.String())); err != nil {
		return nil, errors.Wrap(err, "unmarshal head type")
	}

	header, err := w.cl.HeaderByNumber(ctx, bi.N(bn))
	if err != nil {
		incError(w.chain, endpoint)
		err = errors.Wrap(err, "json-rpc", "endpoint", endpoint)
	}

	return header, err
}

// SetHead sets the current head of the local chain by block number.
// Note, this is a destructive action and may severely damage your chain.
// Use with extreme caution.
func (w Wrapper) SetHead(ctx context.Context, height uint64) error {
	const endpoint = "set_head"
	defer latency(w.chain, endpoint)()

	err := w.cl.Client().CallContext(
		ctx,
		nil,
		"debug_setHead",
		height,
	)
	if err != nil {
		incError(w.chain, endpoint)
		return errors.Wrap(err, "json-rpc", "endpoint", endpoint)
	}

	return nil
}

// PeerCount returns the number of p2p peers as reported by the net_peerCount method.
func (w Wrapper) PeerCount(ctx context.Context) (uint64, error) {
	const endpoint = "peer_count"
	defer latency(w.chain, endpoint)()

	resp, err := w.cl.PeerCount(ctx)
	if err != nil {
		incError(w.chain, endpoint)
		return 0, errors.Wrap(err, "json-rpc", "endpoint", endpoint)
	}

	return resp, nil
}

// EtherBalanceAt returns the current balance in ether of the provided account.
// Note this converts big.Int to float64 so IS NOT accurate.
// Only use if accuracy is not required, i.e., for display/metrics purposes.
func (w Wrapper) EtherBalanceAt(ctx context.Context, addr common.Address) (float64, error) {
	b, err := w.BalanceAt(ctx, addr, nil)
	if err != nil {
		return 0, err
	}

	return bi.ToEtherF64(b), nil
}

// TxReceipt returns the transaction receipt for the given transaction hash.
// It includes additional fields not present in the geth ethclient, such as OP L1 fee info.
func (w Wrapper) TxReceipt(ctx context.Context, txHash common.Hash) (*Receipt, error) {
	const endpoint = "tx_receipt"
	defer latency(w.chain, endpoint)()

	ctx, span := tracer.Start(ctx, spanName(endpoint))
	defer span.End()

	var r *Receipt
	err := w.cl.Client().CallContext(ctx, &r, "eth_getTransactionReceipt", txHash)

	// mirror geth ethclient behavior
	if err == nil && r == nil {
		err = ethereum.NotFound
	}

	if err != nil {
		incError(w.chain, endpoint)
		err = errors.Wrap(err, "json-rpc", "endpoint", endpoint)
	}

	return r, err
}

//nolint:revive // interface{} required by upstream.
func (w Wrapper) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	const endpoint = "raw_call"
	defer latency(w.chain, endpoint)()

	err := w.cl.Client().CallContext(ctx, result, method, args...)
	if err != nil {
		incError(w.chain, endpoint)
		err = errors.Wrap(err, "json-rpc", "endpoint", endpoint)
	}

	return err
}
