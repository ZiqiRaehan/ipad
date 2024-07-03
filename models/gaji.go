package models

import (
	"go-crud/db"
)

type GajiData struct {
	Id_gaji      int64  `json:"id_gaji"`
	Id_pegawai   int64  `json:"id_pegawai"`
	Jumlah_hadir int    `json:"jumlah_hadir"`
	Lembur       int    `json:"lembur"`
	GajiPokok    int    `json:"gaji_pokok"`
	UangTambahan int    `json:"uang_tambahan"`
	TotalGaji    int    `json:"total_gaji"`
	Tanggal      string `json:"tanggal"`
}

func GetAllGajiData() ([]GajiData, error) {
	var gajiData []GajiData

	con := db.CreateCon()

	rows, err := con.Query("SELECT id_gaji, id_pegawai, jumlah_hadir, lembur, gaji_pokok, uang_tambahan, total_gaji, tanggal FROM gaji")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var data GajiData
		err := rows.Scan(&data.Id_gaji,&data.Id_pegawai, &data.Jumlah_hadir,  &data.Lembur, &data.GajiPokok, &data.UangTambahan, &data.TotalGaji, &data.Tanggal)
		if err != nil {
			return nil, err
		}
		gajiData = append(gajiData, data)
	}
	return gajiData, nil
}
