package config

import (
	"os"
	"strconv"
)

var (
	ServerPort                = os.Getenv("PORT")
	MongodbUri                = os.Getenv("MONGODB_URI")
	MongodbDatabase           = os.Getenv("MONGODB_DATABASE")
	JWTSecret                 = os.Getenv("JWT_SECRET")
	JWTIssuer                 = os.Getenv("JWT_ISSUER")
	JWTUserTokenExpiryInHours = os.Getenv("JWT_USER_TOKEN_EXPIRY_IN_HOURS")
)

type Config struct {
	Mongo Mongo
	JWT   JWT
}

type Mongo struct {
	URI      string
	Database string
}

type JWT struct {
	Secret                    string
	Issuer                    string
	JWTUserTokenExpiryInHours int
}

func GetConfig() Config {
	config := Config{}
	mongoConfig := Mongo{
		URI:      MongodbUri,
		Database: MongodbDatabase,
	}

	_expiry, err := strconv.Atoi(JWTUserTokenExpiryInHours)
	if err != nil {
		panic(err)
	}

	jwtConfig := JWT{
		Secret:                    JWTSecret,
		Issuer:                    JWTIssuer,
		JWTUserTokenExpiryInHours: _expiry,
	}

	config.Mongo = mongoConfig
	config.JWT = jwtConfig

	return config
}
