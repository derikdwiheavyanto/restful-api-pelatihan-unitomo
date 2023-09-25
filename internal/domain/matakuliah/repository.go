package matakuliah

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

var (
	queryMatakuliah = struct {
		Create     string
		Select     string
		SelectById string
		Update     string
		Delete     string
	}{
		Create:     `INSERT INTO matakuliah(kode_mk,nama_mk,sks) VALUES (:kode_mk,:nama_mk,:sks)`,
		Select:     `SELECT * FROM matakuliah`,
		SelectById: `SELECT * FROM matakuliah WHERE kode_mk = $1`,
		Update:     `UPDATE matakuliah SET nama_mk=:nama_mk,sks=:sks WHERE kode_mk= :kode_mk`,
		Delete:     `DELETE FROM matakuliah where kode_mk= $1`,
	}
)

type Repository interface {
	Save(Matakuliah) (Matakuliah, error)
	GetAllData() ([]Matakuliah, error)
	GetDataById(kodeMK string) (Matakuliah, error)
	Update(Matakuliah) (Matakuliah, error)
	Delete(kodeMK string) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(input Matakuliah) (dosen Matakuliah, err error) {
	stmt, err := r.db.PrepareNamed(queryMatakuliah.Create)
	if err != nil {
		return
	}
	_, err = stmt.Exec(input)
	if err != nil {
		return
	}
	return input, nil
}

func (r *repository) GetAllData() (dosens []Matakuliah, err error) {
	rows, err := r.db.Queryx(queryMatakuliah.Select)
	if err != nil {
		return
	}

	for rows.Next() {
		var dosen Matakuliah

		err = rows.StructScan(&dosen)
		if err != nil {
			return
		}

		dosens = append(dosens, dosen)
	}

	return
}

func (r *repository) GetDataById(kodeMK string) (matakuliah Matakuliah, err error) {
	rows, err := r.db.Queryx(queryMatakuliah.SelectById, kodeMK)

	if err == sql.ErrNoRows {
		return matakuliah, errors.New("data tidak ditemukan")
	}

	if err != nil {
		return
	}

	if rows.Next() {
		err = rows.StructScan(&matakuliah)
		if err != nil {
			return
		}
	} else {
		return matakuliah, errors.New("data tidak ditemukan")
	}

	return
}

func (r *repository) Update(updateInput Matakuliah) (dosen Matakuliah, err error) {
	stmt, err := r.db.PrepareNamed(queryMatakuliah.Update)
	if err != nil {
		return
	}

	_, err = stmt.Exec(updateInput)
	if err != nil {
		return
	}

	return updateInput, nil
}

func (r *repository) Delete(kodeMK string) error {
	_, err := r.db.Exec(queryMatakuliah.Delete, kodeMK)

	if err != nil {
		return err
	}

	return nil
}
