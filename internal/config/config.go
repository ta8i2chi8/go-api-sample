package config

import (
	"context"
	"fmt"
	"regexp"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Env      string `env:"ENV,default=local"`
	Port     string `env:"PORT,default=8070"`
	APIToken string `env:"API_TOKEN,required"`
}

var (
	c *Config
)

func Load(ctx context.Context) (*Config, error) {
	cfg := Config{}
	if err := envconfig.Process(ctx, &cfg); err != nil {
		return nil, err
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	c = &cfg
	return &cfg, nil
}

func Get() (*Config, error) {
	if c == nil {
		return nil, fmt.Errorf("config is not loaded")
	}

	return c, nil
}

func (r *Config) validate() error {
	if !regexp.MustCompile(`^\d+$`).MatchString(r.Port) {
		return fmt.Errorf("PORT must be a number")
	}

	if r.Env != "local" && r.Env != "dev" && r.Env != "prod" {
		return fmt.Errorf("ENV must be one of: local, dev, prod")
	}

	if r.APIToken == "" {
		return fmt.Errorf("API_TOKEN is required")
	}

	return nil
}
