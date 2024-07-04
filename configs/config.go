package config

type Config struct {
	Env string `env:"ENV" envDefault:"dev"`
	Server Server
	DB DB
}

type Server struct {
	Host string `env:"HOST" envDefault:"127.0.0.1"`
	Port string `env:"PORT" envDefault:"8080"`
}

type DB struct {
	User string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Database string `env:"DB_DATABASE"`
	Host string `env:"DB_HOST" envDefault:"127.0.0.1"`
	Port string `env:"DB_PORT" envDefault:"5432"`
}
