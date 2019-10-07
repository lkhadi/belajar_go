package controllers

import "net/http"

func Route() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", health())
	return mux
}
