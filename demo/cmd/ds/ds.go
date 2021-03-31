// Show the specified directory structure
package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

const (
	INDENT = "  "
)

var (
	rootPath string
)

func init() {
	flag.StringVar(&rootPath, "p", "", "The path of target directory.")
}

func showFiles(basePath string, prefix string, showAll bool) error {
	base, err := os.Open(basePath)
	if err != nil {
		return err
	}
	subs, err := base.Readdir(-1)
	if err != nil {
		return err
	}
	for _, v := range subs {
		fi := v.(os.FileInfo)
		fp := fi.Name()
		if strings.HasPrefix(fp, ".") && !showAll {
			continue
		}
		if fi.IsDir() {
			absFp := path.Join(basePath, fp)
			fmt.Printf("%s/\n", prefix+fp)
			err = showFiles(absFp, INDENT+prefix, showAll)
			if err != nil {
				return err
			}
		} else {
			fmt.Printf("%s\n", prefix+fp)
		}
	}
	return nil
}

func main() {
	flag.Parse()
	if len(rootPath) == 0 {
		defaultPath, err := os.Getwd()
		if err != nil {
			fmt.Println("GetwdError:", err)
			return
		}
		rootPath = defaultPath
	}
	fmt.Printf("%s:\n", rootPath)
	err := showFiles(rootPath, INDENT, false)
	if err != nil {
		fmt.Println("showFilesError:", err)
	}
}
