package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "event_booking_secret_key"

func GenerateToken(email string, userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   email,
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}
func ValidateToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, err
	}
	if !parsedToken.Valid {
		return 0, jwt.ErrTokenInvalidId
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, jwt.ErrTokenInvalidClaims
	}

	// Check expiration
	exp, ok := claims["exp"].(float64)
	if !ok || time.Unix(int64(exp), 0).Before(time.Now()) {
		return 0, jwt.ErrTokenExpired
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, jwt.ErrTokenInvalidClaims
	}

	return int64(userID), nil
}
