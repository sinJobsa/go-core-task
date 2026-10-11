package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestInChan_ReturnValuers(t *testing.T) {
	slice := []int{1, 2, 3}
	out := InChan(slice)

	got := []int{}

	for v := range out {
		got = append(got, v)
	}

	if !reflect.DeepEqual(got, slice) {
		t.Errorf("got %v, want %v", got, slice)
	}
}

func TestFanInAndCollect(t *testing.T) {
	ch1 := InChan([]int{1, 2, 3})
	ch2 := InChan([]int{4, 5, 6})

	result := FanIn(ch1, ch2)
	got := Collect(result)
	sort.Ints(got)

	want := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
