package services

import (
	"NSK_mayoralty_app/database"
	"NSK_mayoralty_app/models"
	"time"
)

func MakeSession(userId uint) (uint, error) {
	var session models.Sessions

	result := database.DB.Create(&session)
	session.UserID = userId
	database.DB.Save(&session)
	return session.ID, result.Error
}

func StopSession(userID uint) error {
	var session models.Sessions

	result := database.DB.Where("user_id = ? AND stop_at = ? ", userID, time.Time{}).First(&session)

	session.Stop_at = time.Now()

	database.DB.Save(&session)

	return result.Error
}
