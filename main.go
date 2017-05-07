package main

import (
	"os"
	"yata-cli/yata"
)

// func SeedFile(m *yata.Manager) {
// 	t1 := yata.NewHighPriorityTask("Task 1", "")
// 	t2 := yata.NewLowPriorityTask("Task 2", "")
// 	t3 := yata.NewSimpleTask("Task 3", "")
// 	m.SaveNewTask(*t1)
// 	m.SaveNewTask(*t2)
// 	m.SaveNewTask(*t3)
// }

func main() {
	m := yata.NewFileManager()
	p := yata.NewCmdParser(m)
	if len(os.Args) > 1 {
		p.Parse(os.Args[1:])
	} else {
		p.Usage()
	}
}
