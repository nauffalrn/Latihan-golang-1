package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"latihan/config"
	"latihan/route"
)

func main() {
	fmt.Println("Memulai server Management Dashboard...")

	// 1. Jalankan koneksi database (dan baca .env)
	config.ConnectDB()

	// 2. Ambil port dari .env (default ke 8080 jika kosong)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 3. Setup Router (dari package route)
	router := route.SetupRouter()

	// 4. Jalankan HTTP Server
	serverAddr := ":" + port
	fmt.Printf("Server berjalan di http://localhost%s\n", serverAddr)
	
	// http.ListenAndServe akan terus menyala (blocking) untuk menerima request
	err := http.ListenAndServe(serverAddr, router)
	if err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}