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

func ValidateToken(token string, expectedEmail string, expectedUserID int64) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}
	if !parsedToken.Valid {
		return nil, jwt.ErrTokenInvalidId
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, jwt.ErrTokenInvalidClaims
	}

	// Validate email
	email, ok := claims["email"].(string)
	if !ok || email != expectedEmail {
		return nil, jwt.ErrTokenInvalidClaims
	}

	// Validate user_id
	userID, ok := claims["user_id"].(float64) // jwt.MapClaims uses float64 for numbers
	if !ok || int64(userID) != expectedUserID {
		return nil, jwt.ErrTokenInvalidClaims
	}

	// Check if the token has expired
	exp, ok := claims["exp"].(float64)
	if !ok || time.Unix(int64(exp), 0).Before(time.Now()) {
		return nil, jwt.ErrTokenExpired
	}

	return parsedToken, nil
}
