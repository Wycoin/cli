package main

import (
	"net/http"
    "io/ioutil"
)

func httpGet(url string) string {
    resp, _ := http.Get("http://google.com")
    body, _ := ioutil.ReadAll(resp.Body);
    return string(body)
}

// func postJSON(url string, jsonStr string) (resp interface, err interface, body interface) {
//     json := []byte(jsonStr)
//     req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
//     req.Header.Set("Content-Type", "application/json")

//     client := &http.Client{}
//     resp, err := client.Do(req)
//     defer resp.Body.Close()

//     if err != nil {
//       return (resp, err, nil)  
//     }

//     body, _ := ioutil.ReadAll(resp.Body)
//     return (resp, err, body)
// }