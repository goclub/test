package xtest

func StringsToAnys(strings []string) (anys []interface{}) {
	for _, s := range strings {
		anys = append(anys, s)
	}
	return
}
func IntsToAnys(ints []int) (anys []interface{}) {
	for _, s := range ints {
		anys = append(anys, s)
	}
	return
}
func Int64sToAnys(int64s []int64) (anys []interface{}) {
	for _, s := range int64s {
		anys = append(anys, s)
	}
	return
}

func Uint64sToAnys(uint64s []uint64) (anys []interface{}) {
	for _, s := range uint64s {
		anys = append(anys, s)
	}
	return
}
