package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

// Config is global configuration for this application
type Config struct {
	DataSources *DataSources `yaml:"datasources"`
	DataBase    *DataBase    `yaml:"database"`
}

type DataSources struct {
	Jira *Jira `yaml:",inline"`
}

type Jira struct {
	JiraConfigs []JiraConfig `yaml:"jira"`
}

type JiraConfig struct {
	Name     string    `yaml:"name"`
	Endpoint string    `yaml:"endpoint"`
	Project  string    `yaml:"project"`
	JiraAuth *JiraAuth `yaml:"auth"`
}

type JiraAuth struct {
	Method   string `yaml:"method"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type DataBase struct {
	Type   string  `yaml:"type"`
	Host   string  `yaml:"host"`
	DBAuth *DBAuth `yaml:"auth"`
}

type DBAuth struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func validateNormalizeConfig(configFilePath string, c *Config) error {
	if c.DataSources == nil {
		return errors.New("datasources must be specified, but got nil")
	}
	if c.DataBase == nil {
		return errors.New("database must be specified, but got nil")
	}
	return nil
}

func LoadFromFile(filePath string) (*Config, error) {
	c := &Config{}
	b, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(b, &c); err != nil {
		return nil, err
	}
	// change all path to absolute path for making things easy
	err = validateNormalizeConfig(filePath, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
