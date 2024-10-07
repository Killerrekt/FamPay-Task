package utils

import (
	"log"

	"github.com/killerrekt/fampay-task/model"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	log.Println("Running Migrations")

	if err := db.AutoMigrate(&model.Video{}); err != nil {
		log.Fatal(err)
	}
}
