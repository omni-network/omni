package appv2

import (
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	statusInvalid  uint8 = 0
	statusPending  uint8 = 1
	statusAccepted uint8 = 2
	statusRejected uint8 = 3
	statusReverted uint8 = 4
	statusFilled   uint8 = 5
	statusClaimed  uint8 = 6
)

var (
	bindingsABI = mustGetABI(bindings.ISolverNetBindingsMetaData)
	inboxABI    = mustGetABI(bindings.SolverNetInboxMetaData)

	// Event log topics (common.Hash).
	topicOpened   = mustGetEventTopic(inboxABI, "Open")
	topicAccepted = mustGetEventTopic(inboxABI, "Accepted")
	topicRejected = mustGetEventTopic(inboxABI, "Rejected")
	topicReverted = mustGetEventTopic(inboxABI, "Reverted")
	topicFilled   = mustGetEventTopic(inboxABI, "Filled")
	topicClaimed  = mustGetEventTopic(inboxABI, "Claimed")

	inputsFillOriginData = mustGetInputs(bindingsABI, "fillOriginData")
)

// eventMeta contains metadata about an event.
type eventMeta struct {
	Topic   common.Hash
	Status  uint8
	ParseID func(contract bindings.SolverNetInboxFilterer, log types.Log) (OrderID, error)
}

var (
	allEvents = []eventMeta{
		{
			Topic:   topicOpened,
			Status:  statusPending,
			ParseID: parseOpened,
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
			Topic:   topicFilled,
			Status:  statusFilled,
			ParseID: parseFilled,
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
	case statusFilled:
		return "filled"
	case statusClaimed:
		return "claimed"
	default:
		return unknown
	}
}

func parseOpened(contract bindings.SolverNetInboxFilterer, log types.Log) (OrderID, error) {
	e, err := contract.ParseOpen(log)
	if err != nil {
		return OrderID{}, errors.Wrap(err, "parse opened")
	}

	return e.OrderId, nil
}

func parseAccepted(contract bindings.SolverNetInboxFilterer, log types.Log) (OrderID, error) {
	e, err := contract.ParseAccepted(log)
	if err != nil {
		return OrderID{}, errors.Wrap(err, "parse accepted")
	}

	return e.Id, nil
}

func parseRejected(contract bindings.SolverNetInboxFilterer, log types.Log) (OrderID, error) {
	e, err := contract.ParseRejected(log)
	if err != nil {
		return OrderID{}, errors.Wrap(err, "parse rejected")
	}

	return e.Id, nil
}

func parseReverted(contract bindings.SolverNetInboxFilterer, log types.Log) (OrderID, error) {
	e, err := contract.ParseReverted(log)
	if err != nil {
		return OrderID{}, errors.Wrap(err, "parse reverted")
	}

	return e.Id, nil
}

func parseFilled(contract bindings.SolverNetInboxFilterer, log types.Log) (OrderID, error) {
	e, err := contract.ParseFilled(log)
	if err != nil {
		return OrderID{}, errors.Wrap(err, "parse filled")
	}

	return e.Id, nil
}

func parseClaimed(contract bindings.SolverNetInboxFilterer, log types.Log) (OrderID, error) {
	e, err := contract.ParseClaimed(log)
	if err != nil {
		return OrderID{}, errors.Wrap(err, "parse claimed")
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

// mustGetInputs returns the inputs for the method with the given name.
func mustGetInputs(abi *abi.ABI, name string) abi.Arguments {
	method, ok := abi.Methods[name]
	if !ok {
		panic("method not found")
	}

	return method.Inputs
}

// parseFillOriginData parses FillOriginData from packed bytes.
func parseFillOriginData(data []byte) (FillOriginData, error) {
	unpacked, err := inputsFillOriginData.Unpack(data)
	if err != nil {
		return FillOriginData{}, errors.Wrap(err, "unpack fill data")
	}

	wrap := struct {
		Data FillOriginData
	}{}
	if err := inputsFillOriginData.Copy(&wrap, unpacked); err != nil {
		return FillOriginData{}, errors.Wrap(err, "copy fill data")
	}

	return wrap.Data, nil
}
