package routes

import (
	"Crud-Gin/feature/user/controller"
	"Crud-Gin/feature/user/repository"
	"Crud-Gin/feature/user/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ImagesRouter(r *gin.Engine, db *gorm.DB) {
	imageRepo := repository.NewImageRepository(db)
	imageService := service.NewImageService(imageRepo)
	imageController := controller.NewImageController(ctx, imageService)

	images := r.Group("/images")
	{
		// images.GET("/", imageController.Getall)
		// images.GET("/:id", imageController.GetByID)
		images.POST("/", imageController.CreateImages)
		// images.PATCH("/:id", userController.Update)
		// images.DELETE("/:id", userController.Delete)
	}

}
