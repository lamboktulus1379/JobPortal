package persistence

import (
	"fmt"
	"jobportal/domain/model"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func NewRepositories() (*gorm.DB, error) {
	cfg := &Config{}

	cfg.User = os.Getenv("MYSQL_USER")
	cfg.Host = os.Getenv("MYSQL_HOST")
	cfg.Port = os.Getenv("MYSQL_PORT")
	cfg.Password = os.Getenv("MYSQL_PASSWORD")
	cfg.Database = os.Getenv("MYSQL_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color
			},
		),
	})
	if err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}
	log.Printf("INFO: Connected to DB")
	db.AutoMigrate(&model.User{})
	return db, nil
}
