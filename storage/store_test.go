package storage

import (
	"testing"
)

func TestInMemoryStore_SaveLoad(t *testing.T) {
	s := NewInMemoryStore()
	if err := s.Save("a", []float64{1, 2, 3}); err != nil {
		t.Fatalf("save: %v", err)
	}
	got, err := s.Load("a")
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	if len(got) != 3 || got[0] != 1 || got[2] != 3 {
		t.Fatalf("unexpected values: %v", got)
	}
}

func TestInMemoryStore_Append_EmptyName(t *testing.T) {
	store := &InMemoryStore{data: make(map[string][]float64)}
	err := store.Append("", 1.0)

	if err == nil || err.Error() != "storage: empty name" {
		t.Errorf("expected error 'storage: empty name', got %v", err)
	}
}

func TestInMemoryStore_Append_ValidName(t *testing.T) {
	store := NewInMemoryStore()
	name := "test"
	value := 1.23

	err := store.Append(name, value)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(store.data[name]) != 1 || store.data[name][0] != value {
		t.Errorf("expected data to contain %v, got %v", value, store.data[name])
	}
}

func TestInMemoryStore_Append_MultipleValues(t *testing.T) {
	store := NewInMemoryStore()
	name := "test"
	values := []float64{1.23, 4.56, 7.89}

	for _, value := range values {
		err := store.Append(name, value)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	}

	if len(store.data[name]) != len(values) {
		t.Errorf("expected data length to be %d, got %d", len(values), len(store.data[name]))
	}

	for i, value := range values {
		if store.data[name][i] != value {
			t.Errorf("expected value at index %d to be %v, got %v", i, value, store.data[name][i])
		}
	}
}
