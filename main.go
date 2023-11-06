package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var generalmap = make(map[string]string)

type Request struct {
	Message string
}

func get(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var requestget Request

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
		return
	}
	r.Body.Close()

	errors := json.Unmarshal(body, &requestget)
	if errors != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Bad Request")
		return
	}
	if requestget.Message == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Empty element is not allowed")
		return
	}

	value, exists := generalmap[requestget.Message]
	if exists {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Get value:", value)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Not Found")
	}
}

func put(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var requestget Request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
		return
	}
	r.Body.Close()

	errors := json.Unmarshal(body, &requestget)
	if errors != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Bad Request")
		return
	}
	if requestget.Message == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Empty element is not allowed")
		return
	}

	generalmap[requestget.Message] = requestget.Message
	w.WriteHeader(http.StatusCreated)
}

func deleted(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var requestget Request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
		return
	}
	r.Body.Close()

	errors := json.Unmarshal(body, &requestget)
	if errors != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Bad Request")
		return
	}
	if requestget.Message == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Empty element is not allowed")
		return
	}

	delete(generalmap, requestget.Message)
	w.WriteHeader(http.StatusNoContent)
}

func post(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var requestget Request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
		return
	}
	r.Body.Close()

	errors := json.Unmarshal(body, &requestget)
	if errors != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Bad Request")
		return
	}
	if requestget.Message == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Empty element is not allowed")
		return
	}

	if generalmap[requestget.Message] == requestget.Message {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "An element is already being selected")

	} else {
		generalmap[requestget.Message] = requestget.Message
		w.WriteHeader(http.StatusAccepted)
	}

}

func main() {
	router := httprouter.New()
	router.GET("/get", get)
	router.PUT("/put", put)
	router.DELETE("/delete", deleted)
	router.POST("/post", post)
	http.ListenAndServe(":8080", router)
}
