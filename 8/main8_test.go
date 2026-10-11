package main

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestWaitGroup_WaitAll(t *testing.T) {
	wg := NewWaitGroup()
	var counter atomic.Int64

	const workers = 10
	for i := 0; i < workers; i++ {
		wg.AddWg(1)
		go func() {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond)
			counter.Add(1)
		}()
	}
	wg.Wait()

	if got := counter.Load(); got != workers {
		t.Errorf("got %d, want %d", got, workers)
	}
}

func TestWaitGroup_AddWgNegative(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("panic at negative: %v", r)
		}
	}()
	wg := NewWaitGroup()
	wg.AddWg(-1)
}

func TestWaitGroup_DoneMoreThanAdd(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic on Done more than Add")
		}
	}()

	wg := NewWaitGroup()
	wg.AddWg(1)
	wg.Done()
	wg.Done()
}
