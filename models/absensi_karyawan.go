package models

import (
	"net/http"

	"go-crud/db"
)



type AbsensiK struct {
	Id_kehadiran       		int64  `json:"id_kehadiran"`
	Id_pegawai    		 	int64 `json:"id_pegawai"`
	CheckIn  				string `json:"check_in"`
	CheckOut 				string `json:"check_out"`
	Status_kehadiran	 	  string `json:"status_kehadiran"`
}

func GetAbsensik() (Response, error) {
	var obj Absensi
	var arrobj []Absensi
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM kehadiran"

	rows, err := con.Query(sqlStatement)

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

	res.Status = http.StatusOK
	res.Message = "Berhasil"
	res.Data = arrobj

	return res, nil
}
