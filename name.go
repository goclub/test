package xtest

import "testing"

func FirstName(t *testing.T) string {
	return PickString(t, seed.FirstName)
}
func LastName(t *testing.T) string {
	return PickString(t, seed.LastName)
}
func Name(t *testing.T) string {
	return FirstName(t) + " " + LastName(t)
}
func FullName(t *testing.T) string {
	return FirstName(t) + " " + FirstName(t) + " " + LastName(t)
}

func CFirstName(t *testing.T) string {
	return PickString(t, seed.ChineseFirstName)
}
func CLastName(t *testing.T) string {
	return PickString(t, seed.ChineseLastName)
}
func CName(t *testing.T) string {
	return CFirstName(t) + CLastName(t)
}
