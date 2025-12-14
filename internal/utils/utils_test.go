package utils

import "testing"

func TestStringWithCharset(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{name: "Тест 1", length: 20},
		{name: "Тест 2", length: 1},
		{name: "Тест 3", length: 5},
		{name: "Тест 4", length: 999},
		{name: "Тест 5", length: 9999},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRandomString(tt.length); len(got) != tt.length {
				t.Errorf("Count StringWithCharset() = %v, want %v", got, tt.length)
			}
		})
	}
}
