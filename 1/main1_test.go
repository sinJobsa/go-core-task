package main

import (
	"reflect"
	"testing"
)

func TestFormatTypes(t *testing.T) {
	values := map[string]any{
		"int":     42,
		"float":   3.14,
		"string":  "Golang",
		"bool":    true,
		"complex": complex64(1 + 2i),
	}

	formated := FormatTypes(values)

	baseTypes := map[string]string{
		"int":     "int",
		"float":   "float64",
		"string":  "string",
		"bool":    "bool",
		"complex": "complex64",
	}

	if !reflect.DeepEqual(formated, baseTypes) {
		t.Error("FormatTypes does not match")
	}

}

func TestToString(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  string
	}{
		{"int", 42, "42"},
		{"int zero", 0, "0"},
		{"int negative", -7, "-7"},
		{"float64", 3.14, "3.14"},
		{"float64 whole", 14.0, "14"},
		{"bool true", true, "true"},
		{"bool false", false, "false"},
		{"string", "Golang", "Golang"},
		{"string empty", "", ""},
		{"complex64", complex64(1 + 2i), "(1+2i)"},
		{"complex128 via default", complex128(1 + 2i), "(1+2i)"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := toString(test.input)
			if got != test.want {
				t.Errorf("toString(%v) = %v, want %v", test.input, got, test.want)
			}
		})
	}
}

func TestCombinedStrings(t *testing.T) {
	values := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
		"d": "4",
		"e": "5",
	}

	order := []string{"a", "b", "c", "d", "e", "f"}

	want := "12345"

	got := CombinedStrings(values, order)

	if got != want {
		t.Errorf("CombinedStrings(%v) = %v, want %v", values, got, want)
	}
}

func TestCombinedStrings_IgnoredKeys(t *testing.T) {
	values := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
		"d": "4",
		"e": "5",
	}
	order := []string{"a", "b", "c"}
	want := "123"
	got := CombinedStrings(values, order)
	if got != want {
		t.Errorf("CombinedStrings(%v) = %v, want %v", values, got, want)
	}
}

func TestMakeRunes(t *testing.T) {
	got := MakeRunes("123")
	want := []rune{'1', '2', '3'}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MakeRunes(%v) = %v, want %v", got, want, got)
	}
}

func TestAddSalt(t *testing.T) {
	runes := []rune("123")
	got := AddSalt(runes, "GGG")
	want := []rune("1GGG23")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("AddSalt(%v) = %v, want %v", runes, got, want)
	}
}

func TestHash(t *testing.T) {
	runes := []rune("123")
	got := Hash(runes)
	got2 := Hash(runes)
	if got != got2 {
		t.Errorf("Hash(%v) = %v, want %v", runes, got, got2)
	}
}

func TestHash_KnownValue(t *testing.T) {
	got := Hash([]rune(""))
	want := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"

	if got != want {
		t.Errorf("Hash(\"\") = %s, want %s", got, want)
	}
}

func TestHash_DifferentRunes(t *testing.T) {
	got1 := Hash([]rune("123"))
	got2 := Hash([]rune("456"))

	if got1 == got2 {
		t.Errorf("Hash is Equal for the different runes")
	}
}

func TestRun_DifferentOrder_CheckHash(t *testing.T) {
	values := map[string]any{
		"a": "1",
		"b": "2",
	}

	run1 := Run(values, []string{"a", "b"}, "XXX")
	run2 := Run(values, []string{"b", "a"}, "YYY")

	if run1.Hash == run2.Hash {
		t.Errorf("Runs should not have the same hash")
	}
}
