package jwtclient

import "fmt"

func jwtError(err error) error {
	return fmt.Errorf("jwtclient: %v", err)
}
