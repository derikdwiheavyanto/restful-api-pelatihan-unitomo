package fakultas

import (
	"time"

	"github.com/gofrs/uuid"
)

type Fakultas struct {
	ID              uuid.UUID `json:"id" db:"id"`
	Name            string    `json:"name" db:"name"`
	JumlahJurusan   int       `json:"jmlh_jurusan" db:"jmlh_jurusan"`
	JumlahMahasiswa int       `json:"jmlh_mahasiswa" db:"jmlh_mahasiswa"`
	JumlahDosen     int       `json:"jmlh_dosen" db:"jmlh_dosen"`
	CreatedAt       time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt       time.Time `json:"updatedAt" db:"updated_at"`
}

type DeleteFakultas struct {
	ID uuid.UUID `json:"id" db:"id" binding:"required"`
}

type Total struct {
	Fakultas       []Fakultas `json:"fakultas"`
	TotalJurusan   int        `json:"total_jurusan"`
	TotalMahasiswa int        `json:"total_mahasiswa"`
	TotalDosen     int        `json:"total_dosen"`
}
