package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kenf1/delegator/src/models"
)

func EncodeJWT(userInfo models.UserInfo, authConfig models.AuthConfig) (string, error) {
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(authConfig.SecretKey)
	if err != nil {
		fmt.Printf("Error signing token: %s", err)
		return "", err
	}

	return tokenString, nil
}

func parseJWT(tokenString string, secretKey []byte) (*jwt.Token, error) {
	//return key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}

func DecodeJWT(tokenString string, secretKey []byte) (map[string]interface{}, error) {
	token, err := parseJWT(tokenString, secretKey)
	if err != nil {
		return nil, nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		result := make(map[string]interface{})
		for key, value := range claims {
			result[key] = value
		}
		return result, nil
	}

	return nil, errors.New("invalid token or failed to extract claims")
}
