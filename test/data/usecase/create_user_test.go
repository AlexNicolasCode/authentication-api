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

type SutSetupTypes struct {
	sut               UseCase
	createUserRepoSpy *mock.CreateUserRepository
}

func MakeSutSetup() SutSetupTypes {
	createUserRepoSpy := new(mock.CreateUserRepository)
	sut := MakeCreateUser(createUserRepoSpy)
	return SutSetupTypes{sut, createUserRepoSpy}
}

func MakeUserRequest() domain.CreateUserParams {
	return domain.CreateUserParams{
		Name:     faker.Name(),
		Email:    faker.Email(),
		Password: faker.Password(),
	}
}

func TestShouldCallCreateUserRepositoryOneTime(t *testing.T) {
	setup := MakeSutSetup()

	setup.sut.CreateUser(MakeUserRequest())

	if setup.createUserRepoSpy.Count != 1 {
		t.Error("CreateUser method from CreateUserRepository was called more than one time or not called")
	}
}

func TestShouldCallCreateUserRepositoryWithCorrectParams(t *testing.T) {
	setup := MakeSutSetup()
	fakeParams := MakeUserRequest()

	setup.sut.CreateUser(fakeParams)

	if setup.createUserRepoSpy.Params.Email != fakeParams.Email ||
		setup.createUserRepoSpy.Params.Name != fakeParams.Name ||
		setup.createUserRepoSpy.Params.Password != fakeParams.Password {
		t.Error("CreateUser method from CreateUserRepository was called more than one time or not called")
	}
}
