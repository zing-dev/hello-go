package basic

import "testing"

func TestS1(t *testing.T) {
	S1()
}

func TestS2(t *testing.T) {
	S2()
}

func TestS3(t *testing.T) {
	t.Log(people)

	var p2 *People
	t.Log(p2)
}

func TestStructClone(t *testing.T) {
	p := &People{Name: "zing"}
	println(p.Name)
	p2 := *p
	p2.Name = "zing2"
	println(p.Name)
	println(p2.Name)
}
