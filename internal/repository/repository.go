package repo

import (
	"pomo/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// The `repository` struct is the concrete implementation of the `Repository` interface.
type repository struct {
	db *gorm.DB
}

// the `NewRepository` function is the constructor of the `repository` struct that
// creates the new instance of the `repository` struct
func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

// The `Repository` interface defines a contract for the repository responsible for handling auth-related operations
type Repository interface {
	Login(input *models.UserModel) (*models.UserModel, int)
	Register(input *models.UserModel) (*models.UserModel, int)
	Logout(ctx *gin.Context) int
	GetUserByUUID(uuid uuid.UUID) (*models.UserModel, int)
}
