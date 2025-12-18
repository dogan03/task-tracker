package main

import (
	"flag"
	"strconv"
	"strings"

	"github.com/fatih/color"
)


func main() {
	add := flag.String("add", "-1", "adding new task to task tracker")
	delete := flag.Int("delete", -1, "delete a task by id")
	update := flag.String("update", "-1", "updates the task by id")
	mark_in_progress := flag.Int("mark_in_progress", -1, "mark a task in progress by id")
	mark_done := flag.Int("mark_done", -1, "mark a task done by id")
	list := flag.String("list", "-1", "List the tasks with: all, done, todo, in progress")
	flag.Parse()
	
	tracker:=Tracker{}
	tracker.Load()
	defer tracker.Save()

	if *add != "-1"{
		t := NewTask(*add, 0)
		tracker.AddTask(t)
		color.Green("Added the task:\n")
		t.PrettyPrint()
	}
	
	if *delete != -1 {
		t:= tracker.FindById(*delete)
		color.Green("Deleting the task:\n")
		t.PrettyPrint()
		tracker.DeleteTask(*delete)
	}
	
	if *update != "-1" {
		o := strings.Split(*update, ":")
		id, task := o[0], o[1]
		id_int,_ := strconv.Atoi(id)
		
		color.Green("Updated the task with id: %d\n", id)
		tracker.UpdateTask(id_int, task)
		t := tracker.FindById(id_int)
		t.PrettyPrint()
	}
	
	if *mark_in_progress != -1{
		color.Green("Marking the task in progress with id: %d\n", *mark_in_progress)
		t:= tracker.FindById(*mark_in_progress)
		t.PrettyPrint()
		tracker.ChangeProgress(*mark_in_progress, "in progress")
	}
	
	if *mark_done != -1{
		color.Green("Marking the task done with id: %d\n", *mark_done)
		t:= tracker.FindById(*mark_done)
		t.PrettyPrint()
		tracker.ChangeProgress(*mark_done, "done")
	}
	
	if *list != "-1" {
		color.Green("Listing %s Tasks!", *list)
		tracker.List(*list)
	}
}
