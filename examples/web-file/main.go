package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

//go:generate fileb0x b0x.yml

type config struct {
	port string
	dir  string
}
type file struct {
	Name    string
	IsDir   bool
	Type    string
	Size    int64
	ModTime string
}

func main() {
	flag.Parse()
	c := config{
		port: flag.Arg(0),
		dir:  flag.Arg(1),
	}
	if c.port == "" {
		c.port = "8081"
	}
	if c.dir == "" {
		c.dir = "./"
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			if request.URL.Path != "/" {
				/*if path.Base(request.URL.Path) == "index.html" {
					data, err := resources.ReadFile("index.html")
					if err != nil {
						return
					}
					_, _ = writer.Write(data)
					return
				}
				writer.WriteHeader(http.StatusInternalServerError)
				f, err := os.ReadFile(path.Base(request.URL.Path))
				if err != nil {
					writer.WriteHeader(http.StatusInternalServerError)
					_, _ = writer.Write([]byte(err.Error()))
					return
				}
				_, _ = writer.Write(f)
				return*/
			}

			dir, err := os.ReadDir(c.dir)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				_, _ = writer.Write([]byte(err.Error()))
				return
			}

			files := make([]file, len(dir))
			for k, v := range dir {
				info, err := v.Info()
				if err != nil {
					continue
				}
				files[k] = file{
					Name:    v.Name(),
					IsDir:   v.IsDir(),
					Type:    v.Type().String(),
					Size:    info.Size(),
					ModTime: info.ModTime().String(),
				}
			}

			t, err := template.ParseFiles("index.html")
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				_, _ = writer.Write([]byte(err.Error()))
				return
			}
			_ = t.Execute(writer, files)
		case http.MethodPost:
			_, _ = writer.Write([]byte("post"))
		default:
		}
	})
	fmt.Println(fmt.Sprintf("server start from :%s...", c.port))
	err := http.ListenAndServe(fmt.Sprintf(":%s", c.port), nil)
	fmt.Println("server end: ", err)

}
