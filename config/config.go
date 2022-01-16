package config

import (
	"os"
)

var (
	ServerPort      = os.Getenv("PORT")
	ENV             = os.Getenv("ENV")
	MongodbUri      = os.Getenv("MONGODB_URI")
	MongodbDatabase = os.Getenv("MONGODB_DATABASE")
)

type Config struct {
	MongoConfig MongoConfig
}

type MongoConfig struct {
	URI            string
	Database       string
	UserCollection string
}

func GetConfig() Config {
	config := Config{}
	mongoConfig := MongoConfig{
		URI:      MongodbUri,
		Database: MongodbDatabase,
	}

	config.MongoConfig = mongoConfig

	return config
}
