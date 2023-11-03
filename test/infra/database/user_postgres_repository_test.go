package database

import (
	"domain/model"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"test/domain/mock"
)

type UserPostgresRepository struct {
	gorm.Model
	database *gorm.DB
}

func (db *UserPostgresRepository) CheckByEmail(email string) (bool, error) {
	var count int64
	db.database.Model(&model.User{}).Where("email = ?", email).Count(&count)
	if count >= 1 {
		return true, nil
	}
	return false, nil
}

func MakeDatabase() (*gorm.DB, error) {
	DATABASE_URI := "postgres://postgres:postgres@localhost:5432/postgres"
	db, err := gorm.Open(postgres.Open(DATABASE_URI), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&model.User{})
	return db, nil
}

func TestShouldReturnTrueIfEmailIsAlreadyUsed(t *testing.T) {
	db, err := MakeDatabase()
	if err != nil {
		t.Error(err)
	}
	user := mock.MakeUser()
	db.Create(user)
	sut := &UserPostgresRepository{database: db}

	exists, _ := sut.CheckByEmail(user.Email)

	if !exists {
		t.Error("CheckByEmail should return true, but return false")
	}
}
