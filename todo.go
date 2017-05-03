package main

import (
	"errors"
	"os"
	//"flag"
	"fmt"
	"reflect"
)

const (
	envTag     = "env"
	defaultTag = "default"
)

type Config struct {
	Home    string `env:"HOME"`
	BaseDir string `default:".todo"`
	TodoDir string `default:"tasks"`
}

func Parse(v interface{}) error {
	ptr := reflect.ValueOf(v)
	if ptr.Kind() != reflect.Ptr {
		return errors.New(fmt.Sprintf("Expected: struct pointer, Got: %s", ptr.Kind()))
	}
	return nil
}

func DisplayTag(v interface{}) {
	t := reflect.TypeOf(v)

	fmt.Println("Type:", t.Name())
	fmt.Println("Kind:", t.Kind())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(envTag)
		fmt.Printf("%d. %v (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
		tag = field.Tag.Get(defaultTag)
		fmt.Printf("%d. %v (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
	}
}

func main() {
	if len(os.Args) <= 1 {
		// TODO Show usage
		os.Exit(1)
	}
	args := os.Args[1:]

	// home := os.Getenv("HOME")
	// if home == "" {
	// 	home = ""
	// }
	// todoRoot := os.Getenv("TODO_ROOT")
	// if todoRoot == "" {
	// 	todoRoot = ""
	// }

	cfg := Config{}
	// fmt.Println("type:", reflect.ValueOf(&cfg).Elem().Type().Field(0).Tag.Get("env"))
	// err := Parse(1)
	// fmt.Println(err)
	DisplayTag(cfg)

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
