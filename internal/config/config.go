package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
)

type Config struct {
	Server   ServerCfg   `yaml:"server"`
	DataBase DataBaseCfg `yaml:"db"`
	Logger   *zap.Logger
}

type DataBaseCfg struct {
	Host     string `yaml:"host"`
	Port     int64  `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
	URL      string `yaml:"url"`
}

type ServerCfg struct {
	Port int64 `yaml:"port"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}

	return MustLoadBuPath(path)
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

func MustLoadBuPath(cfgPath string) *Config {
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		panic("config file does not exists: " + cfgPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}
