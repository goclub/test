package xtest_test

import (
	xtest "github.com/goclub/test"
	"github.com/stretchr/testify/assert"
	"regexp"
	"strconv"
	"testing"
)
func TestUUID(t *testing.T) {
	testRun(100, func(i int) (_break bool) {
		assert.Equal(t,len(xtest.UUID()), 36)
		assert.True(t,testMustBool(regexp.MatchString("[a-z0-9]{8}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{12}", xtest.UUID())))
		return
	})
	countMap := map[string]int{}
	testRun(100, func(i int) (_break bool) {
		countMap[xtest.UUID()] = countMap[xtest.UUID()] +1
		if countMap[xtest.UUID()] >1 {
			t.Log("uuid repeat!")
			t.Fail()
		}
		return
	})
}
func TestIncrID(t *testing.T) {
	
	userIncrID := xtest.IncrID()
	userStringID := xtest.IncrID()
	testRun(100, func(i int) (_break bool) {
		id := i+1
		assert.Equal(t,id, userIncrID.Int())
		return
	})
	testRun(100, func(i int) (_break bool) {
		id := i+1
		assert.Equal(t,strconv.Itoa(id), userStringID.String())
		return
	})
}
func TestNameIncrID(t *testing.T) {
	
	testRun(100, func(i int) (_break bool) {
		id := strconv.Itoa(i+1)
		assert.Equal(t,id, xtest.NameIncrID("34gv43g43gv"))
		return
	})
}
