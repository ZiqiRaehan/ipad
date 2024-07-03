package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"go-crud/controllers"
)

func Init() *echo.Echo {
	e := echo.New()

	// Endpoint untuk halaman utama
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Selamat Datang")
	})

	// Contoh penambahan rute untuk berbagai fitur
	e.POST("/login", controllers.Login)
	e.GET("/cuti", controllers.SemuaCuti)
	e.POST("/cuti", controllers.InsertCuti)
	e.GET("/absensi", controllers.GetAbsensi)
	e.GET("/absensiK", controllers.GetAbsensiK)
	e.GET("/gaji", controllers.GetGajiData)
	e.POST("/register", controllers.RegisterUser)

	// Rute untuk mengambil QR code berdasarkan username
	e.GET("/qrcode/:username", controllers.GetQRCode)

	// Handler untuk melayani file statis
	e.Static("/qrcodes", "path/to/your/qrcodes")

	return e
}
