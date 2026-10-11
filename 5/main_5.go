package main

import "fmt"

func SliceIntersection(slice1, slice2 []int) ([]int, bool) {
	m := make(map[int]struct{})
	result := make([]int, 0, len(slice1)+len(slice2))
	seen := make(map[int]struct{}, len(slice1)+len(slice2))

	if len(slice1) <= len(slice2) {
		for _, v1 := range slice1 {
			m[v1] = struct{}{}
		}
		for _, v2 := range slice2 {
			if _, ok := m[v2]; ok {
				if _, dup := seen[v2]; !dup {
					seen[v2] = struct{}{}
					result = append(result, v2)
				}
			}
		}
	} else {
		for _, v2 := range slice2 {
			m[v2] = struct{}{}
		}

		for _, v1 := range slice1 {
			if _, ok := m[v1]; ok {
				if _, dup := seen[v1]; !dup {
					seen[v1] = struct{}{}
					result = append(result, v1)
				}
			}
		}
	}

	return result, len(result) > 0

}

func main() {
	a := []int{1, 2, 3}
	b := []int{3, 3, 3}
	fmt.Println(SliceIntersection(a, b))
}
