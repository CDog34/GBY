package config

import (
	"fmt"
	"os"
)

var configStaging = ConfigRule{
	"listenAddr": fmt.Sprintf(":%s", string(os.Getenv("GBY_LISTEN_PORT"))),

	"dbUrl":  os.Getenv("MONGODB_CONNECTION"),
	"dbName": os.Getenv("MONGODB_INSTANCE_NAME"),

	"sessionProvider":                "memory",
	"accessControlAllowOriginHeader": "http://web.gby.isues.net",
}
