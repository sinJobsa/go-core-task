package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"все числа", []int{1, 2, 3, 4, 5}, []int{2, 4}},
		{"пустой", []int{}, []int{}},
		{"нечетные", []int{1, 3, 5, 7}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SliceExample(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceExample() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddElement(t *testing.T) {
	got := AddElement([]int{1, 2, 3, 4, 5}, 6)
	want := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("AddElement() = %v, want %v", got, want)
	}
}

func TestAddElement_Capacity(t *testing.T) {
	input := make([]int, 3, 6)
	_ = AddElement(input, 99)

	if len(input) != 3 {
		t.Errorf("AddElement() length = %v, want 3", len(input))
	}
}

func TestCopySlice_RealCopy(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	got := CopySlice(input)

	got[0] = 99

	if input[0] == 99 {
		t.Errorf("CopySlice() copied = %v, want %v", got, input)
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
		index int
	}{
		{"посередине", []int{1, 2, 3, 4, 5}, []int{1, 2, 4, 5}, 2},
		{"начало", []int{1, 2, 3}, []int{2, 3}, 0},
		{"конец", []int{1, 2, 3, 4}, []int{1, 2, 3}, 3},
		{"больше len", []int{1, 2, 3}, []int{1, 2, 3}, 10},
		{"отрицательный индекс", []int{1, 2, 3}, []int{1, 2, 3}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveElement(tt.input, tt.index)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
