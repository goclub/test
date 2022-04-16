package xtest_test

import (
	xtest "github.com/goclub/test"
	"github.com/stretchr/testify/assert"
	"log"
	"regexp"
	"testing"
)

func coreTestInt(t *testing.T, min int, max int, rangeList []int) {
	list := []int{}
	testRun(1000, func(i int) (_break bool) {
		list = append(list, xtest.Int(t, min, max))
		return
	})
	for _, item := range list {
		assert.True(t, item >= min && item <= max)
	}
	for i := 0; i < len(rangeList); i++ {
		number := rangeList[i]
		foundIt := false
		for _, item := range list {
			if item == number {
				foundIt = true
			}
		}
		if !foundIt {
			log.Print(" can not found ", number)
			log.Print(min, max)
			log.Print(list)
			t.Fail()
		}
	}
}
func TestInt(t *testing.T) {
	coreTestInt(t, -3, 6, []int{-3, -2, -1, 0, 1, 2, 3, 4, 5, 6})
	coreTestInt(t, -2, 6, []int{-2, -1, 0, 1, 2, 3, 4, 5, 6})
	coreTestInt(t, -1, 6, []int{-1, 0, 1, 2, 3, 4, 5, 6})
	coreTestInt(t, 0, 6, []int{0, 1, 2, 3, 4, 5, 6})
	coreTestInt(t, 1, 6, []int{1, 2, 3, 4, 5, 6})
	coreTestInt(t, 2, 6, []int{2, 3, 4, 5, 6})
	coreTestInt(t, 3, 6, []int{3, 4, 5, 6})
	coreTestInt(t, 4, 6, []int{4, 5, 6})
	coreTestInt(t, 5, 6, []int{5, 6})
	coreTestInt(t, 6, 6, []int{6})
}
func coreTestBool(t *testing.T, likelihood int, trueCount int, falseCount int) {

	assert.Equal(t, 10000, trueCount+falseCount)
	if trueCount < likelihood*100 && trueCount > likelihood*100 {
		t.Log("trueCount", trueCount, " overflow normal range")
		t.Fail()
	}

}
func TestBool(t *testing.T) {
	{
		trueCount := 0
		falseCount := 0
		testRun(10000, func(i int) (_break bool) {
			if xtest.Bool(t) {
				trueCount++
			} else {
				falseCount++
			}
			return
		})
		coreTestBool(t, 50, trueCount, falseCount)
	}
}
func TestTrueLikelihood(t *testing.T) {
	{
		trueCount := 0
		falseCount := 0
		testRun(10000, func(i int) (_break bool) {
			if xtest.TrueLikelihood(t, 0) {
				trueCount++
			} else {
				falseCount++
			}
			return
		})
		coreTestBool(t, 0, trueCount, falseCount)
	}
	{
		trueCount := 0
		falseCount := 0
		testRun(10000, func(i int) (_break bool) {
			if xtest.TrueLikelihood(t, 10) {
				trueCount++
			} else {
				falseCount++
			}
			return
		})
		coreTestBool(t, 10, trueCount, falseCount)
	}
	{
		trueCount := 0
		falseCount := 0
		testRun(10000, func(i int) (_break bool) {
			if xtest.TrueLikelihood(t, 20) {
				trueCount++
			} else {
				falseCount++
			}
			return
		})
		coreTestBool(t, 20, trueCount, falseCount)
	}
	{
		trueCount := 0
		falseCount := 0
		testRun(10000, func(i int) (_break bool) {
			if xtest.TrueLikelihood(t, 100) {
				trueCount++
			} else {
				falseCount++
			}
			return
		})
		coreTestBool(t, 100, trueCount, falseCount)
	}

}

func TestLetter(t *testing.T) {

	assert.Equal(t, len(xtest.Letter(t, 10)), 10)
	assert.False(t, testMustBool(regexp.MatchString(`[^a-z]`, xtest.Letter(t, 10000))))
}
func TestCapitalLetter(t *testing.T) {

	assert.Equal(t, len(xtest.CapitalLetter(t, 10)), 10)
	assert.False(t, testMustBool(regexp.MatchString(`[^A-Z]`, xtest.CapitalLetter(t, 10000))))
}
func TestAZaz09(t *testing.T) {
	testRun(100, func(i int) (_break bool) {
		s := xtest.AZaz09(t, uint64(xtest.Int64(t, 1, 5)))
		testRange(t, len(s), 1, 5)
		for i := 0; i < len(s); i++ {
			v := s[i]
			assert.Contains(t, "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789", string(v))
		}
		return
	})
}
