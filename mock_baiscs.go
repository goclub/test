package xtest

import (
	"crypto/rand"
	"log"
	"math/big"
)

func randomBig(max int) (random *big.Int, err error) {
	return rand.Int(rand.Reader, big.NewInt(int64(max)))
}

// 返回 min ~ max 之间的数字，包括 min 和 max
func Int(min int, max int) (v int) {
	if min > max {
		min, max = max, min
		log.Print("Int(min int, max int) min can not greater than max")
	}
	if min == max {
		return max
	}
	rangeValue := max - min + 1
	random, err := randomBig(rangeValue)
	mockCheck(err)
	return int(random.Int64()) + min
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
func Bool() bool {
	return TrueLikelihood(50)
}

// TrueLikelihood 根据百分比 返回 true 或 false
// It panics if likelihood < 0 or likelihood > 100 .
func TrueLikelihood(likelihood int) bool {
	if likelihood < 0 {
		panic(mockNewError("BoolLikelihood(likelihood int) likelihood can not less than 0%"))
	}
	if likelihood > 100 {
		panic(mockNewError("BoolLikelihood(likelihood int) likelihood can not greater than 100%"))
	}
	if likelihood == 0 {
		return false
	}
	random := Int(1, 100)
	return bool(random <= likelihood)
}

// RunesBySeed 基于 seed 返回指定数量的 []rune
func RunesBySeed(seed []rune, size int) []rune {
	result := []rune("")
	for i := 0; i < size; i++ {
		randIndex, err := randomBig(len(seed))
		mockCheck(err)
		result = append(result, seed[randIndex.Int64()])
	}
	return result
}

func String(size int) string {
	return string(RunesBySeed([]rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()[]"), size))
}
func Letter(size int) string {
	return string(RunesBySeed([]rune("abcdefghijklmnopqrstuvwxyz"), size))
}
func CapitalLetter(size int) string {
	return string(RunesBySeed([]rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ"), size))
}
