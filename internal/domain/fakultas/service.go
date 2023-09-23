package fakultas

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

type Service interface {
	InputFakultas(InputFakultas) (Fakultas, error)
	GetAllData() ([]Fakultas, error)
	GetDataById(id uuid.UUID) (Fakultas, error)
	Update(id uuid.UUID, fakultas UpdateFakultas) (Fakultas, error)
	Delete(id uuid.UUID) error
	GetTotal() (Total, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) InputFakultas(input InputFakultas) (Fakultas, error) {
	fakultas := Fakultas{}

	newID, _ := uuid.NewV4()
	fakultas.ID = newID
	fakultas.Name = input.Name
	fakultas.JumlahDosen = input.JumlahDosen
	fakultas.JumlahMahasiswa = input.JumlahMahasiswa
	fakultas.JumlahJurusan = input.JumlahJurusan

	newFakultas, err := s.repository.Save(fakultas)
	if err != nil {

		return newFakultas, err
	}

	return newFakultas, nil

}

func (s *service) GetAllData() (data []Fakultas, err error) {

	return s.repository.GetAllData()
}

func (s *service) GetDataById(id uuid.UUID) (data Fakultas, err error) {
	return s.repository.GetDataById(id)
}

func (s *service) Update(id uuid.UUID, input UpdateFakultas) (Fakultas, error) {

	fakultas, err := s.repository.GetDataById(id)
	if err != nil {
		fmt.Println(err)
		return fakultas, err
	}

	fakultas.JumlahDosen = input.JumlahDosen
	fakultas.JumlahJurusan = input.JumlahJurusan
	fakultas.JumlahMahasiswa = input.JumlahMahasiswa
	fakultas.Name = input.Name
	fakultas.UpdatedAt = time.Now()

	fakultas, err = s.repository.Update(fakultas)
	if err != nil {
		fmt.Println(err)
		return fakultas, err
	}

	return fakultas, nil
}

func (s *service) Delete(id uuid.UUID) error {
	_, err := s.repository.GetDataById(id)
	if err != nil {
		return err
	}

	err = s.repository.Delete(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *service) GetTotal() (total Total, err error) {
	fakultas, err := s.repository.GetAllData()
	if err != nil {
		return
	}

	for _, f := range fakultas {
		total.TotalDosen += f.JumlahDosen
		total.TotalJurusan += f.JumlahJurusan
		total.TotalMahasiswa += f.JumlahMahasiswa
	}

	total.Fakultas = fakultas

	return

}
