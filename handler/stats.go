package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Klagvar/demo-stats-service/core"
	"github.com/Klagvar/demo-stats-service/storage"
)

type StatsRequest struct {
	Series string    `json:"series"`
	Values []float64 `json:"values"`
}

type StatsResponse struct {
	Mean float64 `json:"mean"`
	Min  float64 `json:"min"`
	Max  float64 `json:"max"`
}

type StatsHandler struct {
	Store *storage.InMemoryStore
}

func NewStatsHandler(store *storage.InMemoryStore) *StatsHandler {
	return &StatsHandler{Store: store}
}

func (h *StatsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req StatsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	if req.Series == "" {
		http.Error(w, "series required", http.StatusBadRequest)
		return
	}
	if len(req.Values) == 0 {
		http.Error(w, "values required", http.StatusBadRequest)
		return
	}

	if err := h.Store.Save(req.Series, req.Values); err != nil {
		http.Error(w, "storage error", http.StatusInternalServerError)
		return
	}

	mean, err := core.Mean(req.Values)
	if err != nil {
		if errors.Is(err, core.ErrEmpty) {
			http.Error(w, "no values", http.StatusBadRequest)
			return
		}
		http.Error(w, "internal", http.StatusInternalServerError)
		return
	}
	min, _ := core.Min(req.Values)
	max, _ := core.Max(req.Values)

	resp := StatsResponse{
		Mean: core.Round(mean, 4),
		Min:  min,
		Max:  max,
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
