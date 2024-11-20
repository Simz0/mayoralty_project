package controllers

import (
	"NSK_mayoralty_app/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Registration(c *gin.Context) {
	var req struct {
		Username string
		Password string
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Check your registration data"})
		return
	}

	if err := services.Registration(req.Username, req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User register  successed"})
}

func Login(c *gin.Context) {
	var req struct {
		Username string
		Password string
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Uncorrect request, undefined your username or password"})
		return
	}
	splitString := (strings.SplitAfter(req.Password, " "))
	token, loginError := services.Login(req.Username, splitString[0])
	if loginError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Your username or password is uncorrect"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "success", "data": token})
	}
}
