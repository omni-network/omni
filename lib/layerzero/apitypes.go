package layerzero

import (
	"github.com/omni-network/omni/lib/errors"
)

type MsgStatus string

const (
	MsgStatusUnknown       MsgStatus = "UNKNOWN"        // Unknown status
	MsgStatusConfirming    MsgStatus = "CONFIRMING"     // System confirming the source tx
	MsgStatusInFlight      MsgStatus = "INFLIGHT"       // Inflight to destination
	MsgStatusDelivered     MsgStatus = "DELIVERED"      // Successfully delivered on the destination
	MsgStatusFailed        MsgStatus = "FAILED"         // Tx errored and did not complete
	MsgStatusPayloadStored MsgStatus = "PAYLOAD_STORED" // Ran out of gas, needs to be retried
)

func (s MsgStatus) Verify() error {
	switch s {
	case MsgStatusUnknown, MsgStatusConfirming, MsgStatusInFlight, MsgStatusDelivered, MsgStatusFailed, MsgStatusPayloadStored:
		return nil
	default:
		return errors.New("invalid message status", "status", s)
	}
}

func (s MsgStatus) String() string {
	return string(s)
}

// Message is the Message type returned by the LayerZero API.
// Most of the fields are omitted. Full docs here: https://scan.layerzero-api.com/v1/swagger
type Message struct {
	Pathway Pathway `json:"pathway"`
	Status  Status  `json:"status"`
}

func (m Message) IsInFlight() bool {
	return m.Status.Name == MsgStatusInFlight.String()
}

func (m Message) IsDelivered() bool {
	return m.Status.Name == MsgStatusDelivered.String()
}

func (m Message) IsConfirming() bool {
	return m.Status.Name == MsgStatusConfirming.String()
}

func (m Message) IsFailed() bool {
	return m.Status.Name == MsgStatusFailed.String()
}

func (m Message) IsPayloadStored() bool {
	return m.Status.Name == MsgStatusPayloadStored.String()
}

type Pathway struct {
	SrcEid int `json:"srcEid"`
	DstEid int `json:"dstEid"`
}

type Status struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

type MessageResponse struct {
	Data []Message `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
