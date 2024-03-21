package fireblocks

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

// jsonHTTP provides a simple interface for sending JSON HTTP requests.
type jsonHTTP struct {
	apiKey       string
	clientSecret string
	host         string
	http         http.Client
}

// newJSONHTTP creates a new jsonHTTP.
func newJSONHTTP(host string, apiKey string, clientSecret string) jsonHTTP {
	return jsonHTTP{
		host:         host,
		apiKey:       apiKey,
		clientSecret: clientSecret,
	}
}

// Send sends an JSON HTTP request with the json formatted request as body.
// It marshals the response body into the provided response pointer if not nil.
func (c jsonHTTP) Send(ctx context.Context, endpoint string, httpMethod string, request any, headers map[string]string, response any) error {
	endpoint, err := url.JoinPath(c.host, endpoint)
	if err != nil {
		return errors.Wrap(err, "joining endpoint")
	}

	// on get requests even will a nil request, we are passing in a non nil request body as the body marshaled to equal `null`
	// so we just set it to nil if the request is nil
	var reqBytes []byte
	if request != nil {
		reqBytes, err = json.Marshal(request)
		if err != nil {
			return errors.Wrap(err, "marshaling JSON")
		}
	}

	req, err := http.NewRequestWithContext(
		ctx,
		httpMethod,
		endpoint,
		bytes.NewReader(reqBytes),
	)
	if err != nil {
		return errors.Wrap(err, "new http request")
	}

	req.Header = mergeJSONHeaders(headers)

	resp, err := c.http.Do(req)
	if err != nil {
		return errors.Wrap(err, "http do")
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Error(ctx, "Http: closing body failure", err)
		}
	}()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "read response body")
	}

	if resp.StatusCode != http.StatusOK {
		return errors.Wrap(err, "http response code not okay", "status code", resp.StatusCode, "body", string(respBytes))
	}

	if response != nil {
		err = json.Unmarshal(respBytes, response)
		if err != nil {
			return errors.Wrap(err, "unmarshal response")
		}
	}

	return nil
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
