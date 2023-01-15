package main

import (
	"fmt"
	"os"

	"github.com/go-zen-chu/product-measurement/internal/config"
	"github.com/go-zen-chu/product-measurement/internal/log"
	"github.com/go-zen-chu/product-measurement/usecase/kpi"
	"github.com/go-zen-chu/product-measurement/usecase/schedule"
)

func main() {
	rc := NewRuntimeConfig()
	if err := rc.LoadCommandArgs(os.Args); err != nil {
		panic(err)
	}
	if rc.help {
		fmt.Println(HelpString())
		os.Exit(0)
	}
	if err := log.Init(rc.debug); err != nil {
		log.Fatal(err)
	}
	c := config.NewConfig()
	if err := c.LoadFromFile(rc.configFilePath); err != nil {
		log.Fatal(err)
	}

	// handle subcommands
	subCommandArgs := os.Args[1+flgSet.NFlag():] // NFlag is number of flags for root command
	if len(subCommandArgs) == 0 {
		log.Fatal("no subcommand")
	}
	switch subCommand := subCommandArgs[0]; subCommand {
	case "import-excel":
		kpi.ImportExcel(c)
	case "import-jira":
		schedule.ImportJira(c)
	default:
		fmt.Println(HelpString())
		os.Exit(0)
	}
}
