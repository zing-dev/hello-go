//
// Created by zhangrongxiang on 2017/9/30 11:40
// File interface2
//
package main

import "fmt"

type People interface {
	eat()
	sleep()
}

type Student struct {
	id   int
	name string
	age  int
}

func (student Student) eat() Student {
	fmt.Printf("%s eating now....", student.name)
	return student
}
func (student Student) sleep() Student {
	fmt.Printf("%s sleep now....", student.name)
	return student
}

func main() {
	var s Student
	s.name = "zhangrongxiang"
	s.eat()
}