package mengajar

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
)

type Service interface {
	Create(input InputMengajar) (mengajar Mengajar, err error)
	GetAllData() ([]Mengajar, error)
	GetDataById(id uuid.UUID, kodeMK string) (Mengajar, error)
	Update(inputMengajar InputMengajar) (Mengajar, error)
	Delete(id uuid.UUID, kodeMK string) error
}

type service struct {
	r Repository
}

func Newservice(r Repository) *service {
	return &service{r: r}
}

func (s *service) Create(input InputMengajar) (mengajar Mengajar, err error) {
	_, err = s.r.GetDataById(input.IdDosen, input.KodeMK)
	if err == nil {
		return mengajar, errors.New("dosen dan matkul terdaftar mengajar")
	}

	mengajar.IdDosen = input.IdDosen
	mengajar.KodeMK = input.KodeMK
	jadwalMulai, err := time.Parse(time.Kitchen, input.JadwalMulai)
	if err != nil {
		return
	}
	mengajar.JadwalMulai = jadwalMulai
	mengajar.JadwalBerakhir, err = time.Parse(time.Kitchen, input.JadwalBerakhir)
	if err != nil {
		return
	}
	mengajar, err = s.r.Save(mengajar)
	if err != nil {
		return
	}

	return
}

func (s *service) GetAllData() ([]Mengajar, error) {
	return s.r.GetAllData()
}

func (s *service) GetDataById(id uuid.UUID, kodeMK string) (Mengajar, error) {
	return s.r.GetDataById(id, kodeMK)
}

func (s *service) Update(inputDosen InputMengajar) (mengajar Mengajar, err error) {
	mengajar, err = s.r.GetDataById(inputDosen.IdDosen, inputDosen.KodeMK)
	if err != nil {
		return
	}

	mengajar.IdDosen = inputDosen.IdDosen
	mengajar.KodeMK = inputDosen.KodeMK
	jadwalMulai, err := time.Parse(time.Kitchen, inputDosen.JadwalMulai)
	if err != nil {
		return
	}
	mengajar.JadwalMulai = jadwalMulai
	mengajar.JadwalBerakhir, err = time.Parse(time.Kitchen, inputDosen.JadwalBerakhir)
	if err != nil {
		return
	}

	mengajar, err = s.r.Update(mengajar)
	if err != nil {
		return
	}

	return
}

func (s *service) Delete(idDosen uuid.UUID, kodeMK string) error {
	_, err := s.r.GetDataById(idDosen, kodeMK)
	if err != nil {
		return err
	}
	return s.r.Delete(idDosen, kodeMK)
}
