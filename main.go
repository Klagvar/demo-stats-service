package main

import (
	"log"
	"net/http"

	"github.com/Klagvar/demo-stats-service/handler"
	"github.com/Klagvar/demo-stats-service/storage"
)

func main() {
	store := storage.NewInMemoryStore()
	h := handler.NewStatsHandler(store)

	mux := http.NewServeMux()
	mux.Handle("/stats", h)

	addr := ":8080"
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
