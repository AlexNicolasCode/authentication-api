package mock

import (
	"errors"

	"github.com/bxcodec/faker/v3"
)

type Hasher struct {
	Count        int
	Plaintext    string
	ErrorMessage string
	Result       string
}

func (m *Hasher) Hash(plaintext string) (string, error) {
	m.Count += 1
	m.Plaintext = plaintext
	if len(m.ErrorMessage) > 0 {
		return "", errors.New(m.ErrorMessage)
	}
	m.Result = faker.Password()
	return m.Result, nil
}
