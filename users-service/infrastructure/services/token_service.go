package services

import (
	"errors"
	"fmt"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/services"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTService struct {
	config Config
}

type Config struct {
	SigningKey string
	Issuer     string
}

func NewJWTService(c Config) *JWTService {
	return &JWTService{
		config: c,
	}
}

func (j *JWTService) Generate(input *services.GenerateInput) (services.Token, error) {
	if input == nil {
		return "", errors.New("input is required")
	}
	if input.Payload == nil {
		return "", errors.New("invalid payload")
	}

	claims := jwt.MapClaims{
		"exp":     jwt.NewNumericDate(input.ExpiresIn),
		"iat":     time.Now(),
		"iss":     j.config.Issuer,
		"payload": input.Payload,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(j.config.SigningKey))

	if err != nil {
		return "", err
	}

	return services.Token(tokenString), nil
}

func (j *JWTService) Parse(token services.Token) (services.Payload, error) {
	jwtoken, err := jwt.Parse(string(token), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.config.SigningKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := jwtoken.Claims.(jwt.MapClaims); ok && jwtoken.Valid {
		if payload, ok2 := claims["payload"].(map[string]any); ok2 {
			return payload, nil
		}

		return services.Payload{}, nil
	}

	return nil, errors.New("invalid token")
}
