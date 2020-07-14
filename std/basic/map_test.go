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
