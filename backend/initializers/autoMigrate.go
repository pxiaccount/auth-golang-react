package initializers

import "auth-golang-react/models"

func AutoMigrate() {
	db.AutoMigrate(&models.User{})
}
