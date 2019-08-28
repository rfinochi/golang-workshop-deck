package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type data struct {
	Num  int    `json:"Num"`
	Text string `json:"Text"`
}

func main() {
	var newData data

	json.Unmarshal([]byte("{\"Num\":1,\"Text\":\"Test_1\"}"), &newData)

	fmt.Println(newData)

	json.NewEncoder(os.Stdout).Encode(newData)
}
