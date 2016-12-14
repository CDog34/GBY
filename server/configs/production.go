package config

import (
	"fmt"
	"os"
)

var configProduction = ConfigRule{
	"listenAddr": fmt.Sprintf(":%s", string(os.Getenv("GBY_LISTEN_PORT"))),

	"dbUrl":  os.Getenv("MONGODB_CONNECTION"),
	"dbName": os.Getenv("MONGODB_INSTANCE_NAME"),

	"sessionProvider":                "memory",
	"accessControlAllowOriginHeader": "*",
}
