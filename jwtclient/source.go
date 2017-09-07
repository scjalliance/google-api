package jwtclient

import (
	"context"
	"net/http"

	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

// Source is a source of JWT HTTP clients that share a common configuration
// but can differ in impersonation subjects.
type Source struct {
	config jwt.Config
}

// NewSource returns a new source of OAuth2 authenticating HTTP clients. The
// clients returned by the source will use the given key data and scope(s).
//
// The key data must be in JSON format.
func NewSource(keydata []byte, scope ...string) (*Source, error) {
	config, err := google.JWTConfigFromJSON(keydata, scope...)
	if err != nil {
		return nil, jwtError(err)
	}
	return &Source{
		config: *config,
	}, nil
}

// NewSourceFromKeyfile returns a source of OAuth2 authenticating HTTP clients.
// The clients returned by the source will use the key data contained in the
// given key file and will use given scope(s).
//
// The key data contained within the key file must be in JSON format.
func NewSourceFromKeyfile(keyfile string, scope ...string) (*Source, error) {
	config, err := configFromKeyfile(keyfile, scope...)
	if err != nil {
		return nil, err
	}
	return &Source{
		config: *config,
	}, nil
}

// Client returns a new HTTP client with the key data and scope(s) defined by
// the source.
//
// The client will only be valid as long as the given context is valid.
func (s *Source) Client(ctx context.Context) (*http.Client, error) {
	return s.config.Client(ctx), nil
}

// ClientWithSubject returns a new HTTP client for the given subject with the
// key data and scope(s) defined by the source.
//
// The client will only be valid as long as the given context is valid.
func (s *Source) ClientWithSubject(ctx context.Context, subject string) (*http.Client, error) {
	conf := s.config // Copy the config so that we don't modify the source
	conf.Subject = subject
	return conf.Client(ctx), nil
}
