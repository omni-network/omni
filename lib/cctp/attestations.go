package cctp

import (
	"github.com/omni-network/omni/lib/errors"
)

// AttestationStatus is the status of an attestation, 'complete' or 'pending_confirmations'.
type AttestationStatus string

const (
	AttestationStatusComplete             AttestationStatus = "complete"
	AttestationStatusPendingConfirmations AttestationStatus = "pending_confirmations"
)

// Validate checks if the status is a known status.
func (s AttestationStatus) Validate() error {
	switch s {
	case AttestationStatusComplete, AttestationStatusPendingConfirmations:
		return nil
	default:
		return errors.New("unexpected attestation status", "status", s)
	}
}
