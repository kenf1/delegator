package models

type AuthConfig struct {
	SecretKey []byte
	Issuer    string
}
