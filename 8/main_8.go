package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type WaitGroup struct {
	count atomic.Int64
	done  chan struct{}
}

func NewWaitGroup() *WaitGroup {
	return &WaitGroup{
		done: make(chan struct{}),
	}
}

func (wg *WaitGroup) AddWg(delta int) {
	if delta < 0 {
		panic("WaitGroup: negative Add")
	}

	newCount := wg.count.Add(int64(delta))

	if newCount == 0 {
		close(wg.done)
	}
}

func (wg *WaitGroup) Done() {
	newCount := wg.count.Add(-1)
	if newCount < 0 {
		panic("WaitGroup: negative counter")
	}
	if newCount == 0 {
		close(wg.done)
	}
}

func (wg *WaitGroup) Wait() {
	<-wg.done
}

func main() {
	wg := NewWaitGroup()

	for i := 1; i < 10; i++ {
		wg.AddWg(1)
		go func() {
			defer wg.Done()
			fmt.Printf("goroutine: %d started\n", i)
			time.Sleep(200 * time.Millisecond)
			fmt.Printf("goroutine %d finished\n", i)
		}()
	}

	wg.Wait()
	fmt.Println("goroutines done")
}
