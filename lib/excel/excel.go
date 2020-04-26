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

func test3() {
	const (
		one = "通道一"
		two = "通道二"
	)
	file := excelize.NewFile()
	index := file.NewSheet(one)
	file.SetCellValue(one, "A1", "防区")
	file.SetCellValue(one, "A2", "1-1-1-1")
	file.SetCellValue(one, "A3", "1-1-1-2")
	file.SetCellValue(one, "A4", "1-1-1-3")

	file.SetCellValue(one, "B1", "2020-03-12 17:00:00")
	file.SetCellValue(one, "B2", "30.0000")
	file.SetCellValue(one, "B3", "32.0000")
	file.SetCellValue(one, "B4", "33.0000")

	file.NewSheet(two)
	file.SetCellValue(two, "A1", "防区")
	file.SetCellValue(two, "A2", "1-1-1-1")
	file.SetCellValue(two, "A3", "1-1-1-2")
	file.SetCellValue(two, "A4", "1-1-1-3")

	file.SetCellValue(two, "B1", "2020-03-12 17:00:00")
	file.SetCellValue(two, "B2", "30.0000")
	file.SetCellValue(two, "B3", "32.0000")
	file.SetCellValue(two, "B4", "33.0000")
	file.SetActiveSheet(index)

	file.DeleteSheet("Sheet1")
	err := file.SaveAs("data.xlsx")
	if err != nil {
		log.Fatal(err)
	}
}
