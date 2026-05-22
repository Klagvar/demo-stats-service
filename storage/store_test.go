package storage

import "testing"

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
