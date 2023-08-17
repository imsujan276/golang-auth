package services

import (
	"pomo/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *service) GetUserByCustomField(field string, value string) (*models.UserModel, error) {
	return s.repository.GetUserByCustomField(field, value)
}

func (s *service) GetUserByUUID(uuid uuid.UUID) (*models.UserModel, int) {
	return s.repository.GetUserByUUID(uuid)
}

func (s *service) GetMe(ctx *gin.Context) (*models.UserModel, int) {
	currentUser := ctx.MustGet("user").(models.UserModel)
	return s.repository.GetUserByUUID(currentUser.ID)
}

func (s *service) UpdateUser(input *models.UserModel) (*models.UserModel, error) {
	return s.repository.UpdateUser(input)
}
