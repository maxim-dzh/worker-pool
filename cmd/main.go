package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/maxim-dzh/worker-pool/pkg/image"
	"github.com/maxim-dzh/worker-pool/pkg/pool"
)

const (
	chanLen  = 100
	routines = 100
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := pool.NewPool(ctx, routines, chanLen)
	defer p.Close()
	go p.Start(ctx)

	queue := p.Queue()
	for i := 0; i < chanLen; i++ {
		queue <- image.NewImage()
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
	defer func(sig os.Signal) {
		cancel()
	}(<-signals)
}
