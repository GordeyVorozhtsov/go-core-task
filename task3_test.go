package main

import (
	"testing"
)

var StringIntMap = make(map[string]int)

func Add(key string, value int) map[string]int {
	StringIntMap[key] = value
	return StringIntMap
}
func Remove(m map[string]int, key string) map[string]int {
	delete(m, key)
	return m
}
func Copy(mapa map[string]int) map[string]int {
	newMap := make(map[string]int)
	for k, v := range mapa {
		newMap[k] = v
	}
	return newMap
}
func Exists(key string) bool {
	if _, ok := StringIntMap[key]; ok {
		return true
	}
	return false
}
func Get(key string) (int, bool) {

	if val, ok := StringIntMap[key]; ok {
		return val, true
	}
	return 0, false
}

func TestAdd(t *testing.T) {
	testCases := []struct {
		key      string
		value    int
		expected map[string]int
	}{
		{"a", 1, map[string]int{"a": 1}},
		{"b", 2, map[string]int{"a": 1, "b": 2}},
		{"c", 3, map[string]int{"a": 1, "b": 2, "c": 3}},
	}
	for _, test := range testCases { // используйте `tc` для доступа к каждому тестовому случаю
		res := Add(test.key, test.value)

		// Проверяем, согласуются ли все ключи и значения
		for k, v := range test.expected {
			if res[k] != v {
				t.Errorf("for key=%s, expected %v but got %v", test.key, test.expected, res)
			}
		}
		t.Logf("Result: %v", res)
	}
}
func TestRemove(t *testing.T) {
	testCases := []struct {
		StringIntMap map[string]int
		key          string
		expected     map[string]int
	}{
		{map[string]int{"a": 1, "b": 2, "c": 3}, "b", map[string]int{"a": 1, "c": 3}},
		{map[string]int{"a": 1, "b": 2, "c": 3}, "c", map[string]int{"a": 1, "b": 2}},
	}

	for _, tc := range testCases { // Используйте переменную tc для доступа к тестовому случаю
		result := Remove(tc.StringIntMap, tc.key)
		for i, v := range result {
			if v != tc.expected[i] {
				t.Errorf("expected %v but got %v", tc.expected, result)
			}
		}
		t.Logf("Result: %v", result)
	}
}
func TestExists(t *testing.T) {
	testCases := []struct {
		key      string
		expected bool
	}{
		{"a", true},
		{"b", true},
		{"notInMap", false},
	}

	for _, test := range testCases {
		result := Exists(test.key)
		if result != test.expected {
			t.Errorf("Exists(%q) = %v; want %v", test.key, result, test.expected)
		}
		t.Logf("Result: %v", result)
	}

}
func TestGet(t *testing.T) {
	testCases := []struct {
		key      string
		value    int
		expected bool
	}{
		{"a", 1, true},
		{"b", 2, true},
	}

	for _, test := range testCases {
		result, ok := Get(test.key)
		if result != test.value && ok != test.expected {
			t.Errorf("Exists(%q) = %v; want %v", test.key, result, test.expected)
		}
		t.Logf("Result: %v %v", result, ok)
	}

}
