package excel

import (
	"github.com/xuri/excelize/v2"
	"log"
	"testing"
)

var f, err = excelize.OpenFile("./data.xlsx")

func TestSheet(t *testing.T) {
	log.Println(f.GetActiveSheetIndex())
	f.SetActiveSheet(2)
	log.Println(f.GetActiveSheetIndex())

	log.Println(f.SheetCount)
	log.Println(f.GetSheetMap())
	log.Println(f.GetSheetName(1))
	log.Println(f.GetSheetIndex(f.GetSheetName(1)))
	log.Println(f.GetSheetVisible(f.GetSheetName(2)))
	log.Println(f.GetSheetName(0))
	log.Println(f.GetSheetName(1))
	log.Println(f.GetSheetName(2))
}

func TestRow(t *testing.T) {
	rows, err := f.Rows(f.GetSheetName(1))
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		columns, err := rows.Columns()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(columns)
		for k, v := range columns {
			log.Println(k, v)
		}
	}

	row2, err := f.GetRows(f.GetSheetName(1))
	log.Println(row2)
	for _, row := range row2 {
		log.Println(row)
	}
}

func TestRead(t *testing.T) {
	if err != nil {
		log.Fatal(err)
	}
	log.Println(f.GetActiveSheetIndex())
	m := f.GetSheetMap()
	log.Println(m)
}
