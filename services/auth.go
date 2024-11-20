package services

import (
	"NSK_mayoralty_app/database"
	"NSK_mayoralty_app/models"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("zMFxyAM7fj")

func getRole(username string) string {
	var user models.User
	query := database.DB.First(&user, username)
	if query.Error != nil {
		return "undefined"
	}
	result := user.Role
	return result
}

func CreateToken(username string, id uint) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"iss": "NSKMayoralty",
		"aud": getRole(username),
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Login(username string, password string) (string, error) {
	var user models.User
	query := database.DB.Where("username = ?", username).First(&user)
	if query.Error != nil {
		return "", query.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid username or password")
	}

	token, err := CreateToken(user.Username, user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func Registration(username, password string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Username: username,
		Password: string(passwordHash),
	}

	return database.DB.Create(&user).Error
}

func keyFunction(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Invalid signature method: %v", token.Header["alg"])
	}
	return secretKey, nil
}

func CheckUserToken(token string) (bool, uint) {
	checkingToken, err := jwt.Parse(token, keyFunction)
	if err != nil {
		return false, 0
	}

	if claims, ok := checkingToken.Claims.(jwt.MapClaims); ok && checkingToken.Valid {
		// Проверяем, есть ли поле "sub"
		if sub, ok := claims["sub"]; ok {
			switch v := sub.(type) {
			case float64: // Если "sub" закодировано как число
				return true, uint(v)
			case string: // Если "sub" закодировано как строка
				parsedID, err := strconv.ParseUint(v, 10, 64)
				if err != nil {
					return false, 0
				}
				return true, uint(parsedID)
			default:
				return false, 0
			}
		}
	}
	return false, 0
}
