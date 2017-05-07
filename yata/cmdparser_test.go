package yata_test

import (
	"sort"
	"testing"
	"yata-cli/yata"
)

type S struct {
	t *testing.T
}

type SR struct {
	t        *testing.T
	expected []interface{}
}

func Spec(t *testing.T) *S {
	return &S{t: t}
}

func (s *S) Expect(expected ...interface{}) (sr *SR) {
	return &SR{t: s.t, expected: expected}
}

func (sr *SR) ToEqual(actuals ...interface{}) {
	for index, actual := range actuals {
		if sr.expected[index] != actual {
			sr.t.Errorf("expected %+v to equal %+v", sr.expected[index], actual)
		}
	}
}

func (sr *SR) ToNotEqual(actuals ...interface{}) {
	for index, actual := range actuals {
		if sr.expected[index] == actual {
			sr.t.Errorf("expected %+v to not equal %+v", sr.expected[index], actual)
		}
	}
}

type InMemoryManager struct{}

func (m *InMemoryManager) Initialize() {}
func (m *InMemoryManager) GetAllOpenTasks(s sort.Interface) []yata.Task {
	return make([]yata.Task, 0)
}
func (m *InMemoryManager) SaveNewTask(t yata.Task) {}

// TODO Figure out a list of commands to run

func TestTODO(t *testing.T) {
	//spec := Spec(t)
	p := yata.NewCmdParser(&InMemoryManager{})
	p.Usage()
}
