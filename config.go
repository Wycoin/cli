package main

import (
	"os"
  	"encoding/json"
)

type Config struct {
    Address string `json:"address"`
    EntryNodeURL string `json:"entryNodeURL"`
}

func getConfigPath() string {
	return "./wyllet-config.json"
}

func getConfig() Config {
	jsonStr := readFile(getConfigPath())
    res := Config{}
    json.Unmarshal([]byte(jsonStr), &res)
    return res
}

func setConfig(config Config) {
	jsonStr, _ := json.Marshal(config)
	writeFile(getConfigPath(), jsonStr)
}

func rmConfig() {
	os.Remove(getConfigPath())
}