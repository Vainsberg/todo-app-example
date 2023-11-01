package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func get(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "GET")

}
func put(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "PUT")

}

func main() {
	router := httprouter.New()
	router.GET("/get", get)
	router.PUT("/put", put)
	http.ListenAndServe(":8080", router)

}
