package scheduler

import (
	"context"
	"fmt"
	"scheduler/queue"
	"scheduler/task"
	"scheduler/utils"
	"time"
)

type master struct {
	queue queue.Queue
	task chan task.Task
}

func NewMaster(queue queue.Queue, task chan task.Task) *master {
	return &master{queue: queue, task: task}
}

func (m *master) start(ctx context.Context) {

	go func() {
		for {
			select {
			case <- ctx.Done():
				return
			default:
				item := m.queue.Dequeue()
				switch item.(type) {
				case queue.Item:
					fmt.Printf("picked task: %v, priority: %v, current time: %v\n", item.Task().Name(), item.Priority(), utils.TimeInSeconds())
					if item.Priority() <= int(utils.TimeInSeconds()) {
						m.task <- item.Task()
						if item.ExecutionType().Periodic {
							m.queue.Enqueue(
								queue.NewItem(
									item.Task(), item.ExecutionType(),
									item.ExecutionType().TimeInSeconds + int(utils.TimeInSeconds())))
						}
					} else {
						m.queue.Enqueue(item)
						time.Sleep(2*time.Second)
					}
				default:
					fmt.Println("item is nil")
					time.Sleep(1*time.Second)
				}
			}
		}
	}()
}
