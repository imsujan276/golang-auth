package repo

import (
	"fmt"
	"net/http"
	"pomo/internal/models"

	"github.com/google/uuid"
)

func (r *repository) GetUserByCustomField(field string, value string) (*models.UserModel, error) {
	var user models.UserModel
	result := r.db.First(&user, fmt.Sprintf("%s = ?", field), value)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *repository) GetUserByUUID(uuid uuid.UUID) (*models.UserModel, int) {
	var user models.UserModel
	checkAccount := r.db.First(&user, "id = ?", fmt.Sprint(uuid))
	if checkAccount.Error != nil {
		return nil, http.StatusNotFound
	}
	return &user, http.StatusOK
}

func (r *repository) UpdateUser(input *models.UserModel) (*models.UserModel, error) {
	err := r.db.Save(input).Error
	if err != nil {
		return nil, err
	}
	return input, nil
}
