package mock

import "errors"

type BcryptSpy struct {
	Count        int
	Password     []byte
	Cost         int
	Result       []byte
	ErrorMessage string
}

func (m *BcryptSpy) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	m.Count += 1
	m.Password = password
	m.Cost = cost
	if len(m.ErrorMessage) > 0 {
		return nil, errors.New(m.ErrorMessage)
	}
	return m.Result, nil
}
