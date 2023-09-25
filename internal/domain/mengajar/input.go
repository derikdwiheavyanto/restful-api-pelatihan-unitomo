package mengajar

import (
	"github.com/gofrs/uuid"
)

type InputMengajar struct {
	IdDosen        uuid.UUID `json:"id_dosen" db:"id_dosen"  binding:"required"`
	KodeMK         string    `json:"kode_mk" db:"kode_mk"  binding:"required"`
	JadwalMulai    string    `json:"jadwal_mulai" db:"jadwal_mulai" binding:"required"`
	JadwalBerakhir string    `json:"jadwal_berakhir" db:"jadwal_berakhir" binding:"required"`
}
