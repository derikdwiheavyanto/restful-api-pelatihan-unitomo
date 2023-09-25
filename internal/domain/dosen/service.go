package dosen

import (
	"time"

	"github.com/gofrs/uuid"
)

type Service interface {
	Create(input InputDosen) (dosen Dosen, err error)
	GetAllData() ([]Dosen, error)
	GetDataById(id uuid.UUID) (Dosen, error)
	Update(id uuid.UUID, inputDosen InputDosen) (Dosen, error)
	Delete(id uuid.UUID) error
}

type service struct {
	r Repository
}

func Newservice(r Repository) *service {
	return &service{r: r}
}

func (s *service) Create(input InputDosen) (dosen Dosen, err error) {
	newId, _ := uuid.NewV4()
	dosen.IdDosen = newId
	dosen.Nama = input.Nama
	dosen.Alamat = input.Alamat
	dosen.JenisKelamin = input.JenisKelamin

	dosen, err = s.r.Save(dosen)
	if err != nil {
		return
	}

	return
}

func (s *service) GetAllData() ([]Dosen, error) {
	return s.r.GetAllData()
}

func (s *service) GetDataById(id uuid.UUID) (Dosen, error) {
	return s.r.GetDataById(id)
}

func (s *service) Update(id uuid.UUID, inputDosen InputDosen) (dosen Dosen, err error) {
	dosen, err = s.r.GetDataById(id)
	if err != nil {
		return
	}
	dosen.Nama = inputDosen.Nama
	dosen.JenisKelamin = inputDosen.JenisKelamin
	dosen.Alamat = inputDosen.Alamat
	dosen.UpdatedAt = time.Now()

	dosen, err = s.r.Update(dosen)
	if err != nil {
		return
	}

	return
}

func (s *service) Delete(id uuid.UUID) error {
	_, err := s.r.GetDataById(id)
	if err != nil {
		return err
	}
	return s.r.Delete(id)
}
