package dts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

var (
	client = http.Client{}
	host   = "http://192.168.0.215:8084"
)

func get(url string) {
	response, err := client.Get(url)

	if err != nil {
		log.Fatal("get error")
	}

	if response.StatusCode == http.StatusOK {
		all, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatal("ReadAll error")
		}

		data := make(map[string]interface{})

		if all != nil {
			err = json.Unmarshal(all, &data)

			if err != nil {
				log.Println("Unmarshal error")
				log.Fatal(string(all))
			}

			success := data["Success"].(bool)
			if success {
				fmt.Println(data)
			} else {
				log.Fatalln("success is false")
			}

		} else {
			fmt.Println("data nil")
		}

	} else {
		log.Fatal("status error")
	}
}

func HistoryLog() {
	historyLog := "/api/v1/history_log"
	get(fmt.Sprintf("%s%s", host, historyLog))

}
func BenchmarkHistoryLog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HistoryLog()
	}
}
func TestHistoryLog(t *testing.T) {
	HistoryLog()
}

func GetAlarmHistory() {
	getAlarmHistory := "/api/v1/get_alarm_history"
	get(fmt.Sprintf("%s%s", host, getAlarmHistory))
}

func BenchmarkGetAlarmHistory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetAlarmHistory()
	}
}
