package configs

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env 		   string `yaml:"env"`
	MigrationsPath string `yaml:"migrations_path"`
	Server 		   Server `yaml:"server"`
	DB 			   DB	  `yaml:"db"`
	OuterAPI 	   string `yaml:"outer_api"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type DB struct {
	User 	   string `yaml:"user"`
	Password   string `yaml:"password"`
	Name 	   string `yaml:"name"`
	Host 	   string `yaml:"host"`
	Port 	   string `yaml:"port"`
}

func LoadConfig(cfgPath string) (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(cfgPath, &cfg)
	if err != nil {
		return nil, fmt.Errorf("error occurred while loading config: %v", err)
	}

	return &cfg, nil
}
