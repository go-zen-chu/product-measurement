package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config is global configuration for this application
type Config struct {
	Excel *Excel `yaml:"excel"`
	Jira  *Jira  `yaml:"jira"`
}

type Excel struct {
	ExcelConfig []ExcelConfig
}

type ExcelConfig struct {
	Path  string `yaml:"path"`
	Sheet string `yaml:"sheet"`
}

type Jira struct {
	JiraConfig []JiraConfig
}

type JiraConfig struct {
	Endpoint string `yaml:"endpoint"`
	Project  string `yaml:"project"`
}

// NewConfig generate default config
func NewConfig() *Config {
	return &Config{}
}

func (c *Config) LoadFromFile(filePath string) error {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	// TODO: should be overwrite old value, not replace
	if err := yaml.Unmarshal(b, &c); err != nil {
		return err
	}
	return nil
}
