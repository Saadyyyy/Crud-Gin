package dto

import "mime/multipart"

type CreateUserDto struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"email"`
}

type UpdateUserDto struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"email"`
}

type ImageRes struct {
	Judul string `json:"judul"`
	Foto  string `json:"foto"`
}

type ImagesReq struct {
	Judul string `form:"judul"`
	Foto  *multipart.FileHeader
}
