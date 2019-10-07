package main

import "net/http"

func main() {
	mux := controllers.Route()
	http.ListenAndServe("localhost:80", mux)
}
