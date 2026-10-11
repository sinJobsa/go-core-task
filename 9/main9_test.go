package main

import (
	"context"
	"reflect"
	"testing"
)

func TestCollector(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nums := []uint8{1, 2, 3, 4}
	want := []float64{1, 8, 27, 64}

	res := Collector(Cube(ctx, Producer(ctx, nums)))
	if !reflect.DeepEqual(res, want) {
		t.Errorf("want: %v, got: %v", want, res)
	}
}
