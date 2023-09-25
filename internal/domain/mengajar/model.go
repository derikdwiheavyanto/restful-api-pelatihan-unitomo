package mengajar

import (
	"time"

	"github.com/gofrs/uuid"
)

type Mengajar struct {
	IdDosen        uuid.UUID `json:"id_dosen" db:"id_dosen"`
	KodeMK         string    `json:"kode_mk" db:"kode_mk"`
	JadwalMulai    time.Time `json:"jadwal_mulai" db:"jadwal_mulai"`
	JadwalBerakhir time.Time `json:"jadwal_berakhir" db:"jadwal_berakhir"`
}
