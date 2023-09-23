package user

import (
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repository Repository
}

func Newservice(repository Repository) *service {
	return &service{repository: repository}
}

type Service interface {
	RegisterUser(RegisterUserInput) (User, error)
	GetAllData() ([]UserDto, error)
	GetUsersFakultas() ([]UserFakultas, error)
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	newID, _ := uuid.NewV4()
	user.ID = newID
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)
	user.Role = "user"
	user.AvatarFileName = "woke"
	user.CreatedAt = time.Now()
	newUser, err := s.repository.Save(user)
	if err != nil {
		return user, err
	}

	return newUser, err
}

func (s *service) GetAllData() (data []UserDto, err error) {
	return s.repository.GetAllData()
}

func (s *service) GetUsersFakultas() ([]UserFakultas, error) {
	return s.repository.GetUsersFakultas()
}
