package lib

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

func SignedToString[T constraints.Signed](num T) string {
	return strconv.FormatInt(int64(num), 10)
}

func UnsignedToString[T constraints.Unsigned](num T) string {
	return strconv.FormatUint(uint64(num), 10)
}
