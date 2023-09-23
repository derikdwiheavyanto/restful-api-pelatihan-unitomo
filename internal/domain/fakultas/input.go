package fakultas

type InputFakultas struct {
	Name            string `json:"name" db:"name" binding:"required"`
	JumlahJurusan   int    `json:"jmlh_jurusan" db:"jmlh_jurusan" binding:"required"`
	JumlahMahasiswa int    `json:"jmlh_mahasiswa" db:"jmlh_mahasiswa" binding:"required"`
	JumlahDosen     int    `json:"jmlh_dosen" db:"jmlh_dosen" binding:"required"`
}

type UpdateFakultas struct {
	Name            string `json:"name" db:"name" binding:"required"`
	JumlahJurusan   int    `json:"jmlh_jurusan" db:"jmlh_jurusan" binding:"required"`
	JumlahMahasiswa int    `json:"jmlh_mahasiswa" db:"jmlh_mahasiswa" binding:"required"`
	JumlahDosen     int    `json:"jmlh_dosen" db:"jmlh_dosen" binding:"required"`
}


