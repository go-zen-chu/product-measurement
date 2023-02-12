package config

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var (
	tempDirPath string
)

func TestMain(m *testing.M) {
	tempDirPath, err := os.MkdirTemp("", "internal_config_*")
	if err != nil {
		log.Fatalf("create temp dir: %s", err)
	}
	defer func() {
		if err := os.RemoveAll(tempDirPath); err != nil {
			log.Fatalf("remove all tempdir: %s, %s", tempDirPath, err)
		}
	}()
	log.Printf("start testmain generate temp dir: %s", tempDirPath)
	// run tests
	exitVal := m.Run()
	os.Exit(exitVal)
}

func genTempFile(fileName string, content string) (string, error) {
	f, err := os.CreateTemp(tempDirPath, fileName)
	if err != nil {
		return "", fmt.Errorf("gen temp file: %w", err)
	}
	if _, err := f.WriteString(content); err != nil {
		return "", fmt.Errorf("gen temp file: %w", err)
	}
	return f.Name(), nil
}

func genValidYaml() (string, error) {
	return genTempFile("valid.yaml",
		`version: "0.1"
excel:
- path: "/Users/xxxx/aaa.xlsx"
  sheet: "Sheet 1"
jira:
- endpoint: "https://test.atlassian.net/jira/"
  project: "TEST"
`)
}

func TestConfig_LoadFromFile(t *testing.T) {
	validYamlPath, err := genValidYaml()
	if err != nil {
		t.Errorf("create yaml file: %s", err)
	}
	type fields struct {
		Excel *Excel
		Jira  *Jira
	}
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "If valid yaml given, it should parse yaml file successfully",
			fields:  fields{Excel: nil, Jira: nil},
			args:    args{filePath: validYamlPath},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Excel: tt.fields.Excel,
				Jira:  tt.fields.Jira,
			}
			if err := c.LoadFromFile(tt.args.filePath); (err != nil) != tt.wantErr {
				t.Errorf("Config.LoadFromFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
