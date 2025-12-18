package main

import (
	"encoding/json"
	"os"
)

type Tracker struct {
	Task_list []Task `json:"tasks"`
	Last_id int `json:"last_id"`
}

func (t *Tracker)Save() {
	b, _ := json.Marshal(t)
	os.WriteFile("tasks.json", b, 0644)
}

func (t *Tracker) Load(){
	b, _ := os.ReadFile("tasks.json")
	json.Unmarshal(b, t)
}

func (t *Tracker) AddTask(task Task){
	task.Id = t.Last_id
	t.Last_id++
	t.Task_list = append(t.Task_list, task)
}

func (t *Tracker) DeleteTask(id int) {
	i:=0
	for {
		if t.Task_list[i].Id == id {
			t.Task_list = append(t.Task_list[:i], t.Task_list[i+1:]...)
		}
		i++
		if i==len(t.Task_list) {
			return
		}
	}
}

func (t *Tracker) UpdateTask(id int, new_task string) {
	i:=0
	for {
		if t.Task_list[i].Id == id {
			t.Task_list[i].Task = new_task
		}
		i++
		if i==len(t.Task_list) {
			return
		}
	}
}

func (t *Tracker) ChangeProgress(id int, new_progress string) {
	i:=0
	for {
		if t.Task_list[i].Id == id {
			t.Task_list[i].Is_done = new_progress
		}
		i++
		if i==len(t.Task_list) {
			return
		}
	}
}

func (t *Tracker) FindById(id int) Task {
	i:=0
	for {
		if t.Task_list[i].Id == id{
			return t.Task_list[i]
		}
		i++
		if i==len(t.Task_list) {
			break
		}
	}
	return Task{}
}

func (t *Tracker) List(query string) {
	switch query{
	case "all":
		i:=0
		for {
			t.Task_list[i].PrettyPrint()
			i++
			if i==len(t.Task_list) {
			return
		}
		}

	case "done":
		i:=0
		for {
			if t.Task_list[i].Is_done == "Done"{
				t.Task_list[i].PrettyPrint()
			}
			i++
			if i==len(t.Task_list) {
			return
		}
		}
	case "todo":
		i:=0
		for {
			if t.Task_list[i].Is_done == "todo"{
				t.Task_list[i].PrettyPrint()
			}
			i++
			if i==len(t.Task_list) {
			return
		}
		}
	case "in progress":
		i:=0
		for {
			if t.Task_list[i].Is_done == "in progress"{
				t.Task_list[i].PrettyPrint()
			}
			i++
			if i==len(t.Task_list) {
			return
		}
		}

	}
}