package main

import (
	"encoding/json"
	"fmt"
)

type target struct {
	Name string `json:"name"`
	Threshold int `json:"threshold"`
}

type config struct {
	Addr string `json:"addr"`
	Target target `json:"target"`
}

func main() {
	cfg := config{
		Addr: "Japan",
		Target: target{
			Name: "bar",
			Threshold: 4,
		},
	}
	b, _ := json.MarshalIndent(&cfg, "", "  ")
	fmt.Println(string(b))
}
