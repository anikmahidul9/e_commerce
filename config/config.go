package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	SecretKey   string
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("SERVICE_NAME is not set in .env file")
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("VERSION is not set in .env file")
		os.Exit(1)
	}

	httpPortStr := os.Getenv("HTTP_PORT")
	if httpPortStr == "" {
		fmt.Println("HTTP_PORT is not set in .env file")
		os.Exit(1)
	}
	httpPort, err := strconv.Atoi(httpPortStr)
	if err != nil {
		fmt.Println("Error converting HTTP_PORT to integer:", err)
		os.Exit(1)
	}

	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		fmt.Println("SECRET_KEY is not set in .env file")
		os.Exit(1)
	}

	configurations = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    httpPort,
		SecretKey:   secret,
	}

}

func GetConfig() *Config {
	loadConfig()

	return &configurations
}
