package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations *Config

type DbConfig struct {
	DbUsername string
	DbPassword string
	DbHost     string
	DbPort     int
	DbName     string
	EnableSSL  bool
}

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	SecretKey   string
	DB          *DbConfig
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

	dbUsername := os.Getenv("DB_USERNAME")
	if dbUsername == "" {
		fmt.Println("DB_USERNAME is not set in .env file")
		os.Exit(1)
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		fmt.Println("DB_PASSWORD is not set in .env file")
		os.Exit(1)
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		fmt.Println("DB_HOST is not set in .env file")
		os.Exit(1)
	}

	dbPortStr := os.Getenv("DB_PORT")
	if dbPortStr == "" {
		fmt.Println("DB_PORT is not set in .env file")
		os.Exit(1)
	}
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		fmt.Println("Error converting DB_PORT to integer:", err)
		os.Exit(1)
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("DB_NAME is not set in .env file")
		os.Exit(1)
	}
	enableSSLStr := os.Getenv("ENABLE_SSL_MODE")
	enableSSL, err := strconv.ParseBool(enableSSLStr)
	if err != nil {
		fmt.Println("Error converting ENABLE_SSL_MODE to boolean:", err)
		os.Exit(1)
	}

	dbConfig := &DbConfig{
		DbUsername: dbUsername,
		DbPassword: dbPassword,
		DbHost:     dbHost,
		DbPort:     dbPort,
		DbName:     dbName,
		EnableSSL:  enableSSL,
	}
	configurations = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    httpPort,
		SecretKey:   secret,
		DB:          dbConfig,
	}

}

func GetConfig() *Config {
	if configurations == nil {
		loadConfig()

	}

	return configurations
}
