// The inventory service: what is in the warehouse, and how much of it.
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	store := NewStore()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{
			"status": "ok", "service": "inventory", "items": store.Size(),
		})
	})

	mux.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		if below := r.URL.Query().Get("below"); below != "" {
			threshold, err := strconv.Atoi(below)
			if err != nil {
				writeJSON(w, http.StatusBadRequest, map[string]string{"error": "below must be a number"})
				return
			}
			writeJSON(w, http.StatusOK, map[string]any{"items": store.LowStock(threshold)})
			return
		}
		writeJSON(w, http.StatusOK, map[string]any{"items": store.All()})
	})

	mux.HandleFunc("/items/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/items/")
		parts := strings.Split(path, "/")
		sku := parts[0]

		if len(parts) == 2 && parts[1] == "reserve" && r.Method == http.MethodPost {
			var body struct {
				Amount int `json:"amount"`
			}
			_ = json.NewDecoder(r.Body).Decode(&body)
			item, ok := store.Reserve(sku, body.Amount)
			if !ok {
				writeJSON(w, http.StatusConflict, map[string]string{"error": "not enough stock, or no such sku"})
				return
			}
			writeJSON(w, http.StatusOK, item)
			return
		}

		item, ok := store.Find(sku)
		if !ok {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "no such sku"})
			return
		}
		writeJSON(w, http.StatusOK, item)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("inventory listening on %s", port)
	if err := http.ListenAndServe("0.0.0.0:"+port, mux); err != nil {
		log.Fatal(err)
	}
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}
