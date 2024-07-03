package main

import (
	"go-crud/db"
	"go-crud/routes"
)

func main(){
	db.Init()

	e := routes.Init()
	e.Logger.Fatal(e.Start(":9696"))
	
    
    // Tambahkan handler untuk melayani file statis
    e.Static("/qrcodes", "path/to/qrcodes")
}
