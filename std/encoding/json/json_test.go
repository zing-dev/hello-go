package json

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"testing"
)

func TestStructToBytes(t *testing.T) {
	StructToBytes()
}

func TestBytesToStruct(t *testing.T) {
	BytesToStruct()
}

func TestJSON(t *testing.T) {
	JSON()
}

func TestJSON2(t *testing.T) {
	type Route struct {
		Name   string `json:"name"`
		Method string `json:"method"`
	}
	var R = map[Route]string{
		Route{
			Name:   "system/sms",
			Method: "post",
		}: "保存短信配置",
		Route{
			Name:   "system/sms",
			Method: "post",
		}: "保存短信配置",
	}

	data, err := json.Marshal(R)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}

func TestJSON3(t *testing.T) {
	type Version struct {
		Version   string `json:"version"`
		Log       string `json:"log"`
		Status    string `json:"status"`
		GitHash   string `json:"git_hash"`
		updatedAt string
		hash      string
	}

	type Soft struct {
		Name      string    `json:"name"`
		Alias     string    `json:"alias"`
		Author    string    `json:"author"`
		Version   []Version `json:"versions"`
		Copyright string    `json:"copyright"`
		Inherit   bool      `json:"inherit"`
	}

	v := `
{
  "name": "版本控制-软件",
  "alias": "别名",
  "author": "作者",
  "versions": [
 	{
      "version": "0.0.2",
      "log": "add",
      "status": "Base",
      "git_hash": "1f10cb1d57ef6e1ce962634148c95223bd2e4d0f"
    },
    {
      "version": "0.0.1",
      "log": "init",
      "status": "Base",
      "git_hash": "1f10cb1d57ef6e1ce962634148c95223bd2e4d0f"
    }
  ],
  "copyright": "All rights reserved",
  "inherit": true
}`

	s := &Soft{}
	err := json.Unmarshal([]byte(v), s)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
	t.Log(s.Version[0])
	t.Log(s.Version[1])
}

func TestName(t *testing.T) {
	str := "hello,1,1,1,1,1,1,"
	fmt.Println(str[:len(str)-1])
	fmt.Println(fmt.Sprintf("%v", 11.1111))
	fmt.Println(fmt.Sprintf("%v", 11.002))
	fmt.Println("1"[:0])
	fmt.Println("1"[:0])
}

func TestMarshalSlice(t *testing.T) {
	arr := []byte{1, 2, 3, 4}
	data, err := json.Marshal(arr)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(arr)
	t.Log(data)
	t.Log(string(data))

	arr0 := []uint8{1, 2, 3, 4}
	data, err = json.Marshal(arr0)
	t.Log(arr0)
	t.Log(data)
	t.Log(string(data))

	arr2 := []int{1, 2, 3, 4}
	data, err = json.Marshal(arr2)
	t.Log(arr2)
	t.Log(data)
	t.Log(string(data))

	arr3 := [5]int{1, 2, 3, 4}
	data, err = json.Marshal(arr3)
	t.Log(arr3)
	t.Log(data)
	t.Log(string(data))

	arr4 := []uint16{1, 2, 3, 4}
	data, err = json.Marshal(arr4)
	t.Log(arr4)
	t.Log(data)
	t.Log(string(data))
}

func TestMarshalIndent(t *testing.T) {
	data, err := json.MarshalIndent(map[string]interface{}{
		"id":   "1",
		"name": "zing",
		"age":  25,
	}, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
}

func TestName1(t *testing.T) {
	a := [][]float32{
		{1.111111, 2.22220000},
		{.333333, 1111111111111111.4444454444111111111},
		{.333333, 0},
		{.333333, float32(math.Inf(0)), math.Float32frombits(math.MaxUint32)},
	}

	data, err := json.Marshal(a)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))

	a = [][]float32{}
	data, err = json.Marshal(a)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))

	a = nil
	data, err = json.Marshal(a)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
}

func TestName2(t *testing.T) {
	fmt.Println(float32(math.Inf(0)))
	fmt.Println(float32(math.Inf(-1)))
	fmt.Println(float32(math.Inf(1)))

	a := make([][]float32, 10)
	fmt.Println(a[0])
	fmt.Println(a[0] == nil)
	for i := 0; i < 10; i++ {
		a[0] = append(a[0], float32(i))
	}
	a[0] = append(a[0], float32(math.Inf(0)))
	fmt.Println(a)
	fmt.Println(1 > math.Inf(0))
	fmt.Println(1 + math.Inf(0))
	fmt.Println(math.Inf(0) / 2)
	fmt.Println(math.Inf(0) - math.Inf(0))
	fmt.Println(math.MaxFloat32)

	fmt.Println(math.Float32bits(math.MaxFloat32))
	fmt.Println(math.Float32bits(0))
	fmt.Println(math.Float32frombits(math.MaxUint32) + 1)
	fmt.Println(math.Float32frombits(math.MaxUint16 * 2))
	fmt.Println(0x7FF0000000000000)
	fmt.Println(fmt.Sprintf("%b", 0x7FF0000000000000))
	fmt.Println(math.Float64frombits(0x7FF0000000000000))
	fmt.Println(math.Float64frombits(0xFFF0000000000000))
	fmt.Println(fmt.Sprintf("==> %b", math.Float64frombits(0xFFF0000000000000)))
	fmt.Println(math.Float32frombits(0x7FF00000))
	fmt.Println(math.Float32frombits(0xffffffff))
	fmt.Println(math.Float32frombits(0x7F000000))
}

func Decimal(value float32) float32 {
	return float32(math.Trunc(float64(value*1e1+0.5)) / 1e1)
}

func TestJSONOmit(t *testing.T) {
	type User struct {
		Id   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	}
	u := User{
		Id: 1,
	}
	data, _ := json.Marshal(u)
	log.Println(string(data))

	u2 := new(User)
	_ = json.Unmarshal(data, u2)
	log.Println(u2.Id)
	log.Println(u2.Name)
}
