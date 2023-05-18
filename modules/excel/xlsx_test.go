package excel

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"log"
	"testing"
)

func TestXlsx(t *testing.T) {
	file, err := xlsx.OpenFile("13L电磁阀1.xlsx")
	if err != nil {
		t.Fatal(err)
	}

	for _, sheet := range file.Sheet {
		rows := sheet.Rows
		fmt.Println(sheet.Name, len(rows))
		for _, row := range rows {
			fmt.Println(row.Cells)
		}
	}
}

func TestReadXLSX(t *testing.T) {
	rows, _ := f.GetRows(f.GetSheetName(f.GetActiveSheetIndex()))
	log.Println(rows)
	cols, _ := f.GetCols(f.GetSheetName(f.GetActiveSheetIndex()))
	log.Println(cols)
	for i, col := range cols {
		log.Println(i, col)
	}
}
