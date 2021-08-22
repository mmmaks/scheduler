package task

import (
	"context"
	"fmt"
	"scheduler/utils"
	"time"
)

type sleepyTask struct {
	name string
	duration time.Duration
}

func (s sleepyTask) Name() string {
	return s.name
}

func (s sleepyTask) Run(ctx context.Context) {
	fmt.Printf("Running task: %v, current time %v\n", s.name, utils.TimeInSeconds())
	time.Sleep(s.duration)
}

func NewSleepyTask(name string, duration time.Duration) Task {
	return &sleepyTask{name: name, duration: duration}
}

