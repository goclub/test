package xtest

import "testing"

type RunOp struct {
	shouldBreak bool
	err         error
}

func (RunOp) Continue() RunOp {
	return RunOp{shouldBreak: false, err: nil}
}
func (RunOp) Break() RunOp {
	return RunOp{shouldBreak: true, err: nil}
}
func (RunOp) BreakWithErr(err error) RunOp {
	return RunOp{shouldBreak: true, err: err}
}
func Run(t *testing.T, times int, fn func(i int) (op RunOp)) error {
	for i := 0; i < times; i++ {
		op := fn(i)
		if op.err != nil {
			return op.err
		}
		if op.shouldBreak {
			break
		}
	}
	return nil
}
