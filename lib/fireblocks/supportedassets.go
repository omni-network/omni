package fireblocks

import (
	"context"
	"net/http"
)

// GetSupportedAssets returns a list of supported assets.
func (c Client) GetSupportedAssets(ctx context.Context) ([]SupportedAssets, error) {
	jwtToken, err := c.token(supportedAssetsEndpoint, nil)
	if err != nil {
		return nil, err
	}

	var res []SupportedAssets
	err = c.jsonHTTP.Send(
		ctx,
		supportedAssetsEndpoint,
		http.MethodGet,
		nil,
		c.getAuthHeaders(jwtToken),
		&res,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}
