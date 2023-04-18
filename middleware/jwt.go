package middleware

import (
	"errors"
	"keuangan-pribadi/models"
	"keuangan-pribadi/utils"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userId uint, name string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"name": name,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(utils.GetConfig("JWT_SECRET_KEY")))
}

func VerifyToken(tokenString string) (models.User, error) {
    if strings.HasPrefix(tokenString, "Bearer ") {
        tokenString = strings.TrimPrefix(tokenString, "Bearer ")
    }
    var user models.User

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(utils.GetConfig("JWT_SECRET_KEY")), nil
    })

    if err != nil {
        return user, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        user.ID = uint(claims["user_id"].(float64))
    } else {
        return user, errors.New("Invalid token")
    }

    return user, nil
}