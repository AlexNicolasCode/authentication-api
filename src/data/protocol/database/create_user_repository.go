package database

type CreateUserRepository interface {
	CreateUser(params CreateUserRepositoryParams) error
}

type CreateUserRepositoryParams struct {
	Name     string
	Email    string
	Password string
}
