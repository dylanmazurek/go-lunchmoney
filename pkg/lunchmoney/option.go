package lunchmoney

import "github.com/dylanmazurek/go-lunchmoney/pkg/utilities/vault"

type Options struct {
	vaultClient *vault.Client
	apiKey      string
}

func DefaultOptions() Options {
	defaultOptions := Options{}

	return defaultOptions
}

type Option func(*Options)

func WithVaultClient(client *vault.Client) Option {
	return func(opts *Options) {
		opts.vaultClient = client
	}
}

func WithAPIKey(apiKey string) Option {
	return func(opts *Options) {
		opts.apiKey = apiKey
	}
}
