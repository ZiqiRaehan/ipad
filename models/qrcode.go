package models

import (
	"go-crud/db"
)

func GetQRCodePath(username string) (string, error) {
	database := db.CreateCon()

	var qrPath string
	err := database.QueryRow("SELECT qrcode_path FROM pegawai WHERE username = ?", username).Scan(&qrPath)
	if err != nil {
		return "", err
	}
	return qrPath, nil
}
