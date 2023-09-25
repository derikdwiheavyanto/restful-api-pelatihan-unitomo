package dosen

import (
	"time"

	"github.com/gofrs/uuid"
)

type Dosen struct {
	IdDosen      uuid.UUID `json:"id_dosen" db:"id_dosen"`
	Nama         string    `json:"nama" db:"nama"`
	JenisKelamin string    `json:"jenis_kelamin" db:"jenis_kelamin"`
	Alamat       string    `json:"alamat" db:"alamat"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time `json:"updatedAt" db:"updated_at"`
}
