package p2putil_test

import (
	"net"
	"os"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/p2p/enode"

	"github.com/stretchr/testify/require"
)

// TestDiscv4 fetches and prints discovered peers from the specified network execution seeds.
func TestDiscv4(t *testing.T) {
	t.Parallel()
	if !*integration {
		t.Skip("skipping integration test")
	}

	log.SetDefault(log.NewLogger(log.LogfmtHandlerWithLevel(os.Stdout, log.LevelDebug)))

	var bootnodes []*enode.Node
	for _, seed := range netconf.Omega.Static().ExecutionSeeds() {
		bootnode, err := enode.ParseV4(seed)
		require.NoError(t, err)
		bootnodes = append(bootnodes, bootnode)
	}

	nodeKey, err := crypto.GenerateKey()
	require.NoError(t, err)
	db, err := enode.OpenDB("")
	require.NoError(t, err)
	node := enode.NewLocalNode(db, nodeKey)

	addr, err := net.ResolveUDPAddr("udp", ":")
	require.NoError(t, err)
	conn, err := net.ListenUDP("udp", addr)
	require.NoError(t, err)

	discv4, err := discover.ListenV4(conn, node, discover.Config{
		PrivateKey: nodeKey,
		Bootnodes:  bootnodes,
		Log:        log.Root(),
	})
	require.NoError(t, err)

	// TODO(corver): Figure out how to show discovered peers.
	// In the mean time, add logging to github.com/ethereum/go-ethereum@v1.14.8/p2p/discover/v4_udp.go::UDPv4.findnode
	for range 5 {
		iter := discv4.RandomNodes() // This returns 0 always :(
		var ips []string
		for {
			n := iter.Node()
			if n == nil {
				break
			}
			ips = append(ips, n.IP().String())
			if !iter.Next() {
				break
			}
		}

		log.Info("discovered nodes", "count", len(ips), "ips", ips)
		time.Sleep(time.Second * 2)
	}
}
