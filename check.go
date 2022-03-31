package x

func CheckPanic(err error) {
	if err != nil {
		panic(err)
	}
}
