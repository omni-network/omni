package ethp2p

import (
	"context"
	"net"
	"net/netip"
	"time"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/enr"
)

// DNSResolveHostname updates the given node from its DNS hostname.
// This is used to resolve static dial targets.
func DNSResolveHostname(ctx context.Context, n *enode.Node) (*enode.Node, error) {
	if n.Hostname() == "" {
		return n, nil
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	foundIPs, err := net.DefaultResolver.LookupNetIP(ctx, "ip", n.Hostname())
	if err != nil {
		return n, errors.Wrap(err, "dns lookup", "hostname", n.Hostname())
	}

	// Check for IP updates.
	var (
		nodeIP4, nodeIP6   netip.Addr
		foundIP4, foundIP6 netip.Addr
	)
	_ = n.Load((*enr.IPv4Addr)(&nodeIP4))
	_ = n.Load((*enr.IPv6Addr)(&nodeIP6))
	for _, ip := range foundIPs {
		if ip.Is4() && !foundIP4.IsValid() {
			foundIP4 = ip
		}
		if ip.Is6() && !foundIP6.IsValid() {
			foundIP6 = ip
		}
	}

	if !foundIP4.IsValid() && !foundIP6.IsValid() {
		// Lookup failed.
		return n, errors.New("dns lookup failed", "hostname", n.Hostname())
	}

	if foundIP4 == nodeIP4 && foundIP6 == nodeIP6 {
		// No updates necessary.
		return n, nil
	}

	// Update the node. Note this invalidates the ENR signature, because we use SignNull
	// to create a modified copy. But this should be OK, since we just use the node as a
	// dial target. And nodes will usually only have a DNS hostname if they came from a
	// enode:// URL, which has no signature anyway. If it ever becomes a problem, the
	// resolved IP could also be stored into dialTask instead of the node.
	rec := n.Record()
	if foundIP4.IsValid() {
		rec.Set(enr.IPv4Addr(foundIP4))
	}
	if foundIP6.IsValid() {
		rec.Set(enr.IPv6Addr(foundIP6))
	}
	rec.SetSeq(n.Seq()) // ensure seq not bumped by update
	newNode := enode.SignNull(rec, n.ID()).WithHostname(n.Hostname())

	return newNode, nil
}
