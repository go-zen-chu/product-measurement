package main

import (
	"fmt"
	"log"

	"github.com/go-zen-chu/product-measurement/usecase/importer"
	"github.com/xuri/excelize/v2"
)

func main() {
	eh, err := importer.GetExcelHandler()
	if err != nil {
		log.Fatalf("get excel handler: %s", err)
	}
	if err := eh.Handle(func(f *excelize.File) error {
		rows, err := f.GetRows("")
		if err != nil {
			return err
		}
		for idx, row := range rows {
			fmt.Printf("line: %d\n", idx)
			for _, colCell := range row {
				fmt.Print(colCell, "\t")
			}
			fmt.Println()
		}
		return nil
	}); err != nil {
		log.Fatalf("get excel handler: %s", err)
	}
}
