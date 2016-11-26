package config

import (
	"os"
	"fmt"
)

var configStaging = ConfigRule{
	"listenAddr":fmt.Sprintf(":%s", string(os.Getenv("GBY_LISTEN_PORT"))),

	"dbUrl":  "localhost",
	"dbName": "GBY",

	"sessionProvider": "memory",
}
