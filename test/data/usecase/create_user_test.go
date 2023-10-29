package data_test

import (
	"testing"

	"github.com/bxcodec/faker/v3"

	protocol "finances-api/src/data/protocol/database"
	domain "finances-api/src/domain/usecase"
	mock "finances-api/test/domain/usecase"
)

type UseCase struct {
	create_user_repository protocol.CreateUserRepository
}

func (uc *UseCase) CreateUser(params domain.CreateUserParams) error {
	var repositoryParams protocol.CreateUserRepositoryParams
	repositoryParams.Name = params.Name
	repositoryParams.Email = params.Email
	repositoryParams.Password = params.Password
	err := uc.create_user_repository.CreateUser(repositoryParams)
	return err
}

func MakeCreateUser(create_user_repository protocol.CreateUserRepository) UseCase {
	return UseCase{create_user_repository}
}

func MakeUserRequest() domain.CreateUserParams {
	return domain.CreateUserParams{
		Name:     faker.Name(),
		Email:    faker.Email(),
		Password: faker.Password(),
	}
}

func TestCreateUserRepositoryCallTimes(t *testing.T) {
	createUserRepoSpy := new(mock.CreateUserRepository)
	sut := MakeCreateUser(createUserRepoSpy)

	sut.CreateUser(MakeUserRequest())

	if createUserRepoSpy.Count != 1 {
		t.Error("CreateUser method from CreateUserRepository was called more than one time or not called")
	}
}
