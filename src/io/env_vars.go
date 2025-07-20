package io

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kenf1/delegator/src/models"
)

func loadEnvFile(env_file string) error {
	err := godotenv.Load(env_file)
	if err != nil {
		return err
	}

	return nil
}

func importServerAddr() (models.ServerAddr, error) {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	if host == "" && port == "" {
		return models.ServerAddr{}, fmt.Errorf("invalid host or port value")
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
