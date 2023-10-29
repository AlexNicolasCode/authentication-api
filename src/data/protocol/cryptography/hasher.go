package protocol

type Hasher interface {
	Hash(plaintext string) (string, error)
}
