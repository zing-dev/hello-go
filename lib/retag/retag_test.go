package retag

import (
	"encoding/json"
	"fmt"
	"github.com/sevlyar/retag"
	"testing"
)

func TestTag1(t *testing.T) {
	Tag1()
}

func TestTag2(t *testing.T) {
	Tag2()
}

func Example_viewOfData() {
	type UserProfile struct {
		Id          int64  `view:"-"`
		Name        string `view:"*"`
		CardNumber  string `view:"user"`
		SupportNote string `view:"support"`
	}
	profile := &UserProfile{
		Id:          7,
		Name:        "Duke Nukem",
		CardNumber:  "4378 0990 7823 1019",
		SupportNote: "Strange customer",
	}

	userView := retag.Convert(profile, retag.NewView("json", "user"))
	supportView := retag.Convert(profile, retag.NewView("json", "support"))

	// Now profile, userView and supportView point
	// on the same memory but have different types
	// with different tags.

	b, _ := json.MarshalIndent(userView, "", "  ")
	fmt.Println(string(b))
	b, _ = json.MarshalIndent(supportView, "", "  ")
	fmt.Println(string(b))
	// Output:
	// {
	//   "Name": "Duke Nukem",
	//   "CardNumber": "4378 0990 7823 1019"
	// }
	// {
	//   "Name": "Duke Nukem",
	//   "SupportNote": "Strange customer"
	// }
}
