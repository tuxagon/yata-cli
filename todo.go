package main

import (
	"errors"
	"os"
	//"flag"
	"fmt"
	"path"
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

func GetHomeEnv() string {
	val, ok := os.LookupEnv("HOME")
	if ok {
		return val
	}
	panic(errors.New("Expected 'HOME' environment variable"))
}

func GetOr(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

func main() {
	home := GetHomeEnv()
	rootPath := path.Join(home, ".yata")
	_, err := os.Stat(rootPath)
	if err != nil {
		os.Mkdir(rootPath, 0777)
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
