package cryptography

type CryptoAdapter struct {
	salt                 int
	GenerateFromPassword func(password []byte, cost int) ([]byte, error)
}

func (c *CryptoAdapter) Hash(plaintext string) (string, error) {
	bytes, err := c.GenerateFromPassword([]byte(plaintext), c.salt)
	return string(bytes), err
}

func NewCryptoAdapter(
	GenerateFromPassword func(password []byte, cost int) ([]byte, error),
	salt int,
) CryptoAdapter {
	return CryptoAdapter{salt, GenerateFromPassword}
}
