package database

import (
	"context"
        "fmt"
	"log"
	"os"

	"github.com/hashicorp/vault/api"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (*gorm.DB, error) {
	// Пробуем получить из Vault
	vaultAddr := os.Getenv("VAULT_ADDR")
	vaultToken := os.Getenv("VAULT_TOKEN")
	
	var dbConfig map[string]interface{}
	
	if vaultAddr != "" && vaultToken != "" {
		config := api.DefaultConfig()
		config.Address = vaultAddr
		
		client, err := api.NewClient(config)
		if err == nil {
			client.SetToken(vaultToken)
			
			secret, err := client.KVv2("secret").Get(context.Background(), "app/database")
			if err == nil && secret != nil {
				dbConfig = secret.Data
				log.Println("Loaded database config from Vault")
			}
		}
	}
	
	// Fallback на env vars
	host := getConfigValue(dbConfig, "DB_HOST", "DB_HOST", "localhost")
	user := getConfigValue(dbConfig, "DB_USER", "DB_USER", "postgres")
	password := getConfigValue(dbConfig, "DB_PASSWORD", "DB_PASSWORD", "postgres")
	dbname := getConfigValue(dbConfig, "DB_NAME", "DB_NAME", "postgres")
	port := getConfigValue(dbConfig, "DB_PORT", "DB_PORT", "5432")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}

func getConfigValue(vaultData map[string]interface{}, vaultKey, envKey, defaultVal string) string {
	if vaultData != nil {
		if val, ok := vaultData[vaultKey]; ok {
			return fmt.Sprintf("%v", val)
		}
	}
	if val := os.Getenv(envKey); val != "" {
		return val
	}
	return defaultVal
}
