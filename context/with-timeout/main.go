package main

import (
	"context"
	"fmt"
	"time"
)

func process(ctx context.Context) {
	finish := make(chan int)

	go func() {
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Microsecond)
	defer cancel()

	process(ctx)
}
