package domain

type CreateUser interface {
	CreateUser(params CreateUserParams) error
}

type CreateUserParams struct {
	Name     string
	Email    string
	Password string
}
