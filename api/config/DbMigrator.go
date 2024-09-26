package config

import "book-store/api/models"

func MigrateDB() {

	DB.AutoMigrate(&models.Book{})

}
