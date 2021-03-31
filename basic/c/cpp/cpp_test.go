package cpp

import "testing"

//g++ -c foo.cpp
//g++ -c bridge.c
//ar -crs libfoo.a foo.o bridge.o
func TestRun(t *testing.T) {
	Run()
}
