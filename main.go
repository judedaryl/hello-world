package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rs/cors"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{
			"hello": "world",
			"message":  r.URL.Query().Get("message"),
		}
		data, _ := json.MarshalIndent(response, "", "  ")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})

	mux.HandleFunc("GET /echo", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{
			"message": r.URL.Query().Get("message"),
		}
		data, _ := json.MarshalIndent(response, "", "  ")
		statusCode, err := strconv.ParseInt(r.URL.Query().Get("statusCode"), 10, 16)
		if err != nil {
			w.WriteHeader(http.StatusOK)
		} else if statusCode < 100 || statusCode > 999 {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(int(statusCode))
		}
		w.Write(data)
	})

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]bool{
			"healthy": true,
		}
		data, _ := json.MarshalIndent(response, "", "  ")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})

	http.ListenAndServe(":8080", cors.AllowAll().Handler(mux))
	// quitChannel := make(chan os.Signal, 1)
	// signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	// <-quitChannel

}
