package main

import (
	"context"
	"testing"
	"time"
)

func TestGenerateInts_InvalidRange(t *testing.T) {
	tests := []struct {
		name     string
		inputMax int
		inputMin int
	}{
		{"min==max", 5, 5},
		{"min > max", 0, 5},
		{"отрицательные", -5, -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			ch := GenerateInts(ctx, test.inputMin, test.inputMax)

			select {
			case _, ok := <-ch:
				if ok {
					t.Errorf("Channel closed unexpectedly")
				}
			case <-time.After(100 * time.Millisecond):
				t.Errorf("channel is not closed for invalid range")
			}
		})
	}
}

func TestGenerateInts_ValidRange(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := GenerateInts(ctx, 0, 100)

	for i := 0; i < 10; i++ {
		select {
		case v, ok := <-ch:
			if !ok {
				t.Errorf("Channel closed unexpectedly, iteration %d", i)
			}
			if v < 0 || v > 100 {
				t.Errorf("invalid range, iteration %d", i)
			}
		case <-time.After(100 * time.Millisecond):
			t.Errorf("timeout waiting for iteration %d", i)
		}
	}
}

func TestGenerateInts_ChannelBuffered(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := GenerateInts(ctx, 0, 100)
	if cap(ch) != 0 {
		t.Errorf("channel is buffered")
	}
}
