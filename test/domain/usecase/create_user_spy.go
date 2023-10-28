package spy

import (
	"finances-api/src/data/protocol/database"
)

type CreateUserRepository struct {
	Count int
}

func (m *CreateUserRepository) CreateUser(params database.CreateUserRepositoryParams) error {
	m.Count += 1
	return nil
}
