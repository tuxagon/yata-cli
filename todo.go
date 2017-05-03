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

var (
	ErrConfigNotLoaded = errors.New(fmt.Sprintf("Err: Could not load config"))
)

type Config struct {
	Home    string `env:"HOME"`
	BaseDir string `default:".todo"`
	TodoDir string `default:"tasks"`
}

func LoadConfig(v interface{}) error {
	ptrRef := reflect.ValueOf(v)
	if ptrRef.Kind() != reflect.Ptr {
		return ErrConfigNotLoaded
	}

	ref := ptrRef.Elem()
	if ref.Kind() != reflect.Struct {
		return ErrConfigNotLoaded
	}

	refType := ref.Type()
	for i := 0; i < refType.NumField(); i++ {
		field := refType.Field(i)
		fmt.Printf("%d. %v (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name())
	}
	return nil
}

func DisplayTag(v interface{}) {
	t := reflect.TypeOf(v)

	val := reflect.ValueOf(v)
	fmt.Println("ValueOf:", val.Type())

	fmt.Println("Type:", t.Name())
	fmt.Println("Kind:", t.Kind())

	for i := 0; i < val.Type().NumField(); i++ {
		field := val.Type().Field(i)
		tag := field.Tag.Get(envTag)
		fmt.Printf("%d. %v (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
		tag = field.Tag.Get(defaultTag)
		fmt.Printf("%d. %v (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
	}
}

func main() {
	cfg := Config{}
	err := LoadConfig(&cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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

	// fmt.Println("type:", reflect.ValueOf(&cfg).Elem().Type().Field(0).Tag.Get("env"))
	// err := Parse(1)
	// fmt.Println(err)
	//DisplayTag(&cfg)

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
