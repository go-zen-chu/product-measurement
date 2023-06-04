package main

import (
	"os"

	"github.com/go-zen-chu/product-measurement/internal/config"
	"github.com/go-zen-chu/product-measurement/internal/log"
	"github.com/go-zen-chu/product-measurement/internal/runner"
	"github.com/go-zen-chu/product-measurement/usecase/importer"
)

func main() {
	runner := runner.NewRunner("importer")
	if err := runner.LoadFromCommandArgs(os.Args); err != nil {
		panic(err)
	}
	runner.SetCommandHandler(func(c *config.Config) error {
		uci := importer.NewUseCaseImporter(c)
		return uci.UseCaseImportAll()
	})
	if err := runner.Run(); err != nil {
		log.Fatal(err)
	}
}
