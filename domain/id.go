package domain

import (
	"math/rand"
	"strconv"
)

func NewID() string {
	return strconv.Itoa(int(rand.Uint32()))
}
