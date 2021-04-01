package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestN(t *testing.T) {
	n := N{
		Warehouse: map[string]struct {
			Group map[string][]Zone `json:"group"`
		}{
			"w1": {
				map[string][]Zone{
					"g1": {
						{
							Id: 1,
						},
						{
							Id: 2,
						},
					},
					"g2": {
						{
							Id: 1,
						},
						{
							Id: 2,
						},
					},
				},
			},
			"w2": {
				map[string][]Zone{
					"g1": {
						{
							Id: 1,
						},
						{
							Id: 2,
						},
					},
					"g2": {
						{
							Id: 1,
						},
						{
							Id: 2,
						},
					},
				},
			},
		},
	}
	data, _ := json.MarshalIndent(n, "", " ")
	fmt.Println(string(data))

	n2 := N2{
		"w":  {"g1": {{Id: 1}, {Id: 2}, {Id: 3}}, "g2": {{Id: 1}}},
		"w2": {"g1": {{Id: 1}}, "g2": {{Id: 1}}},
	}
	data, _ = json.MarshalIndent(n2, "", " ")
	fmt.Println(string(data))

}
