package config

type Config struct {
	Host        string `env:"SERVER_HOST" env-default:"localhost"`
	Port        string `env:"SERVER_PORT" env-default:"8080"`
	DatabaseURL string `env:"DATABASE_URL" env-default:"host=localhost user=sletkov password=postgres dbname=consultationAppDB sslmode=disable"`
	SessionKey  string `env:"SESSION_KEY" env-default:"consultation-app"`
}

func New(host, port, databaseURL, sessionKey string) *Config {
	return &Config{
		host,
		port,
		databaseURL,
		sessionKey,
	}
}
