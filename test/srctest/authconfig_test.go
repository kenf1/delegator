package srctest

import (
	"os"
	"testing"

	"github.com/kenf1/delegator/src/configs"
	"github.com/kenf1/delegator/src/models"
	"github.com/stretchr/testify/assert"
)

func TestImportAuthConfig_Success(t *testing.T) {
	_ = os.Setenv("SECRET_KEY", "supersecret")
	_ = os.Setenv("ISSUER", "auth-service")

	config, err := configs.ImportAuthConfig()
	assert.NoError(t, err)
	assert.Equal(t, []byte("supersecret"), config.SecretKey)
	assert.Equal(t, "auth-service", config.Issuer)
}

func TestImportAuthConfig_MissingSecretKey(t *testing.T) {
	_ = os.Unsetenv("SECRET_KEY")
	_ = os.Setenv("ISSUER", "auth-service")

	config, err := configs.ImportAuthConfig()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "secret key not found")
	assert.Equal(t, models.AuthConfig{}, config)
}

func TestImportAuthConfig_MissingIssuer(t *testing.T) {
	_ = os.Setenv("SECRET_KEY", "supersecret")
	_ = os.Unsetenv("ISSUER")

	config, err := configs.ImportAuthConfig()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "issuer not found")
	assert.Equal(t, models.AuthConfig{}, config)
}
