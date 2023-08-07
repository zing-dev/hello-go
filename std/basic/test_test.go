//go:build windows
// +build windows

package basic

import (
	"encoding/json"
	"fmt"
	"golang.org/x/sys/windows"
	"log"
	"strconv"
	"testing"
)

func TestName3(t *testing.T) {
	fmt.Println(fmt.Sprintf("%d", 1))
	fmt.Println(fmt.Sprintf("%2d", 1))
	fmt.Println(fmt.Sprintf("%02d", 1))
	fmt.Println(fmt.Sprintf("%02d", 11))
	fmt.Println(fmt.Sprintf("%02d", 111))
	fmt.Println(fmt.Sprintf("%010d", 111))
	fmt.Println(fmt.Sprintf("%02.0f", 1.1))
	fmt.Println(string(rune(1)))
	fmt.Println(strconv.Itoa(int(byte(1))))
}
func TestName(t *testing.T) {

	// 返回UI显示语言的简码数组
	languages, err := windows.GetUserPreferredUILanguages(windows.MUI_LANGUAGE_ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(languages)
	// 返回UI显示语言的名称数组
	uiLanguages, err := windows.GetUserPreferredUILanguages(windows.MUI_LANGUAGE_NAME)
	if err != nil {
		log.Fatal(err)
	}
	//2020/08/04 16:58:31 [zh-CN en-US]
	log.Println(uiLanguages)
}

func TestName2(t *testing.T) {
	str := "{\"TempID\":\"CCBD921B8D304CC68A436BC6F6270800\",\"LineID\":\"1\",\"ChannelID\":\"2\",\"CollectTime\":\"\\/Date(1596590473239+0800)\\/\",\"TempData\":\"6.988,2.191,85.847,62.891,94.341,50.39,94.445,77.678,16.3,73.39,26.675\"}"
	aa := map[string]string{}
	json.Unmarshal([]byte(str), &aa)
	log.Println(aa)

	log.Println(AddSlashes(str))

}

func AddSlashes(str string) string {
	var tmpRune []rune
	strRune := []rune(str)
	for _, ch := range strRune {
		switch ch {
		case []rune{'\\'}[0], []rune{'"'}[0], []rune{'\''}[0]:
			tmpRune = append(tmpRune, []rune{'\\'}[0])
			tmpRune = append(tmpRune, ch)
		default:
			tmpRune = append(tmpRune, ch)
		}
	}
	return string(tmpRune)
}

func TestName1(t *testing.T) {
	fmt.Println(fmt.Sprintf("%v", []byte{1, 2, 3}))
	fmt.Println(fmt.Sprintf("%d", []byte{1, 2, 3, 12}))
	fmt.Println(fmt.Sprintf("%x", []byte{1, 2, 3, 12, 255}))
}

func TestFor(t *testing.T) {
	for range [3]interface{}{} {
		t.Log("")
	}
}
func TestType(t *testing.T) {
	type A string
	type B = string
	for _, a := range []any{A("A"), "a"} {
		switch a.(type) {
		case string:
			fmt.Println(a, "string")
		case A:
			fmt.Println(a, "a")
			//case B:
		}
	}
}
