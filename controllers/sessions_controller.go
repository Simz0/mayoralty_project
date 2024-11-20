package controllers

import (
	"NSK_mayoralty_app/services"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

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
