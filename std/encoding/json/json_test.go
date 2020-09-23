package json

import (
	"encoding/json"
	"fmt"
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
