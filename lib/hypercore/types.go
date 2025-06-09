package hypercore

// SigRSV is a signature in the R, S, V format.
type SigRSV struct {
	R string `json:"r"`
	S string `json:"s"`
	V uint8  `json:"v"`
}

// PhantomAgent is a HL core concept that wraps an "action" for use in EIP-712 messages.
type PhantomAgent struct {
	Source       string   `json:"source"`       // "a" for mainnet, "b" for testnet
	ConnectionID [32]byte `json:"connectionId"` // Hash of action data
}

// ActionPayload is the payload structure for actions sent to the Hypercore API.
type ActionPayload struct {
	Action    any    `json:"action"`    // The action to be performed
	Nonce     uint64 `json:"nonce"`     // Nonce for the action
	Signature SigRSV `json:"signature"` // Signature of the action
}

type ActionEVMUserModify struct {
	Type           string `json:"type"` // "evmUserModify"
	UsingBigBlocks bool   `json:"usingBigBlocks"`
}
