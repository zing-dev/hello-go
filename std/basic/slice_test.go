package basic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
	Slice4()
	s := make([]int, 0)
	s = append(s, 1)
	s = append(s, 2)
	fmt.Println(s)

	var a []byte
	log.Println(append(a, 1, 2, 3))
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

func ReadAll(r io.Reader) ([]byte, error) {
	b := make([]byte, 0, 3)
	for {
		if len(b) == cap(b) {
			// Add more capacity (let append pick how much).
			fmt.Println("==> ", len(b), cap(b))
			a := append(b, 0)
			fmt.Println("a ==> ", len(a), cap(a))
			b = a[:len(b)]
			fmt.Println("===> ", len(b), cap(b))
		}
		n, err := r.Read(b[len(b):cap(b)])
		b = b[:len(b)+n]
		fmt.Println("=> ", len(b), cap(b))
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return b, err
		}
	}
}

func TestSlice12(t *testing.T) {
	b := make([]byte, 0, 4)
	buffer := bytes.NewBuffer([]byte{1, 2, 3, 4})
	n, err := buffer.Read(b[len(b):cap(b)])
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(n, b)
	if len(b) == cap(b) {
		b = append(b, 0)[:len(b)]
		fmt.Println(b)
	}

	fmt.Println(io.ReadAll(bytes.NewBuffer([]byte{1, 2, 3, 4})))
	fmt.Println(ReadAll(bytes.NewBuffer([]byte{1, 2, 3, 4})))
}
func TestSliceAppend(t *testing.T) {
	a := make([]byte, 0, 3)
	fmt.Println(len(a), cap(a))

	a = append(a, 0)
	fmt.Println(len(a), cap(a), a)
	a = append(a, 1)
	a = append(a, 2)
	fmt.Println(len(a), cap(a), a)
	a = append(a, 0)
	fmt.Println(len(a), cap(a), a)
}

func TestRangeSlice(t *testing.T) {
	RangeSlice()
}

func TestRemove(t *testing.T) {
	for _, s := range [][]int{
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{1, 2, 3, 2, 1},
		{3, 2, 1, 2, 3},
	} {
		for i, f := range remove()[0:3] {
			t.Log(i, f(s, 1))
		}
	}
}
