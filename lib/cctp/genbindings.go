package cctp

// Generates bindings for the CCTP's TokenMessenger and MessageTransmitter contracts.
//
//go:generate abigen --abi token-messenger.json --type TokenMessenger --pkg cctp --out token_messenger_bindings.go
//go:generate abigen --abi message-transmitter.json --type MessageTransmitter --pkg cctp --out message_transmitter_bindings.go
