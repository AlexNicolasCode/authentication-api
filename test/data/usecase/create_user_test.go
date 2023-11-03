package data_test

import (
	"testing"

	"github.com/bxcodec/faker/v3"

	usecase "src/data/usecase"
	domain "src/domain/usecase"
	mock "test/infra"
)

type SutSetupTypes struct {
	sut                     usecase.DbCreateUser
	checkUserByEmailRepoSpy *mock.CheckUserByEmailRepository
	createUserRepoSpy       *mock.CreateUserRepository
	hasher                  *mock.Hasher
}

func MakeSutSetup() SutSetupTypes {
	checkUserByEmailRepoSpy := new(mock.CheckUserByEmailRepository)
	createUserRepoSpy := new(mock.CreateUserRepository)
	hasherSpy := new(mock.Hasher)
	sut := usecase.MakeDbCreateUser(checkUserByEmailRepoSpy, createUserRepoSpy, hasherSpy)
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

func TestShouldReturnFalseIfEmailAlreadyUsed(t *testing.T) {
	setup := MakeSutSetup()
	setup.checkUserByEmailRepoSpy.Result = true

	exists, _ := setup.sut.CreateUser(MakeUserRequest())

	if exists {
		t.Error("CreateUser return incorrect error when CheckUserByEmail throws")
	}
}
