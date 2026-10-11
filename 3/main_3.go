package main

import "fmt"

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

func (s *StringIntMap) Add(key string, value int) {
	s.data[key] = value
}

func (s *StringIntMap) Remove(key string) {
	delete(s.data, key)
}

func (s *StringIntMap) Copy() map[string]int {
	m := make(map[string]int, len(s.data))
	for k, v := range s.data {
		m[k] = v
	}
	return m
}

func (s *StringIntMap) Exists(key string) bool {
	if _, ok := s.data[key]; ok {
		return true
	}
	return false
}

func (s *StringIntMap) Get(key string) (int, bool) {
	value, ok := s.data[key]
	return value, ok
}

func main() {
	m := NewStringIntMap()
	m.Add("a", 1)
	m.Add("b", 2)
	m.Add("c", 3)
	m.Remove("a")
	copied := m.Copy()
	fmt.Println(copied)
	fmt.Println(m.Exists("b"))
	fmt.Println(m.Exists("f"))
	fmt.Println(m.Get("c"))
}
