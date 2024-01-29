package controller

import (
	"Crud-Gin/dto"
	"Crud-Gin/feature/user/service"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ImageControllerImpl struct {
	imageService service.ImageService
}

func NewImageController(ctx *gin.Context, imageService service.ImageService) *ImageControllerImpl {
	return &ImageControllerImpl{imageService: imageService}
}

func (ur *ImageControllerImpl) CreateImages(ctx *gin.Context) {
	images := dto.ImagesReq{}
	if err := ctx.ShouldBind(&images); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Massage": "err",
		})
		return
	}
	if images.Foto != nil {
		if err := os.MkdirAll("/public/public", 0755); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Massage": "err",
			})
			return
		}
		//renamme picture

		ext := filepath.Ext(images.Foto.Filename)
		newName := uuid.New().String() + ext

		dst := filepath.Join("public/picture", filepath.Base(newName))

		ctx.SaveUploadedFile(images.Foto, dst)
	}

	// userId := 1

	data, err := ur.imageService.CreateImages(ctx, &images)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"massage": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"massage": "Succes",
		"data":    data,
	})

}
