package config

import (
	"os"
)


var (
	SERVER_PORT = os.Getenv("PORT")
	ENV         = os.Getenv("ENV")
)

