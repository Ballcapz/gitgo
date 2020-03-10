package main

import (
	"fmt"
	"os"
	"strings"

	color "github.com/fatih/color"
	flag "github.com/ogier/pflag"
)

var (
	user string
)

func main() {

	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	users := strings.Split(user, ",")
	fmt.Printf("Searching user(s): %s\n", users)
	fmt.Println("")
	for _, u := range users {
		result := getUsers(u)
		color.Cyan(`Username: %s`, result.Login)
		color.Blue(`URL:	%s`, result.URL)
		color.HiMagenta(`Repos:	%s`, result.ReposURL)
		fmt.Println("")
	}
}

func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users")
}
