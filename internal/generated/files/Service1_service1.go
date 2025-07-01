package main

import "C"
import (
	"encoding/json"
)

const ABType = "Service1ABTest"

type Service1Out struct {
	Service1Param int      `json:"service1Param"`
	ABType        string   `json:"ABType"`
	ABExists      bool     `json:"AbExists"`
	MyArray       []string `json:"MyArray"`
}

func ModifyResponse(data []byte) []byte {
	var out Service1Out
	if err := json.Unmarshal(data, &out); err != nil {
		return data
	}

	ABModify(&out)
	modified, _ := json.Marshal(out)
	return modified
}

func ABModify(in *Service1Out) {
	in.ABType = ABType
	in.ABExists = true
	in.Service1Param = in.Service1Param * 15
	in.ABType = "MyCustomABTest"
	in.ABExists = true
	in.MyArray = append(in.MyArray, []string{"hello1", "hello2"}...)
}

func main() {}
