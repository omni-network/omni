package fireblocks

import (
	"context"
	"net/http"

	"github.com/omni-network/omni/lib/errors"
)

// GetSupportedAssets returns all asset types supported by Fireblocks.
func (c Client) GetSupportedAssets(ctx context.Context) ([]Asset, error) {
	headers, err := c.authHeaders(endpointAssets, nil)
	if err != nil {
		return nil, err
	}

	var res []Asset
	var errRes errorResponse
	ok, err := c.jsonHTTP.Send(
		ctx,
		endpointAssets,
		http.MethodGet,
		nil,
		headers,
		&res,
		&errRes,
	)
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.New("failed to get supported assets", "resp_msg", errRes.Message, "resp_code", errRes.Code)
	}

	return res, nil
}
