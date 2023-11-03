package mock

import (
	"domain/model"

	"github.com/bxcodec/faker/v3"
)

func MakeUser() *model.User {
	return &model.User{
		Name:     faker.Name(),
		Email:    faker.Email(),
		Password: faker.Password(),
	}
}
