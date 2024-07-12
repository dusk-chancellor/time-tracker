package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env 		   string `env:"ENV" envDefault:"dev"`
	MigrationsPath string `env:"MIGRATIONS_PATH" envDefault:"./migrations"`
	Server 		   Server
	DB 			   DB
}

type Server struct {
	Host string `env:"SERVER_HOST" envDefault:"127.0.0.1"`
	Port string `env:"SERVER_PORT" envDefault:"8080"`
}

type DB struct {
	User 	   string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name 	   string `env:"DB_NAME"`
	Host 	   string `env:"DB_HOST" envDefault:"127.0.0.1"`
	Port 	   string `env:"DB_PORT" envDefault:"5432"`
}

func ReadConfig() *Config{
	err := godotenv.Load("./configs/.env")
	if err != nil {
		panic(err)
	}

	return &Config{
		Env: 			os.Getenv("ENV"),
		MigrationsPath: os.Getenv("MIGRATIONS_PATH"),
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},	
		DB: DB{
			User: 	  os.Getenv("DB_USER"),
			Password:   os.Getenv("DB_PASSWORD"),
			Name:   	  os.Getenv("DB_NAME"),
			Host: 	  os.Getenv("DB_HOST"),
			Port: 	  os.Getenv("DB_PORT"),
		},
	}
}
