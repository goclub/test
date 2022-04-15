package xtest

func PickOne[T comparable](slice []T) T {
	listLen := len(slice)
	if listLen == 0 {
		panic(newError("PickOne(list) list can not be empty slice"))
	}
	index := Int(0, listLen-1)
	return slice[index]
}
