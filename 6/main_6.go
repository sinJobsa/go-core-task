package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
)

func GenerateInts(ctx context.Context, min, max int) <-chan int {
	ch := make(chan int)
	if min >= max {
		close(ch)
		return ch
	}

	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- rand.Intn(max-min) + min:
			}
		}
	}()

	return ch
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := GenerateInts(ctx, 0, 100)
	fmt.Println("до:", runtime.NumGoroutine()) // ~1
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("после:", runtime.NumGoroutine()) // ~6 (утечка!)
}
