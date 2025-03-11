package solvernet

import (
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	inboxABI    = mustGetABI(bindings.SolverNetInboxMetaData)
	outboxABI   = mustGetABI(bindings.SolverNetOutboxMetaData)
	bindingsABI = mustGetABI(bindings.ISolverNetBindingsMetaData)

	inputsOrderData      = mustGetInputs(bindingsABI, "orderData")
	inputsFillOriginData = mustGetInputs(bindingsABI, "fillOriginData")

	// Event log topics (common.Hash).
	TopicOpened   = mustGetEventTopic(inboxABI, "Open")
	TopicRejected = mustGetEventTopic(inboxABI, "Rejected")
	TopicClosed   = mustGetEventTopic(inboxABI, "Closed")
	TopicFilled   = mustGetEventTopic(inboxABI, "Filled")
	TopicClaimed  = mustGetEventTopic(inboxABI, "Claimed")
)

// EventMeta contains metadata about an event.
type EventMeta struct {
	Topic   common.Hash
	Status  OrderStatus
	ParseID func(contract bindings.SolverNetInboxFilterer, log types.Log) (OrderID, error)
}

var (
	allEvents = []EventMeta{
		{
			Topic:   TopicOpened,
			Status:  StatusPending,
			ParseID: ParseOpened,
		},
		{
			Topic:   TopicRejected,
			Status:  StatusRejected,
			ParseID: ParseRejected,
		},
		{
			Topic:   TopicClosed,
			Status:  StatusClosed,
			ParseID: ParseClosed,
		},
		{
			Topic:   TopicFilled,
			Status:  StatusFilled,
			ParseID: ParseFilled,
		},
		{
			Topic:   TopicClaimed,
			Status:  StatusClaimed,
			ParseID: ParseClaimed,
		},
	}

	eventsByTopic = func() map[common.Hash]EventMeta {
		resp := make(map[common.Hash]EventMeta, len(allEvents))
		for _, e := range allEvents {
			resp[e.Topic] = e
		}

		return resp
	}()
)

// EventByTopic returns the event metadata for a given topic.
func EventByTopic(topic common.Hash) (EventMeta, bool) {
	e, ok := eventsByTopic[topic]
	return e, ok
}

// AllEventTopics returns all solvernet event topics.
func AllEventTopics() []common.Hash {
	resp := make([]common.Hash, 0, len(allEvents))
	for _, e := range allEvents {
		resp = append(resp, e.Topic)
	}

	return resp
}

// ParseEvent return the order ID and status from the event log.
func ParseEvent(l types.Log) (OrderID, OrderStatus, error) {
	if len(l.Topics) == 0 {
		return OrderID{}, 0, errors.New("no topics")
	}

	e, ok := EventByTopic(l.Topics[0])
	if !ok {
		return OrderID{}, 0, errors.New("unknown event")
	}

	// Safe to use dummy address and backend since we only parse events.
	parser, err := bindings.NewSolverNetInbox(common.Address{}, nil)
	if err != nil {
		return OrderID{}, 0, errors.New("new solver inbox")
	}

	orderID, err := e.ParseID(parser.SolverNetInboxFilterer, l)
	if err != nil {
		return OrderID{}, 0, errors.Wrap(err, "parse id")
	}

	return orderID, e.Status, nil
}

func ParseOpened(contract bindings.SolverNetInboxFilterer, log types.Log) (OrderID, error) {
	e, err := contract.ParseOpen(log)
	if err != nil {
		return OrderID{}, errors.Wrap(err, "parse opened")
	}

	return e.OrderId, nil
}

func ParseRejected(contract bindings.SolverNetInboxFilterer, log types.Log) (OrderID, error) {
	e, err := contract.ParseRejected(log)
	if err != nil {
		return OrderID{}, errors.Wrap(err, "parse rejected")
	}

	return e.Id, nil
}

func ParseClosed(contract bindings.SolverNetInboxFilterer, log types.Log) (OrderID, error) {
	e, err := contract.ParseClosed(log)
	if err != nil {
		return OrderID{}, errors.Wrap(err, "parse closed")
	}

	return e.Id, nil
}

func ParseFilled(contract bindings.SolverNetInboxFilterer, log types.Log) (OrderID, error) {
	e, err := contract.ParseFilled(log)
	if err != nil {
		return OrderID{}, errors.Wrap(err, "parse filled")
	}

	return e.Id, nil
}

func ParseClaimed(contract bindings.SolverNetInboxFilterer, log types.Log) (OrderID, error) {
	e, err := contract.ParseClaimed(log)
	if err != nil {
		return OrderID{}, errors.Wrap(err, "parse claimed")
	}

	return e.Id, nil
}

func PackFillCalldata(orderID OrderID, fillOriginData []byte) ([]byte, error) {
	// fillerData is optional ERC7683 custom filler specific data, unused in our contracts
	fillerData := []byte{}
	return outboxABI.Pack("fill", orderID, fillOriginData, fillerData)
}

func ParseFillOriginData(data []byte) (bindings.SolverNetFillOriginData, error) {
	unpacked, err := inputsFillOriginData.Unpack(data)
	if err != nil {
		return bindings.SolverNetFillOriginData{}, errors.Wrap(err, "unpack fill data")
	}

	wrap := struct {
		Data bindings.SolverNetFillOriginData
	}{}
	if err := inputsFillOriginData.Copy(&wrap, unpacked); err != nil {
		return bindings.SolverNetFillOriginData{}, errors.Wrap(err, "copy fill data")
	}

	return wrap.Data, nil
}

func ParseOrderData(data []byte) (bindings.SolverNetOrderData, error) {
	unpacked, err := inputsOrderData.Unpack(data)
	if err != nil {
		return bindings.SolverNetOrderData{}, errors.Wrap(err, "unpack fill data")
	}

	wrap := struct {
		Data bindings.SolverNetOrderData
	}{}
	if err := inputsOrderData.Copy(&wrap, unpacked); err != nil {
		return bindings.SolverNetOrderData{}, errors.Wrap(err, "copy fill data")
	}

	return wrap.Data, nil
}

func PackOrderData(data bindings.SolverNetOrderData) ([]byte, error) {
	packed, err := inputsOrderData.Pack(data)
	if err != nil {
		return nil, errors.Wrap(err, "pack fill data")
	}

	return packed, nil
}

func PackFillOriginData(data bindings.SolverNetFillOriginData) ([]byte, error) {
	// Replaces nil call values with zero (inputs.Pack panics if value is nil)
	for i := range data.Calls {
		if data.Calls[i].Value == nil {
			data.Calls[i].Value = big.NewInt(0)
		}
	}

	// Same for expenses
	for i := range data.Expenses {
		if data.Expenses[i].Amount == nil {
			data.Expenses[i].Amount = big.NewInt(0)
		}
	}

	packed, err := inputsFillOriginData.Pack(data)
	if err != nil {
		return nil, errors.Wrap(err, "pack fill data")
	}

	return packed, nil
}

func mustGetEventTopic(abi *abi.ABI, name string) common.Hash {
	event, ok := abi.Events[name]
	if !ok {
		panic("event not found")
	}

	return event.ID
}

func mustGetABI(metadata *bind.MetaData) *abi.ABI {
	abi, err := metadata.GetAbi()
	if err != nil {
		panic(err)
	}

	return abi
}

func mustGetInputs(abi *abi.ABI, name string) abi.Arguments {
	method, ok := abi.Methods[name]
	if !ok {
		panic("method not found")
	}

	return method.Inputs
}
