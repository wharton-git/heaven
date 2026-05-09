package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const port = ":3450"

type response map[string]any

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", getOnly(homeHandler))
	mux.HandleFunc("/api/info", getOnly(infoHandler))
	mux.HandleFunc("/api/status", getOnly(statusHandler))

	log.Printf("Backend Go demarre sur http://localhost%s", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		writeJSON(w, http.StatusNotFound, response{
			"error": "route introuvable",
		})
		return
	}

	writeJSON(w, http.StatusOK, response{
		"message": "Bienvenue sur le backend Go",
		"routes": []string{
			"GET /",
			"GET /api/info",
			"GET /api/status",
		},
	})
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, response{
		"nom":         "API Heaven",
		"description": "Murphy Law",
		"version":     "1.0.0",
		"auteur":      "Xeon",
	})
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, response{
		"status":    "ok",
		"service":   "backend-go",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

func getOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeJSON(w, http.StatusMethodNotAllowed, response{
				"error":   "methode non autorisee",
				"allowed": "GET",
			})
			return
		}

		next(w, r)
	}
}

func writeJSON(w http.ResponseWriter, status int, data response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Erreur JSON: %v", err)
	}
}
