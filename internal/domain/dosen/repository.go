package dosen

import (
	"database/sql"
	"errors"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

var (
	dosenQuery = struct {
		Create     string
		Select     string
		SelectById string
		Update     string
		Delete     string
	}{
		Create:     `INSERT INTO dosen(id_dosen,nama,jenis_kelamin,alamat) VALUES(:id_dosen,:nama,:jenis_kelamin,:alamat)`,
		Select:     `SELECT * FROM dosen`,
		SelectById: `SELECT * FROM dosen WHERE id_dosen = $1`,
		Update:     `UPDATE dosen SET nama = :nama,jenis_kelamin = :jenis_kelamin,alamat = :alamat,updated_at = :updated_at WHERE id_dosen = :id_dosen`,
		Delete:     `DELETE FROM dosen WHERE id_dosen = $1`,
	}
)

type Repository interface {
	Save(input Dosen) (dosen Dosen, err error)
	GetAllData() ([]Dosen, error)
	GetDataById(uuid.UUID) (Dosen, error)
	Update(Dosen) (Dosen, error)
	Delete(id uuid.UUID) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Save(input Dosen) (dosen Dosen, err error) {
	stmt, err := r.db.PrepareNamed(dosenQuery.Create)
	if err != nil {
		return
	}
	_, err = stmt.Exec(input)
	if err != nil {
		return
	}
	return input, nil
}

func (r *repository) GetAllData() (dosens []Dosen, err error) {
	rows, err := r.db.Queryx(dosenQuery.Select)
	if err != nil {
		return
	}

	for rows.Next() {
		var dosen Dosen

		err = rows.StructScan(&dosen)
		if err != nil {
			return
		}

		dosens = append(dosens, dosen)
	}

	return
}

func (r *repository) GetDataById(id uuid.UUID) (dosen Dosen, err error) {
	rows, err := r.db.Queryx(dosenQuery.SelectById, id)

	if err == sql.ErrNoRows {
		return dosen, errors.New("data tidak ditemukan")
	}

	if err != nil {
		return
	}

	if rows.Next() {
		err = rows.StructScan(&dosen)
		if err != nil {
			return
		}
	} else {
		return dosen, errors.New("data tidak ditemukan")
	}

	return
}

func (r *repository) Update(updateInput Dosen) (dosen Dosen, err error) {
	stmt, err := r.db.PrepareNamed(dosenQuery.Update)
	if err != nil {
		return
	}

	_, err = stmt.Exec(updateInput)
	if err != nil {
		return
	}

	return updateInput, nil
}

func (r *repository) Delete(id uuid.UUID) error {
	_, err := r.db.Exec(dosenQuery.Delete, id)

	if err != nil {
		return err
	}

	return nil
}
