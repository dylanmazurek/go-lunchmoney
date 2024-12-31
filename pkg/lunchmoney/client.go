package lunchmoney

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/dylanmazurek/go-lunchmoney/pkg/lunchmoney/constants"
	"github.com/dylanmazurek/go-lunchmoney/pkg/lunchmoney/models"
	"github.com/rs/zerolog"
)

var clientLog zerolog.Logger

type Client struct {
	internalClient *http.Client
}

func New(ctx context.Context, opts ...Option) (*Client, error) {
	clientLog = zerolog.Ctx(ctx).With().Str("client", "lunchmoney").Logger()

	clientOptions := DefaultOptions()
	for _, opt := range opts {
		opt(&clientOptions)
	}

	authClient, err := NewAuthClient(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	clientLog.Debug().Msg("client initialized")

	authTransport, err := authClient.InitTransportSession()
	if err != nil {
		return nil, err
	}

	clientLog.Debug().Msg("client authenticated")

	newServiceClient := &Client{
		internalClient: authTransport,
	}

	return newServiceClient, nil
}

func (c *Client) NewRequest(method string, path string, body io.Reader, params *url.Values) (*models.Request, error) {
	urlString := fmt.Sprintf("%s%s", constants.API_BASE_URL, path)
	requestUrl, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	if params != nil {
		requestUrl.RawQuery = params.Encode()
	}

	req, err := http.NewRequest(method, requestUrl.String(), body)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	request := &models.Request{
		HTTPRequest: req,
	}

	return request, nil
}

func (c *Client) Do(req *models.Request, resp interface{}) error {
	httpResponse, err := c.internalClient.Do(req.HTTPRequest)
	if err != nil || httpResponse == nil || httpResponse.StatusCode >= 400 {
		if httpResponse != nil {
			clientLog.Debug().Msgf("http response error: %s", httpResponse.Status)
		}

		return err
	}
	defer httpResponse.Body.Close()

	bodyBytes, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyBytes, &resp)
	if err != nil {
		return err
	}

	return nil
}
