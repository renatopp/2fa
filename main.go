package main

import (
	"fmt"
	"os"
	"time"

	"github.com/pquerna/otp/totp"
	"github.com/renatopp/2fa/internal"
	"github.com/renatopp/go-cli"
	"golang.org/x/term"
)

func main() {
	cli.Name("2fa")
	cli.Description("Two factor authentication in your command line. \n\nConsult https://github.com/renatopp/2fa for more information.")
	cli.Command("add", "Add a new 2FA entry.", cmdAdd)
	cli.Command("list", "List all 2FA entries.", cmdList)
	cli.Command("remove", "Remove a 2FA entry.", cmdRemove)
	cli.Command("show", "Show a 2FA entry.", cmdShow)
	cli.AutoHelp(true)
	cli.Parse()
	cli.ShowHelp()
}

func cmdAdd() {
	cli.Name("add")
	cli.Description("Adds a new 2FA entry in the database.")
	name := cli.Pos("name", "The identification name of the 2FA entry.").AsRequired()
	key := cli.Pos("secret_key", "The secret key of the 2FA entry.")
	cli.Parse()

	var code []byte = []byte(key.Value())
	var err error
	if !key.IsParsed() {
		fmt.Printf("Code: ")
		code, err = term.ReadPassword(int(os.Stdin.Fd()))
		fmt.Println()
		if err != nil {
			fmt.Println("Error:", err)
			cli.Exit(1)
		}
	}

	err = internal.Set(name.Value(), string(code))
	if err != nil {
		fmt.Println("Error:", err)
		cli.Exit(1)
	}
}

func cmdList() {
	names, err := internal.List()
	if err != nil {
		fmt.Println("Error:", err)
		cli.Exit(1)
	}
	for _, name := range names {
		fmt.Println(name)
	}
}

func cmdRemove() {
	name := cli.Pos("name", "The identification name of the 2FA entry to be removed.").AsRequired()
	cli.Parse()

	err := internal.Remove(name.Value())
	if err != nil {
		fmt.Println("Error:", err)
		cli.Exit(1)
	}
	fmt.Printf("Removed %s\n", name.Value())
}

func cmdShow() {
	name := cli.Pos("name", "The identification name of the 2FA entry to be shown.").AsRequired()
	cli.Parse()

	code, err := internal.Get(name.Value())
	if err != nil {
		fmt.Println("Error:", err)
		cli.Exit(1)
	}

	t, err := totp.GenerateCode(code, time.Now())
	if err != nil {
		fmt.Println("Error generating totp code:", err)
		cli.Exit(1)
	}

	fmt.Println(t)
}
