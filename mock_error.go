package xtest

import "errors"

func mockNewError(s string) error {
	return errors.New("goclub/mock: "+ s)
}
