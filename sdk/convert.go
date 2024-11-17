package sdk

import (
	"math/big"
)

func FromHex(input string) *big.Int {
	input = input[2:]

	out := new(big.Int)
	out.SetString(input, 16)
	return out
}

func ToHex(big *big.Int) string {
	return "0x" + big.Text(16)
}
