package excel

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strings"
	"testing"
)

func TestDemo2(t *testing.T) {
	t.Log(strings.Split("1,2,3,4,", ","))
	t.Log(strings.Join(strings.Split("1,2,3,4,", ","), "."))
}

func TestGetCellValue(t *testing.T) {
	file, err := excelize.OpenFile("13L电磁阀.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(file.GetSheetMap())
	rows, err := file.Rows("继电器3")
	if err != nil {
		t.Fatal(err)
	}
	i := 0
	for rows.Next() {
		t.Log(rows.Columns())
		i++
	}
	fmt.Println("继电器4", i)
	t.Log(file.GetCellValue("继电器4", "A1"))
	t.Log(file.GetCellValue("继电器4", "A631"))
	t.Log(file.GetCellValue("继电器4", "A632"))
}

func TestGetSheetMap(t *testing.T) {
	file, err := excelize.OpenFile("13L电磁阀.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(file.GetSheetMap())
	for _, v := range file.GetSheetMap() {
		i := 0
		if strings.HasPrefix(v, "继电器") {
			rows, err := file.Rows(v)
			if err != nil {
				t.Fatal(err)
			}
			for rows.Next() {
				t.Log(rows.Columns())
				i++
			}
			fmt.Println(v, i)
		}
	}
	file.SetActiveSheet(4)

}

func TestDemo(t *testing.T) {
	file, err := excelize.OpenFile("data.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(file.GetSheetMap())

	t.Log(file.GetActiveSheetIndex())
	name := file.GetSheetName(file.GetActiveSheetIndex())
	t.Log(name == "通道一")
	rows, err := file.Rows(file.GetSheetName(file.GetActiveSheetIndex()))
	if err != nil {
		t.Fatal(err)
	}
	for rows.Next() {
		t.Log(rows.Columns())
	}

	//for _, name := range file.GetSheetMap() {
	//	if sheet, ok := file.Sheet[name]; ok {
	//		t.Log(name)
	//		t.Log(len(sheet.Cols.Col))
	//	} else {
	//		t.Fatal(name, " nil")
	//	}
	//}

	t.Log(file.SheetCount)
	//t.Log(file.Sheet)
	//for _, name := range []string{"通道一", "通道二"} {
	//	if sheet, ok := file.Sheet[name]; ok {
	//		t.Log(name)
	//		t.Log(len(sheet.Cols.Col))
	//	} else {
	//		t.Fatal(name, " nil")
	//	}
	//}
}
