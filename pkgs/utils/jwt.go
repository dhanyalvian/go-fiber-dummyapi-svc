package utils

import (
	"time"

	"go-fiber-dummyapi-svc/apps/configs"

	"github.com/golang-jwt/jwt/v5"
)

type TokenData struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

func GenerateToken(cfg *configs.Config, tokenData TokenData, expire int) (string, error) {
	claims := jwt.MapClaims{
		"id":        tokenData.ID,
		"firstname": tokenData.Firstname,
		"lastname":  tokenData.Lastname,
		"email":     tokenData.Email,
		"avatar":    tokenData.Avatar,
		"exp":       time.Now().Add(time.Minute * time.Duration(expire)).Unix(),
		"iat":       time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.Auth.JwtSecret))
}

func DecodeToken(cfg *configs.Config, tokenString string) (*TokenData, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.Auth.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		tokenData := &TokenData{
			ID:        claims["id"].(string),
			Firstname: claims["firstname"].(string),
			Lastname:  claims["lastname"].(string),
			Email:     claims["email"].(string),
			Avatar:    claims["avatar"].(string),
		}
		return tokenData, nil
	}

	return nil, err
}
