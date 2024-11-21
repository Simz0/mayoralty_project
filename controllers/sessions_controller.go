package controllers

import (
	"NSK_mayoralty_app/services"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Summary StartSession
// @Tags Sessions
// @Description Это рут для начала сессии. ВАЖНО: не начинать сессию при уже имеющейся открытой сессии!
// @Produce json
// @Success 200 {object} models.SuccessResponse
// @Failure 500 {object} models.FailResponse
// @Router /sessions/start [get]
func StartSession(c *gin.Context) {
	userToken := strings.Split(c.GetHeader("Authorization"), " ")
	validUser, userID := services.CheckUserToken(userToken[1])
	if validUser {
		fmt.Println("Valid")
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user token is invalid or undefined"})
		return
	}
	sessionId, err := services.MakeSession(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": sessionId, "status": "success"})
		return
	}
}

// @Summary StopSession
// @Tags Sessions
// @Description Это рут для закрытия сессии
// @Produce json
// @Success 200 {object} models.SuccessResponse
// @Failure 500 {object} models.FailResponse
// @Router /sessions/stop [get]
func StopSession(c *gin.Context) {
	userToken := strings.Split(c.GetHeader("Authorization"), " ")
	if validUser, UserID := services.CheckUserToken(userToken[1]); !validUser {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user token is invalid or undefined"})
		return
	} else {
		services.StopSession(UserID)
		c.JSON(http.StatusOK, gin.H{"status": "success"})
		return
	}
}
