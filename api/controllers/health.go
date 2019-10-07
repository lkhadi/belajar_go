package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/lkhadi/belajar_go/tree/master/api/views"
)

func health() http.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
	}
}
