package security

import (
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

type TokenClaims struct {
	jwt.RegisteredClaims
	UserID uuid.UUID `json:"user_id"`
}

func VerifyToken(tokenString string) (uuid.UUID, error) {
	secretKey := getSecretKey()

	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid token: %w", err)
	}

	if !token.Valid {
		return uuid.Nil, fmt.Errorf("token is not valid")
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return uuid.Nil, fmt.Errorf("invalid token claims")
	}

	return claims.UserID, nil
}

func getSecretKey() []byte {
	key := os.Getenv("JWT_SECRET_KEY")
	if key == "" {
		key = "default-secret-key-only-for-development"
	}
	return []byte(key)
}
