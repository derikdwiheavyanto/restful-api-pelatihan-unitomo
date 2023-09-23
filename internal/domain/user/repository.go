package user

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

var (
	userQuery = struct {
		Insert             string
		Select             string
		SelectUserFakultas string
	}{
		Insert: `INSERT INTO users (id,name,occupation,avatar_file_name, role,email,password_hash,created_at) 
		VALUES (:id,:name,:occupation,:avatar_file_name,:role,:email,:password_hash,:created_at)`,

		Select:             `SELECT id, name, occupation, email, password_hash, avatar_file_name, role, created_at,updated_at FROM users`,
		SelectUserFakultas: `SELECT u.id,u.name,u.email,f.name as fakultas FROM users u INNER JOIN fakultas f on u.id_fakultas = f.id;`,
	}
)

type Repository interface {
	Save(User) (User, error)
	GetAllData() ([]UserDto, error)
	GetUsersFakultas() ([]UserFakultas, error)
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
	if err != nil {
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
