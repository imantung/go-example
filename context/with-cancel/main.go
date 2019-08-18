package main

import (
	"context"
	"fmt"
	"time"
)

func process(ctx context.Context) {
	finish := make(chan int)

	go func() {
		// process work need 3 s
		time.Sleep(3 * time.Second)
		finish <- 1
	}()

	select {
	case <-ctx.Done():
		fmt.Println("cancelled")
	case <-finish:
		fmt.Println("finished")
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// trigger cancel after 10 ms
	go func() {
		time.Sleep(10 * time.Microsecond)
		cancel()
	}()

	process(ctx)
}
