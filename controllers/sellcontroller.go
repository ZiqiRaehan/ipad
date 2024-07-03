	package controllers

	import (
		"net/http"
		// "golang.org/x/crypto/bcrypt"
		"github.com/labstack/echo/v4"
		// "strconv"
		"go-crud/models"
	)

	type RegisterRequest struct {
		Username string `json:"username" form:"username"`
		Password string `json:"password" form:"password"`
		FullName string `json:"full_name" form:"full_name"`
	}

	func RegisterUser(c echo.Context) error {
		var req RegisterRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
		}

		newUser := models.UserRegister{
			Username: req.Username,
			Password: req.Password,
			FullName: req.FullName,
		}

		if err := models.CreateUser(newUser.Username, newUser.Password, newUser.FullName); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to register user"})
		}

		return c.JSON(http.StatusCreated, map[string]string{"message": "User registered successfully"})
	}

	type LoginRequest struct {
		Username string `json:"username" form:"username"`
		Password string `json:"password" form:"password"`
	}

	func Login(c echo.Context) error {
		var req LoginRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
		}

		user, err := models.GetUserByUsername(req.Username)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": err.Error()})
		}

		// Periksa apakah password sesuai (gunakan bcrypt untuk implementasi yang aman)
		//  if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		//     return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid credentials"})
		// }
		if user.Password != req.Password {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "salah password"})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Login successful",
			"user": map[string]interface{}{
				"username": user.Username,
				// "full_name": user.FullName,
				"jabatan": user.Jabatan, // Pastikan menambahkan jabatan di struktur user
			},
		})
	}

	func SemuaCuti(c echo.Context) error {
		result, err := models.SemuaCuti()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		return c.JSON(http.StatusOK, result)
	}

	func InsertCuti(c echo.Context) error {
		var cuti models.Cuti

		if err := c.Bind(&cuti); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
		}

		result, err := models.TambahCuti(cuti.Id_percutian,cuti.Id_pegawai, cuti.Tanggal_mulai, cuti.Tanggal_selesai, cuti.Keterangan)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		return c.JSON(http.StatusOK, result)
	}

	// func GetAbsensi(c echo.Context) error {
	// 	month := 1 // Misalnya, kita ingin mendapatkan data absensi untuk bulan Januari (ganti 1 dengan nomor bulan yang sesuai)
	// 	result, err := models.GetAbsensi(month) // Memanggil fungsi GetAbsensiByMonth dengan bulan sebagai argumen
	// 	if err != nil {
	// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	// 	}

	// 	return c.JSON(http.StatusOK, result)
	// }

	func GetAbsensi(c echo.Context) error {
		var query struct {
			Bulan int `query:"bulan"`
		}
		if err := c.Bind(&query); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid query parameters"})
		}
	
		result, err := models.GetAbsensi(query.Bulan)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, result)
	}
	
	
	func GetAbsensiK(c echo.Context) error {
		result, err := models.GetAbsensik()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
	
		return c.JSON(http.StatusOK, result)
	}
	

	func GetGajiData(c echo.Context) error {
		gajiData, err := models.GetAllGajiData()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, gajiData)
	}


	func GetQRCode(c echo.Context) error {
		username := c.Param("username")
		qrPath, err := models.GetQRCodePath(username)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get QR code"})
		}
		return c.JSON(http.StatusOK, map[string]string{"qrcode_path": qrPath})
	}

	func GetProfile(c echo.Context) error {
		username := c.Param("username")
		user, err := models.GetUserByUsername(username)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get user profile"})
		}
	
		return c.JSON(http.StatusOK, map[string]interface{}{
			"username": user.Username,
			"nama": user.Nama,
			"jabatan": user.Jabatan,
			// tambahkan properti lain sesuai kebutuhan
		})
	}
	