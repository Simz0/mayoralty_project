package controllers

import (
	"NSK_mayoralty_app/models"
	"NSK_mayoralty_app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPoints(c *gin.Context) {
	points, err := services.GetAllPoints()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cant handle data from database"})
		return
	}

	c.JSON(http.StatusOK, points)
}

func GetSessionsPoints(c *gin.Context) {
	sessionId, err := strconv.Atoi(c.Param("sessionId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if sessionData, err := services.GetSessionsPoints(sessionId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cant handle data from database"})
		return
	} else {
		c.JSON(http.StatusOK, sessionData)
	}
}

func AddPointToSession(c *gin.Context) {
	var newPoint models.Work_points
	if err := c.ShouldBindBodyWithJSON(&newPoint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad point data"})
	}

	if err := services.AddPointToSession(newPoint); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func RemovePointFromSession(c *gin.Context) {
	pointId, err := strconv.Atoi("pointId")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Find not point id"})
	}

	if err := services.RemovePointFromSession(pointId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
	}
}
