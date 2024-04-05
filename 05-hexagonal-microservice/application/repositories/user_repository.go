package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/pedro-hca/golang-studies/05-rest-api/domain"
)

type UserRepository interface {
	GetByID(id string) (*domain.User, error)
}

type UserRepositoryDB struct {
	DB *gorm.DB
}

func (repo *UserRepositoryDB) GetByID(id string) (*domain.User, error) {
	var user domain.User
	repo.DB.First(&user, id)
	return &user, nil
}
