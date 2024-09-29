package data

import (
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() error {
	conn := os.Getenv("DATABASE_URL")
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	gormdb, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	db = gormdb
	db.AutoMigrate(&Scrape{})
	return err
}

// gorm.Model definition
type Scrape struct {
	ID        uint
	Source    string
	TimeTaken time.Duration

	CreatedAt time.Time
	UpdatedAt time.Time
}
