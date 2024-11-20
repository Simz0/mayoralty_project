package services

import (
	"NSK_mayoralty_app/database"
	"NSK_mayoralty_app/models"
)

func GetAllPoints() ([]models.Work_points, error) {
	var points []models.Work_points

	result := database.DB.Find(&points)

	return points, result.Error
}

func GetSessionsPoints(sessionId int) ([]models.Work_points, error) {
	var points []models.Work_points

	result := database.DB.Where("session_id = ?", sessionId).Find(&points)

	return points, result.Error
}

func AddPointToSession(point models.Work_points) error {
	result := database.DB.Create(&point)

	return result.Error
}

func RemovePointFromSession(pointId int) error {
	result := database.DB.Where("id = ?", pointId).Delete(&models.Work_points{})

	return result.Error
}
