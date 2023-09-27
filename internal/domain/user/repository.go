package user

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

var (
	userQuery = struct {
		Insert             string
		Select             string
		SelectById         string
		SelectUserFakultas string
		Update             string
		Delete             string
	}{
		Insert:             `INSERT INTO users (id,name,occupation,avatar_file_name, role,email,password_hash,created_at) VALUES (:id,:name,:occupation,:avatar_file_name,:role,:email,:password_hash,:created_at)`,
		Select:             `SELECT id, name, occupation, email, password_hash, avatar_file_name, role, created_at,updated_at FROM users`,
		SelectById:         `SELECT * FROM users WHERE id = $1`,
		SelectUserFakultas: `SELECT u.id,u.name,u.email,f.name as fakultas FROM users u INNER JOIN fakultas f on u.id_fakultas = f.id;`,
		Update:             `UPDATE users SET name = :name, occupation = :occupation, email = :email, avatar_file_name = :avatar_file_name,updated_at = :updated_at  WHERE id = :id`,
		Delete:             `DELETE FROM users WHERE id = $1`,
	}
)

type Repository interface {
	Save(User) (User, error)
	GetAllData() ([]UserDto, error)
	GetDataById(id uuid.UUID) (UserDto, error)
	GetUsersFakultas() ([]UserFakultas, error)
	Update(UpdateInput) (UpdateInput, error)
	Delete(id uuid.UUID) error
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (newUser User, err error) {
	stmt, err := r.db.PrepareNamed(userQuery.Insert)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	_, err = stmt.Exec(user)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	return user, nil
}

func (r *repository) GetAllData() (data []UserDto, err error) {
	rows, err := r.db.Queryx(userQuery.Select)
	if err == sql.ErrNoRows {
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var user UserDto
		err = rows.StructScan(&user)

		if err != nil {
			return
		}

		data = append(data, user)
	}
	return
}

func (r *repository) GetDataById(id uuid.UUID) (user UserDto, err error) {
	rows, err := r.db.Queryx(userQuery.SelectById, id)

	if err == sql.ErrNoRows {
		return user, errors.New("data tidak ditemukan")
	}

	if err != nil {
		return
	}

	if rows.Next() {
		err = rows.StructScan(&user)
		if err != nil {
			return
		}
	} else {
		return user, errors.New("data tidak ditemukan")
	}

	return
}

func (r *repository) GetUsersFakultas() ([]UserFakultas, error) {
	rows, err := r.db.Queryx(userQuery.SelectUserFakultas)
	if err != nil {
		return nil, err
	}

	var dataUserFakultas []UserFakultas
	for rows.Next() {
		var userFakultas UserFakultas

		err = rows.StructScan(&userFakultas)
		if err != nil {
			return nil, err
		}

		dataUserFakultas = append(dataUserFakultas, userFakultas)
	}

	return dataUserFakultas, nil

}

func (r *repository) Update(input UpdateInput) (user UpdateInput, err error) {
	stmt, err := r.db.PrepareNamed(userQuery.Update)
	if err != nil {
		return
	}

	_, err = stmt.Exec(input)
	if err != nil {
		return
	}

	return input, nil
}

func (r *repository) Delete(id uuid.UUID) error {
	_, err := r.db.Exec(userQuery.Delete, id)

	if err != nil {
		return err
	}

	return nil
}
