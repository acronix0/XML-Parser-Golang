package config

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	EnvLocal = "local"
	EnvProd  = "prod"
)

type Config struct {
	Env       string         `yaml:"env" env:"ENV" env-default:"local"`
	Database  DatabaseConfig `yaml:"database"`
	GRPC      GRPCConfig     `yaml:"grpc"`
	XMLConfig XMLConfig      `yaml:"xml"`
	MigrationsFolder string `yaml:"migrations_folder"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int `yaml:"port"`
	Name     string `yaml:"name"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type XMLConfig struct {
	FirstFilePath string `yaml:"first_file_path"`
	SecondFilePath string `yaml:"second_file_path"`
}

func Load() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		log.Fatal("Config path not set")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("config file %s does not exist\n", configPath)
	}
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config %s", configPath)
	}
	return &cfg
}

func fetchConfigPath() string {
	var res string
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}
