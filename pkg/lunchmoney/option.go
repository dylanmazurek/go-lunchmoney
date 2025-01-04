package lunchmoney

import (
	"github.com/dylanmazurek/go-lunchmoney/pkg/utilities/vault"
	"github.com/rs/zerolog/log"
)

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
	if apiKey == "" {
		log.Panic().Msg("api key cannot be empty")
	}

	return func(opts *Options) {
		opts.apiKey = apiKey
	}
}
