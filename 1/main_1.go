package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

type TaskResult struct {
	Types       map[string]string
	Combined    string
	Runes       []rune
	RunesSalted []rune
	Hash        string
}

func FormatTypes(values map[string]any) map[string]string {
	result := make(map[string]string, len(values))
	for name, val := range values {
		result[name] = fmt.Sprintf("%T", val)
	}
	return result
}

func toString(value any) string {
	switch v := value.(type) {
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	case string:
		return v
	case complex64:
		return strconv.FormatComplex(complex128(v), 'f', -1, 64)
	default:
		return fmt.Sprintf("%v", value)
	}
}

// не покрываю тестом, схож с проверкой FomatTypes
func MakeStrings(values map[string]any) map[string]string {
	result := make(map[string]string, len(values))
	for name, val := range values {
		result[name] = toString(val)
	}
	return result
}

func CombinedStrings(mapStr map[string]string, order []string) string {
	var result strings.Builder
	for _, name := range order {
		if val, ok := mapStr[name]; ok {
			result.WriteString(val)
		}
	}
	return result.String()
}

func MakeRunes(str string) []rune {
	return []rune(str)
}

func AddSalt(strRunes []rune, salt string) []rune {
	mid := len(strRunes) / 2
	addToMid := []rune(salt)

	newResult := make([]rune, 0, len(strRunes)+len(addToMid))
	newResult = append(newResult, strRunes[:mid]...)
	newResult = append(newResult, addToMid...)
	newResult = append(newResult, strRunes[mid:]...)
	return newResult
}

func Hash(strRunes []rune) string {
	hash := sha256.Sum256([]byte(string(strRunes)))
	return hex.EncodeToString(hash[:])
}

func Run(value map[string]any, order []string, salt string) TaskResult {
	types := FormatTypes(value)
	stringsMaked := MakeStrings(value)
	combinedStr := CombinedStrings(stringsMaked, order)
	runesMaked := MakeRunes(combinedStr)
	saltMaked := AddSalt(runesMaked, salt)
	hashMaked := Hash(saltMaked)
	return TaskResult{
		Types:       types,
		Combined:    combinedStr,
		Runes:       runesMaked,
		RunesSalted: saltMaked,
		Hash:        hashMaked,
	}
}

func main() {
	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	order := []string{"numDecimal", "numOctal", "numHexadecimal", "pi", "name", "isActive", "complexNum"}

	m := map[string]any{
		"numDecimal":     numDecimal,
		"numOctal":       numOctal,
		"numHexadecimal": numHexadecimal,
		"pi":             pi,
		"name":           name,
		"isActive":       isActive,
		"complexNum":     complexNum,
	}
	salt := "go-2024"
	result := Run(m, order, salt)

	fmt.Println(result.Types)
	fmt.Println(result.Hash)

}
