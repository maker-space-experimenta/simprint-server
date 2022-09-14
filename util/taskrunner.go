package util

import (
	"log"
	"time"
)

type Task struct {
	Callback func()
}

type TaskRunner struct {
	config Config
	tasks  []Task
}

func NewTaskRunner(config Config) *TaskRunner {
	return &TaskRunner{
		config: config,
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
		time.Sleep(time.Second * time.Duration(m.config.TaskRunnerDuration))
	}
}

func (m *TaskRunner) Loop() {
	log.Println("TaskRunner - Run TaskRunner Loop")

	for _, element := range m.tasks {
		element.Callback()
	}
}
