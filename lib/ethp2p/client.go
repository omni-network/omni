package ethp2p

import (
	"context"
	"crypto/ecdsa"
	"math/rand"
	"net"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/umath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/protocols/eth"
	"github.com/ethereum/go-ethereum/eth/protocols/snap"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/rlpx"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/holiman/uint256"
)

var (
	capEth  = p2p.Cap{Name: eth.ProtocolName, Version: eth.ETH68}
	capSnap = p2p.Cap{Name: snap.ProtocolName, Version: snap.SNAP1}
)

// Client represents an individual connection with a remote P2P Peer.
// This was adapted from go-ethereum/cmd/devp2p/internal/ethtest.
type Client struct {
	conn *rlpx.Conn
	key  *ecdsa.PrivateKey
}

// Dial dials the provided destination and performs required handshakes and exchanges using the given
// private key returning a Client to and the status of that instance.
func Dial(ctx context.Context, key *ecdsa.PrivateKey, dest *enode.Node) (Client, *eth.StatusPacket, error) {
	tcpEndpoint, ok := dest.TCPEndpoint()
	if !ok {
		return Client{}, nil, errors.New("invalid TCP endpoint", "endpoint", dest.String())
	}

	fd, err := net.Dial("tcp", tcpEndpoint.String())
	if err != nil {
		return Client{}, nil, errors.Wrap(err, "dial TCP connection", "endpoint", tcpEndpoint.String())
	}

	cl := Client{
		conn: rlpx.NewConn(fd, dest.Pubkey()),
		key:  key,
	}

	_, err = cl.conn.Handshake(key)
	if err != nil {
		return Client{}, nil, errors.Wrap(err, "rlpx handshake")
	}

	if err := cl.negotiate(ctx); err != nil {
		return Client{}, nil, errors.Wrap(err, "protocol handshake")
	}

	status, err := cl.statusExchange(ctx)
	if err != nil {
		return Client{}, nil, errors.Wrap(err, "status exchange")
	}

	return cl, status, nil
}

// readMsg attempts to read a devp2p message with a specific code.
func (c Client) readMsg(ctx context.Context, proto proto, code uint64, msg any) error {
	for {
		got, data, _, err := c.conn.Read()
		if err != nil {
			return errors.Wrap(err, "read message from connection")
		}

		if got == proto.MsgCode(code) {
			if err := rlp.DecodeBytes(data, msg); err != nil {
				return errors.Wrap(err, "decode RLP message", "proto", proto, "code", code)
			}

			return nil
		} else if got == protoBase.MsgCode(discMsg) {
			var reasons []p2p.DiscReason
			_ = rlp.DecodeBytes(data, &reasons)

			return errors.New("disconnect received", "reasons", reasons)
		}

		gotProto := parseProto(got)
		gotMsg := umath.SubtractOrZero(got, gotProto.Offset())

		log.Debug(ctx, "Dropping ethp2p message", "proto", gotProto, "code", gotMsg)
	}
}

// writeMsg writes a eth protocol message to the connection.
func (c Client) writeMsg(proto proto, code uint64, msg any) error {
	payload, err := rlp.EncodeToBytes(msg)
	if err != nil {
		return errors.Wrap(err, "encode RLP message", "proto", proto, "code", code)
	}

	_, err = c.conn.Write(proto.MsgCode(code), payload)
	if err != nil {
		return errors.Wrap(err, "write message", "proto", proto, "code", code)
	}

	return nil
}

// negotiate performs a protocol handshake with the node.
func (c Client) negotiate(ctx context.Context) error {
	shake := &protoHandshake{
		Version: 5,
		Caps:    []p2p.Cap{capEth, capSnap},
		ID:      crypto.FromECDSAPub(&c.key.PublicKey)[1:],
	}

	if err := c.writeMsg(protoBase, handshakeMsg, shake); err != nil {
		return errors.Wrap(err, "write handshake")
	}

	var msg protoHandshake
	if err := c.readMsg(ctx, protoBase, handshakeMsg, &msg); err != nil {
		return errors.Wrap(err, "read handshake message")
	}

	if msg.Version >= 5 {
		c.conn.SetSnappy(true)
	}

	var supportEth, supportSnap bool
	for _, c := range msg.Caps {
		if c == capEth {
			supportEth = true
		} else if c == capSnap {
			supportSnap = true
		}
	}

	if !supportEth {
		return errors.New("peer does not support eth protocol", "caps", msg.Caps)
	} else if !supportSnap {
		return errors.New("peer does not support snap protocol", "caps", msg.Caps)
	}

	return nil
}

func (c Client) Disconnect() error {
	if err := c.writeMsg(protoBase, discMsg, []p2p.DiscReason{p2p.DiscQuitting}); err != nil {
		return errors.Wrap(err, "write disconnect")
	}

	if err := c.conn.Close(); err != nil {
		return errors.Wrap(err, "close connection")
	}

	return nil
}

// statusExchange performs a `Status` message exchange with the given node.
// Note that statuses are not requested, but rather just exchanged after handshake.
func (c Client) statusExchange(ctx context.Context) (*eth.StatusPacket, error) {
	var received eth.StatusPacket
	if err := c.readMsg(ctx, protoEth, eth.StatusMsg, &received); err != nil {
		return nil, errors.Wrap(err, "read status request")
	}

	// Echo response
	if err := c.writeMsg(protoEth, eth.StatusMsg, received); err != nil {
		return nil, errors.Wrap(err, "write status")
	}

	return &received, nil
}

// HeadersDownFrom requests a sequence of block headers starting from the specified head hash, going downwards.
// Results are ordered by height descending.
func (c Client) HeadersDownFrom(ctx context.Context, blockHash common.Hash, count uint64) ([]*types.Header, error) {
	headerReq := &eth.GetBlockHeadersPacket{
		RequestId: uint64(rand.Int63()), //nolint:gosec // Weak random request ID
		GetBlockHeadersRequest: &eth.GetBlockHeadersRequest{
			Origin:  eth.HashOrNumber{Hash: blockHash},
			Amount:  count,
			Skip:    0,
			Reverse: true,
		},
	}

	if err := c.writeMsg(protoEth, eth.GetBlockHeadersMsg, headerReq); err != nil {
		return nil, errors.Wrap(err, "write GetBlockHeaders")
	}

	resp := new(eth.BlockHeadersPacket)
	if err := c.readMsg(ctx, protoEth, eth.BlockHeadersMsg, &resp); err != nil {
		return nil, errors.Wrap(err, "read BlockHeaders")
	} else if len(resp.BlockHeadersRequest) == 0 {
		return nil, errors.New("no headers received")
	}

	return resp.BlockHeadersRequest, nil
}

func (c Client) AccountRange(ctx context.Context, root, origin common.Hash, bytes uint64) (*snap.AccountRangePacket, error) {
	accReq := &snap.GetAccountRangePacket{
		ID:     uint64(rand.Int63()), //nolint:gosec // Weak random request ID
		Root:   root,
		Origin: origin,
		Limit:  common.MaxHash,
		Bytes:  bytes,
	}

	if err := c.writeMsg(protoSnap, snap.GetAccountRangeMsg, accReq); err != nil {
		return nil, errors.Wrap(err, "write GetAccountRange")
	}

	accResp := new(snap.AccountRangePacket)
	if err := c.readMsg(ctx, protoSnap, snap.AccountRangeMsg, accResp); err != nil {
		return nil, errors.Wrap(err, "read AccountRange")
	}

	return accResp, nil
}

// SnapshotRange returns the number of snapshots available down from the provided block hash.
func (c Client) SnapshotRange(ctx context.Context, blockHash common.Hash, max uint64) (int, error) {
	headers, err := c.HeadersDownFrom(ctx, blockHash, max)
	if err != nil {
		return 0, errors.Wrap(err, "fetch headers")
	}

	var resp int
	for _, header := range headers {
		acc, err := c.AccountRange(ctx, header.Root, common.Hash{}, 256)
		if err != nil {
			return 0, errors.Wrap(err, "fetch account range")
		} else if len(acc.Accounts) == 0 {
			break
		}
		resp++
	}

	return resp, nil
}

// AllAccountRanges returns all account ranges starting of the provided block hash (not state root).
func (c Client) AllAccountRanges(ctx context.Context, blockHash common.Hash, batchBytes uint64) ([]*snap.AccountRangePacket, error) {
	headers, err := c.HeadersDownFrom(ctx, blockHash, 1)
	if err != nil {
		return nil, errors.Wrap(err, "fetch headers")
	} else if len(headers) == 0 {
		return nil, errors.New("no headers found")
	}
	header := headers[0]

	var resp []*snap.AccountRangePacket
	var next common.Hash
	for {
		acc, err := c.AccountRange(ctx, header.Root, next, batchBytes)
		if err != nil {
			return nil, errors.Wrap(err, "fetch account range")
		} else if len(acc.Accounts) == 0 {
			break
		}

		resp = append(resp, acc)
		next = incHash(acc.Accounts[len(acc.Accounts)-1].Hash) // Increment to the next account key
	}

	if len(resp) == 0 {
		return nil, errors.New("no accounts found")
	}

	return resp, nil
}

// incHash returns the next hash, in lexicographical order (a.k.a plus one).
// Note it rolls over for common.MaxHash, returning zero hash.
func incHash(h common.Hash) common.Hash {
	return new(uint256.Int).AddUint64(
		new(uint256.Int).SetBytes32(h[:]),
		1,
	).Bytes32()
}
