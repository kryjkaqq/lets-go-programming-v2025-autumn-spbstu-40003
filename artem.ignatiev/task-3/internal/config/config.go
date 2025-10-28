package config

import (
	"errors"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputPath  string `yaml:"input-file"`
	OutputPath string `yaml:"output-file"`
}

var (
	ErrMissingInput  = errors.New("input-file not set")
	ErrMissingOutput = errors.New("output-file not set")
)

func Load(path *string) (*Config, error) {
	data, err := ioutil.ReadFile(*path)
	if err != nil {
		return nil, fmt.Errorf("cannot read config file: %w", err)
	}

	cfg := &Config{}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("cannot parse YAML: %w", err)
	}

	if cfg.InputPath == "" {
		return nil, ErrMissingInput
	}
	if cfg.OutputPath == "" {
		return nil, ErrMissingOutput
	}

	return cfg, nil
}
