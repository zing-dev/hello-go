package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {

	// Fprint
	_, _ = fmt.Fprint(os.Stdout, "hello\n")
	_, _ = fmt.Fprint(os.Stderr, "hello\n")
	_, _ = fmt.Fprint(os.Stdin, "hello\n")

	buffer := bytes.NewBuffer([]byte("hello"))
	//追加，返回添加的长度
	n, _ := fmt.Fprint(buffer, " world")
	fmt.Println(n)               //6
	fmt.Println(buffer.String()) //hello world

	bufferString := bytes.NewBufferString("hello")
	_, _ = fmt.Fprint(bufferString, " golang", " world")
	fmt.Println(bufferString) //hello golang world

	//Fprintf
	_, _ = fmt.Fprintf(bufferString, "\n%s is the best", "golang")
	fmt.Println(bufferString)

	_, _ = fmt.Fprintf(os.Stdout, "%s\n", "golang") //golang

	fmt.Println(fmt.Errorf("%s.Errorf", "fmt")) //fmt.Errorf

	//1 1.2 100 97 a 0xc00007a020
	fmt.Printf("%v %v %v %v %v %v\n", 1, 1.2, 1e2, 'a', "a", &bufferString)

	//1 1.2 100 97 a 0xc00007a020
	fmt.Printf("%#v %#v %#v %#v\n", 1, 2.3, 'a', "aa")

	//int float64 int32 string **bytes.Buffer []int32
	fmt.Printf("%T %T %T %T %T %T\n", 1, 2.3, 'a', "aa", &bufferString, []rune(""))

	//true false
	fmt.Printf("%t %t\n", true, false)

	//1010 1100001
	fmt.Printf("%b %b\n", 10, 'a')

	//a d
	fmt.Printf("%c %c %c\n", 'a', 100, '😄')
	str := []byte("zing")
	//&[z i n g] i
	fmt.Printf("%c %c\n", &str, str[1])

	//1 97 127 9223372036854775807
	fmt.Printf("%d %d %d %d\n", 1, 'a', math.MaxInt8, math.MaxInt64)

	//10 60 157
	fmt.Printf("%o %o %o\n", 8, '0', 'o')

	//'a' 'd'
	fmt.Printf("%q %q\n", 'a', 100)

	//61 64 7f
	fmt.Printf("%x %x %x\n", 'a', 100, 127)

	//61 64 7F 1F604
	fmt.Printf("%X %X %X %X\n", 'a', 100, 127, '😄')

	//U+0061 U+5F20 U+007F U+1F604
	fmt.Printf("%U %U %U %U\n", 'a', '张', 127, '😄')

	//5404319552844595p-52 4503599627370496p-52
	fmt.Printf("%b %b\n", 1.2, 1.0)

	//1.200000e+00 1.000000e+00 1.797693e+308
	fmt.Printf("%e %e %e\n", 1.2, 1.0, math.MaxFloat64)

	//1.200000E+00 1.000000E+00 1.797693E+308
	fmt.Printf("%E %E %E\n", 1.2, 1.0, math.MaxFloat64)

	//1.200000 1.000000 179769313486231570814527423731704356798070567525844996598917476803157260780028538760589558632766878171540458953514382464234321326889464182768467546703537516986049910576551282076245490090389328944075868508455133942304583236903222948165808559332123348274797826204144723168738177180919299881250404026184124858368.000000
	fmt.Printf("%f %f %f\n", 1.2, 1.0, math.MaxFloat64)

	//1.2 1e+06 1.7976931348623157e+308
	fmt.Printf("%g %g %g\n", 1.2, 1000000.0, math.MaxFloat64)

	//1.2 1E+06 1.7976931348623157E+308
	fmt.Printf("%G %G %G\n", 1.2, 1000000.0, math.MaxFloat64)

	//zing 😄
	fmt.Printf("%s %s\n", "zing", "😄")

	//"zing" "😄"
	fmt.Printf("%q %q\n", "zing", "😄")

	//7a696e67 f09f9884
	fmt.Printf("%x %x\n", "zing", "😄")

	//7A696E67 F09F9884
	fmt.Printf("%X %X\n", "zing", "😄")

	//0xc0000540e8 0xc0000540ec  地址
	fmt.Printf("%p %p\n", []byte("zing"), make([]byte, 1))

	fmt.Printf("%f\n", math.Pi)     //3.141593
	fmt.Printf("%2.10f\n", math.Pi) // 3.1415926536
	fmt.Printf("%9f\n", math.Pi)    // 3.141593
	fmt.Printf("%9.5f\n", math.Pi)  //  3.14159
	fmt.Printf("%9.f\n", math.Pi)   // 3

	//+100 -100
	fmt.Printf("%+d %+d\n", 100, -100)
	fmt.Printf("%-d %-d\n", 100, -100)

	var i interface{} = 23
	fmt.Printf("%v\n", i) //23

}
