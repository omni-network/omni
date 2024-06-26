package fireblocks

import (
	"context"
	"crypto/rsa"
	"maps"
	"sync"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

const (
	endpointTransactions = "/v1/transactions"
	endpointAssets       = "/v1/supported_assets"
	endpointVaults       = "/v1/vault/accounts_paged"
	endpointPubkeyTmpl   = "/v1/vault/accounts/{{.VaultAccountID}}/{{.AssetID}}/0/0/public_key_info?compressed"

	assetHolesky = "ETH_TEST6"
	assetSepolia = "ETH_TEST5"
	assetMainnet = "ETH"

	hostProd    = "https://api.fireblocks.io"
	hostSandbox = "https://sandbox-api.fireblocks.io"
)

// Client is a JSON HTTP client for the FireBlocks API.
type Client struct {
	cfg        Config
	apiKey     string
	network    netconf.ID
	privateKey *rsa.PrivateKey
	jsonHTTP   jsonHTTP
	cache      *accountCache
}

// New creates a new FireBlocks client.
func New(network netconf.ID, apiKey string, privateKey *rsa.PrivateKey, opts ...Option) (Client, error) {
	if apiKey == "" {
		return Client{}, errors.New("apiKey is required")
	}
	if privateKey == nil {
		return Client{}, errors.New("privateKey is required")
	}

	cfg := defaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}
	if err := cfg.check(); err != nil {
		return Client{}, errors.Wrap(err, "options check")
	}

	return Client{
		apiKey:     apiKey,
		privateKey: privateKey,
		jsonHTTP:   newJSONHTTP(cfg.Host, apiKey),
		cfg:        cfg,
		cache:      newAccountCache(cfg.TestAccounts),
		network:    network,
	}, nil
}

// authHeaders returns the authentication headers for the FireBlocks API.
func (c Client) authHeaders(endpoint string, request any) (map[string]string, error) {
	token, err := c.token(endpoint, request)
	if err != nil {
		return nil, errors.Wrap(err, "generating token")
	}

	return map[string]string{
		"X-API-KEY":     c.apiKey,
		"Authorization": "Bearer " + token,
	}, nil
}

func (c Client) getAssetID() string {
	switch c.network {
	case netconf.Mainnet:
		return assetMainnet
	default:
		return assetHolesky
	}
}

func newAccountCache(init map[common.Address]uint64) *accountCache {
	return &accountCache{
		accountsByAddress: init,
	}
}

type accountCache struct {
	sync.Mutex
	accountsByAddress map[common.Address]uint64
}

func (c *accountCache) MaybePopulate(ctx context.Context, fn func(context.Context) (map[common.Address]uint64, error)) error {
	c.Lock()
	defer c.Unlock()

	if len(c.accountsByAddress) > 0 {
		return nil
	}

	accounts, err := fn(ctx)
	if err != nil {
		return err
	}

	c.accountsByAddress = accounts

	return nil
}

func (c *accountCache) Get(addr common.Address) (uint64, bool) {
	c.Lock()
	defer c.Unlock()

	acc, ok := c.accountsByAddress[addr]

	return acc, ok
}

func (c *accountCache) Set(addr common.Address, id uint64) {
	c.Lock()
	defer c.Unlock()

	c.accountsByAddress[addr] = id
}

func (c *accountCache) Clone() map[common.Address]uint64 {
	c.Lock()
	defer c.Unlock()

	return maps.Clone(c.accountsByAddress)
}
