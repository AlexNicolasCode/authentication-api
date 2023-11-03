package mock

import (
	"errors"

	protocol "data/protocol/database"
)

type CreateUserRepository struct {
	Count        int
	Params       protocol.CreateUserRepositoryParams
	ErrorMessage string
}

func (m *CreateUserRepository) CreateUser(params protocol.CreateUserRepositoryParams) error {
	m.Count += 1
	m.Params = params
	if len(m.ErrorMessage) > 0 {
		return errors.New(m.ErrorMessage)
	}
	return nil
}
