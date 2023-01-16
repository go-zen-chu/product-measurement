package runner

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/go-zen-chu/product-measurement/internal/config"
	"github.com/go-zen-chu/product-measurement/internal/log"
)

// Runner defines interface for running general applications
type Runner interface {
	// TODO: LoadFromEnvVars() error
	LoadFromCommandArgs(args []string) error
	SetCommandHandler(ch CommandHandler)
	SetSubCommandHandler(subCommand string, ch CommandHandler)
	Run() error
}

type CommandHandler func(c *config.Config) error

type runner struct {
	appName            string
	args               []string
	flgSet             *flag.FlagSet
	debug              bool
	help               bool
	configFilePath     string
	cnf                *config.Config
	commandHandler     CommandHandler
	subCommandHandlers map[string]CommandHandler
}

func NewRunner(appName string) Runner {
	flgSet := flag.NewFlagSet(appName, flag.ExitOnError)
	return &runner{
		appName:            appName,
		flgSet:             flgSet,
		debug:              false,
		help:               false,
		configFilePath:     "",
		cnf:                config.NewConfig(),
		commandHandler:     nil,
		subCommandHandlers: make(map[string]CommandHandler),
	}
}

// LoadFromCommandArgs parse command args and load config
func (r *runner) LoadFromCommandArgs(args []string) error {
	r.args = args
	debugVal := r.flgSet.Bool("debug", false, "Enable debug")
	helpVal := r.flgSet.Bool("help", false, "Show help")
	configFilePathVal := &pathValue{
		path: "",
	}
	r.flgSet.Var(configFilePathVal, "config-path", "Set configuration file path")
	if !r.flgSet.Parsed() {
		if len(args) <= 1 {
			return fmt.Errorf("invalid args len: %d", len(os.Args))
		}
		if err := r.flgSet.Parse(args[1:]); err != nil {
			return fmt.Errorf("parse args: %s", err)
		}
		// visit specified flag
		r.flgSet.Visit(func(f *flag.Flag) {
			switch f.Name {
			case "debug":
				r.debug = *debugVal
			case "help":
				r.help = *helpVal
			case "config-path":
				r.configFilePath = configFilePathVal.path
			}
		})
		if r.configFilePath != "" {
			if err := r.cnf.LoadFromFile(r.configFilePath); err != nil {
				return fmt.Errorf("while loading from config file: %w", err)
			}
		}
	}
	return nil
}

func (r *runner) buildHelpString() string {
	var sb strings.Builder
	sb.WriteString("usage: ")
	sb.WriteString(r.appName)
	sb.WriteString(" <flags>\n")
	// set print setting to string builder
	op := r.flgSet.Output()
	r.flgSet.SetOutput(&sb)
	r.flgSet.PrintDefaults()
	r.flgSet.SetOutput(op)
	return sb.String()
}

func (r *runner) SetCommandHandler(ch CommandHandler) {
	r.commandHandler = ch
}

func (r *runner) SetSubCommandHandler(subCommand string, ch CommandHandler) {
	r.subCommandHandlers[subCommand] = ch
}

func (r *runner) Run() error {
	if r.help {
		fmt.Println(r.buildHelpString())
		return nil
	}
	if err := log.Init(r.debug); err != nil {
		return err
	}
	log.Debugf("[Run] config: %+v", r.cnf)
	subCommandArgs := r.args[1+r.flgSet.NFlag():] // NFlag is number of flags for root command
	if len(subCommandArgs) == 0 {
		if r.commandHandler == nil {
			return errors.New("command handler is not set")
		}
		return r.commandHandler(r.cnf)
	}
	if r.subCommandHandlers == nil {
		return errors.New("subcommand handler is not set")
	}
	subCommand := subCommandArgs[0]
	if _, ok := r.subCommandHandlers[subCommand]; !ok {
		return fmt.Errorf("can not find subcommand: %s", subCommand)
	}
	return r.subCommandHandlers[subCommand](r.cnf)
}

// pathValue is defined to handle path type argument
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
