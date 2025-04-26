package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
	"strings"
)

type Root struct {
	App      App
	Postgres Postgres
}

func Load(filenames ...string) Root {
	_ = godotenv.Overload(filenames...)

	r := Root{
		App:      App{},
		Postgres: Postgres{},
	}

	mustLoad("APP", &r.App)
	mustLoad("POSTGRES", &r.Postgres)

	return r
}

func mustLoad(prefix string, spec interface{}) {
	err := envconfig.Process(prefix, spec)
	if err != nil {
		panic(err)
	}
}

func mayLoad(prefix string, spec interface{}) {
	_ = envconfig.Process(prefix, spec)
}

func Getenv(key string) (fallback string) {
	var value string
	if key == "" {
		return ""
	} else {
		value = os.Getenv(key)
		if len(value) == 0 {
			return ""
		}
		return value
	}
}

func GetEnvCors(key string) (fallback []string) {
	value := Getenv(key)
	return strings.Split(value, ",")
}
