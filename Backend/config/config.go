package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host          string
	Port          int
	Name          string //Database name
	User          string
	Password      string
	EnableSSLMode bool
}

var configuration *Config //store env though process ended

type Config struct {
	Version      string //application version
	ServiceName  string //service name
	HttpPort     int    //port
	JwtSecretKey string //jwt token
	DB           *DbConfig
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load the env", err)
		os.Exit(1)
	}
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		log.Fatal("serviceName is required")
		os.Exit(1)

	}

	httpPortStr := os.Getenv("HTTP_PORT")
	if httpPortStr == "" {
		log.Fatal("httpPort is required")
		os.Exit(1)
	}

	//convert string to int
	httpPort, err := strconv.Atoi(httpPortStr)
	if err != nil {
		fmt.Println("conversion error:", err)
		os.Exit(1)

	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		log.Fatal("JwtSecretKey is required")
		os.Exit(1)
	}

	//DB setup

	host := os.Getenv("HOST")
	if host == "" {
		log.Fatal("host is required")
		os.Exit(1)
	}

	dbPortStr := os.Getenv("PORT")
	if dbPortStr == "" {
		log.Fatal("dbPort is required")
		os.Exit(1)
	}
	//convert string to int
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		fmt.Println("conversion error:", err)
		os.Exit(1)

	}

	name := os.Getenv("NAME")
	if name == "" {
		log.Fatal("dbName is required")
		os.Exit(1)
	}

	user := os.Getenv("USER")
	if user == "" {
		log.Fatal("user is required")
		os.Exit(1)
	}

	password := os.Getenv("PASSWORD")
	if password == "" {
		log.Fatal("password is required")
		os.Exit(1)
	}

	enableSSLModeStr := os.Getenv("ENABLE_SSL_MODE")
	enableSSLMode, err := strconv.ParseBool(enableSSLModeStr)
	if err != nil {
		fmt.Println("conversion error:", err)
		os.Exit(1)

	}

	dbConfig := &DbConfig{
		Host:          host,
		Port:          dbPort,
		User:          user,
		Name:          name,
		Password:      password,
		EnableSSLMode: enableSSLMode,
	}
	configuration = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     httpPort,
		JwtSecretKey: jwtSecretKey,
		DB:           dbConfig,
	}

}

func GetConfig() *Config {
	if configuration == nil {
		//Load Config File if it's load First time
		loadConfig()

	}
	return configuration
}
