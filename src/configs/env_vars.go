package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kenf1/delegator/src/models"
)

func loadEnvFile(env_file string) error {
	//clear prior env vars
	os.Unsetenv("HOST")
	os.Unsetenv("PORT")

	err := godotenv.Load(env_file)
	if err != nil {
		return err
	}

	return nil
}

func importServerAddr() (models.ServerAddr, error) {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	if host == "" {
		return models.ServerAddr{}, fmt.Errorf("host not found")
	}
	if port == "" {
		return models.ServerAddr{}, fmt.Errorf("port not found")
	}

	return models.ServerAddr{
		Host: host,
		Port: port,
	}, nil
}

func serverAddrLocal(env_file string) (models.ServerAddr, error) {
	err := loadEnvFile(env_file)
	if err != nil {
		return models.ServerAddr{}, err
	}

	res, err1 := importServerAddr()
	if err1 != nil {
		return models.ServerAddr{}, err1
	}

	return res, nil
}

func serverAddrRemote() (models.ServerAddr, error) {
	res, err1 := importServerAddr()
	if err1 != nil {
		return models.ServerAddr{}, err1
	}

	return res, nil
}

// attempt load remote -> attempt load local
func ImportServerAddrWrapper(env_file string) (models.ServerAddr, error) {
	runningVer, err := serverAddrRemote()
	if err != nil {
		localVer, err := serverAddrLocal(env_file)
		if err != nil {
			return models.ServerAddr{}, err
		}

		return localVer, nil
	}

	return runningVer, nil
}

func ImportAuthConfig() (models.AuthConfig, error) {
	secretKey := os.Getenv("SECRET_KEY")
	issuer := os.Getenv("ISSUER")

	if secretKey == "" {
		return models.AuthConfig{}, fmt.Errorf("secret key not found")
	}
	if issuer == "" {
		return models.AuthConfig{}, fmt.Errorf("issuer not found")
	}

	return models.AuthConfig{
		SecretKey: []byte(secretKey),
		Issuer:    issuer,
	}, nil
}
