package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Host              string   `json:"host"`
	Port              int      `json:"port"`
	AnalyticsFile     string   `json:"analytics_file"`
	StaticFileVersion int      `json:"static_file_version"`
	StaticDir         string   `json:"static_dir"`
	TemplatesDir      string   `json:"templates_dir"`
	SerTcpSocketHost  string   `json:"serTcpSocketHost"`
	SerTcpSocketPort  int      `json:"serTcpSocketPort"`
	Fruits            []string `json:"fruits"`
}

type Other struct {
	SerTcpSocketHost string   `json:"serTcpSocketHost"`
	SerTcpSocketPort int      `json:"serTcpSocketPort"`
	Fruits           []string `json:"fruits"`
}

func main() {

	conf := Config{
		Host:              "http://localhost",
		Port:              8080,
		AnalyticsFile:     "",
		StaticFileVersion: 1,
		StaticDir:         "/c/Users/zhang/workspace/go/learn-go",
		TemplatesDir:      "/tmp",
		SerTcpSocketHost:  ":1212",
		SerTcpSocketPort:  8081,
		Fruits:            []string{"apple", "peach"},
	}
	jsonStr, _ := json.Marshal(conf)

	//{"host":"http://localhost","port":8080,"analytics_file":"","static_file_version":1,"static_dir":"/c/Users/zhang/workspace/go/learn-go","templates_dir":"/tmp","serTcpSocketHost":":1212","serTcpSocketPort":8081,"fruits":["apple","peach"]}
	fmt.Println(string(jsonStr))

	//json str 转 map
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &dat); err == nil {
		fmt.Println("============== json str 转 map =======================")
		fmt.Println(dat)
		fmt.Println(dat["host"])
		fmt.Println(dat["port"])
	}

	//json str 转struct
	var config Config
	if err := json.Unmarshal([]byte(jsonStr), &config); err == nil {
		fmt.Println("================ json str 转 struct ================")
		fmt.Println(config)
		fmt.Println(config.Host)
	}

	//json str 转struct(部份字段)
	var part Other
	if err := json.Unmarshal([]byte(jsonStr), &part); err == nil {
		fmt.Println("================json str 转 struct =================")
		fmt.Println(part)
		fmt.Println(part.SerTcpSocketPort)
	}

	//struct 到json str
	if b, err := json.Marshal(config); err == nil {
		fmt.Println("================struct 到json str==")
		fmt.Println(string(b))
	}

	//map 到json str
	fmt.Println("================map 到json str=====================")
	enc := json.NewEncoder(os.Stdout)
	_ = enc.Encode(dat)

	//array 到 json str
	arr := []string{"hello", "apple", "python", "golang", "base", "peach", "pear"}
	lang, err := json.Marshal(arr)
	if err == nil {
		fmt.Println("================array 到 json str==")
		fmt.Println(string(lang))
	}

	//json 到 []string
	var wo []string
	if err := json.Unmarshal(lang, &wo); err == nil {
		fmt.Println("================json 到 []string==")
		fmt.Println(wo)
	}
}
