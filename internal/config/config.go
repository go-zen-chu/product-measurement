package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config is global configuration for this application
type Config struct {
	DataSources *DataSources `yaml:"datasources"`
}

type DataSources struct {
	Excel *Excel `yaml:",inline"`
	Jira  *Jira  `yaml:",inline"`
}

type Importer struct {
	Path string `yaml:"path"`
}

type Excel struct {
	ExcelConfig []ExcelConfig `yaml:"excel"`
}

type ExcelConfig struct {
	Name     string    `yaml:"name"`
	Path     string    `yaml:"path"`
	Importer *Importer `yaml:"importer"`
}

type Jira struct {
	JiraConfig []JiraConfig `yaml:"jira"`
}

type JiraConfig struct {
	Name     string    `yaml:"name"`
	Endpoint string    `yaml:"endpoint"`
	Project  string    `yaml:"project"`
	JiraAuth *JiraAuth `yaml:"auth"`
}

type JiraAuth struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func validateNormalizeConfig(configFilePath string, c *Config) error {
	if c.DataSources == nil {
		return errors.New("datasources must be specified, but got nil")
	}
	if c.DataSources.Excel != nil {
		for _, exCnf := range c.DataSources.Excel.ExcelConfig {
			if exCnf.Name == "" {
				return errors.New("name must be specified")
			}
			p, err := getAbsPath(configFilePath, exCnf.Path)
			if err != nil {
				return fmt.Errorf("handling path of %s: %w", exCnf.Name, err)
			}
			exCnf.Path = p
			if exCnf.Importer == nil {
				return fmt.Errorf("importer must be specified for %s", exCnf.Name)
			}
			ip, err := getAbsPath(configFilePath, exCnf.Importer.Path)
			if err != nil {
				return fmt.Errorf("handling path of %s: %w", exCnf.Name, err)
			}
			exCnf.Importer.Path = ip
		}
	}
	return nil
}

func getAbsPath(configFilePath string, targetPath string) (string, error) {
	absPath := targetPath
	if !filepath.IsAbs(targetPath) {
		parent := filepath.Dir(configFilePath)
		absPath = filepath.Join(parent, targetPath)
	}
	if _, err := os.Stat(absPath); err != nil {
		return "", err
	}
	return absPath, nil
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
