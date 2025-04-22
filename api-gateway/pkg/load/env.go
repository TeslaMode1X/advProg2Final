package load

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

const GatewayEnvName = "api-gateway"

func LoadDotEnv() error {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}
	fmt.Println("Current working directory:", dir)

	filePath := fmt.Sprintf(".env.%s", GatewayEnvName)

	if _, err := os.Stat(filePath); err == nil {
		return godotenv.Load(filePath)
	}

	filePath = filepath.Join("../..", fmt.Sprintf(".env.%s", GatewayEnvName))
	return godotenv.Load(filePath)
}
