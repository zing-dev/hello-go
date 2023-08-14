package strings_test

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func TestReplacer(t *testing.T) {
	var (
		str  = "hello world"
		buf  = bytes.NewBufferString("")
		name = "TEST-RUN\\\\n\\\\b.2023.02.22.1.2.3.1.6e1cdf.dll"
	)

	replacer := strings.NewReplacer("h", "", "l", "")
	log.Println(replacer.Replace(str))

	n, err := replacer.WriteString(buf, str)
	log.Println(n, err)
	log.Println(buf.String())
	log.Println(str)

	log.Println(strings.NewReplacer("/", ".", "\\", ".", "*", ".", "?", ".", ">", ".", "<", ".", "=", ".").Replace(name))
}
