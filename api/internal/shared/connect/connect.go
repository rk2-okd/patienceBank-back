package connect

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// パスワードは出さない（安全なログ）
	log.Printf("Connecting to DB host=%s port=%s db=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var dbErr error
	for i := 0; i < 10; i++ {
		DB, dbErr = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if dbErr == nil {
			log.Println("Database connection established")
			break
		}
		log.Printf("Database connection failed (attempt %d): %v", i+1, dbErr)
		time.Sleep(2 * time.Second)
	}

	if dbErr != nil {
		log.Fatalf("DB接続エラー: %v", dbErr)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("DBインスタンス取得エラー: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(30 * time.Hour)

	return DB
}
