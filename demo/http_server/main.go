package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"text/template"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("hello world "))
	})

	http.HandleFunc("/file", func(writer http.ResponseWriter, request *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				_, _ = writer.Write([]byte("出错了"))
			}

		}()
		file, e := os.Open("./demo/websocket/main.go")

		if e != nil {
			log.Println("打开文件失败")
			panic("打开文件失败")
		}
		info, _ := file.Stat()
		content := make([]byte, info.Size())
		_, e = file.Read(content)
		if e != nil {
			log.Println("读文件失败")
			panic("读文件失败")
		}
		_, _ = writer.Write(content)
	})

	http.HandleFunc("/json", func(writer http.ResponseWriter, request *http.Request) {

		header := writer.Header()
		header.Set("Content-Type", "application/json")
		hash := md5.New()
		hash.Write([]byte("hello world"))
		sign := hex.EncodeToString(hash.Sum(nil))

		content := make(map[string]string, 10)
		content["language"] = "golang"
		content["version"] = runtime.Version()
		content["goos"] = runtime.GOOS
		content["username"] = os.Getenv("username")
		content["sign"] = sign

		bytes, e := json.Marshal(content)

		if e != nil {
			fmt.Println("err")
		}

		header.Set("sign", sign)
		_, _ = writer.Write(bytes)
	})

	http.HandleFunc("/template", func(writer http.ResponseWriter, request *http.Request) {

		// Define a template.
		const letter = `
Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the wedding.
{{- else}}
It is a shame you couldn't make it to the wedding.
{{- end}}
{{with .Gift -}}
Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Josie
`
		// Prepare some data to insert into the template.
		type Recipient struct {
			Name, Gift string
			Attended   bool
		}
		var recipients = []Recipient{
			{"Aunt Mildred", "bone china tea set", true},
			{"Uncle John", "moleskin pants", false},
			{"Cousin Rodney", "", false},
			{"zing", "airbus", true},
		}

		t := template.Must(template.New("letter").Parse(letter))
		rand.Intn(len(recipients))
		err := t.Execute(writer, recipients[rand.Intn(len(recipients))])
		if err != nil {
			log.Println("executing template:", err)
		}
	})

	fmt.Println("开启web服务。。。。。")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
