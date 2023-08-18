package services

import (
	"pomo/internal/models"
	modelsInput "pomo/internal/models/input"
	repo "pomo/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type service struct {
	repository repo.Repository
}

// the `NewService` function is the constructor of the `service` struct that
// creates the new instance of the `service` struct
func NewService(repository repo.Repository) *service {
	return &service{repository: repository}
}

// The `Service` interface defines a contract for the service responsible for handling auth-related operations
type Service interface {
	// Auth
	Login(input *modelsInput.LoginInput) (*models.UserModel, int)
	Register(input *modelsInput.RegisterInput) (*models.UserModel, int)
	Logout(ctx *gin.Context) int
	DeleteUser(user *models.UserModel) error

	// User
	GetUserByUUID(uuid uuid.UUID) (*models.UserModel, int)
	GetMe(ctx *gin.Context) (*models.UserModel, int)
	GetUserByEmail(email string) (*models.UserModel, error)
	GetUserByCustomField(field string, value string) (*models.UserModel, error)
	UpdateUser(input *models.UserModel) (*models.UserModel, error)
}
