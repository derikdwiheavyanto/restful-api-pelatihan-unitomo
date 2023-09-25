package matakuliah

type Matakuliah struct {
	KodeMK string `json:"kode_mk" db:"kode_mk"`
	NamaMK string `json:"nama_mk" db:"nama_mk"`
	SKS    int    `json:"sks" db:"sks"`
}
