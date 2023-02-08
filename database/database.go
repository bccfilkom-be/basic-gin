package database

import (
	"basic-gin/entity"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// jalankan fungsi ini sekali saja
// Untuk koneksi ke database sesuai kredensianl yang sudah ditulis di .env
// returns *gorm.DB
func InitDB() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("init db failed,", err)
	}
	return db
}

// NOTE: jangan lupa setiap buat entitas, masukkan entitas tsb
// ke dalam fungsi ini biar auto migrasi
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.User{},
		&entity.Post{},
		&entity.Comment{},
	) //masukkan object yg mau dimigrasi ke dlm parameter ini (variadic)
}
