package auth

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kenf1/delegator/src/models"
)

func parseJWT(tokenString string, secretKey []byte) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secretKey, nil
		},
	)

	if err != nil {
		return nil, errors.New("error parsing token")
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}

func extractUserInfo(token *jwt.Token) (*models.UserInfo, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token or unable to extract claims")
	}

	userInfo := &models.UserInfo{}

	if sub, ok := claims["sub"].(string); ok {
		userInfo.Id = sub
	}

	if email, ok := claims["email"].(string); ok {
		userInfo.Email = email
	}

	if roles, ok := claims["roles"].([]interface{}); ok {
		for _, r := range roles {
			if str, ok := r.(string); ok {
				userInfo.Roles = append(userInfo.Roles, str)
			}
		}
	}

	if permissions, ok := claims["permissions"].([]interface{}); ok {
		for _, p := range permissions {
			if str, ok := p.(string); ok {
				userInfo.Permissions = append(userInfo.Permissions, str)
			}
		}
	}

	if oid, ok := claims["org_id"].(float64); ok {
		userInfo.Org_id = int(oid)
	}

	return userInfo, nil
}

func DecodeJWT(tokenString string, authConfig models.AuthConfig) (*models.UserInfo, error) {
	token, err := parseJWT(tokenString, authConfig.SecretKey)
	if err != nil {
		return nil, err
	}

	userInfo, err := extractUserInfo(token)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}
