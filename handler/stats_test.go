package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Klagvar/demo-stats-service/storage"
)

func TestStatsHandler_HappyPath(t *testing.T) {
	h := NewStatsHandler(storage.NewInMemoryStore())

	body, _ := json.Marshal(StatsRequest{Series: "x", Values: []float64{1, 2, 3}})
	req := httptest.NewRequest(http.MethodPost, "/stats", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status: want 200, got %d (body=%s)", rec.Code, rec.Body.String())
	}
	var resp StatsResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if resp.Mean != 2 {
		t.Fatalf("mean: want 2, got %v", resp.Mean)
	}
}
