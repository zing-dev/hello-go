package strings_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func parseHost(opts string) ([]string, error) {
	var res []string
	if opts == "" {
		return nil, errors.New("empty error")
	}

	var (
		hosts     = strings.Split(opts, ",")
		defaultIp = "0.0.0.0"
	)

	for _, host := range hosts {
		v := strings.Split(host, ":")
		if len(v) != 2 {
			return nil, errors.New("length error")
		}
		if v[0] == "" {
			v[0] = defaultIp
		} else {
			defaultIp = v[0]
		}
		res = append(res, fmt.Sprintf("%s:%s", v[0], v[1]))
	}
	return res, nil
}

func TestSplitParse(t *testing.T) {
	fmt.Println(parseHost("0.0.0.0:20000,:20001"))
	fmt.Println(parseHost("20000,:20001"))
}

func TestSplit(t *testing.T) {
	var a = []string{"0.0.0.0:20000,:20001", "0.0.0.0:22222"}
	for _, v := range a {
		r := strings.Split(v, ",")
		for _, s := range r {
			fmt.Println(len(r), s)
			v := strings.Split(s, ":")
			for _, s := range v {
				fmt.Println(len(v), s)
			}
		}
	}
}
