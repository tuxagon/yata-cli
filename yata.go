package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"time"
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

type MaybeTime struct {
	Time  time.Time
	Valid bool
}

type YataConfig struct {
	BaseDir string
}

type YataTask struct {
	Description string //`json:"desc"`
	//created     time.Time `json:"created"`
	//completed   MaybeTime `json:"done"`
	Priority int    //`json:"priority"`
	Project  string //`json:"project"`
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

func CreateRoot(rootPath string) {
	_, err := os.Stat(rootPath)
	if err != nil {
		os.Mkdir(rootPath, 0777)
	}
}

func CreateMainFile(rootPath, filename string) {
	fullPath := path.Join(rootPath, filename)
	_, err := os.Stat(fullPath)
	if err != nil {
		os.Create(fullPath)
	}
}

func main() {
	home := GetHomeEnv()
	rootPath := path.Join(home, ".yata")
	CreateRoot(rootPath)
	CreateMainFile(rootPath, "yata")
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
		simpleTask := &YataTask{"This is my first task", 2, "test"}
		t, err := json.Marshal(simpleTask)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(t))
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
