package common

func Recovery() {
	if err := recover(); err != nil {
		return
	}
	return
}
