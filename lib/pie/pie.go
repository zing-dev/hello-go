package pie

import (
	"fmt"
	"github.com/elliotchance/pie/pie"
	"strings"
)

func pie1() {
	var str pie.Strings = []string{"Bob", "Sally", "John", "Jane"}
	name := str.
		FilterNot(func(name string) bool {
			return strings.HasPrefix(name, "J")
		}).
		Map(strings.ToUpper).
		Last()

	fmt.Println(name) // "SALLY"
}

func pie2() {
	var str pie.Strings = strings.Split("H=NT;W=D1;ZoneX=02;ZoneY=35;ZoneZ=07;", ";")
	name := str.Filter(func(s string) bool {
		return s != ""
	}).Map(func(s string) string {
		return strings.Split(s, "=")[1]
	})
	fmt.Println(name) // "SALLY"
}
