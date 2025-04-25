package cctp

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"slices"
	"sync"
	"testing"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/cctp/testutil"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/stretchr/testify/require"
)

// DevClient is a devnet implementation of the CCTP client.
type DevClient struct {
	mu           sync.RWMutex
	privKey      *ecdsa.PrivateKey
	ethClients   map[uint64]ethclient.Client
	attestations map[common.Hash][]byte
	cursors      map[uint64]uint64
}

var _ Client = (*DevClient)(nil)

// NewDevClient returns a new devnet CCTP client.
// Dev clients are used to sign and serve attestations on devnet.
func NewDevClient(privKey *ecdsa.PrivateKey, ethClients map[uint64]ethclient.Client) *DevClient {
	return &DevClient{
		privKey:      privKey,
		ethClients:   ethClients,
		attestations: make(map[common.Hash][]byte),
		cursors:      make(map[uint64]uint64),
	}
}

// GetAttestation returns the attestation for a message hash.
func (c *DevClient) GetAttestation(_ context.Context, messageHash common.Hash) ([]byte, AttestationStatus, error) {
	attestation, ok := c.getAttestation(messageHash)
	if !ok {
		return nil, AttestationStatusPendingConfirmations, nil
	}

	return attestation, AttestationStatusComplete, nil
}

// AttestForever watches MessageTransmitter events and signs attestations.
func (c *DevClient) AttestForever(ctx context.Context, chains []evmchain.Metadata, xprov xchain.Provider) error {
	ctx = log.WithCtx(ctx, "process", "cctp.DevClient.AttestForever")

	transmitters, addrs, err := newMessageTransmitters(c.ethClients)
	if err != nil {
		return err
	}

	// Init cursors
	for _, chain := range chains {
		height, err := c.getLatestBlock(ctx, chain.ChainID)
		if err != nil {
			return errors.Wrap(err, "init cursor", "chain", chain.Name)
		}

		c.setCursor(chain.ChainID, height)

		log.Info(ctx, "Initialized cursor", "chain", chain.Name, "height", height)
	}

	for _, chain := range chains {
		chainID := chain.ChainID
		chainName := chain.Name
		transmitter := transmitters[chainID]
		addr := addrs[chainID]

		go func() {
			ctx := log.WithCtx(ctx, "chain", chainName)
			log.Info(ctx, "Starting event processor")
			proc := c.newEventProc(chainID, transmitter)
			c.runEventProc(ctx, chainID, addr, proc, xprov)
		}()
	}

	return nil
}

// newEventProc returns an event processor for a chain.
func (c *DevClient) newEventProc(chainID uint64, transmitter *MessageTransmitter) xchain.EventLogsCallback {
	return func(ctx context.Context, header *ethtypes.Header, elogs []ethtypes.Log) error {
		for _, elog := range elogs {
			msg, err := transmitter.ParseMessageSent(elog)
			if err != nil {
				return errors.Wrap(err, "parse message sent")
			}

			messageHash := crypto.Keccak256Hash(msg.Message)

			// Skip if already attested
			if _, ok := c.getAttestation(messageHash); ok {
				continue
			}

			// Sign and store
			sig, err := crypto.Sign(messageHash[:], c.privKey)
			if err != nil {
				return errors.Wrap(err, "sign message")
			}

			// Make v eth compatible (27 or 28)
			sig[64] += 27

			log.Debug(ctx, "Attesting message", "message_hash", messageHash, "sig", hexutil.Encode(sig))

			c.setAttestation(messageHash, sig)
		}

		c.setCursor(chainID, header.Number.Uint64())

		return nil
	}
}

func (c *DevClient) runEventProc(
	ctx context.Context,
	chainID uint64,
	addr common.Address,
	proc xchain.EventLogsCallback,
	xprov xchain.Provider) {
	backoff := expbackoff.New(ctx)
	for {
		from := c.getCursor(chainID)

		req := xchain.EventLogsReq{
			ChainID:         chainID,
			Height:          from,
			ConfLevel:       xchain.ConfLatest, // Use latest for devnet
			FilterAddresses: []common.Address{addr},
			FilterTopics:    []common.Hash{messageSentEvent.ID},
		}

		err := xprov.StreamEventLogs(ctx, req, proc)
		if ctx.Err() != nil {
			return
		}

		log.Warn(ctx, "Failed streaming events (will retry)", err)
		backoff()
	}
}

// getAttestation returns the attestation for a message hash.
func (c *DevClient) getAttestation(hash common.Hash) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	attestation, ok := c.attestations[hash]

	return slices.Clone(attestation), ok
}

// setAttestation stores an attestation for a message hash.
func (c *DevClient) setAttestation(hash common.Hash, attestation []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.attestations[hash] = slices.Clone(attestation)
}

// getCursor returns the cursor for a chain.
func (c *DevClient) getCursor(chainID uint64) uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.cursors[chainID]
}

// setCursor stores a cursor for a chain.
func (c *DevClient) setCursor(chainID, height uint64) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cursors[chainID] = height
}

// getLatestBlock returns the latest block number for a chain.
func (c *DevClient) getLatestBlock(ctx context.Context, chainID uint64) (uint64, error) {
	client, ok := c.ethClients[chainID]
	if !ok {
		return 0, errors.New("no eth client for chain", "chain_id", chainID)
	}

	blockNum, err := client.BlockNumber(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "block number")
	}

	return blockNum, nil
}

func newMessageTransmitters(clients map[uint64]ethclient.Client) (map[uint64]*MessageTransmitter, map[uint64]common.Address, error) {
	addrs := make(map[uint64]common.Address)
	transmitters := make(map[uint64]*MessageTransmitter)

	for chainID, client := range clients {
		contract, addr, err := newMessageTransmitter(chainID, client)
		if err != nil {
			return nil, nil, errors.Wrap(err, "message transmitter", "chain_id", chainID)
		}

		transmitters[chainID] = contract
		addrs[chainID] = addr
	}

	return transmitters, addrs, nil
}

// StartTestClient is starts a dev client for integration tests.
// It creates a new attester account, enables it on all chains, and starts attesting.
// It requires rpcs to be anvil forks run with auto impersonation.
// It returns the client.
func StartTestClient(
	ctx context.Context,
	t *testing.T,
	xprov xchain.Provider,
	chains []evmchain.Metadata,
	clients map[uint64]ethclient.Client,
) Client {
	t.Helper()

	attesterPk, attesterAddr := testutil.NewAccount(t)

	// Create CCTP client, start attesting
	cctpClient := NewDevClient(attesterPk, clients)

	// Enable attester on all chains
	enableAttester(ctx, t, clients, attesterAddr)

	// Start attesting
	err := cctpClient.AttestForever(ctx, chains, xprov)
	require.NoError(t, err)

	return cctpClient
}

// enableAttester enables the attester on each chains MessageTransmitter, and sets the signature threshold to 1.
func enableAttester(ctx context.Context, t *testing.T, clients map[uint64]ethclient.Client, newAttester common.Address) {
	t.Helper()

	for chainID, client := range clients {
		mtAddr, ok := MessageTransmitterAddr(chainID)
		require.True(t, ok)

		mt, err := NewMessageTransmitter(mtAddr, client)
		require.NoError(t, err)

		// AttesterManager is account allowed to enable attesters
		mngr, err := mt.AttesterManager(&bind.CallOpts{Context: ctx})
		require.NoError(t, err)

		// Send unsigned MessageTransmitter.enableAttester tx from attester manager
		// This requires anvil auto impersonation
		calldata, err := messageTransmitterABI.Pack("enableAttester", newAttester)
		require.NoError(t, err)
		txHash, err := sendUnsignedTransaction(ctx, client, txArgs{
			from:  mngr,
			to:    mtAddr,
			value: nil,
			data:  calldata,
		})
		require.NoError(t, err)

		_, err = bind.WaitMinedHash(ctx, client, txHash)
		require.NoError(t, err)

		// Verify attester is enabled
		enabled, err := mt.IsEnabledAttester(&bind.CallOpts{Context: ctx}, newAttester)
		require.NoError(t, err)
		require.True(t, enabled, "attester not enabled on chain %d", chainID)

		log.Info(ctx, "Enabled attester", "chain", chainID, "attester", newAttester)

		// Reduce signature threshold to 1 (number of attesters required)
		calldata, err = messageTransmitterABI.Pack("setSignatureThreshold", bi.One())
		require.NoError(t, err)
		txHash, err = sendUnsignedTransaction(ctx, client, txArgs{
			from:  mngr,
			to:    mtAddr,
			value: nil,
			data:  calldata,
		})
		require.NoError(t, err)

		_, err = bind.WaitMinedHash(ctx, client, txHash)
		require.NoError(t, err)

		threshold, err := mt.SignatureThreshold(&bind.CallOpts{Context: ctx})
		require.NoError(t, err)
		tutil.RequireEQ(t, threshold, bi.One(), "signature threshold not set to 1 on chain %d", chainID)

		log.Info(ctx, "Signature threshold set", "chain", chainID, "threshold", threshold)
	}
}

type txArgs struct {
	from  common.Address
	to    common.Address
	value *big.Int
	data  []byte
}

// This is used to send auto impersonated txs on anvil.
func sendUnsignedTransaction(ctx context.Context, client ethclient.Client, args txArgs) (common.Hash, error) {
	value := bi.Zero()
	if args.value != nil {
		value = args.value
	}

	to := args.to
	msg := ethereum.CallMsg{
		From:  args.from,
		To:    &to,
		Value: value,
		Data:  args.data,
	}

	gas, err := client.EstimateGas(ctx, msg)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "estimate gas")
	}

	jsonArgs := map[string]any{
		"from":  args.from,
		"to":    to,
		"value": (*hexutil.Big)(value),
		"data":  hexutil.Bytes(args.data),
		"gas":   hexutil.Uint64(gas),
	}

	var result common.Hash
	err = client.CallContext(ctx, &result, "eth_sendTransaction", jsonArgs)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "send transaction")
	}

	return result, nil
}
