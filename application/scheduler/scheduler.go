package scheduler

import (
	"log"
	"time"
)

type Task interface {
	Do()
}

type PeriodicTask struct {
	task     Task
	interval time.Duration
	ticker   *time.Ticker
	stop     chan bool
}

type Scheduler struct {
	tasks chan *PeriodicTask
	stop  chan bool
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks: make(chan *PeriodicTask),
		stop:  make(chan bool),
	}
}

func (s *Scheduler) AddTask(task Task, interval time.Duration) {
	periodicTask := &PeriodicTask{
		task:     task,
		interval: interval,
		stop:     make(chan bool),
	}
	s.tasks <- periodicTask
}

func (s *Scheduler) Start() {
	for {
		select {
		case task := <-s.tasks:
			go s.runTask(task)
		case <-s.stop:
			return
		}
	}
}

func (s *Scheduler) Stop() {
	s.stop <- true
}

func (s *Scheduler) runTask(task *PeriodicTask) {
	task.ticker = time.NewTicker(task.interval)
	for {
		select {
		case <-task.ticker.C:
			log.Println("Running task")
			task.task.Do()
		case <-task.stop:
			task.ticker.Stop()
			return
		}
	}
}
