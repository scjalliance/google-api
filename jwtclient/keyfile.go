package jwtclient

import (
	"errors"
	"io/ioutil"

	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

func configFromKeyfile(keyfile string, scope ...string) (*jwt.Config, error) {
	if keyfile == "" {
		return nil, errors.New("no keyfile specified")
	}

	data, err := ioutil.ReadFile(keyfile)
	if err != nil {
		return nil, err
	}

	return google.JWTConfigFromJSON(data, scope...)
}
