package basic

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"
)

func TestSlice1(t *testing.T) {
	Slice1()
}

func TestSlice2(t *testing.T) {
	Slice2()
}

func TestSlice3(t *testing.T) {
	Slice3()
}

func TestSlice4(t *testing.T) {
	s := make([]int, 0)
	s = append(s, 1)
	s = append(s, 2)
	fmt.Println(s)
	Slice4()
}

func TestSlice5(t *testing.T) {
	a := make([]int, 3)
	a[100] = 1
	a[200] = 1
	a[300] = 1
	fmt.Println(a)
}

func TestSlice6(t *testing.T) {
	arr := make([]int, 0)
	for i := 0; i < 10; i++ {
		arr = append(arr, i)
	}
	log.Println(arr)
	arr = append(arr[:3], arr[4:]...)
	log.Println(arr)
	arr = append(arr, 11)
	log.Println(arr)

	m := map[string][]int{}
	m["1"] = []int{1, 2, 3}
	m["1"][0] = 11
	one := m["1"]
	one[0] = 12
	one = []int{}

	m["2"] = arr
	arr = append(arr[:3], arr[9:]...)
	arr = m["2"]
	arr = []int{}

	log.Println(m)
}

func TestSlice7(t *testing.T) {
	t.Log(len([]byte("")))
	t.Log(len([]byte("[]")))
	data, err := json.Marshal(nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))

	data, err = json.Marshal("")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
	t.Log(string([]byte{}))
	t.Log(string([]byte{'a'}))

	data, err = json.Marshal([]byte{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}

func TestSlice8(t *testing.T) {
	arr := "hello"
	t.Log(arr[1:2])
	t.Log(arr[1:5])
	t.Log(arr[:5])
	t.Log(arr[1:])
}

func NumberStrToSlice(str string) {
	var data []byte
	for _, v := range strings.Split(str, ",") {
		i, err := strconv.Atoi(v)
		if err == nil {
			data = append(data, byte(i))
		}
	}
	fmt.Println(data)
}

func TestNumberStrToSlice(t *testing.T) {
	NumberStrToSlice("1,2,3,4,5,10,111,aaa")
}
