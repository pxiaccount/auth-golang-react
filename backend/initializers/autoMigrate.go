package initializers

import "auth-golang-react/models"

func AutoMigrate() {
	DB.AutoMigrate(&models.User{})
}
