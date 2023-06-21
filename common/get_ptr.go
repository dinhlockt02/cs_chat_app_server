package common

func GetPointer[T bool](value T) *T {
	return &value
}
