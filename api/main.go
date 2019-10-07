package main

import (
	"net/http"
	"github.com/lkhadi/belajar_go/tree/master/api/controllers"
)

func main() {
	mux := controllers.Route()
	http.ListenAndServe("localhost:80", mux)
}
