package config

import (
	"flag"
	"os"
)

var (
	// Port of API server
	Port string
    MongoURL string
    MongoDatabase string
)

func init() {
	flag.StringVar(&Port, "p", "4945", "Api port")

	flag.Parse()

	Port = getEnvValue("SWANAGER_PORT", Port)

	MongoURL = getEnvValue("MONGO_URL", "mongodb://127.0.0.1:27017/")
	MongoDatabase = getEnvValue("MONGO_DATABASE", "swanager")
}

func getEnvValue(varName string, currentValue string) string {
	if os.Getenv(varName) != "" {
		return os.Getenv(varName)
	}
	return currentValue
}
