package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	HttpPort int64 `yaml:""`
}
