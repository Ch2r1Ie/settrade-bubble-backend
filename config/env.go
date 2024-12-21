package config

import (
	"strings"
)

var Env string

func enviroment(environment string) {
	Env = environment
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
