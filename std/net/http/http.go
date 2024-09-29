package http

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var temp = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
  <style type="text/css">
      body,html,#container{
        width: 100%;
        height: 100%;
        margin: 0px
      }
    </style>
<body>
   <div id="container" tabindex="0"></div>
<script src="https://webapi.amap.com/maps?v=1.4.15&key=2051ed061c1ed9d07474cf2296599207"></script>
<script type="text/javascript">
    var map = new AMap.Map('container', {
        center:[117.000923,36.675807],
        zoom:11
    });
</script>
</body>
</html>`

func http1() {
	var param string
	http.HandleFunc("/index", func(writer http.ResponseWriter, request *http.Request) {
		file, err := os.Open("./www/index.html")
		if err != nil {
			_, _ = io.WriteString(writer, err.Error())
			return
		}
		data, err := ioutil.ReadAll(file)
		if err != nil {
			_, _ = io.WriteString(writer, err.Error())
			return
		}
		_, _ = io.WriteString(writer, string(data))
	})

	http.HandleFunc("/upload", func(writer http.ResponseWriter, request *http.Request) {
		files := request.MultipartForm.File
		for _, file := range files {
			for _, f := range file {
				f2, err := os.Create(f.Filename)
				if err != nil {
					_, _ = io.WriteString(writer, err.Error())
					return
				}
				f2.Close()
			}
		}
	})
	http.HandleFunc("/json", func(writer http.ResponseWriter, request *http.Request) {
		file, err := os.Open("./www/data.json")
		if err != nil {
			_, _ = io.WriteString(writer, err.Error())
			return
		}
		data, err := ioutil.ReadAll(file)
		if err != nil {
			_, _ = io.WriteString(writer, err.Error())
			return
		}
		_, _ = io.WriteString(writer, string(data))
	})

	http.HandleFunc("/index2", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte(temp))
	})

	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		data, err := ioutil.ReadAll(request.Body)
		if request.Method == "GET" {
			if err != nil {
				_, _ = writer.Write([]byte(err.Error()))
				return
			}
			log.Println(request.RemoteAddr)
			_, _ = writer.Write([]byte("from " + request.RemoteAddr + " success\n"))
			body := request.URL.Query().Get("data")
			param = body
			file, err := os.Create("./www/data.json")
			if err != nil {
				_, _ = writer.Write([]byte(err.Error()))
				return
			}
			_, err = file.WriteString(param)
			if err != nil {
				_, _ = writer.Write([]byte(err.Error()))
				return
			}
			_, _ = writer.Write([]byte(body))
		} else if request.Method == "POST" {
			_, _ = writer.Write(data)
		}
	})

	http.HandleFunc("/param", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte(param))
	})
	http.Handle("/event", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		flusher, ok := w.(http.Flusher)
		if !ok {
			log.Panic("server not support")
		}
		for i := 0; i < 10; i++ {
			fmt.Fprintf(w, "data: ===> %d\n\n", i)
			flusher.Flush()
			time.Sleep(3 * time.Second)
		}
		fmt.Fprintf(w, "event: close\ndata: close\n\n") // 一定要带上data，否则无效
	}))

	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Fatal(err)
	}
}
