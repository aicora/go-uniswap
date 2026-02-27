package utils

import (
	"math/big"
)

var Q96 = new(big.Int).Lsh(big.NewInt(1), 96) // 2^96
var MaxUint160 = new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(160), nil), big.NewInt(1))
var MaxUint256, _ = new(big.Int).SetString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 16)