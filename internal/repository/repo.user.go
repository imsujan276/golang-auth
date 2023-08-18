package repo

import (
	"fmt"
	"net/http"
	"pomo/internal/models"

	"github.com/google/uuid"
)

func (r *repository) GetUserByCustomField(field string, value string) (*models.UserModel, error) {
	var user models.UserModel
	if err := r.db.First(&user, fmt.Sprintf("%s = ?", field), value).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) GetUserByUUID(uuid uuid.UUID) (*models.UserModel, int) {
	var user models.UserModel
	if err := r.db.First(&user, "id = ?", fmt.Sprint(uuid)).Error; err != nil {
		return nil, http.StatusNotFound
	}
	return &user, http.StatusOK
}

func (r *repository) GetUserByEmail(email string) (*models.UserModel, error) {
	var user models.UserModel
	if err := r.db.First(&user, "email = ?", fmt.Sprint(email)).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) UpdateUser(input *models.UserModel) (*models.UserModel, error) {
	if err := r.db.Save(input).Error; err != nil {
		return nil, err
	}
	return input, nil
}
