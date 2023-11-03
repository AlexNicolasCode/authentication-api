package usecase

import (
	cryptography "src/data/protocol/cryptography"
	database "src/data/protocol/database"
	domain "src/domain/usecase"
)

type DbCreateUser struct {
	check_user_by_email_repository database.CheckUserByEmailRepository
	create_user_repository         database.CreateUserRepository
	hasher                         cryptography.Hasher
}

func (uc *DbCreateUser) CreateUser(params domain.CreateUserParams) (bool, error) {
	exists, err := uc.check_user_by_email_repository.CheckByEmail(params.Email)
	if err != nil {
		return false, err
	}
	if exists {
		return false, nil
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

func MakeDbCreateUser(
	check_user_by_email_repository database.CheckUserByEmailRepository,
	create_user_repository database.CreateUserRepository,
	hasher cryptography.Hasher,
) DbCreateUser {
	return DbCreateUser{check_user_by_email_repository, create_user_repository, hasher}
}
