package usercontroller

import (
	"fmt"
	"net/http"
	"pomo/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// The `repository` struct is the concrete implementation of the `Repository` interface.
type repository struct {
	db *gorm.DB
}

// the `NewUserRepository` function is the constructor of the `repository` struct that
// creates the new instance of the `repository` struct
func NewUserRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

// The `Repository` interface defines a contract for the repository responsible for handling auth-related operations
type Repository interface {
	GetUserByUUID(uuid uuid.UUID) (*models.UserModel, int)
}

func (r *repository) GetUserByUUID(uuid uuid.UUID) (*models.UserModel, int) {
	var user models.UserModel
	checkAccount := r.db.First(&user, "id = ?", fmt.Sprint(uuid))
	if checkAccount.Error != nil {
		return nil, http.StatusNotFound
	}
	return &user, http.StatusOK
}

func (r *repository) Logout(input *models.UserModel) (int, error) {
	return http.StatusOK, nil
}
