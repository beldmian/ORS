package config

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

// SeverConfig ...
type SeverConfig struct {
	Port     string `toml:"port"`
	LogLevel string `toml:"log_level"`
	LogType  string `toml:"log_type"`
	LogFile  string `toml:"log_file"`
}

// NewConfig ...
func NewConfig(path string) (SeverConfig, error) {
	var conf SeverConfig

	file, err := os.Open(path)
	if err != nil {
		return SeverConfig{}, err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)

	if _, err := toml.Decode(string(b), &conf); err != nil {
		return SeverConfig{}, err
	}
	return conf, nil
}
