package main

import (
	"bytes"
	"fmt"
)

func main() {

	///////////////////////////////////////
	a := [10]byte{1, 2, 3, 4, 5, 6}

	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	//////////////////////////////////

	b := make([]byte, 5, 10)

	b[0] = 1
	b[1] = 2

	b = append(b, 4)

	fmt.Println(b)
	fmt.Println(len(b)) //6
	fmt.Println(cap(b)) //10

	count := bytes.Count(b, b)
	fmt.Println(count) // 1

	compare := bytes.Compare(b, make([]byte, 5))
	fmt.Println(compare)

	c := make([]byte, 5)
	c[0] = 1
	c[1] = 2
	fmt.Println(bytes.Contains(b, c)) //true

	/////////////////////////////////////////////////
	d := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(d)
	fmt.Println(len(d)) // 9
	fmt.Println(cap(d)) // 9

	d = append(d, 10, 11)
	fmt.Println(d)      //[1 2 3 4 5 6 7 8 9 10 11]
	fmt.Println(len(d)) // 11
	fmt.Println(cap(d)) //32

	fmt.Println(bytes.HasPrefix(d, []byte{1, 2}))   //true
	fmt.Println(bytes.HasSuffix(d, []byte{10, 11})) //true
	fmt.Println(bytes.HasSuffix(d, []byte{11}))     //true
	fmt.Println(bytes.HasSuffix(d, []byte{10}))     //false

	fmt.Println(bytes.Map(func(r rune) rune {
		return r + 1
	}, d))
	//[2 3 4 5 6 7 8 9 10 11 12]

	fmt.Println(bytes.Repeat(d, 2)) //[1 2 3 4 5 6 7 8 9 10 11 1 2 3 4 5 6 7 8 9 10 11]

	e := []byte{'a', 'b', 'c', 'd', 'e'} //[97 98 99 100 101]
	fmt.Println(e)
	fmt.Println(bytes.ContainsRune(e, 'a')) //true

	fmt.Println(bytes.ToLower(e))
	fmt.Printf("%s\n", bytes.ToLower(e)) // abcde
	fmt.Printf("%s\n", bytes.ToUpper(e)) // ABCDE
	fmt.Printf("%s\n", bytes.Title(e))   // Abcde

	fmt.Printf("%c\n", 0x7F)
	fmt.Println(bytes.Trim(e, "a")) //[98 99 100 101]

}
