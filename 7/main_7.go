package main

import (
	"fmt"
	"sync"
)

func FanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}

	for _, ch := range channels {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func InChan(slice []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, v := range slice {
			out <- v
		}
	}()
	return out
}

func Collect(in <-chan int) []int {
	s := make([]int, 0, len(in))
	for v := range in {
		s = append(s, v)
	}
	return s
}

func main() {
	ch1 := InChan([]int{1, 2, 3})
	ch2 := InChan([]int{4, 5, 6})
	ch3 := InChan([]int{7, 8, 9})

	fanIn := FanIn(ch1, ch2, ch3)

	results := Collect(fanIn)

	fmt.Println(results)

}
