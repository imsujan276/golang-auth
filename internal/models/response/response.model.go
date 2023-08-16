package modelsResponse

import (
	"github.com/google/uuid"
)

type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Image    string    `json:"image"`
	Role     string    `json:"-"`
	Verified bool      `json:"verified"`
	Provider string    `json:"provider"`
}

type AccessTokeResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
	User         UserResponse
}
