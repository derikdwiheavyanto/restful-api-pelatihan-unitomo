package user

import (
	"time"

	"github.com/gofrs/uuid"
)

type RegisterUserInput struct {
	Name       string `json:"name" db:"name" binding:"required"`
	Occupation string `json:"occupation" db:"occupation" binding:"required"`
	Email      string `json:"email" db:"email" binding:"required"`
	Password   string `json:"password" db:"password" binding:"required"`
}

// type LoginInput struct {
// 	Email    string `json:"email" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// }

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type UpdateInput struct {
	ID             uuid.UUID `json:"id" db:"id" binding:"required"`
	Name           string    `json:"name" db:"name" binding:"required"`
	Occupation     string    `json:"occupation" db:"occupation" binding:"required"`
	Email          string    `json:"email" db:"email" binding:"required"`
	Password       string    `json:"password" db:"password" binding:"required"`
	AvatarFileName *string   `json:"avatarFileName" db:"avatar_file_name" binding:"required"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}
