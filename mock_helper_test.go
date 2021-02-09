package xtest_test

import (
	xtest "github.com/goclub/test"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPickString(t *testing.T) {
	list := []string{"abc", "efg"}
	oneCount := 0
	twoCount := 0
	err := xtest.Run(100, func(order int) (op xtest.RunOp) {
		s := xtest.PickString(list)
		if s == "abc" {
			oneCount++
		}
		if s == "efg" {
			twoCount++
		}
		if s != "abc" && s !="efg" {
			t.Log("pickString error")
			t.Fail()
		}
		return
	})
	assert.NoError(t, err)
	if twoCount == 0 {
		t.Log("pickString one error")
		t.Fail()
	}
	if oneCount == 0 {
		t.Log("pickString two error")
		t.Fail()
	}
}

func TestRun(t *testing.T) {
	{
		data := []int{}
		err := xtest.Run(10, func(i int) (op xtest.RunOp) {
			data = append(data, i)
			if i==5 {
				return op.Break()
			}
			return
		})
		assert.NoError(t, err)
		assert.Equal(t, data, []int{0,1,2,3,4,5})
	}
	{
		data := []int{}
		err := xtest.Run(10, func(i int) (op xtest.RunOp) {
			data = append(data, i)
			return
		})
		assert.NoError(t, err)
		assert.Equal(t, data, []int{0,1,2,3,4,5,6,7,8,9})
	}
}
