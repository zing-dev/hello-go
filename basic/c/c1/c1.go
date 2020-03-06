package c1

/*
#include <stdio.h>

static void SayHello(const char* s) {
    puts(s);
}
*/
import "C"

func c1() {
	C.puts(C.CString("你好, GopherChina 2020!\n"))
}

func c2() {
	C.SayHello(C.CString("Hello, World\n"))
}
