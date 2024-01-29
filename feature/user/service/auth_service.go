package service

import (
	"Crud-Gin/dto"
	"Crud-Gin/feature/user/repository"
	repo "Crud-Gin/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserService interface {
	GetAll() []repo.User
	GetByID(id int) repo.User
	Update(ctx *gin.Context) (*repo.User, error)
	Delete(ctx *gin.Context) (*repo.User, error)
	Create(ctx *gin.Context, secretKey []byte) (*repo.User, error)
}
type UserServiceImpl struct {
	userRepo repository.UserRespo
}

func (us *UserServiceImpl) GetAll() []repo.User {
	return us.userRepo.FindAll()
}

func (us *UserServiceImpl) GetByID(id int) repo.User {
	return us.userRepo.FindOne(id)
}

func (us *UserServiceImpl) Create(ctx *gin.Context, secreyKey []byte) (*repo.User, error) {
	var input dto.CreateUserDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return nil, err
	}
	user := repo.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	result, err := us.userRepo.Save(user)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (us *UserServiceImpl) Update(ctx *gin.Context) (*repo.User, error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return nil, err
	}

	var input dto.CreateUserDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return nil, err
	}

	user := repo.User{
		ID:       id,
		Name:     input.Name,
		Password: input.Password,
		Email:    input.Email,
	}

	result, err := us.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (us *UserServiceImpl) Delete(ctx *gin.Context) (*repo.User, error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return nil, err
	}

	user := repo.User{
		ID: id,
	}

	result, err := us.userRepo.Delete(user)

	if err != nil {
		return nil, err
	}

	return result, nil
}
func NewUserService(userRepo repository.UserRespo) UserService {
	return &UserServiceImpl{userRepo: userRepo}
}
