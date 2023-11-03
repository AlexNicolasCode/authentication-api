package cryptography

import (
	"math/rand"
	"testing"

	"github.com/bxcodec/faker/v3"

	"src/infra/cryptography"
	"test/infra/mock"
)

func TestShouldThrowIfBcryptThrows(t *testing.T) {
	randomInt := rand.Int()
	bcryptSpy := new(mock.BcryptSpy)
	bcryptSpy.ErrorMessage = "Mocked Error"
	sut := cryptography.NewCryptoAdapter(bcryptSpy.GenerateFromPassword, randomInt)

	_, err := sut.Hash(faker.Password())

	if err.Error() != bcryptSpy.ErrorMessage {
		t.Error("CreateUser method from CreateUserRepository was called more than one time or not called")
	}
}

func TestShouldCallGenerateFromPasswordMethodOnce(t *testing.T) {
	randomInt := rand.Int()
	bcryptSpy := new(mock.BcryptSpy)
	sut := cryptography.NewCryptoAdapter(bcryptSpy.GenerateFromPassword, randomInt)

	sut.Hash(faker.Password())

	if bcryptSpy.Count != 1 {
		t.Error("CreateUser method from CreateUserRepository was called more than one time or not called")
	}
}

func TestShouldReturnHashOnSuccess(t *testing.T) {
	randomInt := rand.Int()
	bcryptSpy := new(mock.BcryptSpy)
	sut := cryptography.NewCryptoAdapter(bcryptSpy.GenerateFromPassword, randomInt)
	fakePassword := faker.Password()

	hash, _ := sut.Hash(fakePassword)

	if hash != string(bcryptSpy.Result) {
		t.Error("CreateUser method from CreateUserRepository was called more than one time or not called")
	}
}
