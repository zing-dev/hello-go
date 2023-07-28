package encoding

import (
	"encoding/csv"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"testing"
)

func TestCSV(t *testing.T) {
	data := `"Image Name","PID","Session Name","Session#","Mem Usage"
"msedge.exe","13988","Console","1","321,284 K"
"msedge.exe","14296","Console","1","3,684 K"
"msedge.exe","20356","Console","1","107,080 K"
"msedge.exe","20156","Console","1","106,820 K"
"msedge.exe","5020","Console","1","29,404 K"

`
	r := csv.NewReader(strings.NewReader(data))
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records[1:] {
		fmt.Println(record[1])
	}
}

func TestRun(t *testing.T) {
	res, err := exec.Command("cmd", "/c", "tasklist", "/FI", "IMAGENAME eq msedge.exe", "/FO", "csv").Output()
	fmt.Println(string(res), err)
	r := csv.NewReader(strings.NewReader(string(res)))
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, record := range records[1:] {
		fmt.Println(record[1])
	}
}
