package utils

import (
	"strconv"
	"time"

	"github.com/ddld93/tsrims/auth/models"
	"github.com/golang-jwt/jwt"
)

var (
	jwtSecrete string
	jwtExpiration time.Duration
)
func init() {
	cfg := LoadEvn()
	jwtSecrete = cfg.AppPort
	jwtExpHours, _ := strconv.Atoi(cfg.JwtEpiration)
	jwtExpiration = time.Duration(jwtExpHours) * time.Hour
}

func GenerateJWTToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"email":    user.Email,
		"userType": user.UserType,
		"exp":      time.Now().Add(time.Hour * jwtExpiration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtSecrete))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
