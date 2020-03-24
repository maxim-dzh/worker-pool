package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/maxim-dzh/worker-pool/pkg/image"
	"github.com/maxim-dzh/worker-pool/pkg/pool"
)

const (
	chanLen = 100
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	signals := make(chan os.Signal, 1)
	go func() {
		signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
		<-signals
		cancel()
	}()

	p := pool.NewPool(
		ctx,
		runtime.GOMAXPROCS(0),
		chanLen,
	)
	queue := p.Queue()
	for i := 0; i < chanLen; i++ {
		queue <- image.NewImage()
	}
	fmt.Scanln()
}
