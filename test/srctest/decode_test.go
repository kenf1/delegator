package srctest

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kenf1/delegator/src/auth"
	"github.com/kenf1/delegator/src/models"
	"github.com/stretchr/testify/assert"
)

func generateTestToken(user models.UserInfo, config models.AuthConfig) (string, error) {
	claims := jwt.MapClaims{
		"sub":         user.Id,
		"email":       user.Email,
		"roles":       user.Roles,
		"permissions": user.Permissions,
		"org_id":      user.Org_id,
		"iss":         config.Issuer,
		"iat":         time.Now().Unix(),
		"exp":         time.Now().Add(time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.SecretKey)
}

func TestDecodeJWT(t *testing.T) {
	validUser := models.UserInfo{
		Id:          "user123",
		Email:       "test@example.com",
		Roles:       []string{"admin"},
		Permissions: []string{"read", "write"},
		Org_id:      240,
	}
	authConfig := models.AuthConfig{
		Issuer:    "my-app",
		SecretKey: []byte("my-secret-key"),
	}

	t.Run("valid JWT", func(t *testing.T) {
		tokenString, err := generateTestToken(validUser, authConfig)
		assert.NoError(t, err)

		userInfo, err := auth.DecodeJWT(tokenString, authConfig)
		assert.NoError(t, err)
		assert.NotNil(t, userInfo)
		assert.Equal(t, validUser.Id, userInfo.Id)
		assert.Equal(t, validUser.Email, userInfo.Email)
		assert.ElementsMatch(t, validUser.Roles, userInfo.Roles)
		assert.ElementsMatch(t, validUser.Permissions, userInfo.Permissions)
		assert.Equal(t, validUser.Org_id, userInfo.Org_id)
	})

	t.Run("invalid signature", func(t *testing.T) {
		tokenString, err := generateTestToken(validUser, authConfig)
		assert.NoError(t, err)

		//incorrect secret
		wrongConfig := models.AuthConfig{
			Issuer:    "my-app",
			SecretKey: []byte("wrong-secret"),
		}

		userInfo, err := auth.DecodeJWT(tokenString, wrongConfig)
		assert.Error(t, err)
		assert.Nil(t, userInfo)
	})

	t.Run("malformed token", func(t *testing.T) {
		malformedToken := "this.is.not.valid"

		userInfo, err := auth.DecodeJWT(malformedToken, authConfig)
		assert.Error(t, err)
		assert.Nil(t, userInfo)
	})
}
