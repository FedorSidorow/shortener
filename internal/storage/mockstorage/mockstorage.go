package mockstorage

type mockStorage struct {
	tempStorage map[string]string
}

func NewStorage() (*mockStorage, error) {
	s := &mockStorage{}
	s.tempStorage = make(map[string]string, 0)
	return s, nil
}

func (s *mockStorage) Set(url string) (string, error) {
	toReturn := "EwHXdJfB"
	if value, ok := s.tempStorage[toReturn]; ok {
		if value == url {
			return toReturn, nil
		}
	}
	s.tempStorage[toReturn] = url
	return toReturn, nil
}

func (s *mockStorage) Get(key string) (string, error) {

	fullURL, ok := s.tempStorage[key]
	if !ok {
		return "", nil
	}

	return fullURL, nil
}
