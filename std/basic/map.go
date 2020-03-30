package basic

import "log"

type Map struct {
	Id   int
	Name string
}

func map1() {
	m := make(map[string]*Map)
	m["1"] = &Map{
		Id:   1,
		Name: "zing",
	}
	m["2"] = &Map{
		Id:   2,
		Name: "zing2",
	}

	for _, v := range m {
		v.Name += "-zing"
	}

	for _, v := range m {
		log.Println(v.Name)
	}

	m2 := make(map[string]Map)
	m2["1"] = Map{
		Id:   1,
		Name: "zing",
	}
	m2["2"] = Map{
		Id:   2,
		Name: "zing2",
	}

	for _, v := range m2 {
		v.Name += "-zing"
	}

	for _, v := range m2 {
		log.Println(v.Name)
	}
}
