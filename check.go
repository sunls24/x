package x

func CheckPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckEmpty(list ...string) bool {
	for i := range list {
		if len(list[i]) == 0 {
			return false
		}
	}
	return true
}
