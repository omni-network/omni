package relayer

import (
	"context"
	"crypto/rand"
	"encoding/hex"

	"github.com/omni-network/omni/lib/xchain"
)

type StreamUpdate struct {
	xchain.StreamID
	MsgTree     xchain.MsgTree
	Attestation xchain.Attestation // Attestation for the xmsgs
	Msgs        []xchain.Msg       // msgs that increment the cursor
}

// CreateFunc is a function that creates one or more submissions from the given stream update.
type CreateFunc func(streamUpdate StreamUpdate) ([]xchain.Submission, error)

// SendFunc sends a submission to the destination chain by invoking "xsubmit" on portal contract.
type SendFunc func(ctx context.Context, submission xchain.Submission) error

// randomHex7 returns a random 7-character hex string.
func randomHex7() string {
	bytes := make([]byte, 4)
	_, _ = rand.Read(bytes)
	hexString := hex.EncodeToString(bytes)

	// Trim the string to 7 characters
	if len(hexString) > 7 {
		hexString = hexString[:7]
	}

	return hexString
}
