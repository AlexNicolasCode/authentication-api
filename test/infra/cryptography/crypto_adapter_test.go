package cryptography

import (
	"math/rand"
	"test/infra/mock"
	"testing"

	"github.com/bxcodec/faker/v3"
)

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

func TestShouldThrowIfBcryptThrows(t *testing.T) {
	randomInt := rand.Int()
	bcryptSpy := new(mock.BcryptSpy)
	bcryptSpy.ErrorMessage = "Mocked Error"
	sut := NewCryptoAdapter(bcryptSpy.GenerateFromPassword, randomInt)

	_, err := sut.Hash(faker.Password())

	if err.Error() != bcryptSpy.ErrorMessage {
		t.Error("CreateUser method from CreateUserRepository was called more than one time or not called")
	}
}

func TestShouldCallGenerateFromPasswordMethodOnce(t *testing.T) {
	randomInt := rand.Int()
	bcryptSpy := new(mock.BcryptSpy)
	sut := NewCryptoAdapter(bcryptSpy.GenerateFromPassword, randomInt)

	sut.Hash(faker.Password())

	if bcryptSpy.Count != 1 {
		t.Error("CreateUser method from CreateUserRepository was called more than one time or not called")
	}
}

func TestShouldReturnHashOnSuccess(t *testing.T) {
	randomInt := rand.Int()
	bcryptSpy := new(mock.BcryptSpy)
	sut := NewCryptoAdapter(bcryptSpy.GenerateFromPassword, randomInt)
	fakePassword := faker.Password()

	hash, _ := sut.Hash(fakePassword)

	if hash != string(bcryptSpy.Result) {
		t.Error("CreateUser method from CreateUserRepository was called more than one time or not called")
	}
}
