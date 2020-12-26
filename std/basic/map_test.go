package basic

import (
	"log"
	"testing"
)

func TestMap1(t *testing.T) {
	map1()
}

func TestMap2(t *testing.T) {
	map2()
}

func TestMap3(t *testing.T) {
	map3()
}

func TestMap4(t *testing.T) {
	map4()
}

func TestMap5(t *testing.T) {
	for k, v := range map[int]string{
		1: "1",
		2: "2",
	} {
		log.Println(k, v)
	}

	for k, v := range map[int][]int{
		1: {1, 2, 3},
		2: {2, 4, 6},
	} {
		log.Println(k, v)
		for _, v := range v {
			log.Println(v)
		}
	}
}
func TestMap6(t *testing.T) {
	type Area struct {
		Id   int
		Name string
	}

	var cache = map[int][]Area{}
	log.Println(cache)
	log.Println(len(cache))
	log.Println(cache[0])      //[]
	log.Println(cache[1])      //[]
	log.Println(len(cache[2])) //0
	v, ok := cache[3]
	log.Println(v, ok) //[] false

	cache[0] = append(cache[0], Area{1, "111"})
	log.Println(cache)
	log.Println(len(cache))
	log.Println(cache[0]) //[]

	cache[0] = nil
	log.Println(cache)
	log.Println(len(cache))
	log.Println(cache[0]) //[]

}
