package config

import (
	"fmt"
	"path"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Mongo  MongoDatabase
	Server Server
	JWT
	Hasher
}

type MongoDatabase struct {
	Uri  string `env:"CONN_URI" env-default:"mongodb://mongo:27017"`
	Name string `env:"MG_DB" env-default:"test_task"`
}

type Server struct {
	Host string `env:"SRV_HOST"`
	Port string `env:"SRV_PORT" env-default:"3003"`
}

type JWT struct {
	SignKey         string        `env-required:"true"                  env:"JWT_SIGN_KEY"`
	JWTokenTTL      time.Duration `env-required:"true" env:"JWT_TOKEN_TTL"`
	RefreshTokenTTL time.Duration `env-required:"true" env:"REFRESH_TOKEN_TTL"`
}

type Hasher struct {
	Cost int `env:"HASH_COST" env-default:"12"`
}

func NewConfig(configPath ...string) (*Config, error) {
	cfg := &Config{}
	var err error

	if len(configPath) == 0 {
		err = cleanenv.ReadEnv(cfg)
	} else {
		err = cleanenv.ReadConfig(path.Join("./", configPath[0]), cfg)
	}
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	err = cleanenv.UpdateEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("error updating env: %w", err)
	}

	return cfg, nil
}
