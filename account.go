package main

import (
  	"fmt"
  	"strings"
)

func checkHasAccount() bool {
	if fileExists("./config.json") {
		config := getConfig()

		if config.Address != "" {
			return true
		} else {
			return false
		}

	} else {
		return false
	}
}

func askOverwriteAccount(){
    fmt.Println(`
You already have an account saved. Continue-ing will replace it. You can log in again
at any time using your username and private key by running "wyllet login"
    `)

    answer := askForInput("Would you like to continue? [y/N] ")
    answer = strings.TrimSpace(answer)

    if answer == "y" {
        createAccount()
    }
}

func createAccount() {
	fmt.Println(`
Welcome to Wyllet!
Great that you would like to start using Wyllet.
To start we need to create you an account. Enter a username
and the path to your private key below and we'll get you set up.
    `)

    username := askForInput("Username: ")
    fmt.Println(username)

    keyPath := askForInput("Path to private key: ")
    keyPath = strings.TrimSpace(keyPath)

    if fileExists(keyPath) {
    	contents := readFile(keyPath)
    	address := hashStr(username + contents)

    	config := Config{address}
    	setConfig(config)
    	
    	if checkHasAccount() {
    		fmt.Println("Thanks! Now type wyllet --help to get started.")
    	} else {
    		fmt.Println("An error occured.")
    	}
    } else {
    	fmt.Println(keyPath, " does not exist")
    }
}