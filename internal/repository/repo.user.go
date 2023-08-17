package repo

import (
	"fmt"
	"net/http"
	"pomo/internal/models"

	"github.com/google/uuid"
)

func (r *repository) GetUserByUUID(uuid uuid.UUID) (*models.UserModel, int) {
	var user models.UserModel
	checkAccount := r.db.First(&user, "id = ?", fmt.Sprint(uuid))
	if checkAccount.Error != nil {
		return nil, http.StatusNotFound
	}
	return &user, http.StatusOK
}
