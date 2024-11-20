package services

import (
	"NSK_mayoralty_app/database"
	"NSK_mayoralty_app/models"
)

func GetUserId(username string) (uint, error) {
	var user models.User

	result := database.DB.Where("username = ?", username).First(&user)

	if result.Error == nil {
		return user.ID, nil
	} else {
		return 0, result.Error
	}
}
