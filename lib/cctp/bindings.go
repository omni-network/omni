package cctp

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// Generates bindings for the CCTP's TokenMessenger and MessageTransmitter contracts.
//
//go:generate abigen --abi token-messenger.json --type TokenMessenger --pkg cctp --out token_messenger_bindings.go
//go:generate abigen --abi message-transmitter.json --type MessageTransmitter --pkg cctp --out message_transmitter_bindings.go

var (
	tokenMessengerABI     = mustGetABI(TokenMessengerMetaData)
	messageTransmitterABI = mustGetABI(MessageTransmitterMetaData)
	depositForBurnEvent   = mustGetEvent(tokenMessengerABI, "DepositForBurn")
	messageSentEvent      = mustGetEvent(messageTransmitterABI, "MessageSent")
)

func mustGetABI(metadata *bind.MetaData) *abi.ABI {
	abi, err := metadata.GetAbi()
	if err != nil {
		panic(err)
	}

	return abi
}

func mustGetEvent(abi *abi.ABI, eventName string) abi.Event {
	event, ok := abi.Events[eventName]
	if !ok {
		panic("event not found: " + eventName)
	}

	return event
}
