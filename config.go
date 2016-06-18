package main

import (
  	"encoding/json"
)

type Config struct {
    Address string `json:"address"`
}

func getConfig() Config {
	jsonStr := readFile("./config.json")
    res := Config{}
    json.Unmarshal([]byte(jsonStr), &res)
    return res
}

func setConfig(config Config) {
	jsonStr, _ := json.Marshal(config)
	writeFile("./config.json", jsonStr)
}