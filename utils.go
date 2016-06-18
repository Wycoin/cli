package main

import (
	  "fmt"
  	"os"
  	"bufio"
  	"crypto/sha256"
  	"encoding/hex"
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