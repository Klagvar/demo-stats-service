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
		t.Errorf("expected error 'storage: empty name', got: %v", err)
	}
}

func TestInMemoryStore_Append_NilValue(t *testing.T) {
	store := NewInMemoryStore()
	name := "test"
	value := float64(0) // Testing with a zero value

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
	values := []float64{1.0, 2.0, 3.0}

	for _, value := range values {
		err := store.Append(name, value)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	}

	if len(store.data[name]) != len(values) {
		t.Errorf("expected %d values, got %d", len(values), len(store.data[name]))
	}

	for i, value := range values {
		if store.data[name][i] != value {
			t.Errorf("expected value %v at index %d, got %v", value, i, store.data[name][i])
		}
	}
}

func TestInMemoryStore_Append_EmptyName_Error(t *testing.T) {
	store := &InMemoryStore{data: make(map[string][]float64)}
	err := store.Append("", 1.0)
	if err == nil || err.Error() != "storage: empty name" {
		t.Errorf("expected error 'storage: empty name', got: %v", err)
	}
}
