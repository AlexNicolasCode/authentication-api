package mock

import (
	protocol "finances-api/src/data/protocol/database"
)

type CreateUserRepository struct {
	Count  int
	Params protocol.CreateUserRepositoryParams
}

func (m *CreateUserRepository) CreateUser(params protocol.CreateUserRepositoryParams) error {
	m.Count += 1
	m.Params = params
	return nil
}
