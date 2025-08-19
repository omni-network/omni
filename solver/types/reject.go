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
	RejectExpenseOverMax        RejectReason = 11
	RejectExpenseUnderMin       RejectReason = 12
	RejectCallNotAllowed        RejectReason = 13
	RejectChainDisabled         RejectReason = 14
	rejectSentinel              RejectReason = 15

	// RejectSameChain          RejectReason = 10 // Same chain orders supported.
)
