package main

import (
	"encoding/json"
	"net/http"
)

type Health struct {
	Code int         "json:code"
	Body interface{} "json:body"
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.WriteHeader(200)
			data := Health{
				Code: 200,
				Body: "Running",
			}
			json.NewEncoder(w).Encode(data)
		} else {
			w.WriteHeader(500)
			data := Health{
				Code: 500,
				Body: "Something Wrong",
			}
			json.NewEncoder(w).Encode(data)
		}
	})
	http.ListenAndServe("localhost:80", mux)
}
