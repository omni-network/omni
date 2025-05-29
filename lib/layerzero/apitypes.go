package layerzero

const (
	MessageStatusFailed     = "FAILED"
	MessageStatusConfirming = "CONFIRMING"
	MessageStatusInFlight   = "INFLIGHT"
	MessageStatusDelivered  = "DELIVERED"
)

// Message is the Message type returned by the LayerZero API.
// Most of the fields are omitted. Full docs here: https://scan.layerzero-api.com/v1/swagger
type Message struct {
	Pathway Pathway `json:"pathway"`
	Status  Status  `json:"status"`
}

func (m Message) IsInFlight() bool {
	return m.Status.Name == MessageStatusInFlight
}

func (m Message) IsDelivered() bool {
	return m.Status.Name == MessageStatusDelivered
}

func (m Message) IsConfirming() bool {
	return m.Status.Name == MessageStatusConfirming
}

func (m Message) IsFailed() bool {
	return m.Status.Name == MessageStatusFailed
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
