package main

import (
	"context"
	"scheduler/scheduler"
	"scheduler/task"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()

	schdlr := scheduler.NewScheduler(ctx, 5)

	t1 := task.NewSleepyTask("A", 2*time.Second)
	t2 := task.NewSleepyTask("B", 2*time.Second)

	schdlr.Schedule(t1, task.ExecutionType{
		TimeInSeconds: 5,
		Periodic:      true,
	})
	schdlr.Schedule(t2, task.ExecutionType{
		TimeInSeconds: 10,
		Periodic:      true,
	})

	<- ctx.Done()
}
