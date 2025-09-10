package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configuration Config //store env though process ended
type Config struct {
	Version      string //application version
	ServiceName  string //service name
	HttpPort     int    //port
	JwtSecretKey string //jwt token
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

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		log.Fatal("httpPort is required")
		os.Exit(1)
	}

	//convert string to int
	port, err := strconv.Atoi(httpPort)
	if err != nil {
		fmt.Println("conversion error:", err)
		os.Exit(1)

	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		log.Fatal("JwtSecretKey is required")
		os.Exit(1)
	}

	configuration = Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     port,
		JwtSecretKey: jwtSecretKey,
	}

}

func GetConfig() Config {
	loadConfig()
	return configuration
}
