package configs

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const defaultCfgPath = "./configs/local.env"

type Config struct {
	Env 			string `env:"ENV,required"`
	MigrationsPath 	string `env:"MIGRATIONS_PATH,required"`
	// server cfg
	ServerHost 		string `env:"SERVER_HOST,required"`
	ServerPort 		string `env:"SERVER_PORT,required"`
	// db cfg
	DBUrl 			string `env:"DB_URL,required"`
	// 
	OuterAPI 		string `env:"OUTER_API,required"`
}

func LoadConfig() (*Config, error) {
	// read cfg path env var
	cfgPath, ok := os.LookupEnv("CONFIG_PATH")
	if !ok {
		cfgPath = defaultCfgPath
	}

	// read .env
	var cfg Config
	err := cleanenv.ReadConfig(cfgPath, &cfg)
	if err != nil {
		return nil, fmt.Errorf("error occurred while loading config: %v", err)
	}

	return &cfg, nil
}
