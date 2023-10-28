package craete_user_test

import (
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"

	database "finances-api/src/data/protocol/database"
	uc "finances-api/src/data/usecase"
	domain "finances-api/src/domain/usecase"
	spy "finances-api/test/domain/usecase"
)

type UseCase struct {
	create_user_repository database.CreateUserRepository
}

func (uc *UseCase) CreateUser(params domain.CreateUserParams) error {
	var repositoryParams database.CreateUserRepositoryParams
	repositoryParams.Name = params.Name
	repositoryParams.Email = params.Email
	repositoryParams.Password = params.Password
	repositoryParams.CreatedAt = int(time.Now().Unix())
	err := uc.create_user_repository.CreateUser(repositoryParams)
	return err
}

func MakeCreateUser(create_user_repository database.CreateUserRepository) UseCase {
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
	createUserRepoSpy := new(spy.CreateUserRepository)
	sut := uc.MakeCreateUser(createUserRepoSpy)

	sut.CreateUser(MakeUserRequest())

	if createUserRepoSpy.Count != 1 {
		t.Error("CreateUser method from CreateUserRepository was called more than one time or not called")
	}
}
