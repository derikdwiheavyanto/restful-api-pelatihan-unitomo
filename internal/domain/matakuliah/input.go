package matakuliah

type InputMatakuliah struct {
	KodeMK string `json:"kode_mk" db:"kode_mk" binding:"required"`
	NamaMK string `json:"nama_mk" db:"nama_mk" binding:"required"`
	SKS    int    `json:"sks" db:"sks" binding:"required"`
}
