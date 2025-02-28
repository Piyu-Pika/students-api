package config

import (
	"log"
	"os"
	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Address            string
	ReadTimeout        int64
	WriteTimeout       int64
	MaxHeaderBytes     int64
	MaxIdleConnections int64
	MaxConnections     int64
	MaxRequestSize     int64
	MaxBodySize        int64
	MaxHeaderSize      int64
}

type Config struct {
	Env         string     `yaml:"env" env:"ENV" env-required:"true env-default:prod"`
	StoragePath string     `yaml:"storage_path"`
	HttpServer  HttpServer `yaml:"http_server"`
}

func MustLoad() *Config {
	var config string

	config = os.Getenv("CONFIG")
	if config == "" {
		config = "config/local.yaml"
	}

	if _, err := os.Stat(config); os.IsNotExist(err) {
		log.Fatalf("config file not found %s", config)
		// panic("config file not found")
	}

	var cfg Config

	err:=cleanenv.ReadConfig(config, &cfg)
	if err != nil {
		log.Fatalf("error loading config file %s: %s", config, err.Error())
	}

	return &cfg

}
