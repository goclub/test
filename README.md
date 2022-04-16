# goclub/test

[![Go Reference](https://pkg.go.dev/badge/github.com/goclub/test.svg)](https://pkg.go.dev/github.com/goclub/test)


> 所有函数第一次入参是`(t *testing.T)` 的目的是防止在非测试代码中使用 xtest.因为 xtest 为了便利性出现错误时候使用 panic 而不是返回error.