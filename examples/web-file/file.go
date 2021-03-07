package main

import (
	"html/template"
	"net/http"
	"os"
	"strings"
	"web-file/resources"
)

type FileInfo struct {
	Name    string
	Path    string
	IsDir   bool
	Type    string
	Size    int64
	ModTime string
}

type Files []FileInfo

func FileHandle(writer http.ResponseWriter, request *http.Request) {
	path := "." + request.URL.Path
	info, err := os.Stat(path)
	if err != nil {
		return
	}
	if info.IsDir() {
		FileList(writer, path)
		return
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return
	}
	_, _ = writer.Write(content)
}

func FileList(writer http.ResponseWriter, path string) {
	dir, err := os.ReadDir(path)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(err.Error()))
		return
	}
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	files := make(Files, len(dir))
	for k, v := range dir {
		info, err := v.Info()
		if err != nil {
			continue
		}
		files[k] = FileInfo{
			Name:    v.Name(),
			Path:    path + v.Name(),
			IsDir:   v.IsDir(),
			Type:    v.Type().String(),
			Size:    info.Size(),
			ModTime: info.ModTime().Format("2006-01-02 15:04:05"),
		}
	}

	file, err := resources.ReadFile("index.html")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(err.Error()))
		return
	}
	parse, err := template.New("test").Parse(string(file))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(err.Error()))
		return
	}
	_ = parse.Execute(writer, files)
}
