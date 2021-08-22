package scheduler

import (
	"context"
	"scheduler/task"
)

type worker struct {
	tasks chan task.Task
}

func newWorker(tasks chan task.Task) *worker {
	return &worker{tasks: tasks}
}

func (w *worker) start(ctx context.Context) {
	go func() {
		for {
			select {
			case <- ctx.Done():
				return
			default:
				item := <- w.tasks
				item.Run(ctx)
			}
		}
	}()
}
