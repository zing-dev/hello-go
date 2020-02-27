package io_test

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestMkdir(t *testing.T) {

	err := os.Mkdir("test", 0777)

	if err != nil {
		log.Fatal("Mkdir", err)
	}

	log.Println("Mkdir Success")
}

func TestMkdirAll(t *testing.T) {
	err := os.MkdirAll(time.Now().Format("2006/01/02")+"/test", 0777)

	if err != nil {
		log.Fatal("MkdirAll", err)
	}

	err = os.MkdirAll(time.Now().Format("2006/01/02")+"/test2", os.ModePerm)

	if err != nil {
		log.Fatal("MkdirAll", err)
	}

	log.Println("MkdirAll Success")

}

func TestOpenFile(t *testing.T) {

	dst, err := os.OpenFile("dst.txt", os.O_RDWR|os.O_CREATE, os.FileMode(0777))

	if err != nil {
		log.Fatalln("OpenFile", err)
	}

	src, err := os.OpenFile("ioutil_test.go", os.O_RDONLY, os.FileMode(0666))

	if err != nil {
		log.Fatalln("OpenFile", err)
	}
	i, err := io.Copy(dst, src)

	if err != nil {
		log.Fatalln("Copy", err)
	}
	log.Println(i)
}

func TestName(t *testing.T) {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		reader, err := request.MultipartReader()
		writer.Header().Set("Content-Type", "application/json")
		if reader == nil {
			_, _ = writer.Write([]byte("reader nil"))
			if err != nil {
				log.Fatalln("")
			}
		} else {
			form, err := reader.ReadForm(20 * 1024 * 2014)
			if err != nil {
				log.Fatalln("ReadForm", err)
			}

			headers := form.File

			for _, v := range headers {
				for _, v1 := range v {
					file, err := v1.Open()
					if err != nil {
						log.Fatalln("Open", err)
					}
					dst, err := os.OpenFile(time.Now().Format("2006-01-02")+"-"+v1.Filename, os.O_CREATE|os.O_RDWR, os.FileMode(0777))
					if err != nil {
						log.Fatalln("OpenFile", err)
					}
					written, err := io.Copy(dst, file)

					if err != nil {
						data, _ := json.Marshal(map[string]interface{}{
							"status":  true,
							"message": "upload error",
						})
						_, _ = writer.Write(data)
						log.Fatalln("Copy", err)
					}

					log.Println("written", written)
				}
			}
			data, _ := json.Marshal(map[string]interface{}{
				"status":  true,
				"message": "upload success",
			})
			_, _ = writer.Write(data)

		}
	})
	err := http.ListenAndServe(":8082", nil)

	if err != nil {
		log.Println("ListenAndServe", err)
	}

}
