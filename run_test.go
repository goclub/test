package xtest_test

import (
	xtest "github.com/goclub/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRun(t *testing.T) {
	{
		data := []int{}
		err := xtest.Run(t, 10, func(i int) (op xtest.RunOp) {
			data = append(data, i)
			if i == 5 {
				return op.Break()
			}
			return
		})
		assert.NoError(t, err)
		assert.Equal(t, data, []int{0, 1, 2, 3, 4, 5})
	}
	{
		data := []int{}
		err := xtest.Run(t, 10, func(i int) (op xtest.RunOp) {
			data = append(data, i)
			return
		})
		assert.NoError(t, err)
		assert.Equal(t, data, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	}
}
