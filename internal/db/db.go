package db

import (
	"muxtemp/internal/entity"
	"muxtemp/pkg/log"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	Instance, err = gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Logger.Error("Cannot connect to DB", zap.Error(err))
	}
	log.Logger.Info("Connected to Database...")
}

func Migrate() {
	Instance.AutoMigrate(&entity.Product{})
	log.Logger.Info("Database Migration Completed...")
}
