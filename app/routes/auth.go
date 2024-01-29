package routes

import (
	"Crud-Gin/feature/user/controller"
	"Crud-Gin/feature/user/repository"
	"Crud-Gin/feature/user/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	ctx *gin.Context
)

func Api(r *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService, ctx)

	auth := r.Group("/user")
	{
		auth.GET("/", userController.Getall)
		auth.GET("/:id", userController.GetByID)
		auth.POST("/", userController.Create)
		auth.PATCH("/:id", userController.Update)
		auth.DELETE("/:id", userController.Delete)
	}

}
