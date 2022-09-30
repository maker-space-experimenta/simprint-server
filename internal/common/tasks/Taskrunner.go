package tasks

import (
	"log"
	"time"
)

type Task struct {
	Callback func()
}

type TaskRunner struct {
	interval int
	tasks    []Task
}

func NewTaskRunner(interval int) *TaskRunner {
	return &TaskRunner{
		interval: interval,
	}
}

func (m *TaskRunner) AddTask(cb func()) {
	m.tasks = append(m.tasks, Task{
		Callback: cb,
	})
}

func (m *TaskRunner) Start() {
	log.Println("TaskRunner - Starting TaskRunner")

	go m.runLoop()
}

func (m *TaskRunner) runLoop() {
	for {
		m.Loop()
		time.Sleep(time.Second * time.Duration(m.interval))
	}
}

func (m *TaskRunner) Loop() {
	log.Println("TaskRunner - Run TaskRunner Loop")

	for _, element := range m.tasks {
		element.Callback()
	}
}
