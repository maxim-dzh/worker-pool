package pool

import "context"

type job interface {
	Process(ctx context.Context)
}

// Pool ...
type Pool interface {
	Start(ctx context.Context)
	Queue() (queue chan job)
	Close()
}

type pool struct {
	queue       chan job
	maxRoutines int
}

// Start ...
func (p *pool) Start(ctx context.Context) {
	for i := 0; i < p.maxRoutines; i++ {
		go func(ctx context.Context) {
			for job := range p.queue {
				job.Process(ctx)
			}
		}(ctx)
	}
}

// Queue ...
func (p *pool) Queue() (queue chan job) {
	return p.queue
}

// Close ...
func (p *pool) Close() {
	close(p.queue)
}

// NewPool ...
func NewPool(ctx context.Context, maxRoutines, chanLen int) Pool {
	return &pool{
		queue:       make(chan job, chanLen),
		maxRoutines: maxRoutines,
	}
}
