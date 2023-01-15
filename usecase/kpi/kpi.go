package kpi

import (
	"fmt"

	"github.com/go-zen-chu/product-measurement/internal/config"
	"github.com/xuri/excelize/v2"
)

func ImportExcel(config *config.Config) error {
	for _, ec := range config.Excel.ExcelConfig {
		f, err := excelize.OpenFile(ec.Path)
		if err != nil {
			return err
		}
		defer func() {
			if err := f.Close(); err != nil {
				fmt.Println(err)
			}
		}()
		// Get value from cell by given worksheet name and cell reference.
		rows, err := f.GetRows("")
		if err != nil {
			return err
		}
		for idx, _ := range rows {
			// for _, colCell := range row {
			// 	fmt.Print(colCell, "\t")
			// }
			fmt.Printf("line: %d\n", idx)
		}
	}
	return nil
}
