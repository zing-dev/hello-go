package retag

import (
	"encoding/json"
	"fmt"
	"github.com/sevlyar/retag"
)

type Language struct {
	Name string `view:"name"`
	Age  int    `view:"age"`
}

type UserProfile struct {
	Id          int64  `view:"-"`
	Name        string `view:"*"`
	CardNumber  string `view:"user"`
	SupportNote string `view:"support"`
}

func Tag1() {
	l := &Language{
		Name: "golang",
		Age:  12,
	}
	userView := retag.Convert(l, retag.NewView("json", "name,age"))
	fmt.Println(userView)
	b, _ := json.MarshalIndent(userView, "", "  ")
	fmt.Println(string(b))
	b, _ = json.MarshalIndent(l, "", "  ")
	fmt.Println(string(b))
}

func Tag2() {
	profile := &UserProfile{
		Id:          7,
		Name:        "Duke Nukem",
		CardNumber:  "4378 0990 7823 1019",
		SupportNote: "Strange customer",
	}

	userView := retag.Convert(profile, retag.NewView("json", "user"))
	supportView := retag.Convert(profile, retag.NewView("json", "support"))

	b, _ := json.MarshalIndent(profile, "", "  ")
	fmt.Println(string(b))
	b, _ = json.MarshalIndent(userView, "", "  ")
	fmt.Println(string(b))
	b, _ = json.MarshalIndent(supportView, "", "  ")
	fmt.Println(string(b))
}
