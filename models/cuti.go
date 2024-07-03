package models

import (
	"fmt"
	"net/http"

	"go-crud/db"
)

type Cuti struct {
	Id_percutian              int64  `json:"id_percutian"`
	Id_pegawai             	 int64  `json:"id_pegawai"`
	Tanggal_mulai   			int64  `json:"tanggal_mulai"`
	Tanggal_selesai 			int64  `json:"tanggal_selesai"`
	Keterangan         			 string `json:"keterangan"`
}

func SemuaCuti() (Response, error) {
	var obj Cuti
	var arrobj []Cuti
	var res Response

	con := db.CreateCon()
	if con == nil {
		return res, fmt.Errorf("failed to create database connection")
	}

	sqlStatement := "SELECT * FROM percutian"

	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id_percutian, &obj.Id_pegawai, &obj.Tanggal_mulai, &obj.Tanggal_selesai, &obj.Keterangan)
		if err != nil {
			return res, err
		}
		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Berhasil"
	res.Data = arrobj

	return res, nil
}

func TambahCuti(id_percutian int64, id_pegawai int64, tanggalMulai int64, tanggalSelesai int64, keterangan string) (Response, error) {
	var res Response

	con := db.CreateCon()
	if con == nil {
		return res, fmt.Errorf("failed to create database connection")
	}

	sqlStatement := "INSERT INTO percutian (id_percutian, id_pegawai, tanggal_mulai, tanggal_selesai, keterangan) VALUES (?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id_percutian,id_pegawai, tanggalMulai, tanggalSelesai, keterangan)
	if err != nil {
		return res, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Data berhasil disimpan"
	res.Data = map[string]interface{}{
		"last_insert_id": lastInsertID,
	}

	return res, nil
}
