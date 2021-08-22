package task

import "context"

type Task interface {
	Name() string
	Run(ctx context.Context)
}

