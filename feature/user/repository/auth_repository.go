package repository

import (
	"Crud-Gin/models"

	"gorm.io/gorm"
)

type UserRespo interface {
	FindAll() []models.User
	FindOne(id int) models.User
	Save(user models.User) (*models.User, error)
	Update(user models.User) (*models.User, error)
	Delete(user models.User) (*models.User, error)
}

type UserRepoImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRespo {
	return &UserRepoImpl{db: db}
}

func (ur *UserRepoImpl) FindAll() []models.User {
	user := []models.User{}
	ur.db.Find(&user)

	return user
}
func (ur *UserRepoImpl) FindOne(id int) models.User {
	user := models.User{}

	ur.db.First(&user, id)

	return user
}
func (ur *UserRepoImpl) Save(user models.User) (*models.User, error) {
	result := ur.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
func (ur *UserRepoImpl) Update(user models.User) (*models.User, error) {
	result := ur.db.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
func (ur *UserRepoImpl) Delete(user models.User) (*models.User, error) {
	result := ur.db.Delete(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
