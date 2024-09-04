// Package config stores all logic for our config knobs.
package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

// Config struct holds configuration knobs.
type Config struct {
	PassRequestHeaders  []string `yaml:"request_headers"`
	PassResponseHeaders []string `yaml:"response_headers"`
}

func New(path string) (Config, error) {
	return loadCfg(path)
}

func loadCfg(path string) (Config, error) {
	cfg := Config{}
	f, err := os.Open(path)
	if err != nil {
		return Config{}, fmt.Errorf("error opening config : %w", err)
	}

	cfgB, err := io.ReadAll(f)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config bytes : %w", err)
	}

	err = yaml.Unmarshal(cfgB, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("error unmarshaling config : %w", err)
	}

	return cfg, nil
}
