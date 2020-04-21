package tree

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Node struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Children []*Node `json:"children"`
}

type Field struct {
	Name   string `json:"name"`
	Alias  string `json:"alias"`
	Fields []*Field
}

func (f *Field) String() string {
	data, err := json.MarshalIndent(f, "", "	")
	if err != nil {
		return ""
	}
	return string(data)
}

func (f *Field) AddChildren(children ...*Field) *Field {
	if f.Fields == nil {
		f.Fields = children
	} else {
		f.Fields = append(f.Fields, children...)
	}
	return f
}

func (f *Field) AddChild(child *Field) *Field {
	if f.Fields == nil {
		f.Fields = make([]*Field, 1)
		f.Fields[0] = child
	} else {
		f.Fields = append(f.Fields, child)
	}
	return child
}

func (f *Field) fieldsName() string {
	names := ""
	if f.Fields != nil {
		for _, v := range f.Fields {
			names += v.Name + ","
			names += v.fieldsName()
		}
	}
	return names
}

func fieldsName(f *Field) string {
	names := ""
	if f.Fields != nil {
		for _, v := range f.Fields {
			names += v.Name + ","
			names += fieldsName(v)
		}
	}
	return names
}

func pp(n *Node) {
	fmt.Println(n.Name)
	if n.Children != nil {
		for _, v := range n.Children {
			pp(v)
		}
	}
}

func tree1() {
	n := &Node{
		Id:   1,
		Name: "Root",
		Children: []*Node{
			{
				Id:   2,
				Name: "2",
				Children: []*Node{
					{
						Id:       4,
						Name:     "4",
						Children: nil,
					},
					{
						Id:       5,
						Name:     "5",
						Children: nil,
					},
				},
			},
			{
				Id:       3,
				Name:     "3",
				Children: nil,
			},
		},
	}

	data, _ := json.MarshalIndent(n, "", "	")
	fmt.Println(string(data))

	pp(n)
}

func tree2() {
	f := &Field{
		Name: "root",
	}
	f.AddChild(&Field{
		Name: "2",
	}).AddChild(&Field{
		Name: "3",
	})

	f.AddChildren(&Field{
		Name: "4",
	}, &Field{
		Name: "5",
	})
	fmt.Println(f.String())

	fmt.Println(fieldsName(f))
	fmt.Println(strings.Split(fieldsName(f), ","))
	fmt.Println(strings.Split(f.fieldsName(), ","))

	f2 := Field{}
	_ = json.Unmarshal([]byte(f.String()), &f2)
	fmt.Println(f2.String())
}
