package pool

import "context"

type job interface {
	Process(ctx context.Context)
}

// Pool ...
type Pool interface {
	Queue() (queue chan job)
}

type pool struct {
	queue chan job
}

// Queue ...
func (p *pool) Queue() (queue chan job) {
	return p.queue
}

// NewPool ...
func NewPool(ctx context.Context, maxRoutines, chanLen int) Pool {
	var p pool
	p.queue = make(chan job, chanLen)
	for i := 0; i < maxRoutines; i++ {
		go func(ctx context.Context) {
			for job := range p.queue {
				job.Process(ctx)
			}
		}(ctx)
	}
	return &p
}
