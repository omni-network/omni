package types

//go:generate stringer -type=RejectReason -trimprefix=Reject
type RejectReason uint8

const (
	RejectNone                  RejectReason = 0
	RejectDestCallReverts       RejectReason = 1
	RejectInvalidDeposit        RejectReason = 2
	RejectInvalidExpense        RejectReason = 3
	RejectInsufficientDeposit   RejectReason = 4
	RejectInsufficientInventory RejectReason = 5
	RejectUnsupportedDeposit    RejectReason = 6
	RejectUnsupportedExpense    RejectReason = 7
	RejectUnsupportedDestChain  RejectReason = 8
	RejectUnsupportedSrcChain   RejectReason = 9
	RejectSameChain             RejectReason = 10
)
