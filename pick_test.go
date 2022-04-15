package xtest_test

import (
	xtest "github.com/goclub/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPickOne(t *testing.T) {
	func() struct{} {
		// -------------
		// PickOne []stirng
		{
			var err error
			count := map[string]int{}
			err = xtest.Run(100, func(_ int) (op xtest.RunOp) {
				item := xtest.PickOne([]string{"a", "b"})
				count[item]++
				return
			})
			assert.NoError(t, err)
			assert.Greater(t, count["a"], 0)
			assert.Greater(t, count["b"], 0)
			assert.Equal(t, count["a"]+count["b"], 100)
		}
		// PickOne []int
		{
			var err error
			count := map[int]int{}
			err = xtest.Run(100, func(_ int) (op xtest.RunOp) {
				item := xtest.PickOne([]int{1, 2})
				count[item]++
				return
			})
			assert.NoError(t, err)
			assert.Greater(t, count[1], 0)
			assert.Greater(t, count[2], 0)
			assert.Equal(t, count[1]+count[2], 100)
		}
		// -------------
		return struct{}{}
	}()

}
