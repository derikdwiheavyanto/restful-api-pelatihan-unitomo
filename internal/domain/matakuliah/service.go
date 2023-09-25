package matakuliah

type Service interface {
	Create(input InputMatakuliah) (matakuliah Matakuliah, err error)
	GetAllData() ([]Matakuliah, error)
	GetDataById(kodeMK string) (Matakuliah, error)
	Update(inputMatakuliah InputMatakuliah) (Matakuliah, error)
	Delete(kodeMK string) error
}

type service struct {
	r Repository
}

func Newservice(r Repository) *service {
	return &service{r: r}
}

func (s *service) Create(input InputMatakuliah) (matakuliah Matakuliah, err error) {
	matakuliah.KodeMK = input.KodeMK
	matakuliah.NamaMK = input.NamaMK
	matakuliah.SKS = input.SKS

	matakuliah, err = s.r.Save(matakuliah)
	if err != nil {
		return
	}

	return
}

func (s *service) GetAllData() ([]Matakuliah, error) {
	return s.r.GetAllData()
}

func (s *service) GetDataById(kodeMK string) (Matakuliah, error) {
	return s.r.GetDataById(kodeMK)
}

func (s *service) Update(inputMatakuliah InputMatakuliah) (matakuliah Matakuliah, err error) {
	matakuliah, err = s.r.GetDataById(inputMatakuliah.KodeMK)
	if err != nil {
		return
	}
	matakuliah.KodeMK = inputMatakuliah.KodeMK
	matakuliah.NamaMK = inputMatakuliah.NamaMK
	matakuliah.SKS = inputMatakuliah.SKS

	matakuliah, err = s.r.Update(matakuliah)
	if err != nil {
		return
	}

	return
}

func (s *service) Delete(kodeMK string) error {
	_, err := s.r.GetDataById(kodeMK)
	if err != nil {
		return err
	}
	return s.r.Delete(kodeMK)
}
