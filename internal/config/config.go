package config

type Config struct {
	ServerHost   string `env:"SERVER_HOST" env-default:"localhost"`
	ServerPort   string `env:"SERVER_PORT" env-default:"8080"`
	DatabaseURL  string `env:"DATABASE_URL" env-default:"host=localhost user=sletkov password=postgres dbname=consultationAppDB sslmode=disable"`
	SessionKey   string `env:"SESSION_KEY" env-default:"consultation-app"`
	SMTPHost     string `env:"SMTP_HOST" env-default:"smtp.gmail.com"`
	SMTPPort     int    `env:"SMTP_PORT" env-default:"587"`
	SMTPEmail    string `env:"SMTP_EMAIL"`
	SMTPPassword string `env:"SMTP_PASSWORD"`
}

func New(host, port, databaseURL, sessionKey, SMTPHost, SMTPEmail, SMTPPassword string, SMTPPort int) *Config {
	return &Config{
		host,
		port,
		databaseURL,
		sessionKey,
		SMTPHost,
		SMTPPort,
		SMTPEmail,
		SMTPPassword,
	}
}
