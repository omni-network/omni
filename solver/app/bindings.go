package app

import (
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	statusInvalid   uint8 = 0
	statusPending   uint8 = 1
	statusAccepted  uint8 = 2
	statusRejected  uint8 = 3
	statusReverted  uint8 = 4
	statusFulfilled uint8 = 5
	statusClaimed   uint8 = 6
)

var (
	inboxABI = mustGetABI(bindings.SolveInboxMetaData)

	// Event log topics (common.Hash).
	topicRequested = mustGetEventTopic(inboxABI, "Requested")
	topicAccepted  = mustGetEventTopic(inboxABI, "Accepted")
	topicRejected  = mustGetEventTopic(inboxABI, "Rejected")
	topicReverted  = mustGetEventTopic(inboxABI, "Reverted")
	topicFulfilled = mustGetEventTopic(inboxABI, "Fulfilled")
	topicClaimed   = mustGetEventTopic(inboxABI, "Claimed")
)

// eventMeta contains metadata about an event.
type eventMeta struct {
	Topic   common.Hash
	Status  uint8
	ParseID func(contract bindings.SolveInboxFilterer, log types.Log) ([32]byte, error)
}

var (
	allEvents = []eventMeta{
		{
			Topic:   topicRequested,
			Status:  statusPending,
			ParseID: parseRequested,
		},
		{
			Topic:   topicAccepted,
			Status:  statusAccepted,
			ParseID: parseAccepted,
		},
		{
			Topic:   topicRejected,
			Status:  statusRejected,
			ParseID: parseRejected,
		},
		{
			Topic:   topicReverted,
			Status:  statusReverted,
			ParseID: parseReverted,
		},
		{
			Topic:   topicFulfilled,
			Status:  statusFulfilled,
			ParseID: parseFulfilled,
		},
		{
			Topic:   topicClaimed,
			Status:  statusClaimed,
			ParseID: parseClaimed,
		},
	}

	// eventsByTopic maps event topics to their metadata.
	eventsByTopic = func() map[common.Hash]eventMeta {
		resp := make(map[common.Hash]eventMeta, len(allEvents))
		for _, e := range allEvents {
			resp[e.Topic] = e
		}

		return resp
	}()

	allEventTopics = func() []common.Hash {
		resp := make([]common.Hash, 0, len(allEvents))
		for _, e := range allEvents {
			resp = append(resp, e.Topic)
		}

		return resp
	}()
)

func statusString(status uint8) string {
	switch status {
	case statusInvalid:
		return "invalid"
	case statusPending:
		return "pending"
	case statusAccepted:
		return "accepted"
	case statusRejected:
		return "rejected"
	case statusReverted:
		return "reverted"
	case statusFulfilled:
		return "fulfilled"
	case statusClaimed:
		return "claimed"
	default:
		return "unknown"
	}
}

func parseRequested(contract bindings.SolveInboxFilterer, log types.Log) ([32]byte, error) {
	e, err := contract.ParseRequested(log)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "parse requested")
	}

	return e.Id, nil
}

func parseAccepted(contract bindings.SolveInboxFilterer, log types.Log) ([32]byte, error) {
	e, err := contract.ParseAccepted(log)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "parse accepted")
	}

	return e.Id, nil
}

func parseRejected(contract bindings.SolveInboxFilterer, log types.Log) ([32]byte, error) {
	e, err := contract.ParseRejected(log)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "parse rejected")
	}

	return e.Id, nil
}

func parseReverted(contract bindings.SolveInboxFilterer, log types.Log) ([32]byte, error) {
	e, err := contract.ParseReverted(log)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "parse reverted")
	}

	return e.Id, nil
}

func parseFulfilled(contract bindings.SolveInboxFilterer, log types.Log) ([32]byte, error) {
	e, err := contract.ParseFulfilled(log)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "parse fulfilled")
	}

	return e.Id, nil
}

func parseClaimed(contract bindings.SolveInboxFilterer, log types.Log) ([32]byte, error) {
	e, err := contract.ParseClaimed(log)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "parse claimed")
	}

	return e.Id, nil
}

// mustGetABI returns the metadata's ABI as an abi.ABI type.
// It panics on error.
func mustGetABI(metadata *bind.MetaData) *abi.ABI {
	abi, err := metadata.GetAbi()
	if err != nil {
		panic(err)
	}

	return abi
}

// mustGetEvent returns the event with the given name from the ABI.
// It panics if the event is not found.
func mustGetEventTopic(abi *abi.ABI, name string) common.Hash {
	event, ok := abi.Events[name]
	if !ok {
		panic("event not found")
	}

	return event.ID
}
