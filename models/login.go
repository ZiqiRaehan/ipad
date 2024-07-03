package models

import (
	"database/sql"
	"errors"

	"go-crud/db"
)

type User struct {
	ID       int64  `json:"id"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Alamat   string `json:"alamat"`
	Telp     string `json:"telp"`
	Jabatan  string `json:"jabatan"`
	IdPegawai int64  `json:"id_pegawai"`
}

func GetUserByUsername(username string) (*User, error) {
	con := db.CreateCon()

	sqlStatement := `SELECT id, nama, username, password, email, alamat, telp, jabatan FROM pegawai WHERE username = ?`
	row := con.QueryRow(sqlStatement, username)

	var user User
	err := row.Scan(&user.ID, &user.Nama, &user.Username, &user.Password, &user.Email, &user.Alamat, &user.Telp, &user.Jabatan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
