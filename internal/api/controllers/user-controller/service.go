package usercontroller

import (
	"pomo/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// The `service` struct is the concrete implementation of the `Service` interface.
type service struct {
	repository Repository
}

// the `NewUserService` function is the constructor of the `service` struct that
// creates the new instance of the `service` struct
func NewUserService(repository Repository) *service {
	return &service{repository: repository}
}

// The `Service` interface defines a contract for the service responsible for handling auth-related operations
type Service interface {
	GetUserByUUID(uuid uuid.UUID) (*models.UserModel, int)
	GetMe(ctx *gin.Context) (*models.UserModel, int)
}

func (s *service) GetUserByUUID(uuid uuid.UUID) (*models.UserModel, int) {
	return s.repository.GetUserByUUID(uuid)
}

func (s *service) GetMe(ctx *gin.Context) (*models.UserModel, int) {
	currentUser := ctx.MustGet("user").(models.UserModel)
	return s.repository.GetUserByUUID(currentUser.ID)
}
