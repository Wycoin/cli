package main

import (
	"fmt"
  	"os"
  	"io/ioutil"
)

func writeFile(path string, str []byte) {
	f, _ := os.Create(path)
	defer f.Close()
	f.Write(str)
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
  		return true
	} else {
		return false
	}
}

func readFile(path string) string {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error while reading ",path)
	}
    return string(contents)
}