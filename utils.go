package main

import (
	  "fmt"
  	"os"
  	"bufio"
  	"crypto/sha256"
  	"encoding/hex"
    "net/http"
    "io/ioutil"
)

func hashStr(str string) string{
    h := sha256.New()
    h.Write([]byte(str))
    hash := hex.EncodeToString(h.Sum(nil))
    return hash
}

func askForInput(msg string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	result, _ := reader.ReadString('\n')
	return result
}

func postJSON(url string, jsonStr string) (resp interface, err interface, body interface) {
    json := []byte(jsonStr)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    defer resp.Body.Close()

    if err != nil {
      return (resp, err, nil)  
    }

    body, _ := ioutil.ReadAll(resp.Body)
    return (resp, err, body)
}