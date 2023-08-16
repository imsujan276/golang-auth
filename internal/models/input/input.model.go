package modelsInput

import "mime/multipart"

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type RegisterInput struct {
	Name     string                `form:"name" validate:"required,min=3,max=100"`
	Email    string                `form:"email" validate:"required,email"`
	Password string                `form:"password" validate:"required,min=6,max=100"`
	Image    *multipart.FileHeader `form:"image"`
	Provider string                `form:"provider" validate:"required,oneof=email facebook google"`
}
