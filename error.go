package xtest

import "github.com/goclub/error"

func newError(s string) error {
	return xerr.New("goclub/test: " + s)
}
