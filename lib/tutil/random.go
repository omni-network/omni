package tutil

import (
	"crypto/rand"
	"fmt"
	mrand2 "math/rand/v2"
	"net"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

// RandomBytes returns a random byte slice of length l.
func RandomBytes(l int) []byte {
	resp := make([]byte, l)
	_, _ = rand.Read(resp)

	return resp
}

// RandomHash returns a random 32-byte 256-bit hash.
func RandomHash() common.Hash {
	var resp common.Hash
	_, _ = rand.Read(resp[:])

	return resp
}

// RandomAddress returns a random 20-byte ethereum address.
func RandomAddress() common.Address {
	var resp common.Address
	_, _ = rand.Read(resp[:])

	return resp
}

func RandomListenAddress(t *testing.T) string {
	t.Helper()
	return fmt.Sprintf("tcp://127.0.0.1:%d", RandomAvailablePort(t))
}

func RandomAvailablePort(t *testing.T) int {
	t.Helper()

	// randPort returns between base and base + spread
	randPort := func() int {
		base, spread := 20000, 20000
		return base + mrand2.IntN(spread)
	}

	// Pick a random port, since we can't "reserve" it.
	for i := 0; i < 5; i++ {
		port := randPort()
		l, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
		if err != nil {
			continue
		}
		if err := l.Close(); err != nil {
			continue
		}

		return port
	}

	t.Fatalf("failed to find available port")

	return 0
}
