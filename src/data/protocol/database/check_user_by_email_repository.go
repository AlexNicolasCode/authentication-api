package protocol

type CheckUserByEmailRepository interface {
	CheckByEmail(email string) (bool, error)
}
