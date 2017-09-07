// Package jwtclient facilitates creation of HTTP clients that use JSON web
// tokens for two-legged OAuth2 authentication with Google's service APIs.
package jwtclient

import (
	"context"
	"net/http"

	"golang.org/x/oauth2/google"
)

// New returns an OAuth2 authenticating HTTP client for the given key data and
// scope(s).
//
// The key data must be in JSON format.
//
// The client will only be valid as long as the given context is valid.
func New(ctx context.Context, keydata []byte, scope ...string) (*http.Client, error) {
	config, err := google.JWTConfigFromJSON(keydata, scope...)
	if err != nil {
		return nil, jwtError(err)
	}
	return config.Client(ctx), nil
}

// NewWithSubject returns an OAuth2 authenticating HTTP client for the given
// key data, subject and scope(s).
//
// The key data must be in JSON format.
//
// The client will only be valid as long as the given context is valid.
func NewWithSubject(ctx context.Context, keydata []byte, subject string, scope ...string) (*http.Client, error) {
	conf, err := google.JWTConfigFromJSON(keydata, scope...)
	if err != nil {
		return nil, jwtError(err)
	}
	conf.Subject = subject
	return conf.Client(ctx), nil
}

// NewFromKeyfile returns an OAuth2 authenticating HTTP client for the given
// key file and scope(s).
//
// The key data contained within the key file must be in JSON format.
//
// The client will only be valid as long as the given context is valid.
func NewFromKeyfile(ctx context.Context, keyfile string, scope ...string) (*http.Client, error) {
	conf, err := configFromKeyfile(keyfile, scope...)
	if err != nil {
		return nil, jwtError(err)
	}
	return conf.Client(ctx), nil
}

// NewFromKeyfileWithSubject returns an OAuth2 authenticating HTTP client for
// the given key file, subject and scope(s).
//
// The key data contained within the key file must be in JSON format.
//
// The client will only be valid as long as the given context is valid.
func NewFromKeyfileWithSubject(ctx context.Context, keyfile, subject string, scope ...string) (*http.Client, error) {
	conf, err := configFromKeyfile(keyfile, scope...)
	if err != nil {
		return nil, jwtError(err)
	}
	conf.Subject = subject
	return conf.Client(ctx), nil
}
