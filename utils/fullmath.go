package utils

import (
	"math/big"

	"github.com/pkg/errors"

	"github.com/holiman/uint256"
)

var (
	// ErrZeroDenominator is returned when a division operation is attempted with a zero denominator.
	ErrZeroDenominator = errors.New("denominator is zero")
)

// MulDiv calculates floor(a * b / denominator) using only uint256.Int.
//
// This function performs full 256-bit precision multiplication and division
// without converting to big.Int. It returns an error if the denominator is zero.
//
// Parameters:
//   - a: the multiplicand as a uint256.Int
//   - b: the multiplier as a uint256.Int
//   - denominator: the divisor as a uint256.Int
//
// Returns:
//   - *uint256.Int: the result of floor(a * b / denominator)
//   - error: ErrZeroDenominator if denominator is zero
func MulDiv(a, b, denominator *uint256.Int) (*uint256.Int, error) {
	if denominator.IsZero() {
		return nil, ErrZeroDenominator
	}

	prod := new(uint256.Int).Mul(a, b)
	res := new(uint256.Int).Div(prod, denominator)

	return res, nil
}

// MulDivRoundingUp calculates ceil(a * b / denominator) using only uint256.Int.
//
// This function performs full 256-bit precision multiplication and division
// and rounds up the result if there is any remainder.
// It returns an error if the denominator is zero.
//
// Parameters:
//   - a: the multiplicand as a uint256.Int
//   - b: the multiplier as a uint256.Int
//   - denominator: the divisor as a uint256.Int
//
// Returns:
//   - *uint256.Int: the result of ceil(a * b / denominator)
//   - error: ErrZeroDenominator if denominator is zero
func MulDivRoundingUp(a, b, denominator *uint256.Int) (*uint256.Int, error) {
	if denominator.IsZero() {
		return nil, ErrZeroDenominator
	}

	prod := new(uint256.Int).Mul(a, b)
	res := new(uint256.Int).Div(prod, denominator)

	mod := new(uint256.Int).Mod(prod, denominator)
	if !mod.IsZero() {
		res = res.Add(res, uint256.NewInt(1))
	}

	return res, nil
}

// MulMod calculates (a * b) % m using only uint256.Int.
//
// Performs 256-bit multiplication followed by modulo operation.
// This function avoids conversion to big.Int, making it suitable for high-performance scenarios.
//
// Parameters:
//   - a: the multiplicand as a uint256.Int
//   - b: the multiplier as a uint256.Int
//   - m: the modulus as a uint256.Int
//
// Returns:
//   - *uint256.Int: the result of (a * b) % m
func MulMod(a, b, m *uint256.Int) *uint256.Int {
	res := new(uint256.Int).Mul(a, b)
	return res.Mod(res, m)
}

// AbsBigInt returns the absolute value of a big.Int.
//
// Parameters:
//   - x: the input *big.Int
//
// Returns:
//   - *big.Int: absolute value of x
func AbsBigInt(x *big.Int) *big.Int {
	if x.Sign() < 0 {
		return new(big.Int).Neg(x)
	}
	return new(big.Int).Set(x)
}