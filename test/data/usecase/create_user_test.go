package data_test

import (
	"testing"

	"github.com/bxcodec/faker/v3"

	cryptography "src/data/protocol/cryptography"
	database "src/data/protocol/database"
	domain "src/domain/usecase"
	mock "test/infra"
)

type UseCase struct {
	check_user_by_email_repository database.CheckUserByEmailRepository
	create_user_repository         database.CreateUserRepository
	hasher                         cryptography.Hasher
}

func (uc *UseCase) CreateUser(params domain.CreateUserParams) (bool, error) {
	_, err := uc.check_user_by_email_repository.CheckByEmail(params.Email)
	if err != nil {
		return false, err
	}
	hashedPassword, err := uc.hasher.Hash(params.Password)
	if err != nil {
		return false, err
	}
	repoParams := database.CreateUserRepositoryParams{
		Name:     params.Name,
		Email:    params.Email,
		Password: hashedPassword,
	}
	err = uc.create_user_repository.CreateUser(repoParams)
	return false, err
}

func MakeCreateUser(
	check_user_by_email_repository database.CheckUserByEmailRepository,
	create_user_repository database.CreateUserRepository,
	hasher cryptography.Hasher,
) UseCase {
	return UseCase{check_user_by_email_repository, create_user_repository, hasher}
}

type SutSetupTypes struct {
	sut                     UseCase
	checkUserByEmailRepoSpy *mock.CheckUserByEmailRepository
	createUserRepoSpy       *mock.CreateUserRepository
	hasher                  *mock.Hasher
}

func MakeSutSetup() SutSetupTypes {
	checkUserByEmailRepoSpy := new(mock.CheckUserByEmailRepository)
	createUserRepoSpy := new(mock.CreateUserRepository)
	hasherSpy := new(mock.Hasher)
	sut := MakeCreateUser(checkUserByEmailRepoSpy, createUserRepoSpy, hasherSpy)
	return SutSetupTypes{sut, checkUserByEmailRepoSpy, createUserRepoSpy, hasherSpy}
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

	_, err := setup.sut.CreateUser(MakeUserRequest())

	if err.Error() != setup.createUserRepoSpy.ErrorMessage {
		t.Error("CreateUser return incorrect error when CreateUserRepoSpy throws")
	}
}

func TestShouldReturnNoneErrorOnSuccess(t *testing.T) {
	setup := MakeSutSetup()

	_, err := setup.sut.CreateUser(MakeUserRequest())

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

func TestShouldCallCheckUserByEmailWithCorrectParam(t *testing.T) {
	setup := MakeSutSetup()
	fakeParams := MakeUserRequest()

	setup.sut.CreateUser(fakeParams)

	if fakeParams.Email != setup.checkUserByEmailRepoSpy.Email {
		t.Error("CreateUser return incorrect error when CheckUserByEmail throws")
	}
}

func TestShouldThrowIfCheckUserByEmailThrows(t *testing.T) {
	setup := MakeSutSetup()
	setup.checkUserByEmailRepoSpy.ErrorMessage = "Mocked Error"

	_, err := setup.sut.CreateUser(MakeUserRequest())

	if err.Error() != setup.checkUserByEmailRepoSpy.ErrorMessage {
		t.Error("CreateUser return incorrect error when CheckUserByEmail throws")
	}
}
