package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// DB adalah variabel global yang menyimpan connection pool ke database.
// Variabel ini diawali huruf besar agar bisa diakses dari package lain (repository).
var DB *sql.DB

// ConnectDB bertugas menginisialisasi koneksi ke MySQL
func ConnectDB() {
	// 1. Muat file .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Peringatan: File .env tidak ditemukan, menggunakan environment system")
	}

	// 2. Ambil nilai dari file .env
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// 3. Rangkai Data Source Name (DSN) format MySQL
	// parseTime=true sangat penting agar tipe DATE/TIMESTAMP di MySQL terbaca sebagai time.Time di Go
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)

	// 4. Buka koneksi ke database
	database, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Gagal membuka koneksi database: %v", err)
	}

	// 5. Ping database untuk memastikan koneksi benar-benar berhasil (XAMPP menyala)
	err = database.Ping()
	if err != nil {
		log.Fatalf("Gagal melakukan ping ke database (Pastikan XAMPP MySQL menyala!): %v", err)
	}

	fmt.Println("✅ Berhasil terhubung ke database MySQL!")

	// 6. Simpan koneksi ke variabel global DB
	DB = database
}