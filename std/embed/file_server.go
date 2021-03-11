package main

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
)

//go:embed *.go *.mod *tmpl
var content embed.FS

func testFileServer() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.FS(content))))
	http.HandleFunc("/html", func(writer http.ResponseWriter, request *http.Request) {
		t, err := template.ParseFS(content, "*.tmpl")
		if err != nil {
			_, _ = writer.Write([]byte(err.Error()))
			return
		}
		goFiles, err := fs.Glob(content, "*.go")
		if err != nil {
			_, _ = writer.Write([]byte(err.Error()))
			return
		}
		_ = t.Execute(writer, goFiles)
	})
	_ = http.ListenAndServe(":8088", nil)
}
