package auth

import (
	"errors"
	"time"

	"github.com/adedaramola/golang-jwt-auth/utils"
	"github.com/golang-jwt/jwt/v4"
)

type Payload struct {
	Email string
}

type JWTClaim struct {
	jwt.StandardClaims
}

func GenerateToken(payload *Payload) (string, error) {
	claims := &JWTClaim{
		StandardClaims: jwt.StandardClaims{
			Subject:   payload.Email,
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(
		[]byte(utils.EnvString("JWT_SECRET")),
	)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(signedToken string) error {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(utils.EnvString("JWT_SECRET")), nil
	}

	jwtToken, err := jwt.ParseWithClaims(signedToken, &JWTClaim{}, keyFunc)
	if err != nil {
		return err
	}

	claims, ok := jwtToken.Claims.(*JWTClaim)
	if !ok {
		return errors.New("invalid token")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return errors.New("token expired")
	}

	return nil
}
