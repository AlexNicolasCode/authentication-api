package mock

import (
	"errors"
)

type CheckUserByEmailRepository struct {
	Count        int
	Email        string
	ErrorMessage string
	Result       bool
}

func (m *CheckUserByEmailRepository) CheckByEmail(email string) (bool, error) {
	m.Count += 1
	m.Email = email
	m.Result = false
	if len(m.ErrorMessage) > 0 {
		return false, errors.New(m.ErrorMessage)
	}
	return m.Result, nil
}
