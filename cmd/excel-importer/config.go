package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	appName = "excel-importer"
)

var (
	flgSet           = flag.NewFlagSet(appName, flag.ExitOnError)
	debugVal         = flgSet.Bool("debug", false, "Enable debug")
	helpVal          = flgSet.Bool("help", false, "Show help")
	excelFilePathVal *pathValue
)

func init() {
	excelFilePathVal = &pathValue{
		path: "",
	}
	flgSet.Var(excelFilePathVal, "excel-file-path", "Set excel file path")
}

func HelpString() string {
	var sb strings.Builder
	sb.WriteString("usage: ")
	sb.WriteString(appName)
	sb.WriteString(" <flags>\n")
	op := flgSet.Output()
	flgSet.SetOutput(&sb)
	// print to string builder
	flgSet.PrintDefaults()
	flgSet.SetOutput(op)
	return sb.String()
}

type config struct {
	debug         bool
	help          bool
	excelFilePath string
}

func NewConfig() *config {
	// return default config
	return &config{
		debug:         false,
		help:          false,
		excelFilePath: "",
	}
}

func (c *config) LoadCommandArgs(args []string) error {
	if !flgSet.Parsed() {
		if len(args) <= 1 {
			return fmt.Errorf("invalid args len: %d", len(os.Args))
		}
		if err := flgSet.Parse(args[1:]); err != nil {
			return fmt.Errorf("parse args: %s", err)
		}
		// visit specified flag
		flgSet.Visit(func(f *flag.Flag) {
			switch f.Name {
			case "debug":
				c.debug = *debugVal
			case "help":
				c.help = *helpVal
			case "excel-file-path":
				c.excelFilePath = excelFilePathVal.path
			}
		})
	}
	return nil
}

type pathValue struct {
	path string
}

// implements Value interface for flag argument
func (pv *pathValue) String() string {
	return pv.path
}

// implements Value interface for flag argument
func (pv *pathValue) Set(path string) error {
	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("not valid path %s: %w", path, err)
	}
	pv.path = path
	return nil
}
