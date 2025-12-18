package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)
type Task struct {
	Id      int    `json:"id"`
	Task    string `json:"task"`
	Is_done string `json:"is_done"`
}

func NewTask(task string, id int) Task {

	is_done := "todo"

	return Task{
		Id:      id,
		Task:    task,
		Is_done: is_done,
	}
}

func (t *Task) PrettyPrint() {
	fmt.Print("\n")
	color.Cyan(strings.Repeat("*", 50))
	color.Red("ID of the task: %d", t.Id)
	color.Red("Description of the task: %s", t.Task)
	color.Red("Progress of the task: %s", t.Is_done)
	color.Cyan(strings.Repeat("*", 50))
	fmt.Print("\n")
}