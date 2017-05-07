package main

import (
	"os"
	"yata-cli/yata"
)

func SeedFile(m *yata.Manager) {
	t1 := yata.NewHighPriorityTask("Task 1", "")
	t2 := yata.NewLowPriorityTask("Task 2", "")
	t3 := yata.NewSimpleTask("Task 3", "")
	m.SaveNewTask(*t1)
	m.SaveNewTask(*t2)
	m.SaveNewTask(*t3)
}

func main() {
	ym := yata.NewManager()
	ym.InitializeDirectory()

	parser := yata.NewCmdParser(ym)
	if len(os.Args) <= 1 {
		parser.Usage()
	} else {
		parser.Parse(os.Args[1:])
	}
}
