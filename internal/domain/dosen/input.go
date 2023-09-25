package dosen

type InputDosen struct {
	Nama         string `json:"nama" db:"nama" binding:"required"`
	JenisKelamin string `json:"jenis_kelamin" db:"jenis_kelamin" binding:"required"`
	Alamat       string `json:"alamat" db:"alamat" binding:"required"`
}
