package main

import (
	"math/rand/v2"
)

func CreateRandSlice(size int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = rand.IntN(100)
	}
	return result
}

func SliceExample(slice []int) []int {
	result := make([]int, 0, len(slice))
	for _, v := range slice {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

func AddElement(slice []int, element int) []int {
	result := make([]int, len(slice)+1)
	result[len(slice)] = element
	copy(result[:len(slice)], slice)
	return result
}

func CopySlice(slice []int) []int {
	result := make([]int, len(slice))
	copy(result, slice)
	return result
}

func RemoveElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	result := append(slice[:index], slice[index+1:]...)
	return result
}

func main() {
	result := SliceExample(CreateRandSlice(10))
	result = AddElement(result, 2000)
	newRes := CopySlice(result)
	result = RemoveElement(result, 1)
	newRes = RemoveElement(newRes, 0)
}
