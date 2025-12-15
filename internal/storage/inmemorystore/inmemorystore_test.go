package inmemorystore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorageGet(t *testing.T) {
	testCases := []struct {
		name        string
		tempStorage map[string]string
		key         string
		want        string
	}{
		{name: "Тест 1 - успех", tempStorage: map[string]string{"testKey": "testValue"}, key: "testKey", want: "testValue"},
		{name: "Тест 2 - возврат пустой строки при неверном ключе", tempStorage: map[string]string{"testKey": "testValue"}, key: "sdfgfdhg", want: ""},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s := &inMemoryStore{
				tempStorage: tt.tempStorage,
			}
			got, _ := s.Get(tt.key)
			assert.Equal(t, tt.want, got, "Результат не совпадает с ожиданием")
		})
	}
}

func Test_inMemoryStore_Set(t *testing.T) {
	tests := []struct {
		name        string
		tempStorage map[string]string
		url         string
		length      int
	}{
		{name: "Тест 1 - успех", tempStorage: map[string]string{}, url: "testKey", length: 1},
		{name: "Тест 2 - уже существует - успех", tempStorage: map[string]string{"testKey": "testValue"}, url: "testValue", length: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &inMemoryStore{
				tempStorage: tt.tempStorage,
			}
			s.Set(tt.url)
			assert.Len(t, s.tempStorage, tt.length, "Длина словаря не совпадает")
		})
	}
}
