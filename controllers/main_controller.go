package controllers

import (
	"NSK_mayoralty_app/models"
	"NSK_mayoralty_app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary GetPoints
// @Tags Points
// @Description Это рут для получения всех точек перемещения работников из базы
// @Success 200 {object} models.Work_points
// @Failure 500 {object} models.FailResponse
// @Router /points [get]
func GetPoints(c *gin.Context) {
	points, err := services.GetAllPoints()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cant handle data from database"})
		return
	}

	c.JSON(http.StatusOK, points)
}

// @Summary GetSessionPoints
// @Tags Points
// @Description Это рут для получения всех точек перемещения работника за конкретную сессию
// @Success 200 {object} models.Work_points
// @Failure 500 {object} models.FailResponse
// @Router /points/:sessionId [get]
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

// @Summary AddPointToSession
// @Tags Points
// @Description Это рут для загрузки точки для сессии
// @Accept json
// @Param input body models.Work_points true "Данные точки работы"
// @Success 200 {object} models.SuccessResponse
// @Failure 500 {object} models.FailResponse
// @Router /points [post]
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

// @Summary RemovePoint
// @Tags Points
// @Description Это рут для удаление тестовых или ошибочных точек
// @Param pointId path int true "ID точки для удаления"
// @Success 200 {object} models.SuccessResponse
// @Failure 500 {object} models.FailResponse
// @Router /points/:pointId [delete]
func RemovePointFromSession(c *gin.Context) {
	pointId, err := strconv.Atoi("pointId")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Find not point id"})
	}

	if err := services.RemovePointFromSession(pointId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
	}
}
