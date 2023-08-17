package services

import (
	"net/http"
	"pomo/internal/models"
	modelsInput "pomo/internal/models/input"
	"pomo/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func (s *service) Login(input *modelsInput.LoginInput) (*models.UserModel, int) {
	userModel := models.UserModel{
		Email:    input.Email,
		Password: input.Password,
	}
	return s.repository.Login(&userModel)
}

func (s *service) Register(input *modelsInput.RegisterInput) (*models.UserModel, int) {
	fileName := ""
	if input.Image != nil {
		fname, err := utils.UploadedFormDataImg(input.Image)
		if err != nil {
			return nil, http.StatusInternalServerError
		}
		fileName = fname
	}

	userModel := models.UserModel{
		Name:     input.Name,
		Email:    strings.ToLower(input.Email),
		Password: input.Password,
		Image:    fileName,
		Provider: input.Provider,
		Role:     "user",
		Verified: true,
	}
	return s.repository.Register(&userModel)
}

func (s *service) Logout(ctx *gin.Context) int {
	return s.repository.Logout(ctx)
}
