// Code generaTed by fileb0x at "2021-03-05 19:26:46.744648907 +0800 CST m=+0.008115148" from config file "b0x.yml" DO NOT EDIT.
// modified(2021-03-05 15:43:06.940002096 +0800 CST)
// original path: ../../README.md

package resources

import (
	"os"
)

// FilePublicREADMEMd is "public/README.md"
var FilePublicREADMEMd = []byte("\x23\x20\x6c\x65\x61\x72\x6e\x2d\x67\x6f\x0a\x0a\x3e\x20\x54\x68\x69\x73\x20\x69\x73\x20\x61\x6e\x20\x6f\x70\x65\x6e\x20\x73\x6f\x75\x72\x63\x65\x20\x70\x72\x6f\x6a\x65\x63\x74\x20\x74\x6f\x20\x6c\x65\x61\x72\x6e\x20\x67\x6f\x6c\x61\x6e\x67\x0a\x0a\x60\x60\x60\x67\x6f\x0a\x0a\x70\x61\x63\x6b\x61\x67\x65\x20\x6d\x61\x69\x6e\x0a\x0a\x69\x6d\x70\x6f\x72\x74\x20\x22\x66\x6d\x74\x22\x0a\x0a\x66\x75\x6e\x63\x20\x6d\x61\x69\x6e\x28\x29\x20\x7b\x0a\x09\x66\x6d\x74\x2e\x50\x72\x69\x6e\x74\x6c\x6e\x28\x22\x48\x65\x6c\x6c\x6f\x2c\x20\x57\x6f\x72\x6c\x64\x21\x22\x29\x0a\x7d\x0a\x0a\x0a\x60\x60\x60")

func init() {

	f, err := FS.OpenFile(CTX, "public/README.md", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(FilePublicREADMEMd)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
