package lo

import (
	"fmt"
	"github.com/samber/lo"
	"testing"
)

type Data struct {
	Code int
	Msg  string
}

func TestUniqBy(t *testing.T) {
	var data = make([]Data, 10)
	for i := range data {
		if i%3 == 0 {
			data[i].Msg = fmt.Sprintf("%d", i)
			data[i].Code = 2
		} else {
			data[i].Msg = fmt.Sprintf("%d", i)
			data[i].Code = i
		}
	}

	fmt.Println(lo.UniqBy[Data, int](data, func(v Data) int { return v.Code }))
}
