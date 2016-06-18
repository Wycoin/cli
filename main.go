package main

import (
    "fmt"
    "os"

    "github.com/urfave/cli"
)

func main() {
    app := cli.NewApp()

    app.Name = "wyllet"
    app.Usage = "wycoin wallet cli"
    app.Version = "0.1.0"

    app.Commands = []cli.Command {
        {
            Name: "create",
            Aliases: []string{"c", "login"},
            Usage: "create a new account",
            Action: func(c *cli.Context) error {
                if checkHasAccount() {
                    askOverwriteAccount()
                } else {
                    createAccount()
                }
                return nil
            },
        },
        {
            Name: "balance",
            Aliases: []string{"b"},
            Usage: "complete a task on the list",
            Action: func(c *cli.Context) error {
                fmt.Println("completed task: ", c.Args().First())
                return nil
            },
        },
        {
            Name: "transfer",
            Aliases: []string{"t"},
            Usage: "create a new transfer",
            ArgsUsage: "<amount> <address>",
            Action: func(c *cli.Context) error {
                if len(c.Args()) == 2 {
                    amount := c.Args()[0]
                    address := c.Args()[1]

                    fmt.Println(amount, address)
                } else {
                    fmt.Println("usage: wyllet transfer <amount> <address>")
                }
                return nil
            },
        },
    }

    app.Run(os.Args)
}