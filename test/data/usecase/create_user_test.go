package data_test

import (
	"testing"

	"github.com/bxcodec/faker/v3"

	cryptography "src/data/protocol/cryptography"
	database "src/data/protocol/database"
	domain "src/domain/usecase"
	mock_data "test/domain/usecase"
	mock_infra "test/infra"
)

type UseCase struct {
	create_user_repository database.CreateUserRepository
	hasher                 cryptography.Hasher
}

func (uc *UseCase) CreateUser(params domain.CreateUserParams) error {
	hashedPassword, err := uc.hasher.Hash(params.Password)
	if err != nil {
		return err
	}
	repoParams := database.CreateUserRepositoryParams{
		Name:     params.Name,
		Email:    params.Email,
		Password: hashedPassword,
	}
	err = uc.create_user_repository.CreateUser(repoParams)
	return err
}

func MakeCreateUser(create_user_repository database.CreateUserRepository, hasher cryptography.Hasher) UseCase {
	return UseCase{create_user_repository, hasher}
}

type SutSetupTypes struct {
	sut               UseCase
	createUserRepoSpy *mock_data.CreateUserRepository
	hasher            *mock_infra.Hasher
}

func MakeSutSetup() SutSetupTypes {
	createUserRepoSpy := new(mock_data.CreateUserRepository)
	hasherSpy := new(mock_infra.Hasher)
	sut := MakeCreateUser(createUserRepoSpy, hasherSpy)
	return SutSetupTypes{sut, createUserRepoSpy, hasherSpy}
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
		setup.createUserRepoSpy.Params.Password != setup.hasher.Result {
		t.Error("CreateUser method from CreateUserRepository was called more than one time or not called")
	}
}

func TestShouldThrowWhenCreateUserRepositoryThrows(t *testing.T) {
	setup := MakeSutSetup()
	setup.createUserRepoSpy.ErrorMessage = "Mocked Error"

	err := setup.sut.CreateUser(MakeUserRequest())

	if err.Error() != setup.createUserRepoSpy.ErrorMessage {
		t.Error("CreateUser return incorrect error when CreateUserRepoSpy throws")
	}
}

func TestShouldReturnNoneErrorOnSuccess(t *testing.T) {
	setup := MakeSutSetup()

	err := setup.sut.CreateUser(MakeUserRequest())

	if err != nil {
		t.Error("CreateUser return some error")
	}
}

func TestShouldCallHasherWithCorrectParam(t *testing.T) {
	setup := MakeSutSetup()
	fakeParams := MakeUserRequest()

	setup.sut.CreateUser(fakeParams)

	if setup.hasher.Plaintext != fakeParams.Password {
		t.Error("CreateUser return some error")
	}
}
