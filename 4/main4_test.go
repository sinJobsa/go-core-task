package main

import (
	"reflect"
	"testing"
)

func TestChangeSlice(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date"}
	slice2 := []string{"banana", "date"}

	got := ChangeSlice(slice1, slice2)

	want := []string{"apple", "cherry"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestChangeSlice_FullIntersection(t *testing.T) {
	got := ChangeSlice([]string{"a", "b", "c"}, []string{"a", "b", "c"})

	if len(got) != 0 {
		t.Errorf("got %v, want empty slice", got)
	}
}


