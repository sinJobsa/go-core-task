package main

import (
	"reflect"
	"testing"
)

func TestNewStringIntMap(t *testing.T) {
	m := NewStringIntMap()
	if m == nil {
		t.Fatalf("NewStringIntMap() returned nil")
	}
	if m.data == nil {
		t.Fatalf("NewStringIntMap() returned nil")
	}
	if len(m.data) != 0 {
		t.Fatalf("NewStringIntMap() returned wrong length")
	}
}

func TestAdd(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 1)
	if v, ok := m.Get("a"); !ok || v != 1 {
		t.Errorf("Get(\"a\") returned wrong value")
	}
}

func TestRemove(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 1)
	m.Remove("a")
	if ok := m.Exists("a"); ok {
		t.Errorf("Get(\"a\") returned wrong value")
	}
}

func TestRemove_NoExistKey(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 1)
	m.Remove("b")
	if !m.Exists("a") {
		t.Error("map broken by existing key")
	}
}

func TestExists(t *testing.T) {
	m := NewStringIntMap()
	m.Add("one", 1)
	m.Add("zero", 0)

	tests := []struct {
		key  string
		want bool
	}{
		{"one", true},
		{"zero", true},
		{"missing", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			if got := m.Exists(tt.key); got != tt.want {
				t.Errorf("Exists(%q) = %v, want %v", tt.key, got, tt.want)
			}
		})
	}
}

func TestCopy(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 1)
	m.Add("b", 2)
	m.Add("c", 3)
	got := m.Copy()
	want := map[string]int{"a": 1, "b": 2, "c": 3}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Copy() = %v, want %v", got, want)
	}
}

func TestGet(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 1)

	if got, ok := m.Get("a"); !ok || got != 1 {
		t.Errorf("Get(\"a\") returned wrong value")
	}
}

func TestGet_NonExistingKey(t *testing.T) {
	m := NewStringIntMap()

	v, ok := m.Get("missing")
	if ok {
		t.Error("ok = true, want false")
	}
	if v != 0 {
		t.Errorf("value = %d, want 0 (zero value)", v)
	}
}
