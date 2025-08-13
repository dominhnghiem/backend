package bootstrap

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewDB: mở kết nối GORM tới Postgres
func NewDB(cfg Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("cannot open db: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("cannot get sql db: %v", err)
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("db ping failed: %v", err)
	}
	return db
}
