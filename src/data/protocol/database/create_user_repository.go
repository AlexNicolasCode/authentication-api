package protocol

type CreateUserRepository interface {
	CreateUser(params CreateUserRepositoryParams) error
}

type CreateUserRepositoryParams struct {
	Name     string
	Email    string
	Password string
}
