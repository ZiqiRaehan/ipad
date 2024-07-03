package models

import (
	"net/http"

	"go-crud/db"
)

type Absensi struct {
	Id_kehadiran     int64  `json:"id_kehadiran"`
	Id_pegawai       int64  `json:"id_pegawai"`
	CheckIn          string `json:"check_in"`
	CheckOut         string `json:"check_out"`
	Status_kehadiran string `json:"status_kehadiran"`
}

func GetAbsensi(month int) (Response, error) {
	var obj Absensi
	var arrobj []Absensi
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM kehadiran WHERE EXTRACT(MONTH FROM check_in) = ?"

	rows, err := con.Query(sqlStatement, month)

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id_kehadiran, &obj.Id_pegawai, &obj.CheckIn, &obj.CheckOut, &obj.Status_kehadiran)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	// Jika tidak ada data yang tersedia, kembalikan sebuah array kosong
	if len(arrobj) == 0 {
		res.Status = http.StatusOK
		res.Message = "Data tidak ditemukan"
		res.Data = arrobj
		return res, nil
	}

	res.Status = http.StatusOK
	res.Message = "Berhasil"
	res.Data = arrobj

	return res, nil
}
