package importer

import (
	"errors"
	"log"
	"path/filepath"
	"runtime"

	"github.com/go-zen-chu/product-measurement/internal/config"
	"github.com/xuri/excelize/v2"
)

var excelHandlers map[string]ExcelHandler

type UseCaseImporter interface {
	UseCaseImportAll() error
}

type useCaseImporter struct {
	cnf *config.Config
}

func NewUseCaseImporter(cnf *config.Config) UseCaseImporter {
	return &useCaseImporter{
		cnf: cnf,
	}
}

// Import from all datasources
func (uci *useCaseImporter) UseCaseImportAll() error {
	if uci.cnf.DataSources.Excel != nil {
		excelHandlers = make(map[string]ExcelHandler)
		for _, exCnf := range uci.cnf.DataSources.Excel.ExcelConfig {
			excelHandlers[exCnf.Importer.Path] = ExcelHandler{
				ExcelConfig: &exCnf,
			}
			// run implemented importer
		}
	}
	return nil
}

type ExcelProcessStore func(f *excelize.File) error

type ExcelHandler struct {
	ExcelConfig *config.ExcelConfig
}

func (eh *ExcelHandler) Handle(eps ExcelProcessStore) error {
	f, err := excelize.OpenFile(eh.ExcelConfig.Path)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("closing excel file: %s", err)
		}
	}()
	return eps(f)
}

func GetExcelHandler() (*ExcelHandler, error) {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.New("could not identify caller")
	}
	fileName := filepath.Base(file)
	// TBD
	importerId := "excel/" + fileName
	eh := excelHandlers[importerId]
	return &eh, nil
}
