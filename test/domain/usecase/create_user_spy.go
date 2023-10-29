package mock

import (
	protocol "finances-api/src/data/protocol/database"
)

type CreateUserRepository struct {
	Count int
}

func (m *CreateUserRepository) CreateUser(params protocol.CreateUserRepositoryParams) error {
	m.Count += 1
	return nil
}
