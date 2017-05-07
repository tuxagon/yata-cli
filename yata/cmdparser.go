package yata

import (
	"flag"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type sharedFlags struct {
	help bool `alias:"h"`
}

type rootArgs struct {
	version bool `alias:"v"`
}

type listArgs struct {
	sortBy string `alias:"s"`
}

var CmdMap = map[string]CmdFunc{
	"":     NewNoCmd,
	"list": NewListCmd,
}

type CmdFunc (func() YataCmd)

// YataCmd declares functions that are implemented by each Yata command
type YataCmd interface {
	ParseFlags(args []string)
	Execute(manager TaskManager)
	Usage() string
}

// CmdList provides a list of all the commands
type CmdList map[string]YataCmd

// ListCmd represents all the information regarding a list command
type ListCmd struct{}
type NoCmd struct{}

// NewListCmd creates a new help command handler
func NewListCmd() YataCmd { return ListCmd{} }
func NewNoCmd() YataCmd   { return NoCmd{} }

// ParseFlags will parse the options for the list command
func (cmd ListCmd) ParseFlags(args []string) {}
func (cmd NoCmd) ParseFlags(args []string)   {}

// Execute will handle the list command
func (cmd ListCmd) Execute(manager TaskManager) {
	tasks := manager.GetAllOpenTasks(&ByPriority{})
	sort.Sort(ByPriority(tasks))
	for _, t := range tasks {
		fmt.Println(t)
	}
}
func (cmd NoCmd) Execute(manager TaskManager) {
	fmt.Println("NoCmd executing")
}

// Usage will display usage for the list command
func (cmd ListCmd) Usage() string {
	return "[list] command usage"
}
func (cmd NoCmd) Usage() string {
	return "[nocmd] command usage"
}

// GetCmd gets a map of the commands with their implementation
func GetCmd(cmd string) (YataCmd, bool) {
	cmdFunc, ok := CmdMap[cmd]
	if !ok {
		return nil, false
	}
	return cmdFunc(), true
}

// CmdParser represents the command-line parser
type CmdParser struct {
	Manager TaskManager
}

// NewCmdParser creates a new CmdParser
func NewCmdParser(manager TaskManager) *CmdParser {
	return &CmdParser{
		Manager: manager,
	}
}

// Parse parses the arguments provided
func (p *CmdParser) Parse(args []string) {
	hd, tl := getHeadAndTail(args)

	if len(hd) > 0 {
		if isFlag(hd) {
			p.parseFlags(args, &rootArgs{})
		} else {
			p.parseCmd(hd, tl)
		}
	} else {
		p.Usage()
	}
	// cmd, ok := GetCmd(hd)
	// if !ok {
	// 	cmd, ok = GetCmd("")
	// 	if !ok {
	// 		fmt.Printf("yata: '%s' is not a yata command\nSee 'yata --help'\n", hd)
	// 		return
	// 	}
	// }

	// cmd.ParseFlags(tl)
	// cmd.Execute(p.Manager)
}

func (p *CmdParser) parseFlags(args []string, flagSet interface{}) {
	ptrRef := reflect.ValueOf(flagSet)
	if ptrRef.Kind() != reflect.Ptr {
		// TODO return err
	}

	ref := ptrRef.Elem()
	if ref.Kind() != reflect.Struct {
		// TODO return err
	}

	refType := ref.Type()
	for i := 0; i < refType.NumField(); i++ {
		field := refType.Field(i)
		field.Tag.Get("alias")
		// TODO check for flags with "flag" package based on alias or full scope
		// TODO anything with a default is checked for and if not found, then default
	}
}

func (p *CmdParser) parseCmd(cmd string, args []string) {

}

func (p *CmdParser) Usage() {

}

func (p *CmdParser) getBaseFlagSet() *flag.FlagSet {
	return flag.NewFlagSet("base", flag.ContinueOnError)
}

// getHeadAndTail will split a list into a head and tail
func getHeadAndTail(list []string) (string, []string) {
	n := len(list)
	switch {
	case n == 1:
		return list[0], make([]string, 0)
	case n > 1:
		return list[0], list[1:]
	}
	return "", make([]string, 0)
}

func isFlag(arg string) bool {
	return arg[0] == '-'
}

func parseTagForFlag(key string) (string, []string) {
	flags := strings.Split(key, ",")
	return flags[0], flags[1:]
}
