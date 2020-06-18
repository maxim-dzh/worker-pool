package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kelseyhightower/envconfig"

	"github.com/maxim-dzh/worker-pool/pkg/image"
	"github.com/maxim-dzh/worker-pool/pkg/pool"
)

type configuration struct {
	QueueCapacity int `envconfig:"QUEUE_CAPACITY" required:"true"`
	WorkersNum    int `envconfig:"WORKERS_NUM" required:"true"`
	JobsNum       int `envconfig:"JOBS_NUM" required:"true"`
}

func main() {
	var cfg configuration
	if err := envconfig.Process("", &cfg); err != nil {
		log.Panicln(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := pool.NewPool(ctx, cfg.WorkersNum, cfg.QueueCapacity)
	defer p.Close()
	go p.Start(ctx)

	queue := p.Queue()
	for i := 0; i < cfg.JobsNum; i++ {
		queue <- image.NewImage()
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
	defer func(sig os.Signal) {
		cancel()
	}(<-signals)
}
