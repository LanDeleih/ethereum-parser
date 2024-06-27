package provider

import (
	"fmt"
	"math/big"
)

func ConvertHexToNumber(hexNumber string) (int, error) {
	if len(hexNumber) <= 2 {
		return 0, ErrInvalidHexLength
	}
	if hexNumber[:2] != "0x" {
		return 0, ErrInvalidHex
	}

	bigInt, ok := new(big.Int).SetString(hexNumber[2:], 16)
	if !ok {
		return 0, ErrInvalidHex
	}

	return int(bigInt.Int64()), nil
}

func ConvertNumberToHex(number int) string {
	return fmt.Sprintf("0x%x", number)
}
