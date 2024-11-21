package controllers

import (
	"NSK_mayoralty_app/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Summary Registration
// @Tags Auth
// @Description Это рут для регистрации. Принимает на вход JSON из двух параметров: username и password
// @Accept json
// @Produce json
// @Success 200 {object} models.SuccessResponse уведомление об успешной регистрации "User register  successed"
// @Failure 400 {object} models.FailResponse уведомление об ошибочности формы регистрации
// @Failure 500 {object} models.FailResponse
// @Router /auth/registration [POST]
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

// @Summary Login
// @Tags Auth
// @Description Это рут для логина. Также как и при регистрации принимает на вход JSON из двух параметров
// @Accept json
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.FailResponse уведомление об ошибочности формы авторизации
// @Failure 500 {object} models.FailResponse
// @Router /auth/registration [post]
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
