package main

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

var ErrTimeout = errors.New("执行超时")
var ErrInterrupt = errors.New("执行被中断")

type Runner struct {
	tasks     []func(int)
	complete  chan error
	timeout   <-chan time.Time
	interrupt chan os.Signal
}

func New(tm time.Duration) *Runner {
	return &Runner{
		tasks:     nil,
		complete:  make(chan error),
		timeout:   time.After(tm),
		interrupt: make(chan os.Signal),
	}
}

func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.isInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (r *Runner) isInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)
	go func() {
		r.complete <- r.run()
	}()
	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

func main() {

}
