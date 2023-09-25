package mengajar

import (
	"database/sql"
	"errors"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

var (
	mengajarQuery = struct {
		Create     string
		Select     string
		SelectById string
		Update     string
		Delete     string
	}{
		Create:     `INSERT INTO mengajar(id_dosen,kode_mk,jadwal_mulai,jadwal_berakhir) VALUES(:id_dosen,:kode_mk,:jadwal_mulai,:jadwal_berakhir)`,
		Select:     `SELECT * FROM mengajar`,
		SelectById: `SELECT * FROM mengajar WHERE id_dosen = $1 AND kode_mk = $2`,
		Update:     `UPDATE mengajar SET jadwal_mulai = :jadwal_mulai, jadwal_berakhir= :jadwal_berakhir WHERE id_dosen = :id_dosen AND kode_mk = :kode_mk`,
		Delete:     `DELETE FROM mengajar WHERE id_dosen = $1 AND kode_mk=$2`,
	}
)

type Repository interface {
	Save(inputMengajar Mengajar) (mengajar Mengajar, err error)
	GetAllData() ([]Mengajar, error)
	GetDataById(id uuid.UUID, kodeMK string) (Mengajar, error)
	Update(Mengajar) (Mengajar, error)
	Delete(id uuid.UUID, kodeMK string) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Save(inputMengajar Mengajar) (mengajar Mengajar, err error) {
	stmt, err := r.db.PrepareNamed(mengajarQuery.Create)
	if err != nil {
		return
	}
	_, err = stmt.Exec(inputMengajar)
	if err != nil {
		return
	}
	return inputMengajar, nil
}

func (r *repository) GetAllData() (mengajars []Mengajar, err error) {
	rows, err := r.db.Queryx(mengajarQuery.Select)
	if err != nil {
		return
	}

	for rows.Next() {
		var mengajar Mengajar

		err = rows.StructScan(&mengajar)
		if err != nil {
			return
		}

		mengajars = append(mengajars, mengajar)
	}

	return
}

func (r *repository) GetDataById(idDosen uuid.UUID, kodeMK string) (mengajar Mengajar, err error) {
	rows, err := r.db.Queryx(mengajarQuery.SelectById, idDosen, kodeMK)

	if err == sql.ErrNoRows {
		return mengajar, errors.New("data tidak ditemukan")
	}

	if err != nil {
		return
	}

	if rows.Next() {
		err = rows.StructScan(&mengajar)
		if err != nil {
			return
		}
	} else {
		return mengajar, errors.New("data tidak ditemukan")
	}

	return
}

func (r *repository) Update(updateInput Mengajar) (dosen Mengajar, err error) {
	stmt, err := r.db.PrepareNamed(mengajarQuery.Update)
	if err != nil {
		return
	}

	_, err = stmt.Exec(updateInput)
	if err != nil {
		return
	}

	return updateInput, nil
}

func (r *repository) Delete(idDosen uuid.UUID, kodeMK string) error {
	_, err := r.db.Exec(mengajarQuery.Delete, idDosen, kodeMK)

	if err != nil {
		return err
	}

	return nil
}
