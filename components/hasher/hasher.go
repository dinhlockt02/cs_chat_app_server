package hasher

type Hasher interface {
	Hash(data string) (string, error)
	Compare(data string, hashedData string) (bool, error)
}
