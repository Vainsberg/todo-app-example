package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var generalmap map[string]string

func get(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "GET")
	value, exists := generalmap["get"]
	if exists {
		fmt.Fprintln(w, "Get value:", value)

	} else {
		fmt.Fprintln(w, "Not found get")

	}

}
func put(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "PUT")

	generalmap["put"] = "This is the PUT handler"

}

func main() {
	router := httprouter.New()
	router.GET("/get", get)
	router.PUT("/put", put)
	http.ListenAndServe(":8080", router)

}
