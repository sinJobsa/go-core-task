package main

import (
	"context"
	"fmt"
)

func Producer(ctx context.Context, nums []uint8) <-chan uint8 {
	out := make(chan uint8)
	go func() {
		defer close(out)
		for _, v := range nums {
			select {
			case <-ctx.Done():
				return
			case out <- v:
			}
		}
	}()
	return out
}

func Cube(ctx context.Context, ch <-chan uint8) <-chan float64 {
	out := make(chan float64)
	go func() {
		defer close(out)
		for v := range ch {
			select {
			case <-ctx.Done():
				return
			case out <- float64(v) * float64(v) * float64(v):
			}
		}
	}()
	return out
}

func Collector(ch <-chan float64) []float64 {
	slice := make([]float64, 0)
	for v := range ch {
		slice = append(slice, v)
	}
	return slice
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nums := []uint8{1, 2, 3, 4}

	res := Collector(Cube(ctx, Producer(ctx, nums)))

	fmt.Println(res)

}
