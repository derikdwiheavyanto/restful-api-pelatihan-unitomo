package fakultas

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

var (
	fakultasQuery = struct {
		Insert     string
		Select     string
		SelectById string
		Update     string
		Delete     string
	}{
		Insert: `INSERT INTO fakultas (id,name,jmlh_mahasiswa,jmlh_jurusan,jmlh_dosen)
		VALUES(:id,:name,:jmlh_mahasiswa,:jmlh_jurusan,:jmlh_dosen)`,
		Select:     `SELECT id,name,jmlh_mahasiswa,jmlh_jurusan,jmlh_dosen,created_at,updated_at FROM fakultas`,
		SelectById: `SELECT * FROM fakultas where id = $1`,
		Update:     `UPDATE fakultas SET name = :name,jmlh_mahasiswa = :jmlh_mahasiswa, jmlh_jurusan = :jmlh_jurusan, jmlh_dosen = :jmlh_dosen, updated_at = :updated_at WHERE id = :id`,
		Delete:     `DELETE FROM fakultas where id = :id `,
	}
)

type Repository interface {
	Save(Fakultas) (Fakultas, error)
	GetAllData() ([]Fakultas, error)
	GetDataById(uuid.UUID) (Fakultas, error)
	Update(Fakultas) (Fakultas, error)
	Delete(id uuid.UUID) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(fakultas Fakultas) (Fakultas, error) {
	stmt, err := r.db.PrepareNamed(fakultasQuery.Insert)

	if err != nil {
		return fakultas, err
	}

	_, err = stmt.Exec(fakultas)

	if err != nil {
		return fakultas, err
	}

	return fakultas, nil
}

func (r *repository) GetAllData() (data []Fakultas, err error) {
	rows, err := r.db.Queryx(fakultasQuery.Select)

	if err == sql.ErrNoRows {
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var fakultas Fakultas

		err = rows.StructScan(&fakultas)

		if err != nil {
			return
		}

		data = append(data, fakultas)
	}

	return
}

func (r *repository) GetDataById(id uuid.UUID) (data Fakultas, err error) {

	rows, err := r.db.Queryx(fakultasQuery.SelectById, id)

	if err == sql.ErrNoRows {
		fmt.Println(err)
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	if rows.Next() {
		err = rows.StructScan(&data)
		if err != nil {
			return
		}
	} else {

		return Fakultas{}, errors.New("id tidak ditemukan")
	}

	return

}

func (r *repository) Update(fakultas Fakultas) (data Fakultas, err error) {
	stmt, err := r.db.PrepareNamed(fakultasQuery.Update)

	if err != nil {
		fmt.Println()
		return
	}

	_, err = stmt.Exec(fakultas)

	if err != nil {
		fmt.Println(err)
		return
	}

	return fakultas, nil
}

func (r *repository) Delete(id uuid.UUID) error {
	stmt, err := r.db.PrepareNamed(fakultasQuery.Delete)
	if err != nil {
		fmt.Println(err)
		return err
	}

	idFakultas := DeleteFakultas{ID: id}

	_, err = stmt.Exec(idFakultas)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
