package xtest

func checkError(err error) {
	if err != nil {
		// 此库仅用于测试，所以可以使用 panic 而不是将错误向上传递
		panic(err)
	}
}
