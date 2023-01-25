package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config is global configuration for this application
type Config struct {
	Excel *Excel `yaml:",inline"`
	Jira  *Jira  `yaml:",inline"`
}

type Excel struct {
	ExcelConfig []ExcelConfig `yaml:"excel"`
}

type ExcelConfig struct {
	Path  string `yaml:"path"`
	Sheet string `yaml:"sheet"`
}

type Jira struct {
	JiraConfig []JiraConfig `yaml:"jira"`
}

type JiraConfig struct {
	Endpoint string    `yaml:"endpoint"`
	Project  string    `yaml:"project"`
	JiraAuth *JiraAuth `yaml:"auth"`
}

type JiraAuth struct {
	Method   string `yaml:"method"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
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
	// TODO: it should overwrite old value, not replace
	if err := yaml.Unmarshal(b, &c); err != nil {
		return err
	}
	return nil
}
