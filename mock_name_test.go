package xtest_test

import (
	xtest "github.com/goclub/test"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var seed = xtest.OnlyTestSeed()

func TestCFirstName(t *testing.T) {
	nameMap := map[string]int{}
	testRun(100, func(i int) (_break bool) {
		name := xtest.CFirstName()
		nameMap[name]++
		return
	})
	testPickOne(t, seed.ChineseFirstName, nameMap)
}
func TestFirstName(t *testing.T) {
	nameMap := map[string]int{}
	testRun(100, func(i int) (_break bool) {
		name := xtest.FirstName()
		nameMap[name]++
		return
	})
	testPickOne(t, seed.FirstName, nameMap)
}

func TestCLastName(t *testing.T) {
	nameMap := map[string]int{}
	testRun(100, func(i int) (_break bool) {
		name := xtest.CLastName()
		nameMap[name]++
		return
	})
	testPickOne(t, seed.ChineseLastName, nameMap)
}

func TestLastName(t *testing.T) {
	nameMap := map[string]int{}
	testRun(100, func(i int) (_break bool) {
		name := xtest.LastName()
		nameMap[name]++
		return
	})
	testPickOne(t, seed.LastName, nameMap)
}

func TestName(t *testing.T) {

	nameMap := map[string]int{}
	testRun(100, func(i int) (_break bool) {
		name := xtest.Name()
		nameMap[name]++
		return
	})
	assert.True(t, len(nameMap) != 0)
	for name, _ := range nameMap {
		names := strings.Split(name, " ")
		firstName := names[0]
		lastName := names[1]
		assert.True(t, testIn(seed.FirstName, firstName))
		assert.True(t, testIn(seed.LastName, lastName))
	}
}

func TestCName(t *testing.T) {

	nameMap := map[string]int{}
	testRun(100, func(i int) (_break bool) {
		name := xtest.CName()
		nameMap[name]++
		return
	})
	assert.True(t, len(nameMap) != 0)
	assert.True(t, len(nameMap) > 10)
}

func TestFullName(t *testing.T) {

	nameMap := map[string]int{}
	testRun(100, func(i int) (_break bool) {
		name := xtest.FullName()
		nameMap[name]++
		return
	})
	assert.True(t, len(nameMap) != 0)
	for name, _ := range nameMap {
		names := strings.Split(name, " ")
		firstName := names[0]
		middleName := names[1]
		lastName := names[2]
		assert.True(t, testIn(seed.FirstName, firstName))
		assert.True(t, testIn(seed.MiddleName, middleName))
		assert.True(t, testIn(seed.LastName, lastName))
	}
}

type mockNames struct {
	FirstName  string `cha:"FirstName()"`
	LastName   string `cha:"LastName()"`
	Name       string `cha:"Name()"`
	FullName   string `cha:"FullName()"`
	CFirstName string `cha:"CFirstName()"`
	CLastName  string `cha:"CLastName()"`
	CName      string `cha:"CName()"`
}
