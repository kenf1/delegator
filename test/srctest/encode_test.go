package srctest

import (
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/kenf1/delegator/src/auth"
	"github.com/kenf1/delegator/src/models"
)

func TestEncodeJWT(t *testing.T) {
	user := models.UserInfo{
		Id:          "user1",
		Email:       "user@example.com",
		Roles:       []string{"admin"},
		Permissions: []string{"read"},
		Org_id:      96,
	}
	config := models.AuthConfig{
		Issuer:    "issuer-1",
		SecretKey: []byte("topsecret"),
	}

	tokenStr, err := auth.EncodeJWT(user, config)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return config.SecretKey, nil
	})
	if err != nil {
		t.Fatalf("token parse failed: %v", err)
	}

	if !token.Valid {
		t.Error("token is invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		t.Fatal("claims type assertion failed")
	}

	if claims["sub"] != user.Id {
		t.Errorf("expected sub %v, got %v", user.Id, claims["sub"])
	}
	if claims["iss"] != config.Issuer {
		t.Errorf("expected iss %v, got %v", config.Issuer, claims["iss"])
	}
	orgIDFromClaim, ok := claims["org_id"].(float64)
	if !ok {
		t.Fatalf("org_id type assertion failed")
	}
	if int(orgIDFromClaim) != user.Org_id {
		t.Errorf("expected org_id %v, got %v", user.Org_id, orgIDFromClaim)
	}
}
