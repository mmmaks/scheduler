package scheduler

import (
	"context"
	"fmt"
	"scheduler/queue"
	"scheduler/task"
	"scheduler/utils"
)

type Scheduler interface {
	Schedule(task.Task, task.ExecutionType)
}

type scheduler struct {
	numWorkers int
	queue queue.Queue
}

func (s *scheduler) Schedule(t task.Task, executionType task.ExecutionType) {
	fmt.Println("Schedule item")
	item := queue.NewItem(t, executionType, int(utils.TimeInSeconds()) + executionType.TimeInSeconds)
	s.queue.Enqueue(item)
}

func NewScheduler(ctx context.Context, numWorkers int) Scheduler {
	s := &scheduler{numWorkers: numWorkers}
	s.queue = queue.NewQueue(nil)

	chn := make(chan task.Task, numWorkers)
	m := NewMaster(s.queue, chn)
	m.start(ctx)
	for i := 0; i < numWorkers; i++ {
		newWorker(chn).start(ctx)
	}
	return s
}
