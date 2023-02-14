package importexcel

import (
	"fmt"
	"math"

	"github.com/go-zen-chu/product-measurement/internal/config"
	"github.com/xuri/excelize/v2"
)

func convertToColAlphabet(col int) (string, error) {
	const (
		asciiAlphabetStart = 65
		alphabetCount      = 26
	)
	if col < 1 {
		return "", fmt.Errorf("argument is out of range [%d]", col)
	}
	var colName string
	tmp := col

	for tmp > 0 {
		index := tmp - 1
		remaining := index / alphabetCount
		charIndex := int(math.Mod(float64(index), alphabetCount))
		colName = string(rune(charIndex+asciiAlphabetStart)) + colName
		tmp = remaining
	}
	return colName, nil
}

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
		rows, err := f.GetRows(ec.Sheet)
		if err != nil {
			return err
		}
		for idx, r := range rows {
			fmt.Printf("line %d: %d %+v\n", idx, len(r), r)
		}
	}
	return nil
}
