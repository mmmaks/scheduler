package queue

import "scheduler/task"

type Item interface {
	Task() task.Task
	ExecutionType() task.ExecutionType
	Priority() int
}

type item struct {
	task          task.Task
	executionType task.ExecutionType
	priority      int
}

func (i *item) Task() task.Task {
	return i.task
}

func (i *item) ExecutionType() task.ExecutionType {
	return i.executionType
}

func (i item) Priority() int {
	return i.priority
}

func NewItem(task task.Task, executionType task.ExecutionType, priority int) Item {
	return &item{task: task, executionType: executionType, priority: priority}
}
