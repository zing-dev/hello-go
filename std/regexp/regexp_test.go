package regexp

import (
	"fmt"
	"mime"
	"regexp"
	"testing"
	"time"
)

func TestRegexp1(t *testing.T) {
	regexp1()
}

func TestRegexp2(t *testing.T) {
	regexp2()
}

// 简单的匹配使时间
func TestMatch(t *testing.T) {
	compile, err := regexp.Compile("^\\d+[mhs]$")
	if err != nil {
		t.Fatal(err)
	}
	for _, s := range []string{"1", "1s", "1ss", "1.1s", "90s", "1m", "1h", "1a", "1m1s", "1h1s"} {
		ok := compile.MatchString(s)
		str := compile.FindAllString(s, -1)
		duration, err := time.ParseDuration(s)
		t.Log(ok, s, str, duration, err)
	}

}

func TestMatch2(t *testing.T) {
	var (
		a1 = "attachment; filename=content.txt"
		a2 = "attachment; filename*=UTF-8''filename.txt"
		a3 = `attachment; filename="EURO rates"; filename*=utf-8''%e2%82%ac%20rates`
		a4 = `attachment; filename="omáèka.jpg"`
		//rule = `filename\*?=((['"])[\s\S]*?\2|[^;\n]*)`
	)
	mediatype, params, err := mime.ParseMediaType(a1)
	fmt.Println(params["filename"])
	mediatype, params, err = mime.ParseMediaType(a2)
	fmt.Println(params["filename"])
	fmt.Println(mediatype, params, err)
	mediatype, params, err = mime.ParseMediaType(a3)
	fmt.Println(params["filename"])
	mediatype, params, err = mime.ParseMediaType(a4)
	fmt.Println(params["filename"])
}
