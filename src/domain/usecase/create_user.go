package domain

type CreateUser interface {
	CreateUser(params CreateUserParams) (bool, error)
}

type CreateUserParams struct {
	Name     string
	Email    string
	Password string
}
