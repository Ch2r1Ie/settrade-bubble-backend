package config

const (
	environment = "APP_ENV"
	timezone    = "APP_TIME_LOCATION"
	app_prefix  = "APP_NAME"
)

const (
	localhost_path = "./config/localhost.configmap.env"
	secret_path    = "./config/secret.configmap.env"
)

const (
	interconnected = "_"
)

const (
	Local string = "LOCAL"
	Dev   string = "DEV"
	UAT   string = "UAT"
	Prod  string = "PROD"
)

const (
	hostname = "HOST_NAME"
	port     = "PORT"
)

const (
	GOMAXPROCS = "APP_GOMAXPROCS"
	GOMEMLIMIT = "APP_GOMEMLIMIT"
)
