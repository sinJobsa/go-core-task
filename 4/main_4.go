package main

import "fmt"

func ChangeSlice(slice1 []string, slice2 []string) []string {
	m := make(map[string]struct{}, len(slice2))
	result := make([]string, 0, len(slice1))
	for _, v := range slice2 {
		m[v] = struct{}{}
	}

	for _, v := range slice1 {
		if _, ok := m[v]; !ok {
			result = append(result, v)
		}
	}

	return result
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	result := ChangeSlice(slice1, slice2)
	fmt.Println(result)
}
