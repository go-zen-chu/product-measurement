package main

import (
	"os"

	"github.com/go-zen-chu/product-measurement/internal/config"
	"github.com/go-zen-chu/product-measurement/internal/log"
	"github.com/go-zen-chu/product-measurement/internal/runner"
	"github.com/go-zen-chu/product-measurement/usecase/kpi"
	"github.com/go-zen-chu/product-measurement/usecase/schedule"
)

func main() {
	runner := runner.NewRunner("importer")
	if err := runner.LoadFromCommandArgs(os.Args); err != nil {
		panic(err)
	}
	runner.SetSubCommandHandler("import-excel", func(c *config.Config) error {
		return kpi.ImportExcel(c)
	})
	runner.SetSubCommandHandler("import-jira", func(c *config.Config) error {
		return schedule.ImportJira(c)
	})
	if err := runner.Run(); err != nil {
		log.Fatal(err)
	}
}
