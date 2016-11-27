package config

import "os"

type ConfigRule map[string]interface{}

var env string

var configBase = ConfigRule{
	"sessionLifeTime": 60 * 60 * 24 * 7,
	"sessionName":     "X-Stella",
}

func init() {
	if s := os.Getenv("GBYENV"); s == "" {
		env = "development"
	} else {
		env = s
	}

}

func Get(key string) interface{} {
	if value, ok := configBase[key]; ok {
		return value
	}
	switch env {
	case "development":
		return configDev[key]
	case "staging":
		return configStaging[key]
	case "production":
		return configProduction[key]
	default:
		return nil
	}

}
