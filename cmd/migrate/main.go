package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"web/internal/app/ds"
	"web/internal/app/dsn"
)

func main() {
	_ = godotenv.Load()
	env, err := dsn.FromEnv()
	if err != nil {
		panic("cant migrate db")
	}
	db, err := gorm.Open(postgres.Open(env), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&ds.Forecasts{}, &ds.Predictions{}, &ds.Preds_Forecs{})
	if err != nil {
		panic("cant migrate db")
	}
}
