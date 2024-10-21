package db

import (
	"fmt"
	"log"
	"os"

	model "algohub.dev/backend/model"
	"algohub.dev/backend/structs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// EraseDB erases all the data from the database
func EraseDB(env *structs.Env) {
	// Get the DB GORM instance
	db, err := gorm.Open(postgres.Open(env.DB_URL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	for _, model := range model.GetModels() {
		// We have to use the Where clause because it refuses to delete without it
		if err := db.Where("id is not null").Delete(model).Error; err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting %s: %v\n", model, err)
			os.Exit(1)
		}
	}

	log.Println("[i] Database values erased")
}
