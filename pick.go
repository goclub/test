package xtest

import "testing"

func PickOne[T comparable](t *testing.T, slice []T) T {
	listLen := len(slice)
	if listLen == 0 {
		panic(newError("PickOne(t,list) list can not be empty slice"))
	}
	index := Int(t, 0, listLen-1)
	return slice[index]
}
