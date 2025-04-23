package cctp

import (
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/evmchain"

	"github.com/ethereum/go-ethereum/common"
)

type msgContracts struct {
	MessageTransmitter        *MessageTransmitter
	MessageTransmitterAddress common.Address
	TokenMessenger            *TokenMessenger
	TokenMessengerAddress     common.Address
}

// Contracts returns a new contracts instance for chainID (TokenMessenger and MessageTransmitter).
func newContracts(chainID uint64, client ethclient.Client) (msgContracts, error) {
	msgr, msgrAddr, err := newMessageTransmitter(chainID, client)
	if err != nil {
		return msgContracts{}, errors.Wrap(err, "new message transmitter")
	}

	tknMsgr, tknMsgrAddr, err := newTokenMessenger(chainID, client)
	if err != nil {
		return msgContracts{}, errors.Wrap(err, "new token messenger")
	}

	return msgContracts{
		MessageTransmitter:        msgr,
		MessageTransmitterAddress: msgrAddr,
		TokenMessenger:            tknMsgr,
		TokenMessengerAddress:     tknMsgrAddr,
	}, nil
}

// newTokenMessenger returns a new TokenMessenger instance for chainID.
func newTokenMessenger(chainID uint64, client ethclient.Client) (*TokenMessenger, common.Address, error) {
	addr, ok := tokenMessengers[chainID]
	if !ok {
		return nil, common.Address{}, errors.New("no messenger", "chain", evmchain.Name(chainID))
	}

	msgr, err := NewTokenMessenger(addr, client)
	if err != nil {
		return nil, common.Address{}, err
	}

	return msgr, addr, nil
}

// newMessageTransmitter returns a new MessageTransmitter instance for chainID.
func newMessageTransmitter(chainID uint64, client ethclient.Client) (*MessageTransmitter, common.Address, error) {
	addr, ok := messageTransmitters[chainID]
	if !ok {
		return nil, common.Address{}, errors.New("no transmitter", "chain", evmchain.Name(chainID))
	}

	transmitter, err := NewMessageTransmitter(addr, client)
	if err != nil {
		return nil, common.Address{}, err
	}

	return transmitter, addr, nil
}
