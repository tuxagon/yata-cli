package yata

import (
	"flag"
	"fmt"
	"sort"
)

var CmdMap = map[string]CmdFunc{
	"":     NewNoCmd,
	"list": NewListCmd,
}

type CmdFunc (func() YataCmd)

// YataCmd declares functions that are implemented by each Yata command
type YataCmd interface {
	ParseFlags(args []string)
	Execute(manager *Manager)
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
func (cmd ListCmd) Execute(manager *Manager) {
	tasks := manager.GetAllOpenTasks()
	sort.Sort(ByPriority(tasks))
	for _, t := range tasks {
		fmt.Println(t)
	}
}
func (cmd NoCmd) Execute(manager *Manager) {
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
	Manager *Manager
}

// NewCmdParser creates a new CmdParser
func NewCmdParser(manager *Manager) *CmdParser {
	return &CmdParser{
		Manager: manager,
	}
}

// Parse parses the arguments provided
func (p *CmdParser) Parse(args []string) {
	// hd, tl := stringSplitHeadFromTail(args)

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

func (p *CmdParser) getBaseFlagSet() *flag.FlagSet {
	return flag.NewFlagSet("base", flag.ContinueOnError)
}

func (p *CmdParser) Usage() {
	fmt.Println("TODO: Show usage")
}

// StringSplitHeadFromTail will split a list into a head and tail
func stringSplitHeadFromTail(list []string) (string, []string) {
	n := len(list)
	switch {
	case n == 1:
		return list[0], make([]string, 0)
	case n > 1:
		return list[0], list[1:]
	}
	return "", make([]string, 0)
}
