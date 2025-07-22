package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kenf1/delegator/src/models"
)

func generateClaims(userInfo models.UserInfo, authConfig models.AuthConfig) (jwt.MapClaims, error) {
	if userInfo.Id == "" {
		return nil, errors.New("missing user ID")
	}

	if userInfo.Email == "" {
		return nil, errors.New("missing email")
	}

	if authConfig.Issuer == "" {
		return nil, errors.New("missing issuer")
	}

	if authConfig.SecretKey == nil {
		return nil, errors.New("missing secret key")
	}

	claims := jwt.MapClaims{
		"sub":         userInfo.Id,
		"email":       userInfo.Email,
		"roles":       userInfo.Roles,
		"permissions": userInfo.Permissions,
		"org_id":      userInfo.Org_id,
		"iss":         authConfig.Issuer,
		"iat":         time.Now().Unix(),
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	}

	return claims, nil
}

func EncodeJWT(userInfo models.UserInfo, authConfig models.AuthConfig) (string, error) {
	claims, err := generateClaims(userInfo, authConfig)
	if err != nil {
		return "", errors.New("error generating claim")
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	newTokenString, err := newToken.SignedString(authConfig.SecretKey)
	if err != nil {
		return "", errors.New("error generating token")
	}

	return newTokenString, nil
}
