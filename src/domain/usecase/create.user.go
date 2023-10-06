package usecase

type CreateUser interface {
	createUser(params CreateUserParams) error
}

type CreateUserParams struct {
	Name     string
	Email    string
	Password string
}
