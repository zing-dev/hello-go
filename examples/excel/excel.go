package excel

import (
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/tealeg/xlsx"
)

func test() {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func test2() {
	file := xlsx.NewFile()

	sheet, err := file.AddSheet("Sheet2")
	if err != nil {
		log.Fatal(err)
	}
	sheet.Name = "hello"
	row := sheet.AddRow()
	row.AddCell().Value = "zone"
	sheet.AddRow().AddCell().Value = "1"
	sheet.AddRow().AddCell().Value = "2"
	err = file.Save("xlsx.xlsx")
	if err != nil {
		log.Fatal(err)
	}
}
