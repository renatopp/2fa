package main

import (
	"fmt"
	"github.com/pquerna/otp/totp"
	"github.com/renatopp/2fa/internal"
	"golang.org/x/term"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		cmdHelp()
		exit(1)
	}

	cmd := os.Args[1]
	switch cmd {
	case "add":
		cmdAdd()
	case "list":
		cmdList()
	case "remove":
		cmdRemove()
	case "show":
		cmdShow()
	case "help":
		cmdHelp()
	}
}

func cmdHelp() {
	fmt.Println("Usage: 2fa <command> [options] <args>")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("- add <name>           Add a new 2FA entry")
	fmt.Println("- list                 List all 2FA entries")
	fmt.Println("- remove <name>        Remove a 2FA entry")
	fmt.Println("- show <name>          Show a 2FA entry")
	fmt.Println("- help                 Show this help message")
}

func cmdAdd() {
	if len(os.Args) != 3 {
		fmt.Println("Error: invalid arguments. Usage: 2fa add <name> <code>")
		exit(1)
	}

	name := os.Args[2]
	fmt.Printf("Code: ")
	code, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	if err != nil {
		fmt.Println("Error:", err)
		exit(1)
	}

	err = internal.Set(name, string(code))
	if err != nil {
		fmt.Println("Error:", err)
		exit(1)
	}
}

func cmdList() {
	names, err := internal.List()
	if err != nil {
		fmt.Println("Error:", err)
		exit(1)
	}
	for _, name := range names {
		fmt.Println(name)
	}
}

func cmdRemove() {
	if len(os.Args) != 3 {
		fmt.Println("Error: invalid arguments. Usage: 2fa remove <name>")
		exit(1)
	}

	name := os.Args[2]
	err := internal.Remove(name)
	if err != nil {
		fmt.Println("Error:", err)
		exit(1)
	}
	fmt.Printf("Removed %s\n", name)
}

func cmdShow() {
	if len(os.Args) != 3 {
		fmt.Println("Error: invalid arguments. Usage: 2fa show <name>")
		exit(1)
	}

	name := os.Args[2]
	code, err := internal.Get(name)
	if err != nil {
		fmt.Println("Error:", err)
		exit(1)
	}

	t, err := totp.GenerateCode(code, time.Now())
	if err != nil {
		fmt.Println("Error generating totp code:", err)
		os.Exit(1)
	}

	fmt.Println(t)
}

func exit(code int) {
	if code != 0 {
		fmt.Println()
	}
	os.Exit(code)
}
