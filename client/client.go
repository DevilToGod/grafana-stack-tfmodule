package gapi

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

var (
	// ErrAuthorizationTokenRequired is the error returned by client
	// calls that require an authorization token.
	//
	// Authorization tokens can be obtained using Install or Init.
	ErrAuthorizationTokenRequired = errors.New("authorization token required")

	// ErrCannotEncodeJSONRequest is the error returned if it's not
	// possible to encode the request as a JSON object. This error
	// should never happen.
	ErrCannotEncodeJSONRequest = errors.New("cannot encode request")

	// ErrUnexpectedResponse is returned by client calls that get
	// unexpected responses, for example some field that must not be
	// zero is zero. If possible a more specific error should be
	// used.
	ErrUnexpectedResponse = errors.New("unexpected response")
)

// Client is a Grafana API client.
type Client struct {
	client      *http.Client
	accessToken string
	baseURL     string
}

/*func (c *Client) request(method, requestPath string, query url.Values, body io.Reader, responseStruct interface{}) error {
	url := c.baseURL
	url.Path = path.Join(url.Path, requestPath)
	url.RawQuery = query.Encode()
	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return req, err
	}

	if c.config.APIKey != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.config.APIKey))
	} else if c.config.OrgID != 0 {
		req.Header.Add("X-Grafana-Org-Id", strconv.FormatInt(c.config.OrgID, 10))
	}

	if os.Getenv("GF_LOG") != "" {
		if body == nil {
			log.Printf("request (%s) to %s with no body data", method, url.String())
		} else {
			log.Printf("request (%s) to %s with body data: %s", method, url.String(), body.(*bytes.Buffer).String())
		}
	}

	req.Header.Add("Content-Type", "application/json")
	return req, err
}*/

// NewClient creates a new client for the Synthetic Monitoring API.
//
// The accessToken is optional. If it's not specified, it's necessary to
// use one of the registration calls to obtain one, Install or Init.
//
// If no client is provided, http.DefaultClient will be used.
func NewClient(baseURL, accessToken string, client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}

	u, err := url.Parse(baseURL + "/api/instances")
	if err != nil {
		return nil
	}

	u.Path = path.Clean(u.Path)

	return &Client{
		client:      client,
		accessToken: accessToken,
		baseURL:     u.String(),
	}
}

func (h *Client) requireAuthToken() error {
	if h.accessToken == "" {
		return ErrAuthorizationTokenRequired
	}

	return nil
}

func (h *Client) NewStack(ctx context.Context) (*http.Response, error) {
	if err := h.requireAuthToken(); err != nil {
		return nil, err
	}

	resp, err := h.stack(ctx, fmt.Sprintf("%s%s", h.baseURL, ""), true)
	if err != nil {
		return resp, fmt.Errorf("Create Stack: %w", err)
	}

	return resp, nil
}

func (h *Client) stack(ctx context.Context, url string, auth bool) (*http.Response, error) {
	return h.do(ctx, url, http.MethodPost, auth, nil, nil)
}

func (h *Client) do(ctx context.Context, url, method string, auth bool, headers http.Header, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("creating new HTTP request: %w", err)
	}

	if headers != nil {
		req.Header = headers
	}

	if auth {
		if req.Header == nil {
			req.Header = make(http.Header)
		}
		req.Header.Set("Authorization", "Bearer "+h.accessToken)
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending HTTP request: %w", err)
	}

	return resp, nil
}
