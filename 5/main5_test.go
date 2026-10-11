package main

import (
	"reflect"
	"testing"
)

func TestSliceIntersection(t *testing.T) {
	tests := []struct {
		name   string
		input1 []int
		input2 []int
		want   []int
		flagOk bool
	}{
		{"Базовый Пример", []int{65, 3, 58, 678, 64}, []int{64, 2, 3, 43}, []int{3, 64}, true},
		{"дубликаты", []int{65, 3, 3, 3, 64}, []int{64, 2, 3, 43}, []int{3, 64}, true},
		{"пустой слайс", []int{1, 2, 3, 4, 5}, []int{6, 7, 8}, []int{}, false},
		{"дубликаты по всем", []int{1, 1, 1, 1}, []int{1, 1, 1}, []int{1}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, flag := SliceIntersection(tt.input1, tt.input2)
			if !reflect.DeepEqual(got, tt.want) || flag != tt.flagOk {
				t.Errorf("SliceIntersection() got = %v, want %v", got, tt.want)
			}
		})
	}
}
