package controller

import (
	"Crud-Gin/feature/user/service"
	"Crud-Gin/middleware"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService, ctx *gin.Context) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Getall(ctx *gin.Context) {
	data := uc.userService.GetAll()

	ctx.JSON(http.StatusOK, gin.H{
		"massage": "Succes",
		"data":    data,
	})
}

func (uc *UserController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"massage": "Error",
			"Data":    err,
		})
	}
	data := uc.userService.GetByID(id)

	ctx.JSON(http.StatusOK, gin.H{
		"massage": "Succes",
		"data":    data,
	})
}

func (uc *UserController) Create(ctx *gin.Context) {

	// Generate Jwt Token
	errE := godotenv.Load()
	if errE != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Massage": "Env Err",
		})
		ctx.Abort()
		return
	}
	secretKey := os.Getenv("SECRET_JWT")

	data, err := uc.userService.Create(ctx, []byte(secretKey))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Massage": err,
		})
		ctx.Abort()
		return
	}

	jwtToken, err := middleware.GenerateToken(data.Name, []byte(secretKey))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Massage": "Env Err",
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"massage": "Succes",
		"data":    data,
		"token":   jwtToken,
	})
}

func (uc *UserController) Delete(ctx *gin.Context) {

	data, err := uc.userService.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Massage": err,
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"massage": "Succes",
		"data":    data,
	})
}

func (uc *UserController) Update(ctx *gin.Context) {

	data, err := uc.userService.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Massage": err,
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"massage": "Succes",
		"data":    data,
	})
}
