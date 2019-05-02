package test

import (
	"fmt"
	"math"
	"testing"
)

func TestTest(t *testing.T) {

	t.Log("hello world")
	t.Error("Error")
	t.Fatal("Fatal")

}

func TestT2(t *testing.T) {

	fmt.Println("hello world")
}

func TestAbs(t *testing.T) {
	got := math.Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %f; want 1", got)
	}
}
