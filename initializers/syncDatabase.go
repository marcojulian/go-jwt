package initializers

import "github.com/marcojulian/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
