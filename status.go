package main

import (
	"fmt"
	"log"

	"github.com/urfave/cli/v2"
)

var Status cli.Command = cli.Command{
	Name:    "status",
	Aliases: []string{"s"},
	Usage:   "Check current authentication status.",
	Action: func(ctx *cli.Context) error {
		token := GetCredentials().Token
		if token == "" {
			log.Fatal(`You didn't auth yourself yet. Run lines-of-gode auth --token "your_token"`)
		}
		context, client := CreateClient(token)
		user := GetUser(context, client)

		fmt.Printf("You are authed as %q (%v).\n", user.GetLogin(), user.GetName())

		return nil
	},
}
