package xtest

import (
	"crypto/rand"
	"math/big"
	"testing"
)

func randomBig(max int64) (random *big.Int, err error) {
	return rand.Int(rand.Reader, big.NewInt(int64(max)))
}

// Int 返回 min ~ max 之间的数字，包括 min 和 max
func Int(t *testing.T, min int, max int) (v int) {
	return int(Int64(t, int64(min), int64(max)))
}

// Int64 返回 min ~ max 之间的数字，包括 min 和 max
func Int64(t *testing.T, min int64, max int64) (v int64) {
	if min > max {
		min, max = max, min
	}
	if min == max {
		return max
	}
	rangeValue := max - min + 1
	random, err := randomBig(rangeValue)
	checkError(err)
	return random.Int64() + min
	// min 6 max 6
	// return 6

	// min 0 max 6
	// random: 0 ~ 6(6-0)
	// return 6 + 0 = 6

	// min 4 max 10
	// random: 0 ~ 6(10-4)
	// return 4 + 4 = 8

	// min -2 max 4
	// random: 0 ~ 6(4-(-2))
	// return 1 + -2 = -1
}

// Bool equal TrueLikelihood(50)
func Bool(t *testing.T) bool {
	return TrueLikelihood(t, 50)
}

// TrueLikelihood 根据百分比 返回 true 或 false
// It panics if likelihood < 0 or likelihood > 100 .
func TrueLikelihood(t *testing.T, likelihood int64) bool {
	if likelihood < 0 {
		panic(newError("BoolLikelihood(likelihood int) likelihood can not less than 0%"))
	}
	if likelihood > 100 {
		panic(newError("BoolLikelihood(likelihood int) likelihood can not greater than 100%"))
	}
	if likelihood == 0 {
		return false
	}
	random := Int64(t, 1, 100)
	return random <= likelihood
}

// RunesBySeed 基于 seed 返回指定数量的 []rune
func RunesBySeed(t *testing.T, seed []rune, size uint64) []rune {
	result := []rune("")
	var i uint64 = 0
	for ; i < size; i++ {
		randIndex, err := randomBig(int64(len(seed)))
		checkError(err)
		result = append(result, seed[randIndex.Int64()])
	}
	return result
}

// AZaz09 seed: ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789
func AZaz09(t *testing.T, size uint64) string {
	return string(RunesBySeed(t, []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"), size))
}

// Letter seed: abcdefghijklmnopqrstuvwxyz
func Letter(t *testing.T, size uint64) string {
	return string(RunesBySeed(t, []rune("abcdefghijklmnopqrstuvwxyz"), size))
}

// CapitalLetter seed: ABCDEFGHIJKLMNOPQRSTUVWXYZ
func CapitalLetter(t *testing.T, size uint64) string {
	return string(RunesBySeed(t, []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ"), size))
}
