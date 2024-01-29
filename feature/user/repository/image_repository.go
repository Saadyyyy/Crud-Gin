package repository

import (
	"Crud-Gin/models"

	"gorm.io/gorm"
)

type ImageRepository interface {
	CreateImages(image models.Image) (*models.Image, error)
}

type ImageRepositoryImpl struct {
	db *gorm.DB
}

// Create implements ImageRepository.
// func (*ImageRepositoryImpl) Create(image models.Image) (*models.Image, error) {
// 	panic("unimplemented")
// }

func NewImageRepository(db *gorm.DB) ImageRepository {
	return &ImageRepositoryImpl{db: db}
}

func (ur *ImageRepositoryImpl) CreateImages(image models.Image) (*models.Image, error) {
	result := ur.db.Create(&image)
	if result.Error != nil {
		return nil, result.Error
	}
	return &image, nil

}
