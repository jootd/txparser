package parser

import "errors"

var (
	ErrAddrInvalid = errors.New("invalid address")
)

type Address string

func parseAddress(addr string) (Address, error) {
	// imitate validation
	if len(addr) == 0 {
		return "", ErrAddrInvalid
	}

	return Address(addr), nil
}
