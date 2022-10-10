package tasks

import (
	"time"

	"github.com/maker-space-experimenta/printer-kiosk/internal/common/logging"
)

type Task struct {
	Callback func()
}

type TaskRunner struct {
	interval int
	tasks    []Task
	logger   *logging.Logger
}

func NewTaskRunner(interval int) *TaskRunner {
	return &TaskRunner{
		interval: interval,
		logger:   logging.NewLogger(),
	}
}

func (m *TaskRunner) AddTask(cb func()) {
	m.tasks = append(m.tasks, Task{
		Callback: cb,
	})
}

func (m *TaskRunner) Start() {
	m.logger.Infof("TaskRunner - Starting TaskRunner")
	go m.runLoop()
}

func (m *TaskRunner) runLoop() {
	for {
		m.Loop()
		time.Sleep(time.Second * time.Duration(m.interval))
	}
}

func (m *TaskRunner) Loop() {
	m.logger.Infof("TaskRunner - Run TaskRunner Loop")

	for _, element := range m.tasks {
		element.Callback()
	}
}
