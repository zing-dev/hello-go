package os

import (
	"os"
	"testing"
)

func TestStat(t *testing.T) {
	info, err := os.Stat("stat_test.go")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Name", info.Name())
	t.Log("Size", info.Size())
	t.Log("IsDir", info.IsDir())
	t.Log("Mode", info.Mode())
	t.Log("ModTime", info.ModTime())
	t.Log("Sys", info.Sys())

	info, err = os.Stat("./stat_test1.go")
	if err != nil {
		t.Log(os.IsExist(err))
		t.Log(os.IsNotExist(err))
	}

	info, err = os.Stat("./stat_test.go")
	t.Log(os.IsExist(err))
	t.Log(os.IsNotExist(err))

	info, err = os.Stat("./www")
	if err != nil {
		t.Log(os.IsExist(err))
		t.Log(os.IsNotExist(err))
		t.Log(info == nil)
	}
}
