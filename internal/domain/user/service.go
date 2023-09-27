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
	Update(input UpdateInput) (UpdateInput, error)
	SaveAvatar(id uuid.UUID, safeLocation string) (UpdateInput, error)
	Delete(id uuid.UUID) error
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

func (s *service) Update(input UpdateInput) (UpdateInput, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return UpdateInput{}, err
	}

	input.UpdatedAt = time.Now()
	input.Password = string(hashPassword)
	input, err = s.repository.Update(input)
	if err != nil {
		return UpdateInput{}, err
	}

	return input, nil

}

func (s *service) SaveAvatar(id uuid.UUID, fileLocation string) (UpdateInput, error) {
	user, err := s.repository.GetDataById(id)
	if err != nil {
		return UpdateInput{}, err
	}

	updateInput := UpdateInput{}
	updateInput.ID = user.ID
	updateInput.Name = *user.Name
	updateInput.Occupation = *user.AvatarFileName
	updateInput.Email = *user.Email
	updateInput.UpdatedAt = user.UpdatedAt
	updateInput.AvatarFileName = &fileLocation

	updatedUser, err := s.repository.Update(updateInput)
	if err != nil {
		return updatedUser, err
	}
	return updatedUser, nil
}

func (s *service) Delete(id uuid.UUID) error {
	_, err := s.repository.GetDataById(id)
	if err != nil {
		return err
	}
	return s.repository.Delete(id)
}
