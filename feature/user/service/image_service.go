package service

import (
	"Crud-Gin/dto"
	repo "Crud-Gin/feature/user/repository"
	"Crud-Gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImageService interface {
	CreateImages(ctx *gin.Context, req *dto.ImagesReq) (*models.Image, error)
}

type imageServiceImpl struct {
	imageRepository repo.ImageRepository
}

func NewImageService(imageRepository repo.ImageRepository) ImageService {
	return &imageServiceImpl{imageRepository: imageRepository}
}

func (ur *imageServiceImpl) CreateImages(ctx *gin.Context, req *dto.ImagesReq) (*models.Image, error) {
	image := models.Image{
		Judul: req.Judul,
	}

	if req.Foto != nil {
		image.Foto = &req.Foto.Filename
	}

	data, err := ur.imageRepository.CreateImages(image)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": err.Error(),
		})
	}
	return data, nil

}
