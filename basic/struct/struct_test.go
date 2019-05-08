package struct_test

import (
	"fmt"
	"testing"
)

func TestStruct(t *testing.T) {
	type Person struct {
		name string
		age  int
	}

	_ = Person{"zing", 25}
	_ = Person{name: "zing", age: 25}
	_ = new(Person)

	var p1 Person
	p1.name, p1.age = "zing", 25
	fmt.Println(p1)
	p2 := Person{"zing", 25}
	fmt.Println(p2)

	fmt.Println(Person{})
	fmt.Println(Person{p1.name, p1.age})
}

func TestAnonymous(t *testing.T) {
	type Skills []string

	type Human struct {
		name   string
		age    int
		weight int
	}

	type Student struct {
		Human      // 匿名字段，那么默认Student就包含了Human的所有字段
		speciality string
	}

	type Student2 struct {
		Skills
		Human // 匿名字段，那么默认Student就包含了Human的所有字段
		int
		speciality string
	}

	mark := Student{Human{
		"Mark", 25, 120},
		"Computer Science",
	}

	// 我们访问相应的字段
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)
	// 修改对应的备注信息
	mark.speciality = "AI"
	fmt.Println("Mark changed his speciality")
	fmt.Println("His speciality is ", mark.speciality)
	// 修改他的年龄信息
	fmt.Println("Mark become old")
	mark.age = 46
	fmt.Println("His age is", mark.age)
	// 修改他的体重信息
	fmt.Println("Mark is not an athlet anymore")
	mark.weight += 60
	fmt.Println("His weight is", mark.weight)

	jane := Student2{Human: Human{"Jane", 35, 100},
		speciality: "Biology", int: 1, Skills: []string{"abc", "😄"}}
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her speciality is ", jane.speciality)
	fmt.Println("Her Skills is ", jane.Skills)
	// 我们来修改他的skill技能字段
	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ", jane.Skills)
	fmt.Println("She acquired two new ones ")
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills now are ", jane.Skills)
	// 修改匿名内置类型字段
	jane.int = 3
	fmt.Println("Her preferred number is", jane.int)

	type Human2 struct {
		name  string
		age   int
		phone string // Human类型拥有的字段
	}

	type Employee struct {
		Human2     // 匿名字段Human
		speciality string
		phone      string // 雇员的phone字段
	}

	Bob := Employee{Human2{"Bob", 34, "777-444-XXXX"}, "Designer", "333-222"}
	fmt.Println("Bob's work phone is:", Bob.phone)
	// 如果我们要访问Human的phone字段
	fmt.Println("Bob's personal phone is:", Bob.Human2.phone)
}
