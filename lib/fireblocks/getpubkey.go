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

	"github.com/ethereum/go-ethereum/crypto"
)

func (c Client) GetPublicKey(ctx context.Context) (*ecdsa.PublicKey, error) {
	endpoint, err := c.pubkeyEndpoint()
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
func (c Client) pubkeyEndpoint() (string, error) {
	tmpl, err := template.New("").Parse(endpointPubkeyTmpl)
	if err != nil {
		return "", errors.Wrap(err, "parsing pubkey endpoint template")
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, struct {
		VaultAccountID string
		AssetID        string
	}{
		VaultAccountID: strconv.FormatUint(c.opts.VaultAccountID, 10),
		AssetID:        c.getAssetID(),
	})
	if err != nil {
		return "", errors.Wrap(err, "executing pubkey endpoint template")
	}

	return buf.String(), nil
}
