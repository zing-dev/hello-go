package template

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		temp := template.New("layer")
		temp, err := temp.ParseFiles("index.html")
		if err != nil {
			fmt.Println("err :", err)
		} else {
			header := writer.Header()
			header.Set("Content-Type", "text/html;charset=utf-8")
			//_ = temp.Execute(writer, []string{"hello", "web"})
			_ = temp.ExecuteTemplate(writer, "layer", []string{"hello", "web"})
		}

	})

	http.HandleFunc("/a", func(writer http.ResponseWriter, request *http.Request) {

		temp := template.New("")
		temp.Funcs(map[string]interface{}{
			"double": func(str string) string {
				return strings.Repeat(str, 2)
			},
		})
		var text = `
<h3>文本</h3>
<h4>{{.hello | double}} </h2>

<h3>结构体</h3>
UserName: {{.user.UserName}} <br>
Address: {{.user.Address}} <br>

<h3>数组</h3>
{{range $i,$e := .array}}
	{{$i}} {{$e}}<br>
{{end}}
`
		temp = template.Must(temp.Parse(text))
		header := writer.Header()
		header.Set("Content-Type", "text/html;charset=utf-8")
		_ = temp.Execute(writer, map[string]interface{}{
			"hello": "hello world ",
			"user": struct {
				UserName string
				Address  string
			}{
				UserName: "zing",
				Address:  "江苏无锡",
			},
			"array": [...]int{10, 20, 30, 40, 50, 60},
		})
	})

	fmt.Println("开启web服务。。。。。")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
func TestNew(t *testing.T) {
	temp := template.New("layer")
	temp, e := temp.ParseFiles("index.html")
	if e == nil {
		_ = temp.Execute(os.Stdout, []string{"Hello", "World"})
	} else {
		fmt.Println(e)
	}
}
func Test(t *testing.T) {

	/*
	   1.声明一个Template对象并解析模板文本
	   func New(name string) *Template
	   func (t *Template) Parse(text string) (*Template, error)

	   2.从html文件解析模板
	   func ParseFiles(filenames ...string) (*Template, error)

	   3.模板生成器的包装
	   template.Must(*template.Template, error )会在Parse返回err不为nil时，调用panic。
	   func Must(t *Template, err error) *Template

	   t := template.Must(template.New("name").Parse("html"))
	*/

	temp, _ := template.ParseFiles("index.html")
	//t,_ := template.ParseGlob("*.tpl")

	_ = temp.Execute(os.Stdout, []string{"zing", "😄"})
}
