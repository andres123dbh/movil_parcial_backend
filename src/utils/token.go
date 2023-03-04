package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var JWTKEY = []byte("MY_SECRET_KEY")

func CreateAccessToken(userID string) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":    userID,
		"notBefore": time.Now(),
		"expire":    time.Now().Add(time.Minute * 30),
	})

	var err error
	accessTokenString, err := accessToken.SignedString([]byte(JWTKEY))
	if err != nil {
		return "", errors.New("Could not create access token")
	}

	return accessTokenString, nil
}

func ValidateAccessToken(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(JWTKEY), nil
	})
	if err != nil {
		return "", errors.New("Invalid access token (parse)")
	}

	// validate token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp, err := time.Parse(time.RFC3339, claims["expire"].(string))

		if err != nil {
			return "", errors.New("Invalid access token (expire format)")
		}

		if exp.Before(time.Now()) {
			return "", errors.New("Access token expired")
		}

		return claims["userid"].(string), nil
	} else {
		return "", errors.New("Invalid access token (claims)")
	}
}
