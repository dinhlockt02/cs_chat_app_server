package common

func Reverse[T any](slice []T) {
	i := 0
	j := len(slice) - 1

	if len(slice) < 2 {
		return
	}

	for i < j {
		slice[i], slice[j] = slice[j], slice[i]
		i++
		j--
	}
	return
}
