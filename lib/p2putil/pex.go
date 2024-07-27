package p2putil

import (
	"context"
	"fmt"

	haloapp "github.com/omni-network/omni/halo/app"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	cmtconfig "github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/crypto/ed25519"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/p2p/conn"
	"github.com/cometbft/cometbft/p2p/pex"
	tmp2p "github.com/cometbft/cometbft/proto/tendermint/p2p"
	"github.com/cometbft/cometbft/version"
)

// FetchPexAddrs fetches a list of cometBFT P2P peers from the provided peer using the PEX protocol.
func FetchPexAddrs(ctx context.Context, network netconf.ID, peer *p2p.NetAddress) ([]*p2p.NetAddress, error) {
	sw, err := makeSwitch(ctx, network)
	if err != nil {
		return nil, err
	}

	receive := make(chan p2p.Envelope, 1)
	sw.AddReactor("pex", pexReactor{receive: receive})

	if err := sw.Start(); err != nil {
		return nil, errors.Wrap(err, "start switch")
	}
	defer sw.Stop() //nolint:errcheck // Not critical.

	if err := sw.DialPeerWithAddress(peer); err != nil {
		return nil, errors.Wrap(err, "dial peer", "peer", peer)
	}

	for {
		select {
		case <-ctx.Done():
			return nil, errors.Wrap(ctx.Err(), "canceled")
		case e := <-receive:
			msg, ok := e.Message.(*tmp2p.PexAddrs)
			if !ok {
				log.Debug(ctx, "Ignoring message", "msg", fmt.Sprintf("%T", e.Message))
				continue
			}

			addrs, err := p2p.NetAddressesFromProto(msg.Addrs)
			if err != nil {
				return nil, errors.Wrap(err, "parse addresses")
			}

			return addrs, nil
		}
	}
}

// makeSwitch creates a new cometBFT p2p switch.
// This was copied from cometbft/p2p/test_util.go.
func makeSwitch(ctx context.Context, network netconf.ID) (*p2p.Switch, error) {
	cfg := cmtconfig.DefaultP2PConfig()
	nodeKey := p2p.NodeKey{PrivKey: ed25519.GenPrivKey()}
	nodeInfo := p2p.DefaultNodeInfo{
		ProtocolVersion: p2p.NewProtocolVersion(
			version.P2PProtocol,
			version.BlockProtocol,
			0,
		),
		DefaultNodeID: nodeKey.ID(),
		ListenAddr:    "127.0.0.1:0",
		Network:       network.Static().OmniConsensusChainIDStr(),
		Version:       "1.2.3",
		Channels:      []byte{pex.PexChannel},
		Moniker:       "omni/p2putil/fetch/pex",
	}

	tranport := p2p.NewMultiplexTransport(nodeInfo, nodeKey, p2p.MConnConfig(cfg))

	cmtlog, err := haloapp.NewCmtLogger(ctx, "debug")
	if err != nil {
		return nil, err
	}

	sw := p2p.NewSwitch(cfg, tranport)
	sw.SetNodeKey(&nodeKey)
	sw.SetNodeInfo(nodeInfo)
	sw.SetLogger(cmtlog)

	return sw, nil
}

// pexReactor is a very simple reactor that sends a PEX request
// to a peer and forwards to response on the receive channel.
type pexReactor struct {
	p2p.Reactor
	receive chan p2p.Envelope
}

func (pexReactor) AddPeer(peer p2p.Peer) {
	// Immediately send a PEX request to the peer.
	peer.Send(p2p.Envelope{
		ChannelID: pex.PexChannel,
		Message:   &tmp2p.PexRequest{},
	})
}

func (p pexReactor) Receive(e p2p.Envelope) {
	select {
	case p.receive <- e:
	default:
	}
}

func (pexReactor) GetChannels() []*conn.ChannelDescriptor {
	return []*conn.ChannelDescriptor{
		{
			ID:                  pex.PexChannel,
			Priority:            1,
			SendQueueCapacity:   10,
			RecvMessageCapacity: 256 * 256,
			MessageType:         &tmp2p.Message{},
		},
	}
}

func (pexReactor) RemovePeer(p2p.Peer, any)        {}
func (pexReactor) SetSwitch(*p2p.Switch)           {}
func (pexReactor) InitPeer(peer p2p.Peer) p2p.Peer { return peer }
func (pexReactor) OnStart() error                  { return nil }
func (pexReactor) Start() error                    { return nil }
func (pexReactor) OnStop()                         {}
func (pexReactor) Stop() error                     { return nil }
