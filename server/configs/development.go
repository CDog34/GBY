package config

var configDev = ConfigRule{
	"listenAddr": ":8080",

	"dbUrl":  "localhost",
	"dbName": "GBY",

	"sessionProvider":                "memory",
	"accessControlAllowOriginHeader": "http://localhost:8000",
}
