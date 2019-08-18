package main

import (
	"context"
	"fmt"
)

func process(ctx context.Context) {
	value := ctx.Value("message")
	fmt.Println(value)
}

func main() {
	ctx := context.WithValue(nil, "message", "Hello World")
	process(ctx)
}
