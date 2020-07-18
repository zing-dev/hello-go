package compress

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestZlib(t *testing.T) {
	var b bytes.Buffer
	src := []byte("hello, world\nhello, world\nhello, world\nhello, world\nhello, world\n")
	fmt.Println("src len ", len(src))
	w := zlib.NewWriter(&b)
	_, _ = w.Write(src)
	_ = w.Close()
	desc := b.Bytes()
	fmt.Println("desc len ", len(desc))
	fmt.Println(desc)

	r, _ := zlib.NewReader(bytes.NewBuffer(desc))
	data, _ := ioutil.ReadAll(r)
	fmt.Println(string(data))

}

func TestZlibStruct(t *testing.T) {
	start := time.Now()
	type Student struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Sex   int    `json:"sex"`
		Grade int    `json:"grade"`
	}
	length := 1000000
	var students = make([]Student, length)
	for i := 0; i < length; i++ {
		students[i] = Student{
			Id:    i + 1,
			Name:  fmt.Sprintf("name-%d", i+1),
			Sex:   rand.Int() % 2,
			Grade: rand.Int() % 12,
		}
	}

	var b bytes.Buffer
	src, _ := json.Marshal(students)
	fmt.Println("src len ", len(src))
	w := zlib.NewWriter(&b)
	_, _ = w.Write(src)
	_ = w.Close()
	desc := b.Bytes()
	fmt.Println("desc len ", len(desc))
	b.Reset()
	r, _ := zlib.NewReader(bytes.NewBuffer(desc))

	//n, _ := b.ReadFrom(r)
	//desc = b.Bytes()

	//data := make([]byte,len(src))
	//n2, _ := r.Read(data)
	//fmt.Println("n len ", n2)

	desc, _ = ioutil.ReadAll(r)
	var students2 []Student

	//ioutil.WriteFile("1.json", desc, 7777)
	//fmt.Println(string(desc))
	//students2 := make([]Student, length)
	err := json.Unmarshal(desc, &students2)
	if err != nil {
		log.Fatal("err ", err)
	}
	end := time.Now()
	fmt.Println(students2[1], students2[2])
	fmt.Println("time ", end.Sub(start))
	time.Sleep(time.Minute)
}
