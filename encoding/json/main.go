package main

import (
	"encoding/json"
	"fmt"
)

// Product 商品信息
type Product struct {
	Name      string
	ProductID int64
	Number    int
	Price     float64
	IsOnSale  bool
}

type Product2 struct {
	Name      string  `json:"name"`
	ProductID int64   `json:"product_id"`
	Number    int     `json:"number"`
	Price     float64 `json:"price"`
	IsOnSale  bool    `json:"is_on_sale"`
}

type Product3 struct {
	Name      string  `json:"name"`
	ProductID int64   `json:"product_id,string"`
	Number    int     `json:"number,string"`
	Price     float64 `json:"price,string"`
	IsOnSale  bool    `json:"is_on_sale,string"`
}

func main() {
	p := &Product{}
	p.Name = "iPhone"
	p.IsOnSale = true
	p.Number = 10000
	p.Price = 2499.00
	p.ProductID = 1
	data, _ := json.Marshal(p)

	//{"Name":"iPhone","ProductID":1,"Number":10000,"Price":2499,"IsOnSale":true}
	fmt.Println(string(data))

	p2 := &Product2{}
	p2.Name = "iPhone"
	p2.IsOnSale = true
	p2.Number = 10000
	p2.Price = 2499.00
	p2.ProductID = 1
	data, _ = json.Marshal(p2)

	//{"name":"iPhone","product_id":1,"number":10000,"price":2499,"is_on_sale":true}
	fmt.Println(string(data))

	p3 := &Product3{}
	p3.Name = "iPhone"
	p3.IsOnSale = true
	p3.Number = 10000
	p3.Price = 2499.00
	p3.ProductID = 1
	data, _ = json.Marshal(p3)

	//{"name":"iPhone","product_id":"1","number":"10000","price":"2499","is_on_sale":"true"}
	fmt.Println(string(data))
	err := json.Unmarshal([]byte(data), p3)

	if err != nil {
		fmt.Println("解析出错")
	} else {
		//{iPhone 1 10000 2499 true}
		fmt.Println(*p, p.Name)
	}

}
