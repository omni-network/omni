package ethclient

import (
	"context"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
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

// HeaderByType returns the block header for the given head type.
func (w Wrapper) HeaderByType(ctx context.Context, typ HeadType) (*types.Header, error) {
	const endpoint = "header_by_type"
	defer latency(w.chain, endpoint)()

	var header *types.Header
	err := w.cl.Client().CallContext(
		ctx,
		&header,
		"eth_getBlockByNumber",
		typ.String(),
		false,
	)
	if err != nil {
		incError(w.chain, endpoint)
		return nil, errors.Wrap(err, "get block")
	}

	return header, nil
}

// PeerCount returns the number of p2p peers as reported by the net_peerCount method.
func (w Wrapper) PeerCount(ctx context.Context) (uint64, error) {
	const endpoint = "peer_count"
	defer latency(w.chain, endpoint)()

	resp, err := w.cl.PeerCount(ctx)
	if err != nil {
		incError(w.chain, endpoint)
		return 0, errors.Wrap(err, "rpc get payload v3")
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

	bf, _ := b.Float64()

	return bf / params.Ether, nil
}
