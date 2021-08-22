package config

import "github.com/BurntSushi/toml"

type Config struct {
	DB *DBConfig `toml:"db"`
}

type DBConfig struct {
	MySQL *MySQLConfig `toml:"mysql"`
}

type MySQLConfig struct {
	URL string `toml:"url"`
}

func New(conf string) (*Config, error) {
	res := &Config{}
	_, err := toml.DecodeFile(conf, res)
	return res, err
}
