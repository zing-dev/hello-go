package main

import (
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"web-file/resources"
)

type Config struct {
	Port string
	Dir  string
}

type FileInfo struct {
	Name    string
	Path    string
	IsDir   bool
	Type    string
	Size    int64
	ModTime string
}

type Files []FileInfo

var config Config

func FileUpload(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseMultipartForm(1 << 32)
	if err != nil {
		return
	}
	root := config.Dir
	if request.URL.Path != "/" {
		root += request.URL.Path
	}
	root += "/"
	for _, files := range request.MultipartForm.File {
		for _, file := range files {
			go func(file *multipart.FileHeader) {
				log.Println("new file upload: ", file.Filename)
				f, err := file.Open()
				if err != nil {
					log.Println("Err:Open ", err)
					return
				}
				defer f.Close()
				f2, err := os.OpenFile(root+file.Filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
				if err != nil {
					log.Println("Err:OpenFile ", err)
					return
				}
				defer f2.Close()
				_, err = io.Copy(f2, f)
				if err != nil {
					log.Println("Err:Copy ", err)
				}
			}(file)

		}
	}
	_, _ = writer.Write([]byte("success upload..."))
}

func FileHandle(writer http.ResponseWriter, request *http.Request) {
	rootPath := config.Dir + request.URL.Path
	info, err := os.Stat(rootPath)
	if err != nil {
		return
	}
	if info.IsDir() {
		FileList(writer, rootPath, request.URL.Path)
		return
	}
	content, err := os.ReadFile(rootPath)
	if err != nil {
		return
	}
	_, _ = writer.Write(content)
}

func FileList(writer http.ResponseWriter, rootPath, urlPath string) {
	dir, err := os.ReadDir(rootPath)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(err.Error()))
		return
	}
	if urlPath == "/" {
		urlPath = "./"
	} else {
		urlPath += "/"
	}
	files := make(Files, len(dir))
	for k, v := range dir {
		info, err := v.Info()
		if err != nil {
			continue
		}
		files[k] = FileInfo{
			Name:    v.Name(),
			Path:    urlPath + v.Name(),
			IsDir:   v.IsDir(),
			Type:    v.Type().String(),
			Size:    info.Size(),
			ModTime: info.ModTime().Format("2006-01-02 15:04:05"),
		}
	}

	if false {
		parse, err := template.ParseFiles("index.html")
		if err != nil {
			return
		}
		_ = parse.Execute(writer, files)
		return
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
