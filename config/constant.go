package config

const (
	app_prefix  = "APP_NAME"
	apiKey      = "API_KEY"
	timezone    = "TIME_LOCATION"
	hostname    = "HOST_NAME"
	port        = "PORT"
	environment = "ENV"
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
