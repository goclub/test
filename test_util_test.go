package xtest_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testRun(times int, cb func(i int) (_break bool)) {
	for i := 0; i < times; i++ {
		if cb(i) {
			break
		}
	}
}

func testMustBool(b bool, err error) bool {
	if err != nil {
		panic(err)
	}
	return b
}
func testRange(t *testing.T, v int, min int, max int) {
	assert.GreaterOrEqual(t, v, min)
	assert.LessOrEqual(t, v, max)
}

func testIn[T comparable](list []T, target T) bool {
	for _, s := range list {
		if s == target {
			return true
		}
	}
	return false
}

func testPickOne[T comparable](t *testing.T, seed []T, counter map[T]int) {
	assert.True(t, len(counter) != 0)
	for name, _ := range counter {
		assert.True(t, testIn(seed, name))
	}
}
