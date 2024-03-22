package fireblocks

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"net/http"
	"strconv"
	"text/template"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Accounts returns all the vault accounts from the account cache, populating it if empty.
func (c Client) Accounts(ctx context.Context) (map[common.Address]uint64, error) {
	if !c.cache.Populated() {
		if err := c.populateAccountCache(ctx); err != nil {
			return nil, errors.Wrap(err, "populating account cache")
		}
	}

	return c.cache.Clone(), nil
}

// getAccount returns the Fireblocks account ID for the given address from the account cache.
// It populates the cache if the account is not found.
func (c Client) getAccount(ctx context.Context, addr common.Address) (uint64, error) {
	accounts, err := c.Accounts(ctx)
	if err != nil {
		return 0, err
	}

	account, ok := accounts[addr]
	if !ok {
		return 0, errors.New("account not found")
	}

	return account, nil
}

// populateAccountCache populates the accounts cache with all vault account addresses.
func (c Client) populateAccountCache(ctx context.Context) error {
	header, err := c.authHeaders(endpointVaults, nil)
	if err != nil {
		return err
	}

	var resp vaultsResponse
	var errResp errorResponse
	ok, err := c.jsonHTTP.Send(
		ctx,
		endpointVaults,
		http.MethodGet,
		nil,
		header,
		&resp,
		&errResp,
	)
	if err != nil {
		return err
	} else if !ok {
		return errors.New("failed to get vaults", "msg", errResp.Message, "code", errResp.Code)
	} else if resp.Paging.After != "" {
		return errors.New("paging not implemented")
	}

	for _, account := range resp.Accounts {
		id, err := strconv.ParseUint(account.ID, 10, 64)
		if err != nil {
			return errors.Wrap(err, "parsing account ID")
		}

		pubkey, err := c.GetPublicKey(ctx, id)
		if err != nil {
			return errors.Wrap(err, "getting public key")
		}

		c.cache.Set(crypto.PubkeyToAddress(*pubkey), id)
	}

	return nil
}

// GetPublicKey returns the public key for the given vault account.
func (c Client) GetPublicKey(ctx context.Context, account uint64) (*ecdsa.PublicKey, error) {
	endpoint, err := c.pubkeyEndpoint(account)
	if err != nil {
		return nil, errors.Wrap(err, "getting pubkey endpoint")
	}

	headers, err := c.authHeaders(endpoint, nil)
	if err != nil {
		return nil, err
	}

	var res pubkeyResponse
	var errRes errorResponse
	ok, err := c.jsonHTTP.Send(
		ctx,
		endpoint,
		http.MethodGet,
		nil,
		headers,
		&res,
		&errRes,
	)
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.New("failed to get public key", "msg", errRes.Message, "code", errRes.Code)
	}

	pk, err := hex.DecodeString(res.PublicKey)
	if err != nil {
		return nil, errors.Wrap(err, "decoding public key")
	}

	resp, err := crypto.DecompressPubkey(pk)
	if err != nil {
		return nil, errors.Wrap(err, "decompressing public key")
	}

	return resp, nil
}

// pubkeyEndpoint returns the public key endpoint by populating the template.
func (c Client) pubkeyEndpoint(account uint64) (string, error) {
	tmpl, err := template.New("").Parse(endpointPubkeyTmpl)
	if err != nil {
		return "", errors.Wrap(err, "parsing pubkey endpoint template")
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, struct {
		VaultAccountID string
		AssetID        string
	}{
		VaultAccountID: strconv.FormatUint(account, 10),
		AssetID:        c.getAssetID(),
	})
	if err != nil {
		return "", errors.Wrap(err, "executing pubkey endpoint template")
	}

	return buf.String(), nil
}
