package archive

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

func TestZipNewReader(t *testing.T) {
	data, err := os.ReadFile("test.zip")
	if err != nil {
		t.Fatal(err)
	}

	z, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range z.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.CopyN(os.Stdout, rc, 10)
		if err != nil {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println()
	}
}

func TestUnZip(t *testing.T) {
	r, err := zip.OpenReader("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.CopyN(os.Stdout, rc, 10)
		if err != nil {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println()
	}
}

func TestZip(t *testing.T) {
	var (
		buf   = new(bytes.Buffer)
		w     = zip.NewWriter(buf)
		files = []struct {
			Name, Body string
		}{
			{"readme.txt", "This archive contains some text files."},
			{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
			{"todo.txt", "Get animal handling licence.\nWrite more examples."},
		}
	)
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	_ = w.SetComment("hello world")
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("test.zip", buf.Bytes(), 0777)

	if err != nil {
		log.Fatal(err)
	}
}
