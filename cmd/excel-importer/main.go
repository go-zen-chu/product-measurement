package main

import (
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

func main() {
	c := NewConfig()
	if err := c.LoadCommandArgs(os.Args); err != nil {
		panic(err)
	}
	if c.help {
		fmt.Println(HelpString())
		os.Exit(0)
	}
	f, err := excelize.OpenFile(c.excelFilePath)
	if err != nil {
		panic(err)
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get value from cell by given worksheet name and cell reference.
	rows, err := f.GetRows("")
	if err != nil {
		fmt.Println(err)
		return
	}
	for idx, _ := range rows {
		// for _, colCell := range row {
		// 	fmt.Print(colCell, "\t")
		// }
		fmt.Printf("line: %d\n", idx)
	}
}
