package xtest

import "testing"

func PickOne(t *testing.T, list []interface{}) interface{} {
	length := len(list)
	if length == 0 {
		panic(newError("PickOne(t,list) list can not be empty slice"))
	}
	index := Int(t, 0, length-1)
	return list[index]
}
func PickString(t *testing.T, list []string) string {
	return PickOne(t, StringsToAnys(list)).(string)
}
func PickInt64(t *testing.T, list []int64) int64 {
	return PickOne(t, Int64sToAnys(list)).(int64)
}
func PickInt(t *testing.T, list []int) int {
	return PickOne(t, IntsToAnys(list)).(int)
}
func PickUint64(t *testing.T, list []uint64) uint64 {
	return PickOne(t, Uint64sToAnys(list)).(uint64)
}
