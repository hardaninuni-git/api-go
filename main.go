package main

import (
	"rhiona-api/config"
	"rhiona-api/models"
	"rhiona-api/routes"
)

func main() {
	// Koneksi ke database
	config.ConnectDatabase()

	// Auto-migrate model Customer
	config.DB.AutoMigrate(&models.Customer{})

	// Setup router dan jalankan server
	r := routes.SetupRouter()
	r.Run(":8080")
}
