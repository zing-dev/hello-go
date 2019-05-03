package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"math/rand"
	"strconv"
	"testing"
)

type User struct {
	gorm.Model

	UserName string
	Age      int
	Address  string
}

var (
	db  *gorm.DB
	err error
)

func openDB() {
	db, err = gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic(err)
	}
}

func closeDB() {

	err = db.Close()
	if err != nil {
		log.Println("close err:", err)
	}
}
func TestOpen(t *testing.T) {
	openDB()
	defer closeDB()
	db.DropTableIfExists(&User{})
	db.AutoMigrate(&User{})
}

func TestCreate(t *testing.T) {
	openDB()
	defer closeDB()
	for i := 0; i < 20; i++ {
		db.Create(&User{
			UserName: "zing_" + strconv.Itoa(i),
			Age:      20 + rand.Intn(6),
			Address:  "wuxi_" + strconv.Itoa(i),
		})
	}
}

func TestFirst(t *testing.T) {
	openDB()
	defer closeDB()

	user := &User{}
	db.First(user, 1)
	fmt.Println(user.UserName, user.Address)

	user = &User{}
	db.First(user, "id = 2")
	fmt.Println(user.UserName, user.Address)

	user = &User{}
	db.First(user, "id = ?", 3)
	fmt.Println(user.UserName, user.Address)

	user = &User{}
	db.FirstOrCreate(user, User{
		UserName: "zing_4",
	})
	fmt.Println(user.UserName, user.Address)

	user = &User{}
	db.FirstOrCreate(user, User{
		UserName: "zing_100",
	})
	fmt.Println(user.UserName, user.Address)

}

func TestSearch(t *testing.T) {
	openDB()
	defer closeDB()
	user := &User{}
	db.First(user)
	fmt.Println(user.ID, user.UserName, user.Address)

	user = &User{}
	db.Last(user)
	fmt.Println(user.ID, user.UserName, user.Address)

	users := make([]User, 0)
	//users := [...]User{}
	//users := make(map[int]User)
	db.Find(&users)
	for _, user := range users {
		fmt.Println(user.ID, user.UserName, user.Address)
	}
	fmt.Println(len(users), cap(users))
}

func TestWhere(t *testing.T) {
	openDB()
	defer closeDB()

	user := &User{}
	db.Where("id = 2").First(user)
	fmt.Println(user.ID, user.UserName, user.Address)
	user2 := User{}
	db.Where("id = 2").First(&user2)
	fmt.Println(user2.ID, user2.UserName, user2.Address)

	users := make([]User, 0)
	db.Where("id >?", 10).Where("id < 15").Find(&users)
	for _, user := range users {
		fmt.Println(user.ID, user.UserName, user.Address)
	}

	db.Where("id in (?)", []int{1, 2, 3}).Find(&users)
	for _, user := range users {
		fmt.Println(user.ID, user.UserName, user.Address)
	}

	db.Where(map[string]interface{}{"Age": 20}).Find(&users)
	for _, user := range users {
		fmt.Println(user.ID, user.UserName, user.Address)
	}

}
