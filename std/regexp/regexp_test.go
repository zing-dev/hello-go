package regexp

import (
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

//简单的匹配使时间
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
