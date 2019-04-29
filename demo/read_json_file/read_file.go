package read_json_file

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type BiliBiliVideo struct {
	CoverWidth           float32 `json:"CoverWidth,float32"`
	CoverHeight          float32 `json:"CoverHeight,float32"`
	IsSelected           bool    `json:"IsSelected,bool"`
	IsLastHit            bool    `json:"IsLastHit,bool"`
	Aid                  string  `json:"Aid"`
	SeasonId             string  `json:"SeasonId"`
	EpisodeId            string  `json:"EpisodeId"`
	Title                string  `json:"Title"`
	Uploader             string  `json:"Uploader"`
	Mid                  string  `json:"Mid"`
	Description          string  `json:"Description"`
	CoverURL             string  `json:"CoverURL"`
	CoverImageSource     string  `json:"CoverImageSource"`
	Tag                  string  `json:"Tag"`
	DownloadTimeRelative int32   `json:"DownloadTimeRelative,int32"`
	DownloadTime         string  `json:"DownloadTime"`
	CreateDate           string  `json:"CreateDate"`
	TotalTime            string  `json:"TotalTime"`
	Parts                string  `json:"Parts"`
	FontSize6            float32 `json:"FontSize6,float32"`
	FontSize8            float32 `json:"FontSize8,float32"`
	FontSize10           float32 `json:"FontSize10,float32"`
	FontSize12           float32 `json:"FontSize12,float32"`
	FontSize14           float32 `json:"FontSize14,float32"`
	FontSize16           float32 `json:"FontSize16,float32"`
	FontSize18           float32 `json:"FontSize18,float32"`
	FontSize20           float32 `json:"FontSize20,float32"`
	FontSize22           float32 `json:"FontSize22,float32"`
	FontSize24           float32 `json:"FontSize24,float32"`
	FontSize30           float32 `json:"FontSize30,float32"`
	FontSize34           float32 `json:"FontSize34,float32"`
}

type BiliVideo struct {
	Title string `json:"Title"`
}

func main() {

	file, e := os.Open("./demo/41446451.json")

	if e != nil {
		log.Fatalf("open file error:%s", e)
		return
	}

	info, e := file.Stat()
	content := make([]byte, info.Size())
	_, e = file.Read(content)

	if e != nil {
		log.Fatalf("read file error:%s", e)
		return
	}

	fmt.Println(string(content))

	video := BiliBiliVideo{}

	e = json.Unmarshal(content, &video)

	if e != nil {
		log.Fatalf("unmarshal content error:%s", e)
		return
	}

	fmt.Println(video.Title)

	var video2 map[string]interface{}

	e = json.Unmarshal(content, &video2)

	if e != nil {
		log.Fatalf("unmarshal content2 error:%s", e)
		return
	}
	fmt.Println(video2["Title"])

	video3 := BiliVideo{}
	e = json.Unmarshal(content, &video3)

	if e != nil {
		log.Fatalf("unmarshal content3 error:%s", e)
		return
	}
	fmt.Println(video3.Title)

}
