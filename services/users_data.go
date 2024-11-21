package services

import (
	"NSK_mayoralty_app/database"
	"NSK_mayoralty_app/models"
)

func GetUserById(ID string) (models.User, error) {
	var user models.User

	result := database.DB.Where("id = ?", ID).First(&user)

	if result.Error == nil {
		return user, nil
	} else {
		return user, result.Error
	}
}
