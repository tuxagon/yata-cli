package main

import (
	"os"
	//"flag"
	"fmt"
)

func main() {
	args := os.Args[1:]

	switch args[0] {
	case "add", "new":
		fmt.Println("New todo!")
	case "config":
		fmt.Println("Configuring")
	case "list", "ls":
		fmt.Println("Listing")
	default:
		fmt.Println("Usage")
	}

	fmt.Println(args)
}
