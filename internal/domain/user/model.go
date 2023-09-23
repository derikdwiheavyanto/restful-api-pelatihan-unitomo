package user

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id" db:"id"`
	Name           string    `json:"name" db:"name"`
	Occupation     string    `json:"occupation" db:"occupation"`
	Email          string    `json:"email" db:"email"`
	PasswordHash   string    `json:"passwordHash" db:"password_hash"`
	AvatarFileName string    `json:"avatarFileName" db:"avatar_file_name"`
	Role           string    `json:"role" db:"role"`
	CreatedAt      time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt      time.Time `json:"updateAt" db:"update_at"`
}

type UserDto struct {
	ID             uuid.UUID `json:"id" db:"id"`
	Name           *string   `json:"name" db:"name"`
	Occupation     *string   `json:"occupation" db:"occupation"`
	Email          *string   `json:"email" db:"email"`
	PasswordHash   *string   `json:"passwordHash" db:"password_hash"`
	AvatarFileName *string   `json:"avatarFileName" db:"avatar_file_name"`
	Role           *string   `json:"role" db:"role"`
	CreatedAt      time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt      time.Time `json:"updatedAt" db:"updated_at"`
}

type UserFakultas struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Name     *string   `json:"name" db:"name"`
	Email    *string   `json:"email" db:"email"`
	Fakultas *string   `json:"fakultas" db:"fakultas"`
}
