package fireblocks

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/omni-network/omni/lib/errors"
)

// jsonHTTP provides a simple interface for sending JSON HTTP requests.
type jsonHTTP struct {
	apiKey string
	host   string
	http   http.Client
}

// newJSONHTTP creates a new jsonHTTP.
func newJSONHTTP(host string, apiKey string) jsonHTTP {
	return jsonHTTP{
		host:   host,
		apiKey: apiKey,
	}
}

// Send sends an JSON HTTP request with the json formatted request as body.
// If the response status code is 2XX, it marshals the response body into the response pointer and returns true.
// Else, it marshals the response body into the errResponse pointer and returns false.
func (c jsonHTTP) Send(ctx context.Context, uri string, httpMethod string, request any, headers map[string]string, response any, errResponse *errorResponse) (bool, error) {
	endpoint, err := url.Parse(c.host + uri)
	if err != nil {
		return false, errors.Wrap(err, "parse")
	}

	// on get requests even will a nil request, we are passing in a non nil request body as the body marshaled to equal `null`
	// so we just set it to nil if the request is nil
	var reqBytes []byte
	if request != nil {
		reqBytes, err = json.Marshal(request)
		if err != nil {
			return false, errors.Wrap(err, "marshaling JSON")
		}
	}

	req, err := http.NewRequestWithContext(
		ctx,
		httpMethod,
		endpoint.String(),
		bytes.NewReader(reqBytes),
	)
	if err != nil {
		return false, errors.Wrap(err, "new http request")
	}

	req.Header = mergeJSONHeaders(headers)

	resp, err := c.http.Do(req)
	if err != nil {
		return false, errors.Wrap(err, "http do")
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, errors.Wrap(err, "read response body")
	}

	if resp.StatusCode/100 != 2 { //nolint:usestdlibvars,nestif // False positive.
		if errResponse != nil {
			// When rate limited, Fireblocks returns http body and not JSON.
			if resp.StatusCode == http.StatusTooManyRequests {
				errResponse.Message = "rate limited"
				errResponse.Code = resp.StatusCode

				return false, nil
			}

			if resp.Header.Get("Content-Type") != "application/json" {
				errResponse.Code = resp.StatusCode

				return false, errors.New("non-JSON error response", "status code", resp.StatusCode, "body", string(respBytes))
			}

			err = json.Unmarshal(respBytes, errResponse)
			if err != nil {
				return false, errors.Wrap(err, "unmarshal error response", "status code", resp.StatusCode, "body", string(respBytes))
			}
		}

		return false, nil
	}

	if response != nil {
		err = json.Unmarshal(respBytes, response)
		if err != nil {
			return false, errors.Wrap(err, "unmarshal response")
		}
	}

	return true, nil
}

// mergeJSONHeaders merges the default JSON headers with the given headers.
func mergeJSONHeaders(m map[string]string) http.Header {
	header := http.Header{}

	header.Set("Accept", "application/json")
	header.Set("Content-Type", "application/json")

	for k, v := range m {
		header.Set(k, v)
	}

	return header
}
