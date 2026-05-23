package storage

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("storage: series not found")

type Store interface {
	Save(name string, values []float64) error
	Load(name string) ([]float64, error)
	Append(name string, value float64) error
}

type InMemoryStore struct {
	mu   sync.RWMutex
	data map[string][]float64
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{data: make(map[string][]float64)}
}

func (s *InMemoryStore) Save(name string, values []float64) error {
	if name == "" {
		return errors.New("storage: empty name")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	copied := make([]float64, len(values))
	copy(copied, values)
	s.data[name] = copied
	return nil
}

func (s *InMemoryStore) Load(name string) ([]float64, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	values, ok := s.data[name]
	if !ok {
		return nil, ErrNotFound
	}
	copied := make([]float64, len(values))
	copy(copied, values)
	return copied, nil
}

func (s *InMemoryStore) Append(name string, value float64) error {
	if name == "" {
		return errors.New("storage: empty name")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[name] = append(s.data[name], value)
	return nil
}
