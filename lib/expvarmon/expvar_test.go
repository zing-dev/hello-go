package expvarmon

import (
	_ "expvar"
	"fmt"
	"net/http"
	"testing"
)

func run() {
	i := 1
	fmt.Println(i)
	_ = http.ListenAndServe(":1234", nil)
}

func TestRun(t *testing.T) {
	run()
}
