package config

import (
	"os"
	"strings"
)

var Env string

func enviroment(prefix string) {
	Env = os.Getenv(connected(prefix, environment))
}

func IsLocalEnv() bool {
	return strings.ToUpper(Env) == Local
}

func IsDevEnv() bool {
	return strings.ToUpper(Env) == Dev
}

func IsUATEnv() bool {
	return strings.ToUpper(Env) == UAT
}

func IsProdEnv() bool {
	return strings.ToUpper(Env) == Prod
}
