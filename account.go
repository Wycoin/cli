package main

import (
  	"fmt"
  	"strings"
)

func checkHasAccount() bool {
	if fileExists(getConfigPath()) {
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

    pvtKeyPath := askForInput("Path to private key: ")
    pvtKeyPath = strings.TrimSpace(pvtKeyPath)

    if fileExists(pvtKeyPath) {

        pblKeyPath := askForInput("Path to public key: ")
        pblKeyPath = strings.TrimSpace(pblKeyPath)

        if fileExists(pblKeyPath) {

            //create address
            //by combining username, public key & private key
            pvtKey := readFile(pvtKeyPath)
            pblKey := readFile(pblKeyPath)
            address := hashStr(username + pvtKey + pblKey)

            //ask for entry node url
            //default to 192.168.2.1
            entryNodeURL := askForInput("Entry node URL [192.168.2.1]: ")
            if entryNodeURL == "" {
                entryNodeURL = "192.168.2.1"
            }

            //create config instance /w address
            config := Config{address, entryNodeURL}

            //create config file
            setConfig(config)
            
            //check if account was saved successfully
            if checkHasAccount() {

                //send public key to entry node
                fmt.Println("Sending public key to entry node ("+entryNodeURL+")...")
                //postJSON(entryNodeURL, "{ \"publicKey\": \""+pblKey+"\" }")

                err := ""

                if err != "" {
                    //if request fails, remove config
                    rmConfig()
                    fmt.Println("The request to the entrynode could not be completed.")
                } else {
                    //show done message
                    fmt.Println("Thanks! Now type wyllet --help to get started.")
                    fmt.Println("Debug: ")
                    fmt.Println(username)
                    fmt.Println(pvtKey)
                    fmt.Println(pblKey)
                    fmt.Println(address)
                    fmt.Println(entryNodeURL)
                }

            } else {
                fmt.Println("An error occured.")
            }

        } else {
            fmt.Println(pblKeyPath, " does not exist")
        }

    } else {
    	fmt.Println(pvtKeyPath, " does not exist")
    }
}