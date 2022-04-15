package xtest

func FirstName() string {
	return PickOne(seed.FirstName)
}
func LastName() string {
	return PickOne(seed.LastName)
}
func Name() string {
	return FirstName() + " " + LastName()
}
func FullName() string {
	return FirstName() + " " + FirstName() + " " + LastName()
}

func CFirstName() string {
	return PickOne(seed.ChineseFirstName)
}
func CLastName() string {
	return PickOne(seed.ChineseLastName)
}
func CName() string {
	return CFirstName() + CLastName()
}
