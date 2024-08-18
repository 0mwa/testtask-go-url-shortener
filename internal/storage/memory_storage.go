package storage

import "fmt"

type memoryStorage struct {
	urlMap map[string]string
}

func NewMemoryStorage() Storage {
	return &memoryStorage{
		urlMap: make(map[string]string),
	}
}

func (m *memoryStorage) Write(sURL string, oURL string) error {
	if _, ok := m.urlMap[sURL]; ok {
		return fmt.Errorf("url already exists")
	}
	m.urlMap[sURL] = oURL
	return nil
}

func (m *memoryStorage) Read(sURL string) (string, error) {
	oURL, ok := m.urlMap[sURL]
	if !ok {
		return "", fmt.Errorf("url not found")
	}
	return oURL, nil
}
