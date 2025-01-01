package lunchmoney

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/dylanmazurek/go-lunchmoney/pkg/lunchmoney/constants"
	"github.com/dylanmazurek/go-lunchmoney/pkg/lunchmoney/models"
	"github.com/dylanmazurek/go-lunchmoney/pkg/utilities/vault"
)

type AuthClient struct {
	internalClient *http.Client

	opts Options

	secrets Secrets
	session models.Session
}

type Secrets struct {
	apiKey string
}

// NewAuthClient creates a new AuthClient with the provided options.
func NewAuthClient(ctx context.Context, opts Options) (*AuthClient, error) {
	authClient := &AuthClient{
		internalClient: &http.Client{Transport: http.DefaultTransport},

		opts:    opts,
		session: models.Session{},
	}

	secrets, err := initVault(ctx, opts.vaultClient)
	if err != nil {
		return nil, err
	}

	authClient.secrets = *secrets

	authClient.session.SetAPIKey(secrets.apiKey)

	return authClient, nil
}

func initVault(ctx context.Context, client *vault.Client) (*Secrets, error) {
	secrets, err := client.GetSecret(ctx, "kv", "finance-sync/lunchmoney")
	if err != nil {
		return nil, err
	}

	secretObj := &Secrets{
		apiKey: secrets["API_KEY"].(string),
	}

	return secretObj, nil
}

type addAuthHeaderTransport struct {
	T       http.RoundTripper
	Session models.Session
}

func (adt *addAuthHeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", adt.Session.GetAPIKey()))
	req.Header.Add("User-Agent", constants.REPO_URL)

	return adt.T.RoundTrip(req)
}

// InitTransportSession initializes a transport session for the AuthClient.
func (c *AuthClient) InitTransportSession() (*http.Client, error) {
	currentAPIKey := c.session.GetAPIKey()
	if currentAPIKey == "" {
		err := fmt.Errorf("api key is not set")

		return nil, err
	}

	user, err := c.getUserData(currentAPIKey)
	if err != nil {
		err := fmt.Errorf("api key not valid")

		return nil, err
	}

	clientLog.Debug().
		Str("username", user.Name).
		Msgf("user data fetched")

	authTransport, err := c.createAuthTransport()

	return authTransport, err
}

// getUserData fetches user data using the provided API key.
func (c *AuthClient) getUserData(apiKey string) (*models.User, error) {
	path := fmt.Sprintf("%s%s", constants.API_BASE_URL, constants.API_PATH_ME)

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("invalid api key")
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var me models.User
	err = json.Unmarshal(bodyBytes, &me)
	if err != nil {
		return nil, err
	}

	return &me, nil
}

func (c *AuthClient) createAuthTransport() (*http.Client, error) {
	authClient := &http.Client{
		Transport: &addAuthHeaderTransport{
			T:       http.DefaultTransport,
			Session: c.session,
		},
	}

	return authClient, nil
}
